package main

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if A > E {
		return computeArea(E, F, G, H, A, B, C, D)
	}

	if B >= H || D <= F || C <= E {
		return abs(A-C)*abs(B-D) + abs(E-G)*abs(F-H)
	}

	// 下边
	down := min(A, E)
	up := min(C, G)
	left := min(B, F)
	right := min(D, H)

	return abs(A-C)*abs(B-D) + abs(E-G)*abs(F-H) - abs(up-down)*abs(left-right)
}

func main() {

}
