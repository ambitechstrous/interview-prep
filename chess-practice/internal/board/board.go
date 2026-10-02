package board

type Board struct {
	board map[int]map[int]any
}

func (b *Board) PlacePiece(x, y int, piece any) {
	existingPiece, ok := b.board[x][y]
	if ok && existingPiece != nil {

	}

	if b.board[x] == nil {
		b.board[x] = make(map[int]any)
	}
	b.board[x][y] = piece
}

func (b *Board) GetPiece(x, y int) Piece {
	if b.board[x] == nil {
		return nil
	}
	return b.board[x][y]
}

func NewBoard(n int) *Board {
	return &Board{
		board: make(map[int]map[int]any),
	}
}
