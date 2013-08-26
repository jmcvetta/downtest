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
	"sort"
)

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("Must specify an import path as an argument")
	}
	p, err := downtest.NewPackage(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
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
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println()
	fmt.Printf("Passed %d / %d downstream tests.\n", total-fail, total)
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
		fmt.Printf("%s \t %s\n", pkg, status)
	}
}
