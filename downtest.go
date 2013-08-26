// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// Package downtest runs the tests on all downstream consumers of a package,
// as known to GoDoc.org.
package downtest

import (
	// vcs "github.com/sourcegraph/go-vcs"
	"github.com/jmcvetta/restclient"
	deps "github.com/sourcegraph/go-deps"
	"io/ioutil"
	"log"
	"os"
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
	Tmpdir     string
	Importers  []string
	context    deps.Context
}

// NewPackage prepares a package for downstream testing by looking up its
// importers.
func NewPackage(importPath string) (*Package, error) {
	p := Package{
		ImportPath: importPath,
	}
	dirName, err := ioutil.TempDir("", "downtest")
	if err != nil {
		return nil, err
	}
	// defer os.RemoveAll(dirName)
	p.Tmpdir = dirName
	p.context = deps.Default
	p.context.GOPATH = p.Tmpdir
	err = p.LookupImporters()
	if err != nil {
		return &p, err
	}
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
	p.Importers = importers
	return nil
}

func (p *Package) RunTests() error {
	log.Println(p.Tmpdir)
	os.Setenv("GOPATH", p.Tmpdir)
	for _, pkg := range p.Importers {
		err := p.context.GoGet(pkg, deps.Verbose)
		if err != nil {
			return err
		}
	}
	return nil
}
