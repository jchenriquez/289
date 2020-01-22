package main

import "fmt"

const (
  PRESENTLY_DEAD = iota
  PRESENTLY_ALIVE
  PREVIOUSLY_ALIVE_NOW_DEAD
  PREVIOUSLY_DEAD_NOW_ALIVE
)

func getAliveNeighboreCount (i, j int, grid[][]int) (count int) {
  if i > 0 && (grid[i-1][j] == PRESENTLY_ALIVE || grid[i-1][j] == PREVIOUSLY_ALIVE_NOW_DEAD) {
    count++
  }

  if i < len(grid) -1 && (grid[i+1][j] == PRESENTLY_ALIVE || grid[i+1][j] == PREVIOUSLY_ALIVE_NOW_DEAD) {
    count++
  }

  if j > 0 && (grid[i][j-1] == PRESENTLY_ALIVE || grid[i][j-1] == PREVIOUSLY_ALIVE_NOW_DEAD) {
    count++
  }

  if j < len(grid[0])-1 && (grid[i][j+1] == PRESENTLY_ALIVE || grid[i][j+1] == PREVIOUSLY_ALIVE_NOW_DEAD) {
    count++
  }

  if i > 0 && j > 0 && (grid[i-1][j-1] == PRESENTLY_ALIVE || grid[i-1][j-1] == PREVIOUSLY_ALIVE_NOW_DEAD) {
    count++
  }

  if j < len(grid[0])-1 && i < len(grid)-1 && (grid[i+1][j+1] == PRESENTLY_ALIVE || grid[i+1][j+1] == PREVIOUSLY_ALIVE_NOW_DEAD) {
    count++
  }

  if i > 0 && j < len(grid[0])-1 && (grid[i-1][j+1] == PRESENTLY_ALIVE || grid[i-1][j+1] == PREVIOUSLY_ALIVE_NOW_DEAD) {
    count++
  }

  if i < len(grid)-1 && j > 0 && (grid[i+1][j-1] == PRESENTLY_ALIVE || grid[i+1][j-1] == PREVIOUSLY_ALIVE_NOW_DEAD) {
    count++
  }

  return
}

func udpateLife (i, j int, grid [][]int) {
  neighboreCount := getAliveNeighboreCount(i, j, grid)

  switch {
    case grid[i][j] == 0 && (neighboreCount == 3): 
      grid[i][j] = PREVIOUSLY_DEAD_NOW_ALIVE
    case grid[i][j] == 1 && neighboreCount < 2:
grid[i][j] = PREVIOUSLY_ALIVE_NOW_DEAD
    case grid[i][j] == 1 && neighboreCount > 3:
      grid[i][j] = PREVIOUSLY_ALIVE_NOW_DEAD
  }

}

func correct (i, j int, board [][]int) {
  if board[i][j] == PREVIOUSLY_ALIVE_NOW_DEAD {
    board[i][j] = PRESENTLY_DEAD
  } else if board[i][j] == PREVIOUSLY_DEAD_NOW_ALIVE {
    board[i][j] = PRESENTLY_ALIVE
  }
}

func gameOfLife(board [][]int)  {
    for i := 0; i < len(board); i++ {
      for j := 0; j < len(board[i]); j++ {
        udpateLife(i, j, board)
      }
    }

    for i := 0; i < len(board); i++ {
      for j := 0; j < len(board[i]); j++ {
        correct(i, j, board)
      }
    }
}

func main() {
  board := [][]int{{0,1,0}, {0,0,1}, {1,1,1}, {0,0,0}}
  gameOfLife(board)
  fmt.Printf("gameOfLife %v\n", board)
}