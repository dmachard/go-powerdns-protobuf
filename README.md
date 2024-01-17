# go-powerdns-protobuf

PowerDNS encoder and decoder protobuf implementation in Golang 

## Installation

```go
go get -u github.com/dmachard/go-powerdns-protobuf
```

## Usage example

Example to use the PowerDNS protobuf decoder

```go
var dm_ref = []byte{8, 1, 26, 10, 112, 111, 119, 101, 114, 100, 110, 115, 112, 98, 32, 1, 40, 5, 160, 1, 144, 78, 168, 1, 53}


dm := &PBDNSMessage{}

err := proto.Unmarshal(dm_ref, dm)
if err != nil {
    fmt.Println("error on decode powerdns protobuf message %s", err)
}
```

Example to use the owerDNS protobuf encoder

```go
dm := &PBDNSMessage{}

dm.Reset()

dm.ServerIdentity = []byte("powerdnspb")
dm.Type = PBDNSMessage_DNSQueryType.Enum()

dm.SocketProtocol = PBDNSMessage_DNSCryptUDP.Enum()
dm.SocketFamily = PBDNSMessage_INET.Enum()
dm.FromPort = proto.Uint32(10000)
dm.ToPort = proto.Uint32(53)

wiremessage, err := proto.Marshal(dm)
if err != nil {
    fmt.Println("error on encode powerdns protobuf message %s", err)
}
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
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/dmachard/go-powerdns-protobuf
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkUnmarshal-4     3656814               293.5 ns/op
BenchmarkMarshal-4       3185277               344.2 ns/op
PASS
ok      github.com/dmachard/go-powerdns-protobuf        2.892s
```

## Development

Add the proto schema as git submodule

```bash
git submodule add https://github.com/PowerDNS/dnsmessage
```

Export GOBIN

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
```

Update golang version

```bash
go mod edit -go=1.21
go mod tidy
```

Download the latest release of protoc and protoc-gen-go

```bash
export PROTOC_VER=25.2
export GITHUB_URL=https://github.com/protocolbuffers
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
wget $GITHUB_URL/protobuf/releases/download/v$PROTOC_VER/protoc-$PROTOC_VER-linux-x86_64.zip
unzip protoc-$PROTOC_VER-linux-x86_64.zip
```

Edit and past the following line in the dnsmessage.proto

```bash
option go_package = "github.com/dmachard/go-powerdns-protobuf;powerdns_protobuf";
```

Generate the golang package

```bash
cd dnsmessage/
../bin/protoc --proto_path=. --go_out=../ --go_opt=paths=source_relative --plugin protoc-gen-go=${GOBIN}/protoc-gen-go dnsmessage.proto 
```
