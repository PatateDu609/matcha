package main

import (
	_ "github.com/PatateDu609/matcha/init"
	"github.com/PatateDu609/matcha/routes"
)

func main() {
	_ = routes.Setup()
}
