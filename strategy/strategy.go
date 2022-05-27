package strategy

import (
	"fmt"
)

type FileSystem interface {
	Open(path string) ([]byte, error)
	Directory() string
	Name() string
}

type FileManager struct {
	mainFileSystem  FileSystem
	minorFileSystem FileSystem
}

func NewFileManager(mainFS string, fs ...FileSystem) *FileManager {
	fm := &FileManager{}
	for _, system := range fs {
		if system.Name() == mainFS {
			fm.mainFileSystem = system
			continue
		}
		fm.minorFileSystem = system
	}
	return fm
}

func (f *FileManager) OpenImage(name string) ([]byte, error) {
	return f.mainFileSystem.Open(fmt.Sprintf("%s/%s", f.mainFileSystem.Directory(), name))
}

func (f *FileManager) SetMainFileSystem(name string) error {
	switch name {
	case f.mainFileSystem.Name():

	case f.minorFileSystem.Name():
		f.mainFileSystem, f.minorFileSystem = f.minorFileSystem, f.mainFileSystem
	default:
		return fmt.Errorf("не найденно нужной файловой системы")
	}
	return nil
}
