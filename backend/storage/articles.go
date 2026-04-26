package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
)

type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	PublishedAt string `json:"published_at"`
}

const articlesDir = "./articles"

// return all articles sorted per date
func GetAll() ([]Article, error) {
	entries, err := os.ReadDir(articlesDir)

	if err != nil {
		return nil, err
	}

	var articles []Article
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(articlesDir, entry.Name()))
		if err != nil {
			continue
		}
		var a Article
		if err := json.Unmarshal(data, &a); err == nil {
			articles.append(articles, a)
		}
	}

	// sorted, the most recent first
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].ID > articles[j].ID
	})

	return articles, nil
}
