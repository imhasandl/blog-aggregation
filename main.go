package main

import (
	"fmt"
	"log"
	"os"

	"module/internal/config"
)

type state struct {
	cfg *config.Config
}

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error handling config: %v", err) 
	}
	fmt.Printf("Read Config: %+v\n", cfg)

	err = cfg.SetUser("Hasan")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}