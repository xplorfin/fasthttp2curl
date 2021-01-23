package fasthttp2curl

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

// convert a fast http request object to a curl command
func GetCurlCommandFastHttp(req *fasthttp.Request) (command *CurlCommand, err error) {
	command = getBaseCommand()

	command.Append("-X", bashEscape(string(req.Header.Method())))

	if req.Body() != nil {
		command.addBodyBytes(req.Body())
	}

	command.addHeaders(FastHttpHeaderToHttp(req))

	command.Append(bashEscape(req.URI().String()))

	return command, nil
}

// transofrm a fast http request to a list of http headers
func FastHttpHeaderToHttp(req *fasthttp.Request) (headers http.Header) {
	headers = http.Header{}
	req.Header.VisitAll(func(rawKey, rawValue []byte) {
		key := string(rawKey)
		value := string(rawValue)
		if headers[key] == nil {
			headers[key] = []string{}
		}
		headers[key] = append(headers[key], value)
	})
	return headers
}
