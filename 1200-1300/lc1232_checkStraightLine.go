package main

func checkStraightLine2(coordinates [][]int) bool {
	// (y1-y2)(x3-x2) = (x1-x2)(y3-y2)
	if len(coordinates) < 2 {
		return true
	}
	var a, b = coordinates[0], coordinates[1]
	for i := 2; i < len(coordinates); i++ {
		var t = coordinates[i]
		if (a[0]-b[0])*(t[1]-b[1]) != (a[1]-b[1])*(t[0]-b[0]) {
			return false
		}
	}

	return true
}

func checkStraightLine(coordinates [][]int) bool {
	// y = ax + b
	if len(coordinates) < 2 {
		return true
	}
	var x1, x2, y1, y2 = coordinates[0][0], coordinates[1][0], coordinates[0][1], coordinates[1][1]
	var a, b float64
	var row bool
	if x1 == x2 {
		b = float64(x1)
	} else if y1 == y2 {
		b = float64(y1)
		row = true
	} else {
		a = float64(y2-y1) / float64(x2-x1)
		b = float64(y1) - a*float64(x1)
	}

	for i := 2; i < len(coordinates); i++ {
		x1, y1 = coordinates[i][0], coordinates[i][1]
		if a == 0 {
			if (row && int(y1) != int(b)) || (!row && int(x1) != int(b)) {
				return false
			}
		} else {
			if float64(y1) != a*float64(x1)+b {
				return false
			}
		}
	}

	return true
}
