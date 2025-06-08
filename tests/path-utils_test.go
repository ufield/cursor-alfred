package libs_test

import (
	"os"
	"strings"
	"testing"

	"github.com/ufield/cursor-alfred/libs"
)

func TestPathIsUNC(t *testing.T) {
	if !libs.PathIsUNC(`\\server\share`) {
		t.Error("UNCパスを正しく判定できません")
	}
	if libs.PathIsUNC(`/home/user`) {
		t.Error("通常パスを誤ってUNCと判定しています")
	}
}

func TestCompactHomePath(t *testing.T) {
	home, _ := os.UserHomeDir()
	path := home + "/project"
	compacted := libs.CompactHomePath(path)
	if !strings.HasPrefix(compacted, libs.HomePathVariable) {
		t.Errorf("ホームディレクトリが$homeに置換されていません: %s", compacted)
	}
	// ホームディレクトリ外は変換しない
	other := "/tmp/project"
	if libs.CompactHomePath(other) != other {
		t.Error("ホームディレクトリ外のパスが変換されています")
	}
}

func TestExpandHomePath(t *testing.T) {
	home, _ := os.UserHomeDir()
	path := libs.HomePathVariable + "/project"
	expanded := libs.ExpandHomePath(path)
	if !strings.HasPrefix(expanded, home) {
		t.Errorf("$homeがホームディレクトリに展開されていません: %s", expanded)
	}
	// $homeで始まらない場合は変換しない
	other := "/tmp/project"
	if libs.ExpandHomePath(other) != other {
		t.Error("$homeで始まらないパスが変換されています")
	}
}

func TestExpandHomePaths(t *testing.T) {
	home, _ := os.UserHomeDir()
	items := []libs.Item{
		{Description: libs.HomePathVariable + "/foo"},
		{Description: "/tmp/bar"},
	}
	result := libs.ExpandHomePaths(items)
	if !strings.HasPrefix(result[0].Description, home) {
		t.Error("$homeが展開されていません")
	}
	if result[1].Description != "/tmp/bar" {
		t.Error("$homeで始まらないパスが変換されています")
	}
}
