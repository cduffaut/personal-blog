package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
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

// create or upodate un article
func Save(a *Article) error {
	if a.ID == 0 {
		// gen a ID based on timestamp
		a.ID = int(time.Now().Unix())
		a.PublishedAt = time.Now().Format("2006-01-02T15:04:05Z")
	}

	data, err := json.MarshalIndent(a, "", "	")
	if err != nil {
		return err
	}

	path := filepath.Join(articlesDir, strconv.Itoa(a.ID)+".json")
	return os.WriteFile(path, data, 0644)
}

func Delete(id int) error {
	path := filepath.Join(articlesDir, fmt.Sprintf("%d.json", id))
	return os.Remove(path)
}
