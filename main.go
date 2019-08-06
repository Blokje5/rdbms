package main

import (
	"fmt"

	"github.com/blokje5/rdbms/pkg/constants"
	"github.com/blokje5/rdbms/pkg/storage"
)

func main() {
	storage.OpenDB("test.db")
	fmt.Println(constants.MaxNumPages)
}
