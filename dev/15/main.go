package main

func createHugeString(size int) string {
	// Создаем строку из рун.  Это более эффективно, чем создавать из байт,
	// особенно для больших строк + правильная обработка Unicode.
	return string(make([]rune, size))
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// Создаем копию подстроки, чтобы избежать утечки памяти
	justString = string([]rune(v[:100]))
}

func main() {
	someFunc()
}
