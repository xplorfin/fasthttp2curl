package fasthttp2curl

import (
	"net/http"
)

// GetCurlCommand returns a CurlCommand corresponding to an http.Request
func GetCurlCommand(req *http.Request) (command *CurlCommand, err error) {
	command = getBaseCommand()

	command.Append("-X", bashEscape(req.Method))

	err = command.addBody(req.Body)
	if err != nil {
		return nil, err
	}

	command.addHeaders(req.Header)
	command.Append(bashEscape(req.URL.String()))

	return command, nil
}
