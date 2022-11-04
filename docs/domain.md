# Minesweeper Domain

## Minesweeper Game

### Gameboard
The first order of business is to create the representation of a minesweeper game board. A minesweeper game consists of a collection of tiles, creating a x by y sized board. We can represent this in two ways within go, either using a slice or an array. 

An array in go has a set size and you cannot extend/shrink the number of elements within one. On the flip side, a slice can be appended to and decreased. Another consideration is that this board is a 2D construct. We can only create such a struct by using slices. 

```
board := make([][]int, y) // slices created are initialized with nothing
for i := range board {
    board[i] = make([]int, x)
}
```
Code snippet [source](https://stackoverflow.com/questions/57482282/multi-dimensional-arrays-in-go)

Side note: go uses a [garbage collector](https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8) for memory management
