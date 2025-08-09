package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kashmii/fizzbuzz-go/internal/fizzbuzz" // fizzbuzzパッケージをインポート
	// "internal/fizzbuzz" パッケージの Convert 関数を使用するために必要
	// ここでは、fizzbuzzパッケージの関数を利用して、コマンドライン引数からの入力を処理する
)

func main() {
	// 1. コマンドラインから引数を受け取り、バリデーションを実施
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please provide a number as an argument.")
		return
	}

	target := args[0]
	n, err := strconv.Atoi(target)
	if err != nil {
		fmt.Printf("Invalid input: %s is not a number\n", target)
		return
	}

	if !validateInput(n) {
		return
	}

	// 2. 1から引数までをループして適切な文字列を出力する
	//    3の倍数 → "fizz"
	// 	  5の倍数 → "buzz"
	// 	  3と5両方の倍数 → "fizzbuzz"
	for i := 1; i <= n; i++ {
		result := fizzbuzz.Convert(i)
		fmt.Println(result)
	}
}

// validateInput validates that the input number is within acceptable range
func validateInput(n int) bool {
	if n < 1 {
		fmt.Printf("Invalid input: %d must be greater than 0\n", n)
		return false
	}
	if n > 100 {
		fmt.Printf("Invalid input: %d must be 100 or less\n", n)
		return false
	}
	return true
}