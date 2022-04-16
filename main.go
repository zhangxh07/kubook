package main

import (
	"kubook/pkg/server"
	"log"
	"os"
)

func main()  {
	host := os.Getenv("CMDB_SERVER")
	username := os.Getenv("CMDB_USERNAME")
	password := os.Getenv("CMDB_PASSWORD")


	if username == ""{
		log.Fatal("用户名不能为空")
	}
	if password == ""{
		log.Fatal("密码不能为空")
	}
	//var u *auth.Client
	svc := server.New(server.Config{
		CmdbHost: host,
		CmdbUsername: username,
		CmdbPassword: password,
	})
	svc.Run()
}
