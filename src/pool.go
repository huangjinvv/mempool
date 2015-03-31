package main

import (
	"unsafe"
)

type MemPool struct {
	size     int
	pageList []*Page
}

type Block struct {
	next unsafe.Pointer
	page *Page
}

type Page struct {
	size int32
}
