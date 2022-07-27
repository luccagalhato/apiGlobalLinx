package main

import (
	"MANCHESTER/API-GLOBAL-LINX/config"
	controller "MANCHESTER/API-GLOBAL-LINX/controllers"

	"flag"
	"log"
	"os"
)

func initConfig() {
	var createFlag bool
	flag.BoolVar(&createFlag, "c", false, "create an yaml config file")
	flag.Parse()
	if createFlag {
		if err := config.NewYaml("config.yaml"); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}

func main() {
	initConfig()
	controller, err := controller.NewController("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(controller.ListenAndServe())
}
