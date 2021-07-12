package main

import (
	"flag"
	"fmt"
	"generate-certificat/cert"
	"generate-certificat/pdf"
	"os"
)

func main() {
	file := flag.String("file", "", "CSV File input")
	flag.Parse()

	if len(*file) <= 0{
		fmt.Printf("Invalid file. got=%v\n", *file)
		os.Exit(1)
	}

	var saver cert.Saver
	var err error
	saver, err = pdf.New("output")
	if err != nil {
		fmt.Printf("Error during PDF generation: %v", err)
		os.Exit(1)
	}

	certs, err := cert.ParseCSV(*file)
	if err != nil {
		fmt.Printf("Could not parse CSV file. %v\n", err)
		os.Exit(1)
	}

	for _, c := range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save Cert. %v\n", err)
			os.Exit(1)
		}
	}
}
