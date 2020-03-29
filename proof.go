package mkltree

import (
	"bytes"
	"hash"
	"log"
)

// Proof the original leaf Block is in the blockIndex of the merkle tree whose root hash
// is known by user already.
func Proof(leafBlock []byte, leafBlockIndex int,
	merkleRoot []byte, merkleHashPath [][]byte, hasher hash.Hash) bool {

	hasher.Reset()

	thisHash := hasher.Sum(leafBlock)
	thisIndex := leafBlockIndex
	for i, broHash := range merkleHashPath {

		log.Printf("merkle-tree index[%v, %v] thisHash: %v\n", i, thisIndex, thisHash)

		//TODO: thisHash is left or right? according to thisIndex
		if thisIndex % 2 == 0 { // this is Left

			thisHash = hasher.Sum(append(thisHash, broHash...))
		} else {
			thisHash = hasher.Sum(append(broHash, thisHash...))
		}
		thisIndex = thisIndex / 2
	}

	return bytes.Equal(thisHash, merkleRoot)
}

// return relevant merkle tree hash path, from leaf to root.
func (m *MklTree) Path(blockIndex int) [][]byte {
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

func (m *MklTree) pathBrother(blockIndex int) int {
	if blockIndex % 2 != 0 {
		return blockIndex - 1
	}  else {
		return blockIndex + 1
	}
}