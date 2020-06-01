package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {

	ServerIP string
	ServerName string
}

type Serverslice struct {
	Servers []Server
}

func main()  {
	var s Serverslice
	str := `{"servers":[{"ServerName":"Shanghai","serverIP":"127.0.0.1"},{"serverName":"Beijing","serverIP":"128.9.0.1"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}