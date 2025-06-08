package libs

import (
	"os"
	"strings"
)

var (
	homeDir, _       = os.UserHomeDir()
	HomePathVariable = "$home"
)

type Item struct {
	Description string
	// 必要に応じて他のフィールドを追加
}

// pathIsUNC: パスがUNCパスかどうか判定
func PathIsUNC(path string) bool {
	return strings.HasPrefix(path, `\\`)
}

// compactHomePath: ホームディレクトリ配下なら$homeに置換
func CompactHomePath(path string) string {
	if !strings.HasPrefix(path, homeDir) {
		return path
	}
	return strings.Replace(path, homeDir, HomePathVariable, 1)
}

// expandHomePath: $homeを実際のホームディレクトリに展開
func ExpandHomePath(path string) string {
	if !strings.HasPrefix(path, HomePathVariable) {
		return path
	}
	return strings.Replace(path, HomePathVariable, homeDir, 1)
}

// expandHomePaths: Item配列のdescriptionを展開
func ExpandHomePaths(items []Item) []Item {
	result := make([]Item, len(items))
	for i, item := range items {
		item.Description = ExpandHomePath(item.Description)
		result[i] = item
	}
	return result
}
