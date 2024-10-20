// Package board represents chess board with table of cells and pieces on them.
package board

import (
	"fmt"
	"strings"
)

const boardSize int = 8

// A Board represents chess board. Player faces white pieces by default.
// [0][0] is "a8", [7][7] is "h1"
type Board [boardSize][boardSize]*Piece

// MoveResult is a type returned by MovePiece()
type MoveResult struct {
	Legal    bool   // false if piece cannot go to desired position
	Captured *Piece // non-nil if piece is captured by move
}

// Constructs new Board
func New() *Board {
	b := new(Board)
	b.setDefaultPieces()
	return b
}

// MovePiece function moves piece from one cell to another.
// turn sets current player in the game
// s is a string in long algebraic form, i.e. "e2e4"
func (b *Board) MovePiece(turn Color, s string) MoveResult {
	return MoveResult{}
}

// Places all pieces to default positions
func (b *Board) setDefaultPieces() {
	// place white pieces
	b.setPieceStr(Piece{t: Rook, color: White}, "a1")
	b.setPieceStr(Piece{t: Knight, color: White}, "b1")
	b.setPieceStr(Piece{t: Bishop, color: White}, "c1")
	b.setPieceStr(Piece{t: Queen, color: White}, "d1")
	b.setPieceStr(Piece{t: King, color: White}, "e1")
	b.setPieceStr(Piece{t: Bishop, color: White}, "f1")
	b.setPieceStr(Piece{t: Knight, color: White}, "g1")
	b.setPieceStr(Piece{t: Rook, color: White}, "h1")
	for i := 0; i < boardSize; i++ {
		b.setPiece(Piece{t: Pawn, color: White}, 6, i)
	}
	// place black pieces
	b.setPieceStr(Piece{t: Rook, color: Black}, "a8")
	b.setPieceStr(Piece{t: Knight, color: Black}, "b8")
	b.setPieceStr(Piece{t: Bishop, color: Black}, "c8")
	b.setPieceStr(Piece{t: Queen, color: Black}, "d8")
	b.setPieceStr(Piece{t: King, color: Black}, "e8")
	b.setPieceStr(Piece{t: Bishop, color: Black}, "f8")
	b.setPieceStr(Piece{t: Knight, color: Black}, "g8")
	b.setPieceStr(Piece{t: Rook, color: Black}, "h8")
	for i := 0; i < boardSize; i++ {
		b.setPiece(Piece{t: Pawn, color: Black}, 1, i)
	}
}

// Place piece p by pos string, e.g. "e2"
func (b *Board) setPieceStr(p Piece, pos string) {
	row, col := str2coord(pos)
	b.setPiece(p, row, col)
}

// Converts algebraic notation "e5" to internal coordinates [4][3]
func str2coord(s string) (row int, col int) {
	// "a1"
	// col: a -> 0
	// row: 1 -> 7

	// "e5"
	// col: e -> 4
	// row: 5 -> 3
	col = int(s[0] - 'a')
	row = int(-((s[1] - '0') - 8))
	return
}

// Converts internal coordinates [4][3] to algebraic notation "e5"
func coord2str(row int, col int) string {
	b := make([]byte, 2)
	b[0] = byte(col + 'a')
	b[1] = byte(8 - row + '0') // 7 -> 1; 6 -> 2; etc.
	return string(b)
}

// Place piece p by coordinates specified by row and col
func (b *Board) setPiece(p Piece, row int, col int) {
	isValid := func(n int) bool {
		return n >= 0 && n < boardSize
	}
	if !isValid(row) || !isValid(col) {
		panic(fmt.Sprintf("setPiece: wrong piece coordinates (row=%v, col=%v)", row, col))
	}
	b[row][col] = &p
}

// Returns Board in human-readable form
func (b *Board) String() string {
	var sb strings.Builder
	for i := 0; i < boardSize+1; i++ {
		for j := 0; j < boardSize+1; j++ {
			if i == 0 && j == 0 {
				fmt.Fprintf(&sb, " ")
				continue
			}
			if i == 0 { // top "a", "b", "c", ...
				s := coord2str(i+1, j-1)
				fmt.Fprintf(&sb, " %v ", s[:1])
				continue
			}
			if j == 0 { // left "8", "7", "6", ...
				s := coord2str(i-1, j+1)
				fmt.Fprintf(&sb, "%v", s[1:])
				continue
			}
			if b[i-1][j-1] == nil {
				fmt.Fprintf(&sb, " - ")
			} else {
				fmt.Fprintf(&sb, " %v ", b[i-1][j-1])
			}
		}
		fmt.Fprintf(&sb, "\n")
	}
	return sb.String()
}
