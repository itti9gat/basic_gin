package main

import (
	"log"

	"iiujapp.tech/basic-gin/conf"
	"iiujapp.tech/basic-gin/mysql"
	"iiujapp.tech/basic-gin/repo"
	"iiujapp.tech/basic-gin/server"
	"iiujapp.tech/basic-gin/service"
)

func main() {

	db, err := mysql.Start()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Connecting to database... connnected")
	defer db.Close()

	s := service.Service{
		DBData:        db,
		QueryDataFunc: repo.QueryData,
		WriteDataFunc: repo.WriteData,
	}

	log.Println("Start Server", conf.ServerPort)

	serv := server.NewServer(s)
	serv.Start()
}
