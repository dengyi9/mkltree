package mkltree

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"
)

const (
	block_0 = "coin_0: A->B"
	block_1 = "coin_1: B->C"
	block_2 = "coin_2: B->D"
)

func buildTree() *MklTree {
	// input data, compute hashes, and build merkle hash tree
	blocks := [][]byte{
		[]byte(block_0),
		[]byte(block_1),
		[]byte(block_2),
	}
	return NewMklTree(blocks, true)
}

func TestBuild(t *testing.T) {
	// input data, compute hashes, and build merkle hash tree
	mtree := buildTree()

	// get merkle root
	rootBytes := mtree.Root()
	if len(rootBytes) == 0 {
		t.Error()
	}
	rootBase64 := base64.StdEncoding.EncodeToString(rootBytes)
	fmt.Printf("rootBase64: %v\n\n", rootBase64)

	// show tree
	treeStr := mtree.String()
	if treeStr == "" {
		t.Error()
	}
	fmt.Println(treeStr)

	// show tree in bytes
	treeBytesStr := mtree.StringBytes()
	if treeBytesStr == "" {
		t.Error()
	}
	fmt.Println(treeBytesStr)

}

/**
blocks:
	block_0 = "coin_0: A->B"
	block_1 = "coin_1: B->C"
	block_2 = "coin_2: B->D"
blocks bytes (separated by comma):
	[99 111 105 110 95 48 58 32 65 45 62 66],[99 111 105 110 95 49 58 32 66 45 62 67],[99 111 105 110 95 50 58 32 66 45 62 68]

merkle tree (rawBytesEncode, separated by comma):
	[99 111 105 110 95 48 58 32 65 45 62 66 227 176 196 66 152 252 28 20 154 251 244 200 153 111 185 36 39 174 65 228 100 155 147 76 164 149 153 27 120 82 184 85],[99 111 105 110 95 49 58 32 66 45 62 67 227 176 196 66 152 252 28 20 154 251 244 200 153 111 185 36 39 174 65 228 100 155 147 76 164 149 153 27 120 82 184 85],[99 111 105 110 95 50 58 32 66 45 62 68 227 176 196 66 152 252 28 20 154 251 244 200 153 111 185 36 39 174 65 228 100 155 147 76 164 149 153 27 120 82 184 85],
	[99 111 105 110 95 49 58 32 66 45 62 67 227 176 196 66 152 252 28 20 154 251 244 200 153 111 185 36 39 174 65 228 100 155 147 76 164 149 153 27 120 82 184 85 90 160 6 236 67 153 221 68 29 92 10 68 103 160 79 159 137 46 99 119 79 38 75 196 28 93 112 164 83 10 148 98],[181 148 218 234 237 156 240 120 168 23 137 149 231 4 206 38 210 69 30 171 182 145 85 240 87 113 208 136 86 76 53 62],
	[181 148 218 234 237 156 240 120 168 23 137 149 231 4 206 38 210 69 30 171 182 145 85 240 87 113 208 136 86 76 53 62 193 13 124 102 96 91 121 212 165 10 6 255 201 144 49 153 12 2 45 251 118 145 156 29 127 6 50 159 167 131 57 217],

path (mt-X-Y is the X-row Y-column element in merkle tree ):
	path_0: block_0 [mt-0-0] -> expectedPath [[mt-0-1], [mt-1-1]]
	path_1: block_1 [mt-0-1] -> expectedPath [[mt-0-0], [mt-1-1]]
	path_2: block_2 [mt-0-2] -> expectedPath [[mt-0-3], [mt-1-0]]
*/
func TestMklTree_Path(t *testing.T) {
	mtree := buildTree()
	path := mtree.Path(0)
	expectedPath := [][]byte{
		[]byte{99, 111, 105, 110, 95, 49, 58, 32, 66, 45, 62, 67, 227, 176, 196, 66, 152, 252, 28, 20, 154, 251, 244, 200, 153, 111, 185, 36, 39, 174, 65, 228, 100, 155, 147, 76, 164, 149, 153, 27, 120, 82, 184, 85},
		[]byte{181, 148, 218, 234, 237, 156, 240, 120, 168, 23, 137, 149, 231, 4, 206, 38, 210, 69, 30, 171, 182, 145, 85, 240, 87, 113, 208, 136, 86, 76, 53, 62},
	}
	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("path is not correct. Expected:%v, Got:%v\n", expectedPath, path)
	}

	path1 := mtree.Path(1)
	expectedPath1 := [][]byte{
		[]byte{99, 111, 105, 110, 95, 48, 58, 32, 65, 45, 62, 66, 227, 176, 196, 66, 152, 252, 28, 20, 154, 251, 244, 200, 153, 111, 185, 36, 39, 174, 65, 228, 100, 155, 147, 76, 164, 149, 153, 27, 120, 82, 184, 85},
		[]byte{181, 148, 218, 234, 237, 156, 240, 120, 168, 23, 137, 149, 231, 4, 206, 38, 210, 69, 30, 171, 182, 145, 85, 240, 87, 113, 208, 136, 86, 76, 53, 62},
	}
	if !reflect.DeepEqual(path1, expectedPath1) {
		t.Errorf("path is not correct. Expected:%v, Got:%v\n", expectedPath1, path1)
	}

	path2 := mtree.Path(2)
	expectedPath2 := [][]byte{
		[]byte{},
		[]byte{99, 111, 105, 110, 95, 49, 58, 32, 66, 45, 62, 67, 227, 176, 196, 66, 152, 252, 28, 20, 154, 251, 244, 200, 153, 111, 185, 36, 39, 174, 65, 228, 100, 155, 147, 76, 164, 149, 153, 27, 120, 82, 184, 85, 90, 160, 6, 236, 67, 153, 221, 68, 29, 92, 10, 68, 103, 160, 79, 159, 137, 46, 99, 119, 79, 38, 75, 196, 28, 93, 112, 164, 83, 10, 148, 98},
	}
	if !reflect.DeepEqual(path2, expectedPath2) {
		t.Errorf("path is not correct. Expected:%v, Got:%v\n", expectedPath2, path2)
	}

}

func TestProof(t *testing.T) {
	mklTree := buildTree()

	// after tranctions happened,
	// receiver try to proof/verify if sender really send him a coin.

	// receiver C knows
	root := mklTree.Root() // root hash is public

	// sender B tells receiver C. Comes from to a full node.
	b := []byte(block_1)
	bIndex := 1
	path := mklTree.Path(bIndex)

	// receiver C uses Proof function to verify
	ok := Proof(b, bIndex, root, path, mklTree.hasher)
	if !ok {
		t.Error("proof fails.")
	}
}
