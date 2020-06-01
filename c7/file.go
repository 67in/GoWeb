package main

import (
	"fmt"
	"os"
)

//func main()  {
//	//os.Mkdir("astaxie",0777)
//	//os.MkdirAll("astaxie/test1/test2",0777)
//	err := os.Remove("astaxie")
//	if err != nil{
//		fmt.Println(err)
//	}
//	os.RemoveAll("astaxie")
//}
//写文件
//func main()  {
//	userFile := "astaxie.txt"
//	fout, err := os.Create(userFile)
//	if err != nil{
//		fmt.Println(userFile, err)
//		return
//	}
//	defer fout.Close()
//	for i :=0; i<10; i++{
//		fout.WriteString("Just a test!\r\n")
//		fout.Write([]byte("Just a test1\r\n"))
//	}
//}
//读文件
func main()  {
	userFile := "astaxie.txt"
	fl, err := os.Open(userFile)
	if err != nil{
		fmt.Println(userFile, err)
	}
	defer fl.Close()
	buf := make([]byte,1024)
	for {
		n, _ := fl.Read(buf)
		if n == 0{
			break
		}
		os.Stdout.Write(buf[:n])
	}
	os.Remove(userFile)
}
