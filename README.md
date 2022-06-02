# WASM Component

## Summary

The purpose of WASM Component is to provide a simple, opinionated demo of a project layout using protobufs,
frontend GRPC over websockets, Go's WASM support, and components through the standard library 
[html/template](https://pkg.go.dev/html/template). Less novel concepts which are often part of a website are 
out of scope for this demonstration. This includes any sort of datastore, authentication, authorization, or 
advanced logging capabilities.

In order to provide style without adding noise to the html, [Pico.css](https://picocss.com/) is utilized. 
Communication via [GRPC](https://grpc.io/docs/languages/go/quickstart/) over wasm uses 
the [tarndt/wasmws](https://github.com/tarndt/wasmws) library, and the 
[version 2 of dominikh/go-js-dom](https://github.com/dominikh/go-js-dom) provides dom convenience functions.

## Dependencies

### Basic

 - [Go](https://golang.org): language and compiler
 - [Git](https://git-scm.com/): version control system

### Specific

See links for alternative installation methods or to verify the latest instructions and versions. 
Otherwise, run the commands listed.

 - [Task](https://taskfile.dev/): a friendly, simple task runner
   - `go install github.com/go-task/task/v3/cmd/task@latest`
 - [protobuf and Go plugins](https://grpc.io/docs/languages/go/quickstart/): protoc, a protobuf generator for use in GRPC
   - `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
   - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
 - (optional) [NSS Tools](https://knowledge.digicert.com/fr/fr/quovadis/end-user-certificates/how-do-i-set-up-client-authentication-for-google-chrome-on-linux.html): 
   to make Chrome accept the CA used to sign the dev certs
   - On Debian/Ubuntu: `sudo apt-get install libnss3-tools`, otherwise see [How do I set up Client Authentication for Google Chrome on Linux?](https://knowledge.digicert.com/fr/fr/quovadis/end-user-certificates/how-do-i-set-up-client-authentication-for-google-chrome-on-linux.html)

## Setup

Setup assumes that your dependencies are available in your PATH. If you installed 
the specific dependencies via go install, for instance `go install github.com/go-task/task/v3/cmd/task@latest` then you 
need to make sure your GOBIN is on your path. You can confirm the location that the Go command currently uses 
with `go env GOBIN`. If you would like to use a different location, you can set this environmental variable using 
your system's tools. On Ubuntu, for instance, you might add the text `export GOBIN=/opt/gobin` to your `~/.bashrc` file.

 1. 

## Project Binaries

### certman

https://shaneutt.com/blog/golang-ca-and-signed-cert-go/
https://serverfault.com/questions/946756/ssl-certificate-in-system-store-not-trusted-by-chrome

## License and Purpose

This project is licensed permissively with a standard [MIT license](LICENSE.txt) which allows both commercial and 
personal use. However, the purpose of this project is for demonstration purposes, so depending upon this code 
for production usage is not recommended. If any of the ideas here are used in a broader project, links will be listed
below.
