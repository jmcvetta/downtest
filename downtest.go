// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// Package downtest runs the tests on all downstream consumers of a package,
// as known to GoDoc.org.
package downtest

import (
	"fmt"
	"github.com/jmcvetta/restclient"
	"os"
	"os/exec"
	"sort"
)

var apiUrl = "http://api.godoc.org/importers/"

type apiResponse struct {
	Results []struct {
		Path     string
		Synopsis string
	}
}

type apiError struct {
	Error struct {
		Message string
	}
}

// A Package is a module of Go code identified by its import path.
type Package struct {
	ImportPath string
	Importers  []string
	Passed     map[string]bool
	Verbose    bool
}

// NewPackage prepares a package for downstream testing by looking up its
// importers.
func NewPackage(importPath string) (*Package, error) {
	p := Package{
		ImportPath: importPath,
	}
	err := p.LookupImporters()
	if err != nil {
		return &p, err
	}
	p.Passed = make(map[string]bool, len(p.Importers))
	return &p, nil
}

// LookupImporters gets the import paths of all downstream packages known by
// GoDoc.org to import this Package.
func (p *Package) LookupImporters() error {
	url := apiUrl + p.ImportPath
	var e apiError
	var res apiResponse
	var importers []string
	rr := restclient.RequestResponse{
		Url:            url,
		Method:         "GET",
		Result:         &res,
		Error:          &e,
		ExpectedStatus: 200,
	}
	_, err := restclient.Do(&rr)
	if err != nil {
		return err
	}
	for _, r := range res.Results {
		importers = append(importers, r.Path)
	}
	sort.Strings(importers)
	p.Importers = importers
	return nil
}

// RunTests runs "go test" on downstream packages.
func (p *Package) RunTests() error {
	for _, pkg := range p.Importers {
		p.Passed[pkg] = false
		c := exec.Command("go", "get", "-u", "-v", pkg)
		if p.Verbose {
			fmt.Fprintln(os.Stderr, "+++ Running tests for", pkg, "+++")
			fmt.Fprintln(os.Stderr)
			fmt.Fprintln(os.Stderr, "> go get -v", pkg)
			fmt.Fprintln(os.Stderr)
			c.Stderr = os.Stderr
			c.Stdout = os.Stderr
		}
		err := c.Run()
		if err != nil {
			continue
		}
		c = exec.Command("go", "test", "-v", pkg)
		if p.Verbose {
			fmt.Fprintln(os.Stderr)
			fmt.Fprintln(os.Stderr, "> go test -v", pkg)
			fmt.Fprintln(os.Stderr)
			c.Stderr = os.Stderr
			c.Stdout = os.Stderr
		}
		err = c.Run()
		if err != nil {
			continue
		}
		p.Passed[pkg] = true
	}
	return nil
}
