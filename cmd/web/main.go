package main

import "github.com/adwip/e-wallet-tlab/internal/interfaces/containers"

func main() {
	err := containers.SetupServiceContainer()
	if err != nil {
		panic(err)
	}
}
