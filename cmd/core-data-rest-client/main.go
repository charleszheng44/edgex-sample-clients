package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/go-resty/resty/v2"
)

type CoreServiceClientConfig struct {
	Host string
	Port int
}

func main() {

	host := flag.String("host", "localhost", "TODO...")
	port := flag.Int("port", 48080, "TODO...")

	flag.Parse()

	restPath := fmt.Sprintf("http://%s:%d/api/v1/valuedescriptor", *host, *port)

	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(restPath)
	if err != nil {
		panic(err)
	}
	vds := []models.ValueDescriptor{}
	if json.Unmarshal(resp.Body(), &vds); err != nil {
		panic(err)
	}
	fmt.Println(vds)
}
