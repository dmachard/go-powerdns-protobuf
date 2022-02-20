package powerdns_protobuf

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	dm := &PBDNSMessage{}

	dm.Reset()
	dm.ServerIdentity = []byte("powerdnspb")
}

func TestUnmarshal(t *testing.T) {

}
