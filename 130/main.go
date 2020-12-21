package main


func solve(board [][]byte)  {
	if len(board) == 0   || len(board[0]) == 0 {
		return
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if i == 0 || i == (len(board) -1) || j == 0 || j == (len(board[i]) -1) {
				dfs(board,i,j)
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}else if board[i][j] == 'O'{
				board[i][j] = 'X'
			}
		}
	}

}

func dfs(board [][]byte,i int,j int)  {
	if i < 0 || j < 0 || i >=len(board)  || j >= len(board[0]) || board[i][j] != 'O'{
		// board[i][j] == '#' 说明已经搜索过了.
		return
	}

	board[i][j] = '#'
	dfs(board,i+1,j)
	dfs(board,i-1,j)
	dfs(board,i,j+1)
	dfs(board,i,j-1)
}

func solve2(board [][]byte)  {
	if len(board) == 0   || len(board[0]) == 0 {
		return
	}

	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}

	var queue [][2]int

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			isEdge :=  i == 0 || i == (len(board) -1) || j == 0 || j == (len(board[i]) -1)
			if isEdge && board[i][j] == 'O'{
				queue = append(queue,[2]int{i,j})
			}
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x,y := node[0],node[1]

		board[x][y] = '#'

		for i := 0; i < 4 ; i++ {
			newX := x + dx[i]
			newY := y + dy[i]

			if newX < 0 || newY < 0 || newX >= len(board) -1 || newY >= len(board[0]) -1 || board[newX][newY] != 'O' {
				continue
			}

			//入队
			queue  = append(queue,[2]int{newX,newY})
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}else if board[i][j] == 'O'{
				board[i][j] = 'X'
			}
		}
	}

}
