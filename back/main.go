package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PatateDu609/matcha/config"
	_ "github.com/PatateDu609/matcha/init"
	"github.com/PatateDu609/matcha/routes/private"
	"github.com/PatateDu609/matcha/routes/public"
	"github.com/PatateDu609/matcha/utils/log"
)

func main() {
	publicRouter := public.Setup()
	privateRouter := private.Setup()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		privateAddr := fmt.Sprintf("%s:%d", config.Conf.API.Host, config.Conf.API.InternalPort)
		log.Logger.Infof("Starting private router with addr `%s`", privateAddr)
		if err := http.ListenAndServe(privateAddr, privateRouter); err != nil {
			log.Logger.Fatalf("error: couldn't launch server: %+v", err)
		}
		wg.Done()
	}()

	go func() {
		publicAddr := fmt.Sprintf("%s:%d", config.Conf.API.Host, config.Conf.API.Port)
		log.Logger.Infof("Starting public router with addr `%s`", publicAddr)
		if err := http.ListenAndServe(publicAddr, publicRouter); err != nil {
			log.Logger.Fatalf("error: couldn't launch server: %+v", err)
		}
		wg.Done()
	}()

	wg.Wait()
}
