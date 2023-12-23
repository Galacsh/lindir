package types

import (
	"bufio"
	"os"
	"path/filepath"
)

type Path string

func (p Path) String() string {
	return string(p)
}

func (p Path) Join(strs ...string) Path {
	pre := p.String()
	post := filepath.Join(strs...)

	return Path(filepath.Join(pre, post))
}

func (p Path) Abs() (Path, error) {
	abs, err := filepath.Abs(p.String())
	if err != nil {
		return "", err
	}
	return Path(abs), nil
}

func (p Path) Dir() Path {
	return Path(filepath.Dir(p.String()))
}

func (p Path) exists() (bool, error) {
	_, err := os.Stat(p.String())

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (p Path) NotExists() (bool, error) {
	exists, err := p.exists()
	if err != nil {
		return false, err
	}

	return !exists, nil
}

func (p Path) CreateDir() error {
	return os.MkdirAll(p.String(), os.ModePerm)
}

func (p Path) Create() error {
	dirNotExists, err := p.Dir().NotExists()
	if err != nil {
		return err
	}

	if dirNotExists {
		err = p.Dir().CreateDir()
		if err != nil {
			return err
		}
	}

	file, err := os.Create(p.String())
	if err != nil {
		return err
	}

	return file.Close()
}

func (p Path) Remove() error {
	return os.RemoveAll(p.String())
}

func (p Path) Read() (PathSet, error) {
	file, err := os.Open(p.String())
	if err != nil {
		return nil, err
	}

	paths := PathSet{}

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		text := scanner.Text()
		if text == "" || text[0] == '#' {
			continue
		}
		paths.Add(text)
	}

	return paths, file.Close()
}

func (p Path) Write(paths PathSet) error {
	file, err := os.Create(p.String())
	if err != nil {
		return err
	}

	for path := range paths {
		_, err = file.WriteString(path + "\n")
		if err != nil {
			return err
		}
	}

	return file.Close()
}
