package main

import (
	"fmt"
	"net/http"

	"github.com/PatateDu609/matcha/config"
	_ "github.com/PatateDu609/matcha/init"
	"github.com/PatateDu609/matcha/routes"
	"github.com/PatateDu609/matcha/utils/log"
)

func main() {
	router := routes.Setup()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Conf.APIPort), router); err != nil {
		log.Logger.Fatalf("error: couldn't launch server: %+v", err)
	}
}
