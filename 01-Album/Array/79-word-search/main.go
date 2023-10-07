package _9_word_search



func exist(board [][]byte, word string) bool {
	// dfs
	// trie 树加速
	// 对23，28行的利用回溯剪枝和复原的讲解非常到位！
	// The word can be constructed from letters of sequentially adjacent cell,
	// where "adjacent" cells are those horizontally or vertically neighboring.
	// 那个小岛的题目（200）在思路上很类似
	// 防止revisit太精妙了

	// 什么时候 return true
	// 什么时候，停止搜索
	// 什么时候，继续深入递归

	row := len(board)
	col := len(board[0])

	// for 循环打错❌了一个条件
	// 导致很久没有 debug 出来
	for i:=0; i < row ; i++{
		for j :=0; j < col; j++{
			if search(i, j, 0, word, board) {
				return true
			}
		}
	}
	return false
}

func search(row, col , d int, word string, board [][]byte) bool {
	if row >= len(board) ||
		col >= len(board[0]) ||
		row < 0 ||
		col < 0 ||
		word[d] != board[row][col] {
		return false
	}

	if d == len(word) -1 {
		return true
	}

	tmp := board[row][col]
	board[row][col] = byte('%')

	find := search(row-1, col, d+1, word, board) ||
		search(row+1, col, d+1, word, board) ||
		search(row, col-1, d+1, word, board) ||
		search(row, col+1, d+1, word, board)

	board[row][col] = tmp
	return find
}

// 解决这题超时的办法之一
// 进行剪枝
// temp:=board[x][y]//这题注意，如果把四个可能方向都直接并列写出来全部可能性都试完，
// 会超时，用||意味着一个个方向去试，有一个方向成了就OK，这样运算步骤少了不会超时
// board[x][y]='$'

