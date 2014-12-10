package main

import (
	"github.com/dropbox/godropbox/cinterop"
	"io"
	"log"
)

func processBatch(data []byte) []byte {
	log.Print("Server got:", string(data))
	return data
}

func nop([]byte, []byte) {

}
func makeEcho() (func(data []byte) []byte, func([]byte, []byte)) {
	return processBatch, nop
}

const numBatchWorkUnits = 4
const workSize = 2

func main() {
	processData := func(r io.ReadCloser, w io.Writer) {
		cinterop.ProcessBatchedData(r, w, makeEcho, numBatchWorkUnits*workSize, workSize)
	}
	cinterop.StartServer(processData)
}
