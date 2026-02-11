package main

import (
	"fmt"

	"github.com/adwip/e-wallet-tlab/internal/interfaces/containers"
)

func main() {
	err := containers.Migrations()
	if err != nil {
		panic(fmt.Sprintf("Error running migrations: %v", err))
	}

	fmt.Println("Migrations completed successfully")
}
