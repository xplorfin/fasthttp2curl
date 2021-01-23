package fasthttp2curl

import (
	"io"
	"strings"
)

// CurlCommand contains exec.Command compatible slice + helpers
type CurlCommand []string

// Append appends a string to the CurlCommand
func (command *CurlCommand) Append(newSlice ...string) {
	*command = append(*command, newSlice...)
}

// String returns a ready to copy/paste command
func (command *CurlCommand) String() string {
	return strings.Join(*command, " ")
}

// nopCloser is used to create a new io.ReadCloser for req.Body
type nopCloser struct {
	io.Reader
}

func bashEscape(str string) string {
	return `'` + strings.ReplaceAll(str, `'`, `'\''`) + `'`
}

func (nopCloser) Close() error { return nil }
