# WASM Component

## Summary

The purpose of WASM Component is to provide a simple, opinionated demo of a project layout using [protobufs (v3)](https://developers.google.com/protocol-buffers/docs/proto3),
frontend GRPC over websockets, [Go's WASM support](https://pkg.go.dev/syscall/js), and components through the standard library 
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
   to make Chrome accept the CA used to sign the dev certs.
   - On Debian/Ubuntu: `sudo apt-get install libnss3-tools`, otherwise see [How do I set up Client Authentication for Google Chrome on Linux?](https://knowledge.digicert.com/fr/fr/quovadis/end-user-certificates/how-do-i-set-up-client-authentication-for-google-chrome-on-linux.html)
   - This optional dependency will be removed once the *certman* command is moved to a more appropriate home.

## Setup

Setup assumes that your **dependencies** are available in your PATH. If you installed 
the specific dependencies via go install, for instance `go install github.com/go-task/task/v3/cmd/task@latest` then you 
need to make sure your GOBIN is on your path. You can confirm the location that the Go command currently uses 
with `go env GOBIN`. If you would like to use a different location, you can set this environmental variable using 
your system's tools. On Ubuntu, for instance, you might add the text `export GOBIN=/opt/gobin` to your `~/.bashrc` file.

 1. `git clone `

## Project Binaries

### server

An example server with embedded asset (single-file production build) or live reload mode (uses filesystem as 
editing occurs). This could use a third party tool like [air](https://github.com/cosmtrek/air), but the 
intent is to be able to live reload templates (only components that change and their subordinate tree) 
without restarting the server and reloading the frontend (which would lose state) unless the Go code 
changes, which would always cause a server restart as well as a client restart. Pieces of this example should 
be moved to a dedicated repo over time.

### client

An example companion application which gets compiled to wasm and added to the static files embedded into the 
server. binary.

Additional resources:

 - [Tutorial: wasm using Go](https://golangbot.com/webassembly-using-go/)

### certman

The intent of *certman* is to provide locally acceptable certificates for serving https in a way that tools
like a browser or curl accept as valid for dev domains. This should be moved out of this repo to a more 
appropriate home (something less specific to components).

Resources for accomplishing this task:

 - [making a CA in Go](https://shaneutt.com/blog/golang-ca-and-signed-cert-go/)
 - [getting Chrome to accept a CA](https://serverfault.com/questions/946756/ssl-certificate-in-system-store-not-trusted-by-chrome)

Once *certman* is removed from this repo, the concept will be expanded to managing `/etc/hosts`, a 
reverse proxy from 

I could use [libpam0g-dev to authenticate a machine user](https://forum.golangbridge.org/t/check-linux-user-and-password-authenticate/6220/4) 
via [msteinert/pam](https://github.com/msteinert/pam) to authenticate into a control panel served at 
local.rhyvu.com installed and from there a user could list domains to add. Each domain would be 
added to `/etc/hosts` (via a setuid utility), have a certificate generated from the CA, and then the 
user would be able to add base href (for instance `/` or `/api`) and the upstream url. In the future, the web
interface could also provide additional features such as stripping or transforming parts of the url 
before sending the data upstream, or there could be convenience utilities for serving static files.

## License and Purpose

This project is licensed permissively with a standard [MIT license](LICENSE.txt) which allows both commercial and 
personal use. However, the purpose of this project is for demonstration purposes, so depending upon this code 
for production usage is not recommended. If any of the ideas here are used in a broader project, links will be listed
below.
