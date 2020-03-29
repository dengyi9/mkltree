# mkltree

**merkle tree** library implemented in Go language.   

## Feature
- [x] build tree when initialized
- [x] proof whether a data block belongs to a merkle tree  
- [ ] add block one by one, then build tree. (Useful when reads data incrementally from big files.)
- [ ] customize tree depth. (Default is binary tree. Depth depends on number of leaves.)
- [ ] consider whether use double hash

## How to Use?
please see [example.go](./example/example.go)

## Reference 
- [wikipedia: merkle tree](https://en.wikipedia.org/wiki/Merkle_tree)

## Development environment
- Go: 1.12.6

