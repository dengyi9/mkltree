package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/dengyi9/mkltree"
)

func main() {
	block_0 := "coin_0: A->B"
	block_1 := "coin_1: B->C"
	block_2 := "coin_2: B->D"
	blocks := [][]byte{
		[]byte(block_0),
		[]byte(block_1),
		[]byte(block_2),
	}

	mtree := mkltree.NewMklTree(blocks, true)

	fmt.Printf("%v\n", mtree)        // show tree
	fmt.Println(mtree.StringBytes()) 		// show tree in bytes


	// after tranctions happened,
	// receiver try to proof/verify if sender really send him a coin.

	// receiver C knows beforehand. Trustable info
	root := mtree.Root()   // root hash is public
	hasher := sha256.New() // hash method is public known. It is a contrast in a specific system like bitcoin

	// receiver C gets. Could come from untrustable source, e.g. sender B tells him.
	b := []byte(block_1)
	bIndex := 1
	path := mtree.Path(bIndex)

	// receiver C uses Proof function to verify if block b is in this merkle tree
	bExisted := mkltree.Proof(b, bIndex, root, path, hasher)
	fmt.Printf("Block b '%s' is existed in %v leaf of merkle tree: %v\n", b, bIndex, bExisted)

	// receiver C proof a fake block, not existed in merkle tree. Only one more space added.
	bFake := []byte(block_1 + " ")
	bNotExisted := mkltree.Proof(bFake, bIndex, root, path, hasher)
	fmt.Printf("Block bFake '%s' is existed in %v leaf of merkle tree: %v\n", bFake, bIndex, bNotExisted)
}
