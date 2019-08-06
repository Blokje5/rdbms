package storage

import (
	"io"

	"github.com/blokje5/rdbms/pkg/constants"
	"github.com/blokje5/rdbms/pkg/util"
)

// PageType represents the type of Page
type PageType uint8

const (
	DirectoryHeaderPageType PageType = iota
	DefaultPageType
)

const (
	headerSize = constants.Uint8ByteSize + constants.Uint16ByteSize + constants.Uint8ByteSize
)

type header struct {
	pageType PageType
	pageSize uint16
	nextPage uint8
}

func parseHeader(b []byte) header {
	raw := b[:headerSize]
	// Read pageType
	size := constants.Uint8ByteSize
	ptRaw := util.Uint8(raw[:size])
	pt := PageType(ptRaw)
	// Read pageSize
	ps := util.Uint16(raw[size:(size + constants.Uint16ByteSize)])
	size += constants.Uint16ByteSize
	// Read nextPage
	np := util.Uint8(raw[size:(size + constants.Uint8ByteSize)])

	return header{
		pt,
		ps,
		np,
	}
}

func (h header) writeBytes(w io.Writer) error {
	if err := util.WriteBytes(w, h.pageType); err != nil {
		return err
	}

	if err := util.WriteBytes(w, h.pageSize); err != nil {
		return err
	}

	if err := util.WriteBytes(w, h.nextPage); err != nil {
		return err
	}

	return nil
}
