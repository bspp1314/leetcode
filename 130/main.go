package main

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m := len(board)
	n := len(board[0])
	var dfs = func(i, j int) {}

	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}

		board[i][j] = '#'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := 0; i < m ;i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1  || j == 0 || j == n-1 {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}

func solve2(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}
	m := len(board)
	n := len(board[0])

	var queue [][2]int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || i == (len(board)-1) || j == 0 || j == (len(board[i])-1)
			if isEdge && board[i][j] == 'O' {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x, y := node[0], node[1]

		board[x][y] = '#'

		for i := 0; i < 4; i++ {
			newX := x + dx[i]
			newY := y + dy[i]

			if newX < 0 || newY < 0 || newX >= m-1 || newY >= n-1 || board[newX][newY] != 'O' {
				continue
			}

			//入队
			queue = append(queue, [2]int{newX, newY})
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}
