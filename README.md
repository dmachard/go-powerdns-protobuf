# go-powerdns-protobuf

PowerDNS encoder and decoder protobuf implementation in Golang 

## Installation

```go
go get -u github.com/dmachard/go-powerdns-protobuf
```

## Usage example

Example to use the PowerDNS protobuf decoder

```go
```

Example to use the owerDNS protobuf encoder

```go
```

## Testing

```bash
$ go test -v
=== RUN   TestMarshal
--- PASS: TestMarshal (0.00s)
=== RUN   TestUnmarshal
--- PASS: TestUnmarshal (0.00s)
PASS
ok      github.com/dmachard/go-powerdns-protobuf        0.002s
```

## Benchmark

```bash
```

## Development

Sync the dnsmessage.proto

    git submodule add https://github.com/PowerDNS/dnsmessage
    wget https://raw.githubusercontent.com/PowerDNS/dnsmessage/master/dnsmessage.proto

Download the latest release of protoc and protoc-gen-go

    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip
    unzip protoc-3.19.4-linux-x86_64.zip

Export GOBIN

    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export GOBIN=$GOPATH/bin

Copy the dnsmessage.proto

    cp dnsmessage/dnsmessage.proto .

Past the following line in the dnsmessage.proto

    option go_package = "github.com/dmachard/go-powerdns-protobuf;powerdns_protobuf";

Generate the golang package

    git submodule update --remote dnsmessage
    cd dnsmessage/
    ../bin/protoc --proto_path=. --go_out=../ --go_opt=paths=source_relative --plugin protoc-gen-go=${GOBIN}/protoc-gen-go dnsmessage.proto 