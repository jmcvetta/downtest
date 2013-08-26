downtest
========

Run tests for the known downstream consumers of a [Go](http://golang.org)
package.

Given the import path of a Go package, `downtest` queries GoDoc.org for the
list of all other packages known to import that package.  Each of these
downstream consumer packages is tested by running `go get` then `go test`.

When the tests are complete a summary of results is printed.  Optionally, the
summary can be output as JSON.  If all tests passed, `downtest` quits with exit
code 0; if there were any failures, it quits with exit code 1.


## Status

Early development.  Works okay.  Needs tests.


## Usage

```bash
$ go get github.com/jmcvetta/downtest/downtest
$ downtest github.com/jmcvetta/restclient # Import path of any Go package

Passed 6 / 8 downstream tests:

pass  github.com/apeacox/txtatus-cli
FAIL  github.com/jmcvetta/heroku
pass  github.com/jmcvetta/neo4j
pass  github.com/jmcvetta/neoism
pass  github.com/jmcvetta/srom/srom
FAIL  github.com/jmcvetta/stormpath
pass  github.com/mostafah/mandrill
pass  github.com/postmaster/postmaster-go
```

`downtest` knows nothing about the test requirements of downstream packages.
In the example above some tests are failing because they require an environment
variable to be set.  You may need to setup the environment, databases, etc
before running `downtest`.  Use flag `-v` to see the output of the tests as
they run.

Tests are run in the context of the current `$GOPATH`.  


## Documentation

Automatically generated [API
documentation](http://godoc.org/github.com/jmcvetta/downtest) for package
`downtest` can be found at
[GoDoc](http://godoc.org/github.com/jmcvetta/downtest) or [Go
Walker](http://gowalker.org/github.com/jmcvetta/downtest).

Options for the command line tool can be listed with the standard help flag:

```bash
$ downtest -h
Usage of downtest:
downtest [options] import_path
  (where import_path is the full import path of a Go package)
  -j=false: JSON output
  -u=true: Update packages with "go get -u"
  -v=false: Verbose
```


## License

This is Free Software, released under the terms of the [GPL
v3](http://www.gnu.org/copyleft/gpl.html).
