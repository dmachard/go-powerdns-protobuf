package powerdns_protobuf

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

var dm_ref = []byte{8, 1, 26, 10, 112, 111, 119, 101, 114, 100, 110, 115, 112, 98, 32, 1, 40, 5, 160, 1, 144, 78, 168, 1, 53}

func TestMarshal(t *testing.T) {
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
		t.Errorf("error on encode powerdns protobuf message %s", err)
	}
	fmt.Println(wiremessage)

	if len(wiremessage) != len(dm_ref) {
		t.Errorf("size of the encoded message is different from reference")
	}
}

func TestUnmarshal(t *testing.T) {
	// init
	dm := &PBDNSMessage{}

	// Unmarshal parses a wire-format message and places the decoded results in dt.
	err := proto.Unmarshal(dm_ref, dm)
	if err != nil {
		t.Errorf("error on decode powerdns protobuf message %s", err)
	}

	if string(dm.GetServerIdentity()) != "powerdnspb" {
		t.Errorf("mismatch identity %s", string(dm.GetServerIdentity()))
	}
}
