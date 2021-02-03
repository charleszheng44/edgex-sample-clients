package main

import (
	"flag"
	"fmt"
	"github.com/go-resty/resty/v2"
	"path/filepath"
)

type CoreServiceClientConfig struct {
	Host string
	Port int
}

func main() {

	host := flag.String("host", "localhost", "TODO...")
	port := flag.Int("port", 48080, "TODO...")
	path := flag.String("path", "", "TODO...")

	flag.Parse()

	restPath := fmt.Sprintf("http://%s:%d%s", host, port, path)

	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(restPath)

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}
