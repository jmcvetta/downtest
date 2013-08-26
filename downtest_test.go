// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package downtest

import (
	"encoding/json"
	"github.com/bmizerany/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testImporterPath = "github.com/username/repo"

var testApiResponse = apiResponse{
	Results: []importer{
		importer{
			Path:     testImporterPath,
			Synopsis: "foo",
		},
		importer{
			Path:     downtestPackage,
			Synopsis: "bar",
		},
	},
}

func apiMock(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(&testApiResponse)
}

func TestLookupImporters(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(apiMock))
	defer srv.Close()
	apiUrl = "http://" + srv.Listener.Addr().String() + "/"
	p, err := NewPackage("unimportant_package_name")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(p.Importers))
	assert.Equal(t, testImporterPath, p.Importers[0])
}

func TestLookupImportersBadUrl(t *testing.T) {
	apiUrl = "foo bar baz"
	_, err := NewPackage("unimportant_package_name")
	assert.NotEqual(t, nil, err)
}
