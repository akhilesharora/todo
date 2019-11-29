# Todo App

Simple todo gRPC app to add, update and delete notes for daily use


## Usage

### Build

Run `make build`

It will create a binary file in `build` directory

#### Configuration options:

|     ENV     	|       Description      	| Required 	|  Default  	|
|-------------	|------------------------	|----------	|-----------	|
| HOST     	    | Application host          | true      | 127.0.0.1 	|
| PORT     	    | Application port          | true    	| 8000      	|

> As per RFC3339 
> **Date format = "2006-01-02"**

#### gRPC UI playground

You can use the `go` tool to install `grpcui`:
```shell
go get github.com/fullstorydev/grpcui
go install github.com/fullstorydev/grpcui/cmd/grpcui
```

This installs the command into the `bin` sub-folder of wherever your `$GOPATH`
environment variable points. If this directory is already in your `$PATH`, then
you should be good to go.

>  When you run `grpcui`, it will show you a URL to put into a browser in order to access
the web UI.

```
$ grpcui -plaintext 127.0.0.1:8000
gRPC Web UI available at http://127.0.0.1:60551/...

```
