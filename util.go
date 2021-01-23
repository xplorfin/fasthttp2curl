package fasthttp2curl

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

func getBaseCommand() *CurlCommand {
	command := CurlCommand{}

	command.Append("curl")
	return &command
}

// add body to curl
func (command *CurlCommand) addBody(inBody io.ReadCloser) error {
	if inBody != nil {
		body, err := ioutil.ReadAll(inBody)
		if err != nil {
			return err
		}
		command.addBodyBytes(body)
	}
	return nil
}

// add body to curl from bytes
func (command *CurlCommand) addBodyBytes(body []byte) {
	if len(string(body)) > 0 {
		bodyEscaped := bashEscape(string(body))
		command.Append("-d", bodyEscaped)
	}
}

// add body to curl from bytes
func (command *CurlCommand) addHeaders(headers http.Header) {
	var keys []string
	for k := range headers {
		keys = append(keys, k)

	}
	sort.Strings(keys)

	for _, k := range keys {
		command.Append("-H", bashEscape(fmt.Sprintf("%s: %s", k, strings.Join(headers[k], " "))))
	}
}
