package main

import (
	"github.com/4afS/cli-clean-architecture/infrastructure"
	"fmt"
	"log"
	"github.com/4afS/cli-clean-architecture/interface/controller"
	"flag"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("invalid arguments")
	}

	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	registerName := registerCmd.String("name", "", "register username")
	registerEmail := registerCmd.String("email", "", "register email")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteName := deleteCmd.String("name", "", "delete user")

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchName := searchCmd.String("name", "", "search user by username")

	handler := infrastructure.NewDbHandler(".db")
	repo := infrastructure.NewUserRepository(handler)
	controller := controller.NewUserController(repo)

	switch os.Args[1] {
	case "register":
		registerCmd.Parse(os.Args[2:])
		err := controller.Register(*registerName, *registerEmail)
		if err != nil {
			log.Fatal(err)
		}

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		err := controller.Delete(*deleteName)
		if err != nil {
			log.Fatal(err)
		}

	case "search":
		searchCmd.Parse(os.Args[2:])
		user, err := controller.Search(*searchName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("username: %v, email: %v\n", user.Name, user.Email)
	}
}
