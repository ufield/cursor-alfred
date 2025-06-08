package libs

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

const ProjectsFile = "projects.json"
const DefaultCursorProjectsFile = "Library/Application Support/Cursor/User/globalStorage/alefragnani.project-manager/projects.json"
const CursorProjectsEnv = "CURSOR_PROJECTS_FILE"

// FileExists checks if a file exists
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// Project represents a project object
// 必要に応じてフィールドを追加
// 例: Name, Group, RootPath, Paths
// GroupやPathsは使われる関数で必要なら追加
// Pathsは[]string型
// RootPathはstring型
// Nameはstring型
// Groupはstring型
// IconPathはstring型
type Project struct {
	Name     string   `json:"name"`
	Group    string   `json:"group,omitempty"`
	RootPath string   `json:"rootPath"`
	Paths    []string `json:"paths,omitempty"`
}

// GetTitle returns the project title
func GetTitle(p Project) string {
	if p.Group == "" {
		return p.Name
	}
	return p.Name + " » " + p.Group
}

// GetSubtitle returns the expanded home path for the project
func GetSubtitle(p Project) string {
	return ExpandHomePath(p.RootPath)
}

// GetIcon returns the first icon.png found in project paths, or "icon.png" if not found
func GetIcon(p Project) string {
	for _, projectPath := range p.Paths {
		iconPath := filepath.Join(projectPath, "icon.png")
		if FileExists(iconPath) {
			return iconPath
		}
	}
	return "icon.png"
}

// GetArgument returns the expanded home path for the project
func GetArgument(p Project) string {
	return ExpandHomePath(p.RootPath)
}

// InputMatchesData: シンプルな部分一致フィルタ（fuzzysortの代替）
func InputMatchesData(list []Project, input string, keys []string) []Project {
	if input == "" || len(list) == 0 || len(keys) == 0 {
		return list
	}
	var result []Project
	inputLower := strings.ToLower(input)
	for _, p := range list {
		for _, key := range keys {
			var value string
			switch key {
			case "name":
				value = p.Name
			case "group":
				value = p.Group
			case "rootPath":
				value = p.RootPath
			}
			if strings.Contains(strings.ToLower(value), inputLower) {
				result = append(result, p)
				break
			}
		}
	}
	return result
}

// ParseProjects: nameとrootPathがあるものだけ返す
func ParseProjects(data []Project) []Project {
	var parsed []Project
	for _, p := range data {
		if p.Name != "" && p.RootPath != "" {
			parsed = append(parsed, p)
		}
	}
	return parsed
}

// Fetch: JSONファイルを読み込む（キャッシュやmaxAgeは未実装）
func Fetch(url string, transform func([]byte) ([]Project, error)) ([]Project, error) {
	b, err := os.ReadFile(url)
	if err != nil {
		return nil, err
	}
	if transform != nil {
		return transform(b)
	}
	var projects []Project
	if err := json.Unmarshal(b, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// GetCursorProjectsFilePath: Cursorのprojects.jsonのパスを返す
func GetCursorProjectsFilePath() string {
	// 環境変数があれば優先
	if env := os.Getenv(CursorProjectsEnv); env != "" {
		return ExpandHomePath(env)
	}
	home := os.Getenv("HOME")
	if home == "" {
		home, _ = os.UserHomeDir()
	}
	return filepath.Join(home, DefaultCursorProjectsFile)
}
