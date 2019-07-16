package snoo

import (
	"io/ioutil"

	"github.com/jarcoal/httpmock"
)

func mockResponseFromFile(url string, filepath string) {
	httpmock.Activate()
	response, _ := ioutil.ReadFile(filepath)
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, string(response)))
}
