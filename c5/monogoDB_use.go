package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name string
	Phone string
}

func main()  {
	session, err := mgo.Dial("sever1.example.com, server2.example.com")
	if err != nil{
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c:= session.DB("test").C("people")
	err = c.Insert(&Person{"Ale","+555381169"},&Person{"Cla","+5553893928392"})
	if err!=nil{
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name":"Ale"}).One(&result)
	if err != nil{
		panic(err)
	}
	fmt.Println("Phone:", result.Phone)
}
