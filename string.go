package mkltree

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"reflect"
	"runtime"
)

// print Hash Tree structure. Hash bytes are encoded in BASE64, separated by character ""
func (m *mklTree) String() string {
	return m.stringProc(base64Encode)
}

func (m *mklTree) StringBytes() string {
	return m.stringProc(rawBytesEncode)
}

// print Hash Tree structure. Hash bytes are encoded in encodeFunc, separated by character ""
func (m *mklTree) stringProc(encodeFunc func([]byte) []byte) string {
	bf := bytes.NewBuffer(nil)
	encodeFuncName := runtime.FuncForPC(reflect.ValueOf(encodeFunc).Pointer()).Name()

	if len(m.blocks) > 0 {
		bf.Write([]byte("Original blocks, encoded by " + encodeFuncName + ", separated by comma: \n"))
		m.writeBlocks(encodeFunc, bf, m.blocks)
	}

	bf.Write([]byte("Merkle tree bottom-up. Leaves first, root last. Each level in one line. Encoded by " + encodeFuncName + ", separated by comma: \n"))
	for _, levelHashes := range m.hashes {
		m.writeBlocks(encodeFunc, bf, levelHashes)
	}

	return bf.String()
}

func (m *mklTree) writeBlocks(encodeFunc func([]byte) []byte, bf *bytes.Buffer, blocks [][]byte) {
	for _, h := range blocks {
		b64hash := encodeFunc(h)

		bf.Write([]byte{' ', ' '}) // space indent
		bf.Write(b64hash)
		bf.Write([]byte{','})
	}
	bf.Write([]byte{'\n'})
}

func base64Encode(block []byte) []byte {
	b64Block := make([]byte, base64.StdEncoding.EncodedLen(len(block)))
	base64.StdEncoding.Encode(b64Block, block)
	return b64Block
}

func rawBytesEncode(block []byte) []byte {
	return []byte(fmt.Sprintf("%v", block))
}
