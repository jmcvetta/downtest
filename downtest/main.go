// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// downtest is a command line tool to run the tests on all downstream consumers
// of a package, as known to GoDoc.org.
package main

import (
	"flag"
	"fmt"
	"github.com/jmcvetta/downtest"
	"log"
	"os"
	"sort"
)

func main() {
	log.SetFlags(log.Lshortfile)
	verbose := flag.Bool("v", false, "Verbose output")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Must specify an import path as an argument.")
		os.Exit(-1)
	}
	p, err := downtest.NewPackage(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	p.Verbose = *verbose
	err = p.RunTests()
	if err != nil {
		log.Fatal(err)
	}
	total := len(p.Passed)
	fail := 0
	for _, pass := range p.Passed {
		if !pass {
			fail++
		}
	}
	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println()
	fmt.Printf("Passed %d / %d downstream tests:\n", total-fail, total)
	fmt.Println()
	packages := p.Importers
	sort.Strings(packages)
	for _, pkg := range packages {
		var status string
		if p.Passed[pkg] {
			status = "pass"
		} else {
			status = "FAIL"
		}
		fmt.Printf("%s  %s\n", status, pkg)
	}
	if fail != 0 {
		os.Exit(1)
	}
}
