package internal

type ChessPiece interface {
	Move() bool
}

type Rook struct{}

type Knight struct{}

type Bishop struct{}

type Queen struct{}

type King struct{}

type Pawn struct{}

func (r *Rook) Move() bool {
	return true
}

func (k *Knight) Move() bool {
	return true
}
