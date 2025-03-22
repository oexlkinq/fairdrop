package storage

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type Storage struct {
	Path string
}

type Folder []string

func Create(path string) *Storage {
	// #3 TODO: можно чекать, существует и доступна ли для записи папка storage
	return &Storage{path}
}

func (s *Storage) CreateFolder(password string) {
	err := os.Mkdir(filepath.Join(s.Path, password), 0700)
	if err != nil {
		panic(fmt.Errorf("create folder mkdir: %w", err))
	}
}

func (s *Storage) ListFolder(password string) *Folder {
	dir, err := os.ReadDir(filepath.Join(s.Path, password))
	if err != nil {
		panic(fmt.Errorf("list folder: %w", err))
	}

	folder := make(Folder, len(dir))
	for i, entry := range dir {
		folder[i] = entry.Name()
	}

	return &folder
}

func (s *Storage) DeleteFolder(password string) {
	err := os.RemoveAll(path.Join(s.Path, password))
	if err != nil {
		panic(fmt.Errorf("storage: delete folder: %w", err))
	}
}
