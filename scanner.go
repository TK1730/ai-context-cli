package main

import (
	"io/fs"
	"path/filepath"
)

// 探索から除外するディレクトリのパス
var ignoreDirs = map[string]struct{}{
	".git":         {},
	"node_modules": {},
	".next":        {},
	".venv":        {},
}

// ListFiles は指定されたディレクトリを探索し、対象ファイルのパス一覧を返す
func ListFiles(rootDir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // 権限エラーなどがあれば終了
		}

		// ディレクトリ判定
		if d.IsDir() {
			// 除外リストに含まれているかチェック
			if _, exists := ignoreDirs[d.Name()]; exists {
				return filepath.SkipDir
			}
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
