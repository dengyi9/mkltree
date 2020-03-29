# mkltree

**merkle tree** library implemented by Go language.   

## Feature
- [x] build tree when initialized
- [x] proof a data block belonging to a merkle tree  
- [ ] add block one by one, then build tree. (Useful when reads data incrementally from big files.)
- [ ] customize tree depth. (Default is binary tree. Depth depends on number of leaves.)
- [ ] consider whether use double hash

## Reference 
- [wikipedia: merkle tree](https://en.wikipedia.org/wiki/Merkle_tree)

## Development environment
- Go: 1.12.6

