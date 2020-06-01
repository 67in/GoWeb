package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Server  []server `xml:"server"`
	Description string `xml:",innerxml"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main()  {
	f, err := os.Open("a.xml")
	if err != nil{
		fmt.Printf("error %v",err)
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil{
		fmt.Printf("error %v",err)
		return
	}
	v := Servers{}
	err = xml.Unmarshal(data,&v)
	if err != nil{
		fmt.Printf("error %v",err)
		return
	}
	fmt.Println(v)
}