package main

import (
	"flag"
	corecli "github.com/edgexfoundry/core-data-go/clients"
	"log"
)

func main() {
	host := flag.String("host", "localhost", "TODO...")
	port := flag.Int("port", 48080, "TODO...")
	timeout := flag.Int("timeout", 10, "TODO...")
	dbName := flag.String("db-name", "", "TODO...")
	userName := flag.String("user", "", "TODO...")
	passwd := flag.String("passwd", "", "TODO...")

	dbConfig := corecli.DBConfiguration{
		Host:         *host,
		Port:         *port,
		Timeout:      *timeout,
		DatabaseName: *dbName,
		Username:     *userName,
		Password:     *passwd,
	}

	cli, err := corecli.NewDBClient(dbConfig)
	if err != nil {
		log.Fatalf("fail to generate client fot the core-data service: %v", err)
	}

	vds, err := cli.ValueDescriptors()
	if err != nil {
		log.Fatalf("fail to list valuedescriptors: %v", err)
	}

	log.Println("Got valuedescriptors:")
	for _, v := range vds {
		log.Println(v.Name)
	}
}
