package git

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
) 

func GetAllRepositories(basePath string) ([]string, error) {
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("base path does not exist: %s", basePath)
	}

	var repos []string

	walkErr := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && (info.Name() == ".git" || filepath.Ext(path) == ".git") {
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
			
			_, err := git.PlainOpen(path)
			if err == nil {
				repos = append(repos, path)
			}
		}

		return nil
	})

	if walkErr != nil {
		return nil, fmt.Errorf("error walking the path: %w", walkErr)
	}
	return repos, nil
}

func LoadRepository(repoPath string) (*git.Repository, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		panic(err)
	}

	_, err = repo.Head()
	if err != nil {
		panic(err)
	}
	
	return repo, nil
}