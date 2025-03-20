package storage

import (
	"fmt"
	"os"
	"path"
)

type Storage struct {
	path string
}

type Folder []string

func Create(path string) *Storage {
	// TODO: можно чекать, существует и доступна ли для записи папка storage
	return &Storage{path}
}

func (s *Storage) CreateFolder(password string) {
	err := os.Mkdir(path.Join(s.path, password), 0700)
	if err != nil {
		panic(fmt.Errorf("create folder mkdir: %w", err))
	}
}

func (s *Storage) ListFolder(password string) *Folder {
	dir, err := os.ReadDir(path.Join(s.path, password))
	if err != nil {
		panic(fmt.Errorf("list folder: %w", err))
	}

	folder := make(Folder, len(dir))
	for i, entry := range dir {
		folder[i] = entry.Name()
	}

	return &folder
}
