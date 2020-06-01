package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Println("key:", k)
		fmt.Println("val：", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "Hello LiuQin")
}

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method)
	if r.Method == "GET"{
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	}else{
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		//fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		//fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		//template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
	//v := url.Values{}
	//v.Set("name","Ava")
	//v.Add("friend", "Jess")
	//v.Add("friend", "Sarah")
	//v.Add("friend","Zoe" )
	//fmt.Println(v.Get("name"))
	//fmt.Println(v.Get("friend"))
	//fmt.Println(v["friend"])
}

func main()  {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}
