package main

import "encoding/json"

func init() {
	print("ssssssssss")
	// orm.Debug = true
}

func main() {
	//as := "{\"应用\":\"lll\",\"应用环境\":[{\"正式VPC环境\":{\"删除主机列表\":[\"vpc-test-发布流程-baidi02(172.20.62.137)\"],\"新增主机列表\":[\"vpc-prod-canal-mate01(172.28.91.228)\"]}},{\"模拟VPC环境\":{\"删除主机列表\":[],\"新增主机列表\":[\"vpc-test-jenkins-ci02(172.20.62.140)\",\"vpc-prod-rocketmq-nameServer02(172.28.91.221)\"]}}]}"

	x := &struct {
		AA string
		BB string
	}{AA: "ddd", BB: "dskjfdklsj"}
	y, _ := json.Marshal(&x)
	print(string(y))
}
