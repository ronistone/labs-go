package main

import (
	"fmt"
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/ronistone/labs-go/controller"
	"github.com/ronistone/labs-go/repository"
	"github.com/ronistone/labs-go/service"
	"github.com/ronistone/labs-go/trie"
)

func cliTest() {
	data := trie.CreateTrie()
	var input string
	for {
		fmt.Printf("Adding Word: ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			break
		}
		data.AddWord(input)
	}

	for {
		fmt.Printf("\nConsulting words: ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			break
		}
		resultQuery := data.QueryWord(input)
		fmt.Println(input, "isFound: ", resultQuery)

		fmt.Println("\nListing By Prefix")
		results := data.ListWordsByPrefix(input)
		for _, word := range results {
			fmt.Println(word)
		}
	}

	fmt.Println("Printing all words:")
	for _, word := range data.PrintTrie() {
		fmt.Println(word)
	}
}

func main() {
	db, err := dbr.Open("postgres", "host=localhost port=5432 user=lab password='lab' dbname=lab sslmode=disable timezone=UTC", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)

	personRepository := repository.CreatePersonRepository(db)
	personService := service.CreatePersonService(personRepository)
	personController := controller.CreatePersonController(personService)

	e := echo.New()

	personController.Register(e)

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
