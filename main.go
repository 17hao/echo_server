package main

import (
	"github.com/17hao/echo_server/kitex_gen/api/echoservice"
	"log"
)

func main() {
	svr := echoservice.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
