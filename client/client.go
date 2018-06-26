package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"time"
)

const RECV_BUF_LEN = 1024

func main() {

	client, err := consulapi.NewClient(consulapi.DefaultConfig())

	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	for {

		time.Sleep(time.Second * 3)
		var services map[string]*consulapi.AgentService
		var err error

		services, err = client.Agent().Services() //获取注册到consul的服务
		if nil != err {
			log.Println("in consual list Services:", err)
			continue
		}
		for k, v := range services {
			fmt.Println(k, " : ", v.Address)
		}
	}
}
