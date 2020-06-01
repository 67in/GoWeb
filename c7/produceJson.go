package main

import (
	"encoding/json"
	"fmt"
)

type Server1 struct {
	ServerName string `json:"server_name"`
	ServerIP string `json:"server_ip"`
}

type Serverslice1 struct {
	Servers1 []Server1 `json:"servers_1"`
}

func main()  {
	var s Serverslice1
	s.Servers1 = append(s.Servers1, Server1{
		ServerName: "Shanghai_VPN",
		ServerIP:   "127.0.0.2",
	})
	s.Servers1 = append(s.Servers1,Server1{
		ServerName: "Beijing_VPN",
		ServerIP:   "127.0.0.1",
	})
	b,err := json.Marshal(s)
	if err != nil{
		fmt.Println("json err:",err)
	}
	fmt.Println(string(b))
}