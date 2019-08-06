package storage

import (
	"errors"
	"io"

	"github.com/blokje5/rdbms/pkg/constants"
)

// PageID is a Integer representing the PageId
type PageID uint32

type Page struct {
	pageID     PageID
	pageHeader *header
	data       []byte
}

const (
	// WriteablePageSize represents the size after the page is filled with the header
	WriteablePageSize = constants.PageSize - headerSize
)

// InitPage initialises an empty page with the given ID
func InitPage(pageID PageID, pageType PageType) *Page {
	return &Page{
		pageID: pageID,
		pageHeader: &header{
			pageType: pageType,
			pageSize: headerSize,
			nextPage: 0,
		},
		data: make([]byte, WriteablePageSize),
	}
}

// LinkToNextPage links the page to another page
func (p *Page) LinkToNextPage(pageID PageID) {
	p.pageHeader.nextPage = pageID
}

// GetNextPage returns the page this page links to, or 0 otherwise
func (p *Page) GetNextPage() PageID {
	return p.pageHeader.nextPage
}

// WriteDataRaw writes the data from the given byte slice to the page data. Note that this 
// method uses copying so it will not be the most efficient 
func (p *Page) WriteDataRaw(data []byte) error {
	l := len(data)
	if l > WriteablePageSize {
		return errors.New("Writing more data then maximum writeable page size")
	}

	copy(p.data, data)
	p.pageHeader.pageSize += uint16(l)

	return nil
}

// ReadPage reads a page from an io.Reader
func ReadPage(r io.Reader) (*Page, error) {
	headerBuffer := make([]byte, headerSize)
	n, err := r.Read(headerBuffer)
	if err != nil {
		return nil, errors.New("Failure reading header from reader")
	}

	if n < headerSize {
		return nil, errors.New("Reader did not contain full header")
	}

	header := parseHeader(headerBuffer)

	data := make([]byte, WriteablePageSize)
	_, err = r.Read(headerBuffer)
	if err != nil {
		return nil, errors.New("Failure reading data from reader")
	}

	return &Page{
		pageID: 0, //TODO should include pageID in header or ensure it is always passed properly
		pageHeader: &header,
		data: data,
	}, nil
}

// WritePage Writes a page to the io.writer
func (p *Page) WritePage(w io.Writer) error {
	err := p.pageHeader.writeBytes(w)
	if err != nil {
		return errors.New("Failed to write header to writer")
	}

	_, err = w.Write(p.data)
	if err != nil {
		return errors.New("Failed to write data to writer")
	}

	return nil
}

// type Page interface {
// 	GetSize() uint16
// 	//WriteHeader(io.Writer)
// }

// const (
// 	// DirectoryHeaderPageHeaderSize returns the size in bytes of the DirectoryHeaderPageHeader
// 	DirectoryHeaderPageHeaderSize = constants.Uint8ByteSize + constants.Uint16ByteSize + constants.Uint8ByteSize
// 	// DirectoryHeaderPageSlots returns the number of available slots per DirectoryHeaderPage
// 	DirectoryHeaderPageSlots = (constants.PageSize - DirectoryHeaderPageHeaderSize)/constants.Uint16ByteSize
// )

// // DirectoryHeaderPage implements the header pages of the Heap File
// // It contains pointers to the next DirectoryHeaderPage. It also contains a directory of pages containing the page size
// type DirectoryHeaderPage struct {
// 	data []byte
// }

// // func CreateDirectoryHeaderPage() *DirectoryHeaderPage {

// // }

// // GetSize returns the size of the header page
// func (dhp *DirectoryHeaderPage) GetSize() uint16 {
// 	return uint16(len(dhp.data))
// }

// func (dhp *DirectoryHeaderPage) writeHeader(w io.Writer) error {
// 	ph := CreateDirectoryHeaderPageHeader(dhp)

// 	if err := ph.WriteBytes(w); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (dhp *DirectoryHeaderPage) readSlotDirectory() []uint16 {
// 	b := dhp.data[DirectoryHeaderPageHeaderSize-1:]
// 	slots := make([]uint16, DirectoryHeaderPageSlots)
// 	util.ReadBytesFromSlice(b, slots)
// 	return slots
// }

// // func (dhp *DirectoryHeaderPage) WritePage(w io.Writer) {

// // }
