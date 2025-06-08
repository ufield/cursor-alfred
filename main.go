package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ufield/cursor-alfred/libs"
)

type AlfredItem struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle,omitempty"`
	Arg      string `json:"arg,omitempty"`
	Icon     *struct {
		Path string `json:"path"`
	} `json:"icon,omitempty"`
	Valid bool              `json:"valid,omitempty"`
	Text  map[string]string `json:"text,omitempty"`
}

type AlfredOutput struct {
	Items []AlfredItem `json:"items"`
}

func main() {
	// Alfredの入力は環境変数"query"から取得
	input := os.Getenv("query")

	file := libs.GetCursorProjectsFilePath()
	projects, err := libs.Fetch(file, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load projects: %v\n", err)
		os.Exit(1)
	}

	matched := libs.InputMatchesData(projects, input, []string{"name", "group"})
	sort.Slice(matched, func(i, j int) bool {
		return strings.Compare(matched[i].Name, matched[j].Name) < 0
	})

	var items []AlfredItem
	for _, project := range matched {
		item := AlfredItem{
			Title:    libs.GetTitle(project),
			Subtitle: libs.GetSubtitle(project),
			Arg:      libs.GetArgument(project),
			Valid:    true,
			Text:     map[string]string{"copy": libs.GetArgument(project)},
			Icon: &struct {
				Path string `json:"path"`
			}{Path: libs.GetIcon(project)},
		}
		items = append(items, item)
	}

	if len(items) == 0 {
		items = append(items, AlfredItem{Title: "No projects found"})
	}

	output := AlfredOutput{Items: items}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(output); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to encode output: %v\n", err)
		os.Exit(1)
	}
}
