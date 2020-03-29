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

	thisHash := hashProc(hasher, leafBlock)
	thisIndex := leafBlockIndex
	for i, broHash := range merkleHashPath {

		log.Printf("merkle-tree level:%v, thisIndex: %v, thisHash: %v, broHash: %v\n", i, thisIndex, thisHash, broHash)

		//TODO: thisHash is left or right? according to thisIndex
		if thisIndex % 2 == 0 { // this is Left, same mechanism as method (m *MklTree) blockBrotherIndex
			thisHash = hashProc(hasher, append(thisHash, broHash...))
		} else {
			thisHash = hashProc(hasher, append(broHash, thisHash...))
		}
		thisIndex = thisIndex / 2
	}

	return bytes.Equal(thisHash, merkleRoot)
}


// return relevant merkle tree hash path, from leaf to root.
func (m *MklTree) Path(leafBlockIndex int) [][]byte {
	path := [][]byte{}

	thisIndex := leafBlockIndex
	for i, hashes := range m.hashes {
		if i == len(m.hashes)-1 { // last row is ROOT, no need to get
			break
		}

		broIndex := m.blockBrotherIndex(thisIndex)
		log.Printf("thisBlockIndex:%v, get merkle-tree index [%v, %v]\n", thisIndex, i, broIndex)

		if broIndex < len(hashes) {
			path = append(path, hashes[broIndex])
		} else {
			path = append(path, []byte{})
		}
		thisIndex  = thisIndex / 2
	}


	return path
}

func (m *MklTree) blockBrotherIndex(blockIndex int) int {
	if blockIndex % 2 == 0 {
		return blockIndex + 1
	}  else {
		return blockIndex - 1
	}
}