package main

import (
	"github.com/PatateDu609/matcha/auth"
	_ "github.com/PatateDu609/matcha/init"
	"github.com/PatateDu609/matcha/utils/log"
)

func main() {
	// router := routes.Setup()
	//
	// if err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Conf.API.Host, config.Conf.API.Port), router); err != nil {
	// 	log.Logger.Fatalf("error: couldn't launch server: %+v", err)
	// }

	log.Logger.Infof("trying to confirm email")
	if err := auth.ConfirmEmail("b76a6de9-e9d4-4ca5-b473-6b03c1afb166", "Teyber", "boucettaghali@gmail.com"); err != nil {
		log.Logger.Fatalf("couldn't confirm email: %s", err)
	}
}
