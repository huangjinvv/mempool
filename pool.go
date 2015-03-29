package main

import (
	"unsafe"
)

type Block struct {
	next unsafe.Pointer
	page *Page
}

type Page struct {
	size int32
}
