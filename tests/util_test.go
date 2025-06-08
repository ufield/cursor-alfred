package libs_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ufield/cursor-alfred/libs"
)

// TestFileExists は FileExists 関数のテストを行います。
// このテストでは以下の2つのケースを確認します：
// 1. 実際に存在するファイルに対して true を返すか
// 2. 存在しないファイルに対して false を返すか
func TestFileExists(t *testing.T) {
	// 一時ファイルを作成してテスト
	f, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	// テスト終了時に一時ファイルを削除
	defer os.Remove(f.Name())

	// ケース1: 存在するファイル
	if !libs.FileExists(f.Name()) {
		t.Error("FileExists should return true for existing file")
	}

	// ケース2: 存在しないファイル
	if libs.FileExists("/no/such/file/exists") {
		t.Error("FileExists should return false for non-existing file")
	}
}

func TestGetTitle(t *testing.T) {
	p := libs.Project{Name: "foo", Group: "bar"}
	if libs.GetTitle(p) != "foo » bar" {
		t.Error("GetTitle failed with group")
	}
	p.Group = ""
	if libs.GetTitle(p) != "foo" {
		t.Error("GetTitle failed without group")
	}
}

func TestGetSubtitle(t *testing.T) {
	p := libs.Project{RootPath: libs.HomePathVariable + "/test"}
	subtitle := libs.GetSubtitle(p)
	home, _ := os.UserHomeDir()
	if filepath.Dir(subtitle) != filepath.Join(home, "test")[:len(filepath.Dir(subtitle))] {
		t.Error("GetSubtitle did not expand home path correctly")
	}
}

func TestGetIcon(t *testing.T) {
	dir := t.TempDir()
	iconPath := filepath.Join(dir, "icon.png")
	os.WriteFile(iconPath, []byte("dummy"), 0644)
	p := libs.Project{Paths: []string{dir}}
	if libs.GetIcon(p) != iconPath {
		t.Error("GetIcon did not find icon.png")
	}
	p = libs.Project{Paths: []string{"/no/such/path"}}
	if libs.GetIcon(p) != "icon.png" {
		t.Error("GetIcon should return default icon.png if not found")
	}
}

func TestGetArgument(t *testing.T) {
	p := libs.Project{RootPath: libs.HomePathVariable + "/test"}
	arg := libs.GetArgument(p)
	home, _ := os.UserHomeDir()
	if filepath.Dir(arg) != filepath.Join(home, "test")[:len(filepath.Dir(arg))] {
		t.Error("GetArgument did not expand home path correctly")
	}
}

func TestInputMatchesData(t *testing.T) {
	list := []libs.Project{
		{Name: "foo", Group: "bar", RootPath: "/a"},
		{Name: "baz", Group: "qux", RootPath: "/b"},
	}
	result := libs.InputMatchesData(list, "foo", []string{"name"})
	if len(result) != 1 || result[0].Name != "foo" {
		t.Error("InputMatchesData did not filter by name")
	}
	result = libs.InputMatchesData(list, "qux", []string{"group"})
	if len(result) != 1 || result[0].Group != "qux" {
		t.Error("InputMatchesData did not filter by group")
	}
	result = libs.InputMatchesData(list, "", []string{"name"})
	if len(result) != 2 {
		t.Error("InputMatchesData should return all if input is empty")
	}
}

func TestParseProjects(t *testing.T) {
	list := []libs.Project{
		{Name: "foo", RootPath: "/a"},
		{Name: "", RootPath: "/b"},
		{Name: "bar", RootPath: ""},
	}
	parsed := libs.ParseProjects(list)
	if len(parsed) != 1 || parsed[0].Name != "foo" {
		t.Error("ParseProjects did not filter correctly")
	}
}

func TestFetch(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "projects.json")
	projects := []libs.Project{{Name: "foo", RootPath: "/a"}}
	b, _ := json.Marshal(projects)
	os.WriteFile(file, b, 0644)
	result, err := libs.Fetch(file, nil)
	if err != nil || !reflect.DeepEqual(result, projects) {
		t.Error("Fetch did not read or unmarshal correctly")
	}
}

func TestGetCursorProjectsFilePath(t *testing.T) {
	// 環境変数未設定時のデフォルトパス
	os.Unsetenv("CURSOR_PROJECTS_FILE")
	home, _ := os.UserHomeDir()
	want := filepath.Join(home, "Library/Application Support/Cursor/User/globalStorage/alefragnani.project-manager/projects.json")
	got := libs.GetCursorProjectsFilePath()
	if got != want {
		t.Errorf("GetCursorProjectsFilePath default: got %s, want %s", got, want)
	}
}
