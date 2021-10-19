package main

import (
	"github.com/gorilla/mux"
	"github.com/moroz-matros/TP_Android_Backend_Moviezam/application/delivery"
	"github.com/moroz-matros/TP_Android_Backend_Moviezam/application/repository"
	"github.com/moroz-matros/TP_Android_Backend_Moviezam/application/usecase"
	"github.com/moroz-matros/TP_Android_Backend_Moviezam/config"
	"log"
	"net/http"
)

func main() {
	m := mux.NewRouter()
	server := &http.Server{
		Handler:     m,
		Addr: 		"localhost"+config.Port,
	}

	/*
	pool, err := pgxpool.Connect(context.Background(), config.DBConnect)
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}


	 */
	repo := repository.CreateRepo(/*pool*/)
	uc := usecase.CreateUsecase(repo)
	_ = delivery.CreateHandler(m, uc)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}


}