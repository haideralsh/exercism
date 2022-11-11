package chessboard


// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupiedCount := 0
	f, found := cb[file]

	if !found {
		return 0
	}

	for _, occupied := range f {
		if occupied {
			occupiedCount++
		}
	}

	return occupiedCount
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	occupiedCount := 0

	if rank <= 0 { return 0 }

	r := rank - 1

	for _, file := range cb {
		if r >= len(file) {
			return 0
		}

		if file[r] { occupiedCount++ }

	}

	return occupiedCount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		for range v {
			count++
		}
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		for _, s := range v {
			if s  { count++ }
		}
	}

	return count
}
