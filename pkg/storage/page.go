package storage

// import (
// 	"io"

// 	"github.com/blokje5/rdbms/pkg/constants"
// 	"github.com/blokje5/rdbms/pkg/util"
// )

// PageID is a Integer representing the PageId
type PageID int

// // PageType represents the type of Page
// type PageType uint8

// const (
// 	DirectoryHeaderPageType PageType = iota
// 	DefaultPageType
// )

// // PageHeader represents the header of the page, containing some metadata
// // Pages should be self containing in this RDBMS
// type PageHeader interface {
// 	WriteBytes(w io.Writer)
// }

// // DefaultPageHeader represents the header of the default page
// type DefaultPageHeader struct {
// 	pageSize uint16
// }

// // CreateDefaultPageHeader creates a DefaultPageHeader for regular pages
// func CreateDefaultPageHeader(page Page) *DefaultPageHeader {
// 	return &DefaultPageHeader{
// 		pageSize: page.GetSize(),
// 	}
// }

// // WriteBytes writes the page header to the io.Writer
// func (ph *DefaultPageHeader) WriteBytes(w io.Writer) error {
// 	if err := binary.Write(w, binary.LittleEndian, DefaultPageType); err != nil {
// 		return err
// 	}

// 	if err := binary.Write(w, binary.LittleEndian, ph.pageSize); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // DirectoryHeaderPageHeader implements PageHeader and creates the header for the DirectoryHeaderPage
// type DirectoryHeaderPageHeader struct {
// 	pageSize uint16
// 	// Theoretically nextPage could be any of the constants.MaxNumPages pages.
// 	// However, the DirectoryHeaderPages are initialised when creating the DB and do not grow dynamically
// 	nextPage uint8
// }

// // CreateDirectoryHeaderPageHeader creates the Header for the DirectoryHeaderPage. It assumes the pageID for the next directory page is sub 256
// func CreateDirectoryHeaderPageHeader(page Page) *DirectoryHeaderPageHeader {
// 	return &DirectoryHeaderPageHeader{
// 		pageSize: page.GetSize(),
// 		nextPage: uint8(page.GetPageID() + 1),
// 	}
// }

// // WriteBytes writes the page header to the io.Writer
// func (ph *DirectoryHeaderPageHeader) WriteBytes(w io.Writer) error {
// 	if err := binary.Write(w, binary.LittleEndian, DirectoryHeaderPageType); err != nil {
// 		return err
// 	}

// 	if err := binary.Write(w, binary.LittleEndian, ph.pageSize); err != nil {
// 		return err
// 	}

// 	if err := binary.Write(w, binary.LittleEndian, ph.nextPage); err != nil {
// 		return err
// 	}

// 	return nil
// }

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
