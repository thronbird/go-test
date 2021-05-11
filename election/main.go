package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type electionKey string

var (
	engineKey           = electionKey("engine")
	defaultElectionName = "themisLeader" //无所谓的字段，只是用来标识选举类型
	ErrElectionNoEngine = errors.New("election: no engine")
)

type Election struct {
	LeaderName string
}

type ElectionRecord struct {
	Id           uint32    `xorm:"INT notnull autoincr pk"`
	ElectionName string    `xorm:"VARCHAR(32) notnull default '' unique"` //无所谓的字段，只是用来标识选举类型
	LeaderName   string    `xorm:"VARCHAR(32) notnull default ''"`
	LastUpdate   time.Time `xorm:"TIMESTAMP notnull default 'CURRENT_TIMESTAMP'"`
}

func NewElection(name string) *Election {
	return &Election{LeaderName: name}
}

// Campaign puts a value as eligible for the election.
// It blocks until it is elected, an error occurs, or the context is cancelled.
func (e *Election) Campaign(ctx context.Context) error {

	_, ok := ctx.Value(engineKey).(*xorm.Engine)
	if !ok {
		return ErrElectionNoEngine
	}

	for {
		res, err := campaign(ctx, e.LeaderName)
		if err != nil {
			return err
		}

		if res {
			// update record every 5 seconds if we became leader.
			go func(ctx context.Context, leader string) error {
				for {
					_, err := campaign(ctx, leader)
					if err != nil {
						return err
					}
					select {
					case <-ctx.Done():
						return ctx.Err()
					case <-time.After(6 * time.Second):
						continue
					}
				}
			}(ctx, e.LeaderName)
			return nil
		} else {
			// wait 20 seconds and campaign again
			time.Sleep(20 * time.Second)
		}
	}
}

// IsLeader query engine if we are the Leader.
func (e *Election) IsLeader(ctx context.Context) (bool, error) {

	engine, ok := ctx.Value(engineKey).(*xorm.Engine)
	if !ok {
		return false, ErrElectionNoEngine
	}

	sql := `SELECT COUNT(*) as is_leader FROM election_record where election_name=? and leader_name=?`

	res, err := engine.Query(sql, defaultElectionName, e.LeaderName)

	if err != nil {
		return false, err
	}

	if len(res) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func campaign(ctx context.Context, leader string) (bool, error) {
	engine, ok := ctx.Value(engineKey).(*xorm.Engine)
	if !ok {
		return false, ErrElectionNoEngine
	}

	sql := `INSERT IGNORE INTO election_record (election_name, leader_name, last_update) VALUES (?, ?, ?)
			ON DUPLICATE KEY UPDATE
			leader_name = IF(last_update < DATE_SUB(VALUES(last_update), INTERVAL 30 SECOND), VALUES(leader_name), leader_name),
			last_update = IF(leader_name = VALUES(leader_name), VALUES(last_update), last_update)`
	res, err := engine.Exec(sql, defaultElectionName, leader, time.Now())

	if err != nil {
		return false, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if affected >= 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func RegisterModel(engine *xorm.Engine) {

	err := engine.Sync2(new(ElectionRecord))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: %s <name>", os.Args[0])
		os.Exit(-1)
	}
	leader_name := os.Args[1]

	engine, err := xorm.NewEngine("mysql", "root:o0o0o0@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer engine.Close()

	RegisterModel(engine)

	election := NewElection(leader_name)
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, engineKey, engine)

	for {
		if err := election.Campaign(ctx); err != nil {
			fmt.Println(err)
			cancel()
			break
		}
		log.Printf("Now %s am the leader.\n", leader_name)
		for {
			isLeader, err := election.IsLeader(ctx)
			if err != nil {
				cancel()
				log.Fatal(err)
			}
			if !isLeader {
				break
			}
			// do something
			log.Println("do something... sleep 20")
			time.Sleep(20 * time.Second)
		}
	}
}
