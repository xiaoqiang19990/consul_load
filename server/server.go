package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"net/http"
)

const RECV_BUF_LEN = 1024

func consulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "check")
}

func RegisterServer(Id, Name, Address string, Port int) {

	var tags []string
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = Id
	registration.Name = Name
	registration.Port = Port
	registration.Tags = tags
	registration.Address = Address
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务
	}

	//获取consul的api
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	//注册服务到consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}

}

func main() {
	Id := "002"
	Name := "002name"
	Port := 9002
	Address := "127.0.0.2"
	RegisterServer(Id, Name, Address, Port)
	http.HandleFunc("/check", consulCheck)
	http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)

}
