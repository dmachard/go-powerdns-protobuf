package powerdns_protobuf

import (
	"bufio"
	"bytes"
	"net"
	"testing"
	"time"
)

func TestProtoStream(t *testing.T) {
	client, server := net.Pipe()

	// init framestream receiver
	fs_client := NewProtobufStream(bufio.NewReader(client), client, 5*time.Second)

	go func() {
		var buf bytes.Buffer
		buf.Write([]byte{0x00, 0x2, 0x0f, 0xff})
		w := bufio.NewWriter(server)
		if _, err := buf.WriteTo(w); err == nil {
			w.Flush()
		}
	}()

	// receive payload
	_, err := fs_client.RecvPayload(true)
	if err != nil {
		t.Errorf("error to receive payload: %s", err)
	}
}
