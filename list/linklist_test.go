package list

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {

}
func gcd(x, y int) int {
	for y != 0 {
		fmt.Println(x, y)
		x, y = y, x%y
	}
	return x
}
