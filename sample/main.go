package main

import (
	"fmt"
	"os"
)

func main() {
	k1 := os.Getenv("K1")
	k2 := os.Getenv("K2")
	k3 := os.Getenv("K3")

	if k1 == "" || k2 == "" {
		fmt.Println("Missing required environment variables.")
		os.Exit(1)
	}

	fmt.Printf("API Token: %s\n", k1)
	fmt.Printf("Project ID: %s\n", k2)
	fmt.Printf("Base Language: %s\n", k3)

	output := "some output value"
	fmt.Println(output)
}
