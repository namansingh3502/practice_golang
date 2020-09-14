package main

import "fmt"

func main() {

	var t int

	fmt.Scan(&t)
	for ;t > 0; t-- {
		
		var n,k int;

		fmt.Scan( &n, &k )

		if (n % 2 == 1) {
			
			i := 2

			for ; i < n/2; i++ {
				if n % i == 0 {
					break
				} 	
			}

			if n % i == 0 {
				n += i + 2 * ( k - 1 )
			} else {
				n += n + 2 * ( k - 1 )
			}

		} else {
			n += 2 * k
		}

		fmt.Print(n, "\n" )
	}
}