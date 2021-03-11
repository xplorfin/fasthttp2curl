package fasthttp2curl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/moul/http2curl"
)

// taken from https://github.com/moul/http2curl/blob/master/http2curl_test.go
func TestGetCurlCommand(t *testing.T) {
	form := url.Values{}
	form.Add("age", "10")
	form.Add("name", "Hudson")
	body := form.Encode()

	req, _ := http.NewRequest(http.MethodPost, "http://foo.com/cats", ioutil.NopCloser(bytes.NewBufferString(body)))
	req.Header.Set("API_KEY", "123")

	libCommand, _ := http2curl.GetCurlCommand(req)
	command, _ := GetCurlCommand(req)
	if libCommand.String() != command.String() {
		t.Errorf("expected library command: %s and command: %s to match", libCommand, command)
	}

	// Output:
	// curl -X 'POST' -d 'age=10&name=Hudson' -H 'Api_key: 123' 'http://foo.com/cats'
}

func TestGetCurlCommand_json(t *testing.T) {
	req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", bytes.NewBufferString(`{"hello":"world","answer":42}`))
	req.Header.Set("Content-Type", "application/json")

	libCommand, _ := http2curl.GetCurlCommand(req)
	command, _ := GetCurlCommand(req)
	if libCommand.String() != command.String() {
		t.Errorf("expected library command: %s and command: %s to match", libCommand, command)
	}

	// Output:
	// curl -X 'PUT' -d '{"hello":"world","answer":42}' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func TestGetCurlCommand_slice(t *testing.T) {
	t.Skip("TODO") // this isn't working for some reason
	// See https://github.com/moul/http2curl/issues/12
	req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", bytes.NewBufferString(`{"hello":"world","answer":42}`))
	req.Header.Set("Content-Type", "application/json")

	libCommand, _ := http2curl.GetCurlCommand(req)
	command, _ := GetCurlCommand(req)
	if libCommand.String() != command.String() {
		//t.Errorf("expected library command: %s and command: %s to match", strings.Join(*libCommand, " \\\n  "), strings.Join(*command, " \\\n  "))
	}

	// Output:
	// curl \
	//   -X \
	//   'PUT' \
	//   -d \
	//   '{"hello":"world","answer":42}' \
	//   -H \
	//   'Content-Type: application/json' \
	//   'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func TestGetCurlCommand_noBody(t *testing.T) {
	req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", nil)
	req.Header.Set("Content-Type", "application/json")

	libCommand, _ := http2curl.GetCurlCommand(req)
	command, _ := GetCurlCommand(req)
	if libCommand.String() != command.String() {
		t.Errorf("expected library command: %s and command: %s to match", libCommand, command)
	}

	// Output:
	// curl -X 'PUT' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func TestGetCurlCommand_emptyStringBody(t *testing.T) {
	t.Skip("broken right now because lib is out of date (master vs 2.2.0)")
	req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", bytes.NewBufferString(""))
	req.Header.Set("Content-Type", "application/json")

	libCommand, _ := http2curl.GetCurlCommand(req)
	command, _ := GetCurlCommand(req)
	if libCommand.String() != command.String() {
		t.Errorf("expected library command: %s and command: %s to match", libCommand, command)
	}

	// Output:
	// curl -X 'PUT' -d '' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func TestGetCurlCommand_newlineInBody(t *testing.T) {
	req, _ := http.NewRequest("POST", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", bytes.NewBufferString("hello\nworld"))
	req.Header.Set("Content-Type", "application/json")

	libCommand, _ := http2curl.GetCurlCommand(req)
	command, _ := GetCurlCommand(req)
	if libCommand.String() != command.String() {
		t.Errorf("expected library command: %s and command: %s to match", libCommand, command)
	}

	// Output:
	// curl -X 'POST' -d 'hello
	// world' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func TestGetCurlCommand_specialCharsInBody(t *testing.T) {
	req, _ := http.NewRequest("POST", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", bytes.NewBufferString(`Hello $123 o'neill -"-`))
	req.Header.Set("Content-Type", "application/json")

	libCommand, _ := http2curl.GetCurlCommand(req)
	command, _ := GetCurlCommand(req)
	if libCommand.String() != command.String() {
		t.Errorf("expected library command: %s and command: %s to match", libCommand, command)
	}

	// Output:
	// curl -X 'POST' -d 'Hello $123 o'\''neill -"-' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func TestGetCurlCommand_other(t *testing.T) {
	uri := "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu"
	payload := new(bytes.Buffer)
	payload.Write([]byte(`{"hello":"world","answer":42}`))
	req, err := http.NewRequest("PUT", uri, payload)
	if err != nil {

	}
	req.Header.Set("X-Auth-Token", "private-token")
	req.Header.Set("Content-Type", "application/json")

	libCommand, _ := http2curl.GetCurlCommand(req)
	command, _ := GetCurlCommand(req)
	if libCommand.String() != command.String() {
		t.Errorf("expected library command: %s and command: %s to match", libCommand, command)
	}
	// Output: curl -X 'PUT' -d '{"hello":"world","answer":42}' -H 'Content-Type: application/json' -H 'X-Auth-Token: private-token' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}
