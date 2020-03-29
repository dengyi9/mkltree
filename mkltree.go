package mkltree

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"reflect"
	"runtime"
)

// Merkle Tree. It can only be build by using New* method, for example: mkltree.NewMklTree()
// You can choose whether to store blocks in this merkle tree, if yes, it
// will cause more memory. You can only build a tree once.
// If you want a new tree, try to build a new one.
type mklTree struct {
	hasher hash.Hash

	hashes [][][]byte
	blocks [][]byte // optional: store original blocks
}

// NewMklTree return an merkle tree with default hash method sha256
func NewMklTree(blocks [][]byte, storeBlocks bool) *mklTree {
	return NewMklTreeCustomHash(blocks, storeBlocks, sha256.New())
}

// NewMklTreeCustomHash return an empty merkle tree with inputted hash method
func NewMklTreeCustomHash(blocks [][]byte, storeBlocks bool, h hash.Hash) *mklTree {
	m := &mklTree{
		h,
		[][][]byte{},
		[][]byte{},
	}

	// build hash tree // TODO: move to Add method
	leafLevelHashes := [][]byte{}
	for _, b := range blocks {
		if storeBlocks {
			m.blocks = append(m.blocks, b)
		}
		bhash := m.hasher.Sum(b)
		leafLevelHashes = append(leafLevelHashes, bhash)
	}
	m.hashes = append(m.hashes, leafLevelHashes)

	// TODO: move to Build method
	// hash into tree structure, leaf to root, bottom-up
	for i := 0; i < len(m.hashes); i++ {
		hashes := m.hashes[i]
		if len(hashes) > 1 {
			nextLevelHashes := [][]byte{}
			for i := 0; i < len(hashes); i = i + 2 {
				leftI, rightI := i, i+1
				left := hashes[leftI]
				right := []byte{}
				if rightI < len(hashes) {
					right = hashes[i+1]
				}

				m.hasher.Write(left)
				newHash := m.hasher.Sum(right)
				m.hasher.Reset()
				nextLevelHashes = append(nextLevelHashes, newHash)
			}
			m.hashes = append(m.hashes, nextLevelHashes)
		}
	}

	return m
}

//TODO: implement Add and Build methods, to allow incrementally reading big file's block to build a hash tree.

func (m *mklTree) Root() []byte {
	// last level element, which should only contain one, if not empty.
	if (len(m.hashes)) > 0 {
		return m.hashes[len(m.hashes)-1][0]
	}
	return []byte{}
}

// print Hash Tree structure. Hash bytes are encoded in BASE64, separated by character ""
func (m *mklTree) String() string {
	return m.stringProc(base64Encode)
}

func (m *mklTree) StringBytes() string {
	return m.stringProc(rawBytesEncode)
}

// print Hash Tree structure. Hash bytes are encoded in BASE64, separated by character ""
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
