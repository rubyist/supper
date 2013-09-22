package main

import (
	"fmt"
	"github.com/rubyist/go-dnsimple"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	myIP, err := getMyIP()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("My IP: %s\n", myIP)

	// Right now it just uses the domain token
	apiToken := "YOURAPITOKEN"
	email := "foo@example.com"
	domainName := "example.com"
	subdomainName := "www"
	client := dnsimple.NewClient(apiToken, email)

	records, err := client.Records(domainName, subdomainName, "")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if len(records) == 0 {
		log.Fatal("Record not found")
		os.Exit(1)
	}
	record := records[0]

	fmt.Printf("Domain IP: %s\n", record.Content)

	if myIP != record.Content {
		fmt.Println("Updating domain")
		record.UpdateIP(client, myIP)
	} else {
		fmt.Println("Domain is up to date")
	}
}

func getMyIP() (string, error) {
	resp, err := http.Get("http://icanhazip.com")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(body)), nil
}
