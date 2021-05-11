#!/bin/bash

if [ $# -ne 1 ]; then
    echo "usage:$0 <node name>"
    exit
fi
node=$1

get_leader() {
    mysql -e "select node from test.election where name='monitor'" | grep -v node
}

get_time() {
    mysql -e "select now()" | grep -v now
}

for i in $(seq 1 300)
do
    old_leader=$(get_leader)
    now=$(get_time)
    mysql -e "insert ignore into test.election (name, node, last_seen) values ('monitor', '${node}', '${now}') on duplicate key update node = if(last_seen < DATE_SUB('${now}', INTERVAL 30 SECOND), '${node}', '${old_leader}'), last_seen = if(node = values(node), values(last_seen), last_seen)"
    new_leader=$(get_leader)
    if [ "${old_leader}" != "${new_leader}" ]; then
        echo "leader switched to: ${new_leader}"
    fi
    sleep 1
done