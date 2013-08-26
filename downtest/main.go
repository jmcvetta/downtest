// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package main

import (
	"flag"
	"github.com/jmcvetta/downtest"
	"log"
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
	log.Println(p)
	log.Println(p.RunTests())
	/*	is, err := p.Importers()
		if err != nil {
			log.Fatal(err)
		}
		for _, path := range is {
			fmt.Println(path)
		}*/
}
