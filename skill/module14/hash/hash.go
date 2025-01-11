package main

import "fmt"

func hashint64(val int64) uint64 {
	return uint64(val % 4294967)
}

func hashstr(val string) uint64 {
	var sum int64 = 0
	for _, char := range val {
		sum += int64(char) // Суммируем ASCII значения символов
	}
	return uint64(sum % 4294967295) // Возвращаем остаток от деления на 4294967295 (2^32-1)
}

func main() {
	valHash64 := 188463345437839030
	valHashStr := "awguqwrfrs;aDFOIEUYlqwiufsg;seoioirguhselriguehrulaweoiirstu"
	fmt.Println(hashint64(int64(valHash64)))
	fmt.Println(hashstr(valHashStr))
}
