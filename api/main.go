package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go-clean-arch/api/controller"
	"go-clean-arch/repository"
	"go-clean-arch/service/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	dbUser := "fuu"
	dbPass := "fuu_pass"
	dbName := "go_clean_arch"
	//dbHost := "localhost"
	dataSourceName := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	userRepo := repository.NewUserMySQL(db)
	userService := user.NewService(userRepo)

	//bookRepo := repository.NewBookMySQL(db)
	//bookService := book.NewService(bookRepo)
	//
	//loanUseCase := loan.NewService(userService, bookService)

	userController := controller.NewUserController(userService)

	// create router
	r := mux.NewRouter()

	userController.Route(r)

	srv := &http.Server{
		Addr:         "0.0.0.0:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
}
