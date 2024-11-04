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
	f, err := os.OpenFile(os.Getenv("GITHUB_OUTPUT"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening GITHUB_OUTPUT: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	fmt.Fprintf(f, "output_variable_name=%s\n", output)
}
