package main

import (
	"flag"
	"fmt"
)

func main() {
	// -dir という引数を定義 (デフォルトは現在のディレクトリ ".")
	targetDir := flag.String("dir", ".", "解析するディレクトリのパス")
	flag.Parse()

	// 受け取ったパスを出力して確認
	fmt.Printf("対象ディレクトリ: %s\n", *targetDir)
}
