package storage

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

const storageFolderName = "storage"

type Storage struct {
	Path string
}

type Folder []string

func Create(dataFolderPath string) *Storage {
	storagePath := path.Join(dataFolderPath, storageFolderName)

	// проверка на существование и создание storage
	testfilePath := path.Join(storagePath, "perm.test")
	file, err := os.Create(testfilePath)
	if err != nil {
		log.Println(fmt.Errorf("cant create perm.test file in storage folder: %w", err))

		log.Println("trying to create storage folder")

		err = os.Mkdir(storagePath, 0700)
		if err != nil {
			log.Fatal(fmt.Errorf("cant create storage folder: %w", err))
		}
		log.Println("storage folder created")
	} else {
		file.Close()
		err = os.Remove(testfilePath)
		if err != nil {
			panic(err)
		}
	}

	return &Storage{storagePath}
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
