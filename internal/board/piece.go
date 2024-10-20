package board

// Enumeration of all pieces types
type PieceType uint8

const (
	King PieceType = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

// Chess color, used to determine pieces and player color
type Color bool

const (
	White Color = false
	Black Color = true
)

// A Piece represents chess's piece (pawn, rook, etc.)
type Piece struct {
	t     PieceType
	color Color
}

func (p *Piece) CanMove(board *Board, src string, dst string) bool {

	return false
}

func (p Piece) String() string {
	if p.t == King && p.color == White {
		return "♔"
	}
	if p.t == Queen && p.color == White {
		return "♕"
	}
	if p.t == Rook && p.color == White {
		return "♖"
	}
	if p.t == Bishop && p.color == White {
		return "♗"
	}
	if p.t == Knight && p.color == White {
		return "♘"
	}
	if p.t == Pawn && p.color == White {
		return "♙"
	}
	if p.t == King && p.color == Black {
		return "♚"
	}
	if p.t == Queen && p.color == Black {
		return "♛"
	}
	if p.t == Rook && p.color == Black {
		return "♜"
	}
	if p.t == Bishop && p.color == Black {
		return "♝"
	}
	if p.t == Knight && p.color == Black {
		return "♞"
	}
	if p.t == Pawn && p.color == Black {
		return "♟︎"
	}
	return "?"
}
