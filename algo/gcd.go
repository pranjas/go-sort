package algo

//GCDNaive returns the gcd of two numbers, which can
//either be positive or negative. This is a very slow
//algorithm to calculate GCD.
func GCDNaive(a, b int) int {
	//If either of a or b is negative just
	//make them positive.
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	//The Naive algorithm will loop until
	//maxLoop. The idea is to try and divide
	//using all numbers and at the end of the
	//loop we shall have the biggest number that
	//divides both a and b.
	var maxLoop int
	gcd := 1

	if a > b {
		maxLoop = b
	} else {
		maxLoop = a
	}
	for i := 1; i < maxLoop; i += 1 {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return gcd
}

//GCD returns the gcd of two numbers, which can either
//be positive or negative. This is an implementation of
//Euclidean Algorithm.
func GCD(a, b int) int {
	//Turn negative number into a positive one
	//Same as we did in Naive algorithm.
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	//Things are interesting here, we change a and b in case
	//the lower of the two doesn't divides the larger one.
	//we only keep the remainder and iterate until one of a or
	//b becomes 0. Note that this is bound to happen.
	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	//The number which is non-zero is our GCD!
	if a != 0 {
		return a
	}
	return b
}
