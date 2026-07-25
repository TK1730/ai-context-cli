package main

import (
	"os"
	"path/filepath"
	"testing"
)

// 所有者の書き込み + グループ/その他の読み取り
const defaultFileMode = 0644

// 所有者の読み書き + グループ/その他の読み取り
const defaultDirMode = 0755

func TestListFiles(t *testing.T) {
	// ファイルやディレクトリを作成する

	// テスト用の一時ディレクトリを作成
	tmpDir := t.TempDir()

	// テストディレクトリ直下
	normalFile_1 := filepath.Join(tmpDir, "file.txt")
	err := os.WriteFile(normalFile_1, []byte("hello"), defaultFileMode)

	// テストディレクトリのサブディレクトリ直下
	normalDir := filepath.Join(tmpDir, "test")
	err = os.Mkdir(normalDir, defaultDirMode)
	normalFile_2 := filepath.Join(normalDir, "file.txt")
	err = os.WriteFile(normalFile_2, []byte("hello"), defaultFileMode)

	// 除外したいディレクトリを作成する
	gitDir := filepath.Join(tmpDir, ".git")
	os.Mkdir(gitDir, defaultDirMode)
	fatalFile := filepath.Join(gitDir, ".gitignore")
	err = os.WriteFile(fatalFile, []byte("hello"), defaultFileMode)

	// ファイル数
	const fileLen = 2

	if err != nil {
		t.Fatalf("%v", err)
	}

	// テストの実行
	files, err := ListFiles(tmpDir)

	if err != nil {
		t.Fatalf("ListFiles が失敗しました: %v", err)
	}

	if len(files) != fileLen {
		t.Errorf("期待するファイル数: %v, 実際のファイル数: %d", fileLen, len(files))
	}
}
