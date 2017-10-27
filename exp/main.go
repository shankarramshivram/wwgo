package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"wwgo/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	//user := models.User{
	//	Name:  "Shankar Shivram",
	//	Email: "shankarshivram@gmail.com",
	//}
	//if err := us.Create(&user); err != nil {
	//	panic(err)
	//}

	defer us.Close()

	foundUser, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)
	us.DestructiveReset()

}
