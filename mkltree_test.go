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
  [197 68 188 221 192 2 160 225 67 110 72 240 144 247 19 5 208 169 99 179 89 123 122 31 14 38 255 44 148 158 207 241],  [161 165 11 180 22 163 114 30 251 130 139 161 101 207 11 117 18 21 167 77 197 151 36 183 120 3 24 227 125 250 63 32],  [124 47 19 12 24 227 85 79 48 25 120 202 225 209 250 240 41 237 184 159 169 148 172 43 45 47 13 11 118 214 28 107],
  [248 132 2 204 112 235 8 92 5 64 164 95 196 182 17 183 148 167 64 254 16 29 25 113 173 230 70 91 145 58 100 173],  [51 254 179 44 159 105 192 191 88 197 231 157 175 14 40 87 7 47 151 241 194 243 246 3 237 25 8 170 210 211 197 130],
  [214 204 103 121 8 26 54 76 98 8 238 16 114 231 3 179 22 66 61 76 188 121 232 131 82 97 216 133 206 127 255 39],

path (mt-X-Y is the X-row Y-column element in merkle tree ):
	path_0: block_0 [mt-0-0] -> expectedPath [[mt-0-1], [mt-1-1]]
	path_1: block_1 [mt-0-1] -> expectedPath [[mt-0-0], [mt-1-1]]
	path_2: block_2 [mt-0-2] -> expectedPath [[mt-0-3], [mt-1-0]]
*/
func TestMklTree_Path(t *testing.T) {
	mtree := buildTree()

	fmt.Println(mtree.StringBytes())

	path := mtree.Path(0)
	expectedPath := [][]byte{
		[]byte{161,165,11,180,22,163,114,30,251,130,139,161,101,207,11,117,18,21,167,77,197,151,36,183,120,3,24,227,125,250,63,32},
		[]byte{51,254,179,44,159,105,192,191,88,197,231,157,175,14,40,87,7,47,151,241,194,243,246,3,237,25,8,170,210,211,197,130},
	}
	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("path is not correct. Expected:%v, Got:%v\n", expectedPath, path)
	}

	path1 := mtree.Path(1)
	expectedPath1 := [][]byte{
		[]byte{197,68,188,221,192,2,160,225,67,110,72,240,144,247,19,5,208,169,99,179,89,123,122,31,14,38,255,44,148,158,207,241},
		[]byte{51,254,179,44,159,105,192,191,88,197,231,157,175,14,40,87,7,47,151,241,194,243,246,3,237,25,8,170,210,211,197,130},
	}
	if !reflect.DeepEqual(path1, expectedPath1) {
		t.Errorf("path is not correct. Expected:%v, Got:%v\n", expectedPath1, path1)
	}

	path2 := mtree.Path(2)
	expectedPath2 := [][]byte{
		[]byte{},
		[]byte{248,132,2,204,112,235,8,92,5,64,164,95,196,182,17,183,148,167,64,254,16,29,25,113,173,230,70,91,145,58,100,173},
	}
	if !reflect.DeepEqual(path2, expectedPath2) {
		t.Errorf("path is not correct. Expected:%v, Got:%v\n", expectedPath2, path2)
	}

}

func TestProof(t *testing.T) {
	mklTree := buildTree()
	fmt.Println(mklTree.StringBytes())
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


func TestProofFake(t *testing.T) {
	mklTree := buildTree()
	fmt.Println(mklTree.StringBytes())
	// after tranctions happened,
	// receiver try to proof/verify if sender really send him a coin.

	// receiver C knows
	root := mklTree.Root() // root hash is public

	// sender B tells receiver C. Comes from to a full node.
	b := []byte(block_1 + " ") // incorrect block info
	bIndex := 1
	path := mklTree.Path(bIndex)

	// receiver C uses Proof function to verify
	shouldFalse := Proof(b, bIndex, root, path, mklTree.hasher)
	if shouldFalse == true {
		t.Error("proof fails.")
	}
}