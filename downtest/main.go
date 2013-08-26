// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// downtest is a command line tool to run the tests on all downstream consumers
// of a package, as known to GoDoc.org.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jmcvetta/downtest"
	"log"
	"os"
	"time"
)

type pkgResult struct {
	Package string
	Passed  bool
}

type results struct {
	Package   string
	Timestamp time.Time
	Importers []pkgResult
}

func main() {
	log.SetFlags(log.Lshortfile)
	verbose := flag.Bool("v", false, "Verbose")
	jsonOutput := flag.Bool("j", false, "JSON output")
	update := flag.Bool("u", true, "Update on go get")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Must specify an import path as an argument.")
		os.Exit(-1)
	}
	p, err := downtest.NewPackage(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	if len(p.Importers) == 0 {
		fmt.Printf("Package %s is not imported by any known package.\n", p.ImportPath)
		os.Exit(0)
	}
	p.Verbose = *verbose
	p.Update = *update
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
	if *jsonOutput {
		rs := results{
			Package:   p.ImportPath,
			Timestamp: time.Now(),
		}
		for _, pkg := range p.Importers {
			pr := pkgResult{
				Package: pkg,
				Passed:  p.Passed[pkg],
			}
			rs.Importers = append(rs.Importers, pr)
		}
		b, err := json.MarshalIndent(&rs, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))

	} else {
		fmt.Println()
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println()
		fmt.Printf("Passed %d / %d downstream tests:\n", total-fail, total)
		fmt.Println()
		for _, pkg := range p.Importers {
			var status string
			if p.Passed[pkg] {
				status = "pass"
			} else {
				status = "FAIL"
			}
			fmt.Printf("%s  %s\n", status, pkg)
		}
	}
	if fail != 0 {
		os.Exit(1)
	}
}
