package main

import (
	"fmt"
	"log"
)

func phiFunc(n int) int {
	result := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}

func isPrime(n int) {
	if n > 1 {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				log.Fatal("Число не простое")
			}
		}
	} else {
		log.Fatal("Число не простое")
	}
}

func eFind(phi int) int {
	for e := 2; e < phi; e++ {
		if gcd(e, phi) == 1 {
			return e
		}
	}
	return 1
}

func gcd(m int, n int) int {
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	if m == n {
		return m
	}
	if m == 1 && n != 1 {
		return m
	}
	if m != 1 && n == 1 {
		return n
	}

	var result int

	if m%2 == 0 && n%2 == 0 {
		result = 2 * gcd(m/2, n/2)
	} else if m%2 == 0 && n%2 != 0 {
		result = gcd(m/2, n)
	} else if m%2 != 0 && n%2 == 0 {
		result = gcd(m, n/2)
	} else if m%2 != 0 && n%2 != 0 {
		if m > n {
			result = gcd((m-n)/2, n)
		} else if m < n {
			result = gcd(m, (n-m)/2)
		}
	}
	return result

}

func dFind(e int, phi int) int {
	d := 1
	for ; mod(d*e, phi) != 1; d++ {
	}
	return d
}

func mod(n int, m int) int {
	n %= m
	if n < 0 {
		n += m
	}
	return n
}

func encrypt(s string, e int, m int) {
	result := make([]int, len(s))

	for i, b := range s {
		if int(b) >= m {
			log.Fatal("Модуль меньше числа")
		}
		result[i] = degree(int(b), e, m)
	}

	fmt.Println("Зашифрованный текст: \n", result)
}

func decrypt(s []int, d int, m int) {
	result := make([]byte, len(s))

	for i, b := range s {
		if b >= m {
			log.Fatal("Модуль меньше числа")
		}
		result[i] = byte(degree(b, d, m))
	}

	fmt.Println("Дешифрованный текст: \n", string(result))
}

func degree(n int, d int, m int) int {
	result := 1
	for i := 0; i < d; i++ {
		result = mod(result*n, m)
	}
	return result
}
