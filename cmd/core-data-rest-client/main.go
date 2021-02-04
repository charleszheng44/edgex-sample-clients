package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"

	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/go-resty/resty/v2"
)

type CoreServiceClientConfig struct {
	Host string
	Port int
}

type CoreServiceClient struct {
	*resty.Client
	Host string
	Port int
}

const (
	ValueDescriptorPath = "/api/v1/valuedescriptor"
	ValueDescriptorJson = `
{
	"name": "humidity",
	"description": "Ambient humidity in percent",
	"min": "0",
	"max": "100",
	"type": "Int64",
	"uomLabel": "humidity",
	"defaultValue": "0",
	"formatting": "%s",
	"labels": [
		"environment",
		"humidity"
	]
}`
)

func addValueDescriptor(vd models.ValueDescriptor, cli CoreServiceClient) error {
	vdJson, err := json.Marshal(&vd)
	if err != nil {
		return err
	}
	postPath := fmt.Sprintf("http://%s:%d%s",
		cli.Host, cli.Port, ValueDescriptorPath)
	resp, err := cli.R().
		SetBody(vdJson).Post(postPath)
	if err != nil {
		return err
	}
	fmt.Println(string(resp.Body()))
	return nil
}

func getValueDescriptorByName(name string,
	cli CoreServiceClient) (models.ValueDescriptor, error) {
	var vd models.ValueDescriptor
	getURL := fmt.Sprintf("http://%s:%d%s/name/%s",
		cli.Host, cli.Port, ValueDescriptorPath, name)
	resp, err := cli.R().Get(getURL)
	if err != nil {
		return vd, err
	}
	err = json.Unmarshal(resp.Body(), &vd)
	return vd, err
}

func listValueDescriptors(cli CoreServiceClient) (
	[]models.ValueDescriptor, error) {
	vds := []models.ValueDescriptor{}
	listURL := fmt.Sprintf("http://%s:%d%s",
		cli.Host, cli.Port, ValueDescriptorPath)
	resp, err := cli.R().Get(listURL)
	if err != nil {
		return vds, err
	}
	err = json.Unmarshal(resp.Body(), &vds)
	return vds, err
}

func deleteValueDescriptorByName(name string, cli CoreServiceClient) error {
	delURL := fmt.Sprintf("http://%s:%d%s/name/%s",
		cli.Host, cli.Port, ValueDescriptorPath, name)
	resp, err := cli.R().Delete(delURL)
	if err != nil {
		return err
	}
	if string(resp.Body()) != "true" {
		return errors.New(string(resp.Body()))
	}
	return nil
}

func main() {
	host := flag.String("host", "", "TODO...")
	flag.Parse()

	cscli := CoreServiceClient{
		Client: resty.New(),
		Host:   *host,
		Port:   48080,
	}
	var vd models.ValueDescriptor
	if err := json.Unmarshal([]byte(ValueDescriptorJson), &vd); err != nil {
		panic(err)
	}

	if err := addValueDescriptor(vd, cscli); err != nil {
		panic(err)
	}
	fmt.Println("successfully add the valuedescriptor")

	vd, err := getValueDescriptorByName("humidity", cscli)
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully get the valuedescriptor")

	vdJson, err := json.Marshal(&vd)
	if err != nil {
		panic(err)
	}
	fmt.Println("The ValueDescriptor Json:")
	fmt.Println(string(vdJson))

	if err := deleteValueDescriptorByName("humidity", cscli); err != nil {
		panic(err)
	}
	fmt.Println("successfully delete the valuedescriptor(humidity)")
}
