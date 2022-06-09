package main

import (
	userbase "dy/cmd/user_base/kitex_gen/userbase/userbaseservice"
	"log"
)

func main() {
	svr := userbase.NewServer(new(UserBaseServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
