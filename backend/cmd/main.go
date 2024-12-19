package main

import (
	"eCommerce/internal/app"
	"eCommerce/internal/setup/constructor"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Fullstack e-commerce project")
	getDependencies, err := app.GetDependencies()
	if err != nil {
		log.Fatal(err)
		return
	}
	constructor.Build(*getDependencies)
	runServer := fmt.Sprintf("%s:%s",
		getDependencies.Config.HttpConfig.HttpHost,
		getDependencies.Config.HttpConfig.HttpPort)

	newApp := app.NewApp(getDependencies)
	if err = newApp.Listen(runServer); err != nil {
		log.Fatal(err)
		return
	}

}
