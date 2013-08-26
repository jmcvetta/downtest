downtest
========

Run tests for the known downstream consumers of a [Go](http://golang.org)
package.


## Status

Early development.  Works okay.  Needs tests.


## Usage

```bash
$ go get github.com/jmcvetta/downtest/downtest
$ downtest github.com/username/package # Import path of a Go package
```


## Documentation

[Automatically generated API
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
  -u=true: Update on go get
  -v=false: Verbose
```


## License

This is Free Software, released under the terms of the [GPL
v3](http://www.gnu.org/copyleft/gpl.html).
