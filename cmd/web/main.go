package main

import "github.com/adwip/aj-teknik-backend-admin/internal/interfaces/containers"

func main() {
	err := containers.SetupServiceContainer()
	if err != nil {
		panic(err)
	}
}
