package main

import (
	useraction "dy/cmd/user_action/kitex_gen/useraction/useractionservice"
	"log"
)

func main() {
	svr := useraction.NewServer(new(UserActionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
