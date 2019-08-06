package util

import (
	"bytes"
	"encoding/binary"
	"io"
)

// WriteBytes ensures consistent endian usage
func WriteBytes(w io.Writer, data interface{}) error {
	return binary.Write(w, binary.BigEndian, data)
}

// ReadBytes ensures consistent endian usage
func ReadBytes(r io.Reader, data interface{}) error {
	return binary.Read(r, binary.BigEndian, data)
}

// ReadBytesFromSlice simplifys reading bytes from a slice
func ReadBytesFromSlice(b []byte, data interface{}) error {
	r := bytes.NewReader(b)
	return ReadBytes(r, data)
}

// Uint8 returns an uint8 from a byte
func Uint8(b []byte) uint8 {
	return b[0]
}

// Uint16 returns an uint16 from a byte
func Uint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

// Uint32 returns an uint32 from a byte
func Uint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}