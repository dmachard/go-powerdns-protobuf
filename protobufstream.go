package powerdns_protobuf

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"time"
)

const DATA_FRAME_LENGTH_MAX = 65535

var ErrPayloadTooLarge = errors.New("protobuf payload too large error")

type ProtoPayload struct {
	data []byte
}

func (pl ProtoPayload) Len() int {
	return len(pl.data)
}

func (pl ProtoPayload) Data() []byte {
	return pl.data
}

type ProtoStream struct {
	buf         []byte
	reader      *bufio.Reader
	conn        net.Conn
	readtimeout time.Duration
}

func NewProtobufStream(reader *bufio.Reader, conn net.Conn, readtimeout time.Duration) *ProtoStream {
	return &ProtoStream{
		buf:         make([]byte, DATA_FRAME_LENGTH_MAX),
		reader:      reader,
		conn:        conn,
		readtimeout: readtimeout,
	}
}

func (ps ProtoStream) RecvPayload(timeout bool) (*ProtoPayload, error) {
	// enable read timeaout
	if timeout && ps.readtimeout != 0 {
		ps.conn.SetReadDeadline(time.Now().Add(ps.readtimeout))
	}

	// read payload len (2 bytes)
	var n uint16
	if err := binary.Read(ps.reader, binary.BigEndian, &n); err != nil {
		return nil, err
	}

	// checking data to read according to the size of the buffer
	if n > uint16(len(ps.buf)) {
		return nil, ErrPayloadTooLarge
	}

	// read  binary data and push it in the buffer
	if _, err := io.ReadFull(ps.reader, ps.buf[0:n]); err != nil {
		return nil, err
	}

	payload := &ProtoPayload{
		data: make([]byte, n),
	}
	copy(payload.data, ps.buf[0:n])

	// disable read timeaout
	if timeout && ps.readtimeout != 0 {
		ps.conn.SetDeadline(time.Time{})
	}

	return payload, nil
}

func (ps ProtoStream) ProcessStream(ch chan []byte) (err error) {
	for {
		payload, err := ps.RecvPayload(false)
		if err != nil {
			break
		}
		ch <- payload.data
	}
	return err
}
