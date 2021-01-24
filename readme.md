# Fast Http To Curl


[![Coverage Status](https://coveralls.io/repos/github/xplorfin/fasthttp2curl/badge.svg?branch=master)](https://coveralls.io/github/xplorfin/fasthttp2curl?branch=master)
[![Renovate enabled](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://app.renovatebot.com/dashboard#github/xplorfin/fasthttp2curl)
[![Build status](https://github.com/xplorfin/fasthttp2curl/workflows/test/badge.svg)](https://github.com/xplorfin/fasthttp2curl/actions?query=workflow%3Atest)
[![Build status](https://github.com/xplorfin/fasthttp2curl/workflows/goreleaser/badge.svg)](https://github.com/xplorfin/fasthttp2curl/actions?query=workflow%3Agoreleaser)
[![](https://godoc.org/github.com/xplorfin/fasthttp2curl?status.svg)](https://godoc.org/github.com/xplorfin/fasthttp2curl)
[![Go Report Card](https://goreportcard.com/badge/github.com/xplorfin/fasthttp2curl)](https://goreportcard.com/report/github.com/xplorfin/fasthttp2curl)

Supports [net/http](https://godoc.org/net/http) and [fasthttp](https://github.com/valyala/fasthttp)

Allows you to transform [net/http](https://godoc.org/net/http) and [fasthttp](https://github.com/valyala/fasthttp) to curl for debugging. This package is based on moul's [http to curl](https://github.com/moul/http2curl) command except it supports fast http and http. This package is tested for parity against moul/http

# License

The original project is licensed by [Manfred Touron](https://manfred.life) under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details. Licenses are included in the directory 
