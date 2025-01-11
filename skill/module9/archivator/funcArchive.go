package archivator

import "fmt"

func Archive(s string) string {
	var curr rune
	currLen := 0
	res := ""
	for _, v := range s {
		if v != curr {
			if currLen > 0 {
				res += fmt.Sprintf("%c%d", curr, currLen)
			}
			curr = v
			currLen = 1
		} else {
			currLen++
		}
	}
	res += fmt.Sprintf("%c%d", curr, currLen)
	return res
}
