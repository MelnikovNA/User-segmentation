package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/app"
)

const configPath = "configs/config.yml"

func main() {
	err := app.Run(configPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("service work!!!\n")
	os.Exit(0)
}
