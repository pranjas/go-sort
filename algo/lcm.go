package algo

//LCMNaive is a slow algorithm which returns
//the LCM of the two numbers provided.
func LCMNaive(a, b int) int {
	//Attempt to find out the lowest number
	//which is divided by both a and b.
	result := a * b
	if result == 0 { //one of a or b is 0.
		return 0
	}
	for i := a * b; i >= 1; i -= 1 {
		if i%a == 0 && i%b == 0 {
			result = i
		}
	}
	return result
}

//LCMFast is a fast algorithm which gives
//LCM of two numbers a and b by exploiting the
//fact that a * b = LCM(a, b) * GCD (a, b)
func LCMFast(a, b int) int {
	gcd := GCD(a, b)
	return (a * b) / gcd
}
