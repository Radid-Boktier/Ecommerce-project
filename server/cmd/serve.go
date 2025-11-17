package cmd

import (
	"ecommerce-server/config"
	"ecommerce-server/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
	
}