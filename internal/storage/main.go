package storage

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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
	err := os.Mkdir(filepath.Join(s.path, password), 0700)
	if err != nil {
		panic(fmt.Errorf("create folder mkdir: %w", err))
	}
}

func (s *Storage) ListFolder(password string) *Folder {
	dir, err := os.ReadDir(filepath.Join(s.path, password))
	if err != nil {
		panic(fmt.Errorf("list folder: %w", err))
	}

	folder := make(Folder, len(dir))
	for i, entry := range dir {
		folder[i] = entry.Name()
	}

	return &folder
}

func (s *Storage) GetFilepath(password string, filename string) string {
	fullname := filepath.Join(s.path, password, filepath.FromSlash(path.Clean("/"+filename)))

	return fullname
}
