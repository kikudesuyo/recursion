package intermediate

func gcd(x, y int32) int32 {
	gcdX, gcdY := x, y
	for {
		if gcdX == gcdY {
			return gcdX
		}
		if gcdX > gcdY {
			gcdX -= gcdY
		} else {
			gcdY -= gcdX
		}
	}
}

func CountSquare(x int32, y int32) int32 {
	gcd := gcd(x, y)
	return (x / gcd) * (y / gcd)
}
