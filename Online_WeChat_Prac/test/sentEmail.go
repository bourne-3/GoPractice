package main

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func helper() {
	e := email.NewEmail()
	e.From = "595962708@qq.com"
	e.To = []string{"zebinli@gdufs.edu.cn"}
	e.Subject = "Test how to Sent Email"
	e.Text = []byte("OK Successfully")
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "595962708@qq.com", "vcefbpmjardzbeeb", "smtp.126.com"))
	//err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "595962708@qq.com", "vcefbpmjardzbeeb", "smtp.126.com"))
	if err != nil {
		log.Fatal(err)
	}
}