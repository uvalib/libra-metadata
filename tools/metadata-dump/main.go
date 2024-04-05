package main

import (
	"flag"
	"fmt"
	"github.com/uvalib/libra-metadata"
	"os"
	"strings"
)

// main entry point
func main() {

	var infile string
	var schema string

	flag.StringVar(&infile, "infile", "", "Input file name")
	flag.StringVar(&schema, "schema", "", "Metadata schema (libraopen|libraetd)")
	flag.Parse()

	// validate required parameters
	if len(infile) == 0 ||
		len(schema) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	buf, err := os.ReadFile(infile)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	switch schema {
	case "libraopen":
		dumpAsOpen(buf)
	case "libraetd":
		dumpAsEtd(buf)
	default:
		fmt.Printf("ERROR: unsupported schema (%s)\n", schema)
	}

	fmt.Printf("INFO: terminating normally\n")
}

func dumpAsOpen(buf []byte) {

	m, err := librametadata.OAWorkFromBytes(buf)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	nonEmpty(" version:       %s\n", m.Version)
	nonEmpty(" resource type: %s\n", m.ResourceType)
	nonEmpty(" title:         %s\n", m.Title)
	for ix, a := range m.Authors {
		if ix == 0 {
			fmt.Printf(" authors:\n")
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
	nonEmpty(" languages:     %s\n", strings.Join(m.Languages, ", "))
	nonEmpty(" keywords:      %s\n", strings.Join(m.Keywords, ", "))
	for ix, c := range m.Contributors {
		if ix == 0 {
			fmt.Printf(" contributors:\n")
		}
		nonEmpty("   cid:         %s\n", c.ComputeID)
		nonEmpty("   first name:  %s\n", c.FirstName)
		nonEmpty("   last name:   %s\n", c.LastName)
		nonEmpty("   department:  %s\n", c.Department)
		nonEmpty("   institution: %s\n", c.Institution)
	}
	nonEmpty(" publisher:     %s\n", m.Publisher)
	nonEmpty(" citation:      %s\n", m.Citation)
	nonEmpty(" pub date:      %s\n", m.PublicationDate)
	nonEmpty(" urls:          %s\n", strings.Join(m.RelatedURLs, ", "))
	nonEmpty(" sponsors:      %s\n", strings.Join(m.Sponsors, ", "))
	nonEmpty(" notes:         %s\n", m.Notes)
}

func dumpAsEtd(buf []byte) {

	m, err := librametadata.ETDWorkFromBytes(buf)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	nonEmpty(" version:       %s\n", m.Version)
	nonEmpty(" degree:        %s\n", m.Degree)
	nonEmpty(" title:         %s\n", m.Title)

	fmt.Printf(" author:\n")
	nonEmpty("   cid:         %s\n", m.Author.ComputeID)
	nonEmpty("   first name:  %s\n", m.Author.FirstName)
	nonEmpty("   last name:   %s\n", m.Author.LastName)
	nonEmpty("   program:     %s\n", m.Author.Program)
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
