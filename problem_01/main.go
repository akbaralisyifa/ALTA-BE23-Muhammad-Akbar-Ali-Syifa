package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	lenA := len(a);
	result := "";

	for i := 0; i <= lenA; i++{
		for j := i+ 1; j <= lenA; j++{
			substr := a[i:j]

			if strings.Contains(b, substr) && len(substr) > len(result) {
				result = string(substr)
			}
		}
	}

	return result;
}

func main() {
	fmt.Println(Compare("KANGORO", "KANG"))
	fmt.Println(Compare("WANG", "WANGI"))
}