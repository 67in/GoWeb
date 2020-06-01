package main

import (
	"fmt"
	"strconv"
)

//func main()  {
//	str := make([]byte, 0,100)
//	str = strconv.AppendInt(str,4567,10)
//	str = strconv.AppendBool(str,false)
//	str = strconv.AppendQuote(str,"abcdefg")
//	str = strconv.AppendQuoteRune(str, '单')
//	fmt.Println(string(str))
//}

//Format 系列函数把其他类型的转换为字符串
//func main() {
//	a := strconv.FormatBool(false)
//	b := strconv.FormatFloat(123.23, 'g', 12, 64)
//	c := strconv.FormatInt(1234, 10)
//	d := strconv.FormatUint(12345, 10)
//	e := strconv.Itoa(1023)
//	fmt.Println(a, b, c, d, e)
//}
//Parse 系列函数把字符串转换为其他类型
func checkError(e error){
	if e != nil{
		fmt.Println(e)
	}
}
func main() {
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e)
}