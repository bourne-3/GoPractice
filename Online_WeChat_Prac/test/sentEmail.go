package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"log"
	"net/smtp"
)

func helper() {
	// 读取自己的数据
	data,errRead := ioutil.ReadFile("F:\\Go_dir\\CodePrac\\GoPractice\\Online_WeChat_Prac\\test\\script.txt")
	if errRead!=nil{
		fmt.Println("File reading error", errRead)
	}


	e := email.NewEmail()
	e.From = "595962708@qq.com"
	e.To = []string{"zebinli@gdufs.edu.cn"}
	//e.To = []string{"junteng.ma@qq.com"}
	e.Subject = "Test how to Sent Email"
	//e.Text = []byte("OK Successfully")
	e.Text = data

	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "595962708@qq.com", "vcefbpmjardzbeeb", "smtp.qq.com"))
	//err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "595962708@qq.com", "vcefbpmjardzbeeb", "smtp.126.com"))
	if err != nil {
		log.Fatal(err)
	}
}

func readFileTest() {
	data,err := ioutil.ReadFile("F:\\Go_dir\\CodePrac\\GoPractice\\Online_WeChat_Prac\\test\\script.txt")
	if err!=nil{
		fmt.Println("File reading error", err)
	}
	fmt.Println(string(data))

	//f, err := os.Open("F:\\Go_dir\\CodePrac\\GoPractice\\Online_WeChat_Prac\\test\\script.txt")
	//if err!=nil{
	//	fmt.Println("File reading error", err)
	//}
	//fmt.Println(f)
}