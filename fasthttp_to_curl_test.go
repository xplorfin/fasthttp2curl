package fasthttp2curl

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/go-http-utils/headers"
	"github.com/valyala/fasthttp"
)

var encodingHeader = []byte(headers.AcceptEncoding)
var authorizationHeader = []byte(headers.Authorization)

// Test converting a fast http request to a curl command
func TestFastHttpRequest(t *testing.T) {
	// mock it up
	testUserAgent := []byte(gofakeit.UserAgent())
	testHeader := []byte("gzip, brotli")

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI("https://entropy.rocks/")
	req.Header.SetUserAgentBytes(testUserAgent)
	req.Header.SetBytesKV(encodingHeader, testHeader)
	req.Header.SetBytesKV(authorizationHeader, []byte("authorizationHeader"))

	res, err := GetCurlCommandFastHttp(req)
	if err != nil {
		t.Error(err)
	}
	expectedString := fmt.Sprintf("curl -X 'GET' -H 'Accept-Encoding: gzip, brotli' -H 'Authorization: authorizationHeader' -H 'User-Agent: %s' 'https://entropy.rocks/'", testUserAgent)
	if res.String() != expectedString {
		t.Error("expected well formed curl")
	}
}
