package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// get command line arguments from command line
	args := os.Args[1:]

	// check if there is at least one argument
	if len(args) < 1 {
		fmt.Printf("Domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")
		fmt.Println("Please provide an email address")
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			checkDomain(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal("Erorr: could not read from input: %v\n", err)
		}
	} else {
		fmt.Printf("Domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")
		for _, arg := range args {
			checkDomain(arg)
		}
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
		}
		hasSPF = true
		spfRecord = record
		break
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
