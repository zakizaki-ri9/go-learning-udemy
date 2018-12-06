package mylib

import "fmt"

// アクセス可能な変数
var Public string = "Public"

// アクセス不可能な変数
var private string = "Private"

// 人クラス
type Person struct {
	// 名前
	Name string
	// 年齢
	Age int
}

// "Human!"を出力する
func Say() {
	fmt.Println("Human!")
}
