package constants

// Pages

const (
	// PageSize of OS is a multiple of the BlockSize of the Os
	PageSize = 1 << 12
	// MaxDbSize is the maximum size of the DB Heapfile
	MaxDbSize = 1 << 30
	// MaxNumPages is the maximum number of pages in the DB Heapfile
	MaxNumPages = MaxDbSize / PageSize
)

// Byte Sizes

const (
	Uint8ByteSize = 1
	Uint16ByteSize = 2
	Uint32ByteSize = 4
)