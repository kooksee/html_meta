package tests

import (
	"github.com/kooksee/html_meta/internal/utils"
	"io/ioutil"
	"net/http"
	"testing"
)

func url(p string) string {
	return "http://localhost:8080" + p
}

func TestHealth(t *testing.T) {
	resp, err := http.Get(url("/health"))
	utils.MustNotError(err)

	dt, err := ioutil.ReadAll(resp.Body)
	utils.MustNotError(err)

	utils.P(resp.Status, string(dt))
}

func TestGetPattern(t *testing.T) {
	resp, err := http.Get(url("/patterns/common"))
	utils.MustNotError(err)

	dt, err := ioutil.ReadAll(resp.Body)
	utils.MustNotError(err)

	utils.P(resp.Status, string(dt))
}
