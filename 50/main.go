package main


func myPow(x float64, n int) float64 {
	if n == 0 {
		return  1
	}

	if n > 0 {
		return myPowHelp(x,n)
	}

	return 1/myPowHelp(x,-n)

}

func myPowHelp(x float64, n int) float64 {
	if n == 0 {
		return  1
	}

	if n == 1 {
		return x
	}

	v := myPowHelp(x,n >> 1)
	v = v * v
	if (n & 1) == 1 {
		return v * x
	}else{
		return v
	}
}


func main() {
}
