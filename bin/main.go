//go 1.10.4

package main
import "fmt"

func row(k int, i int, c int) {
	if k > i {
		return
	}
        fmt.Printf("%d, ", c)
	d := c * (i-k)/(k+1)
	row(k+1, i, d)
}

func col(n int, i int) {
	if i > n {
		return
	}
	row(0, i, 1)
      	fmt.Println()
	col(n, i+1)
 }

func main() {
	N := 10
	fmt.Printf("[INPUT] %d\n", N)
	fmt.Println("[OUTPUT]")
 	col(N, 0)
}
