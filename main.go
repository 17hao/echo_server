package main

import (
	api "github.com/17hao/echo_server/kitex_gen/api/echo"
	"log"
)

func main() {
	svr := api.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
