package utils

func Binom (n int, k int) (binom int){
	if n == k || k == 1 {
		return 1
	} else {
		return int(fac(n) / (fac(k) * fac(n - k)))
	}
}

func fac(n int) (fac int) {
	if n == 0 || n == 1 {
		return 1
	} else {
		var fac = 1
		for i := 2; i <= n; i++ {
			fac *= i	
		}
		return fac
	}
}

