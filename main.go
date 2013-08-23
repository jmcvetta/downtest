// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package main

import (
	"flag"
	"fmt"
	"github.com/jmcvetta/restclient"
	"log"
)

const baseUrl = "http://api.godoc.org/importers/"

type Importers struct {
	Results []struct {
		Path     string
		Synopsis string
	}
}

type ApiError struct {
	Error struct {
		Message string
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("Must specify an import path as an argument")
	}
	importPath := flag.Args()[0]
	url := baseUrl + importPath
	var e ApiError
	var res Importers
	rr := restclient.RequestResponse{
		Url:    url,
		Method: "GET",
		Result: &res,
		Error:  &e,
	}
	_, err := restclient.Do(&rr)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range res.Results {
		fmt.Println(r.Path)
	}

}
