package main

import (
	"fmt"
)

func Html(s string) string {
	return s
}

func main() {
	res := Html(`
		<h1>123</h1>
	`)
	fmt.Println(res)
}
