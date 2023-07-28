package main

func player(board selected) sign {
	if len(board) == 0 {
		return X
	}

	x, o := 0, 0
	for _, element := range board {
		switch element {
		case X:
			x++
		case O:
			o++
		}
	}

	if x >= o {
		return O
	} else {
		return X
	}
}
