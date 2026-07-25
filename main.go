package main

import (
	"flag"
	"fmt"
)

func main() {
	// -dir という引数を定義 (デフォルトは現在のディレクトリ ".")
	targetDir := flag.String("dir", ".", "解析するディレクトリのパス")
	flag.Parse()

	// ファイル探索
	files, err := ListFiles(*targetDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 受け取ったパスを出力して確認
	fmt.Printf("対象ディレクトリ: %s\n", *targetDir)

	// ファイル探索で取得したファイル一覧
	for _, file := range files {
		fmt.Printf("ファイル: %s\n", file)
	}
}
