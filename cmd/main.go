package main

import (
	"fmt"

	sqlDb "github.com/masisiliani/bitBybit/db"
	"github.com/masisiliani/bitBybit/pkg/user"
)

func main() {
	fmt.Println("Hello bit.By.bit")

	err := sqlDb.Connect_SQL()

	if err != nil {
		fmt.Println(err)
	}

	userRepo := user.NewSqlServerRepository(sqlDb.DB)
	userService := user.NewService(userRepo)

	user, err := userService.Find(1)
	fmt.Println(user.Login)

}
