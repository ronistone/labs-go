package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/ronistone/labs-go/controller"
	"github.com/ronistone/labs-go/repository"
	"github.com/ronistone/labs-go/service"
	"github.com/ronistone/labs-go/trie"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(30)))

	err = personController.Register(e, context.Background())
	if err != nil {
		e.Logger.Fatal("Fail to initialize trie")
		return
	}

	go func() {
		if err := e.Start("0.0.0.0:8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("Shutting down the server!")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	} else {
		e.Logger.Fatal("Graceful Shutting down the server!")
	}
}
