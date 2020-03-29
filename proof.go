package mkltree

import (
	"hash"
	"log"
)

// Proof the original leaf Block is in the blockIndex of the merkle tree whose root hash
// is known by user already.
func Proof(leafBlock []byte,
	leafBlockIndex int,
	merkleRoot []byte,
	merkleHashPath [][]byte,
	hasher hash.Hash) bool {
	hasher.Reset()

	return false
}

// return relevant merkle tree hash path, from leaf to root.
// TODO
func (m *mklTree) Path(blockIndex int) [][]byte {
	path := [][]byte{}

	for i, hashes := range m.hashes {
		if i == len(m.hashes)-1 { // last row is ROOT, no need to get
			break
		}

		broIndex := m.pathBrother(blockIndex)
		log.Printf("blockIndex:%v, get merkle-tree index [%v, %v]\n", blockIndex, i, broIndex)

		if broIndex < len(hashes) {
			path = append(path, hashes[broIndex])
		} else {
			path = append(path, []byte{})
		}
		blockIndex = blockIndex / 2
	}


	return path
}

func (m *mklTree) pathBrother(blockIndex int) int {
	if blockIndex % 2 != 0 {
		return blockIndex - 1
	}  else {
		return blockIndex + 1
	}
}