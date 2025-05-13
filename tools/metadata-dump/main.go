package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	librametadata "github.com/uvalib/libra-metadata"
)

// main entry point
func main() {

	var infile string
	var format string

	flag.StringVar(&infile, "infile", "", "Input file name")
	flag.StringVar(&format, "format", "json", "Output format (json, text)")
	flag.Parse()

	// validate required parameters
	if len(infile) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	buf, err := os.ReadFile(infile)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	m, err := librametadata.ETDWorkFromBytes(buf)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	switch format {
	case "json":
		dumpAsJson(m)
	case "text":
		dumpAsText(m)
	default:
		fmt.Printf("ERROR: format %s not supported\n", format)
		os.Exit(1)
	}
	fmt.Printf("INFO: terminating normally\n")
}

func dumpAsJson(m *librametadata.ETDWork) {
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(b))
}

func dumpAsText(m *librametadata.ETDWork) {

	nonEmpty(" version:       %s\n", m.Version)
	nonEmpty(" degree:        %s\n", m.Program)
	nonEmpty(" degree:        %s\n", m.Degree)
	nonEmpty(" title:         %s\n", m.Title)

	fmt.Printf(" author:\n")
	nonEmpty("   cid:         %s\n", m.Author.ComputeID)
	nonEmpty("   first name:  %s\n", m.Author.FirstName)
	nonEmpty("   last name:   %s\n", m.Author.LastName)
	nonEmpty("   institution: %s\n", m.Author.Institution)

	for ix, a := range m.Advisors {
		if ix == 0 {
			fmt.Printf(" advisors:\n")
		}
		nonEmpty("   cid:         %s\n", a.ComputeID)
		nonEmpty("   first name:  %s\n", a.FirstName)
		nonEmpty("   last name:   %s\n", a.LastName)
		nonEmpty("   department:  %s\n", a.Department)
		nonEmpty("   institution: %s\n", a.Institution)
	}

	nonEmpty(" abstract:      %s\n", m.Abstract)
	nonEmpty(" license:       %s\n", m.License)
	nonEmpty(" license url:   %s\n", m.LicenseURL)
	nonEmpty(" keywords:      %s\n", strings.Join(m.Keywords, ", "))
	nonEmpty(" language:      %s\n", m.Language)
	nonEmpty(" urls:          %s\n", strings.Join(m.RelatedURLs, ", "))
	nonEmpty(" sponsors:      %s\n", strings.Join(m.Sponsors, ", "))
	nonEmpty(" notes:         %s\n", m.Notes)
}

func nonEmpty(format string, value string) {
	if len(strings.Trim(value, " ")) != 0 {
		fmt.Printf(format, value)
	}
}

//
// end of file
//
