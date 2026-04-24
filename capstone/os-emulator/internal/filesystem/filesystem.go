package filesystem

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type FileType int

const (
	TypeFile FileType = iota
	TypeDirectory
)

type FSNode struct {
	Name       string
	Type       FileType
	Content    []byte
	Children   map[string]*FSNode
	Parent     *FSNode
	CreatedAt  time.Time
	ModifiedAt time.Time
	Mu         sync.RWMutex
}

type FileSystem struct {
	root *FSNode
	Mu   sync.RWMutex
}

func NewFileSystem() *FileSystem {
	root := &FSNode{
		Name:       "/",
		Type:       TypeDirectory,
		Children:   make(map[string]*FSNode),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
	return &FileSystem{root: root}
}

func (fs *FileSystem) CreateDirectory(path string) error {
	fs.Mu.Lock()
	defer fs.Mu.Unlock()

	parts := strings.Split(strings.Trim(path, "/"), "/")
	current := fs.root

	for _, part := range parts {
		if part == "" {
			continue
		}

		current.Mu.Lock()
		if child, exists := current.Children[part]; exists {
			current.Mu.Unlock()
			if child.Type != TypeDirectory {
				return fmt.Errorf("path exists but is not a directory")
			}
			current = child
		} else {
			newDir := &FSNode{
				Name:       part,
				Type:       TypeDirectory,
				Children:   make(map[string]*FSNode),
				Parent:     current,
				CreatedAt:  time.Now(),
				ModifiedAt: time.Now(),
			}
			current.Children[part] = newDir
			current.Mu.Unlock()
			current = newDir
		}
	}

	return nil
}

func (fs *FileSystem) CreateFile(path string, content []byte) error {
	fs.Mu.Lock()
	defer fs.Mu.Unlock()

	parts := strings.Split(strings.Trim(path, "/"), "/")
	current := fs.root

	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		if part == "" {
			continue
		}

		current.Mu.Lock()
		if child, exists := current.Children[part]; exists {
			current.Mu.Unlock()
			if child.Type != TypeDirectory {
				return fmt.Errorf("path is not a directory")
			}
			current = child
		} else {
			current.Mu.Unlock()
			return fmt.Errorf("directory not found")
		}
	}

	filename := parts[len(parts)-1]
	current.Mu.Lock()
	defer current.Mu.Unlock()

	if _, exists := current.Children[filename]; exists {
		return fmt.Errorf("file already exists")
	}

	newFile := &FSNode{
		Name:       filename,
		Type:       TypeFile,
		Content:    content,
		Parent:     current,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
	current.Children[filename] = newFile

	return nil
}

func (fs *FileSystem) ReadFile(path string) ([]byte, error) {
	fs.Mu.RLock()
	defer fs.Mu.RUnlock()

	node, err := fs.getNode(path)
	if err != nil {
		return nil, err
	}

	node.Mu.RLock()
	defer node.Mu.RUnlock()

	if node.Type != TypeFile {
		return nil, fmt.Errorf("not a file")
	}

	return node.Content, nil
}

func (fs *FileSystem) WriteFile(path string, content []byte) error {
	fs.Mu.RLock()
	defer fs.Mu.RUnlock()

	node, err := fs.getNode(path)
	if err != nil {
		return err
	}

	node.Mu.Lock()
	defer node.Mu.Unlock()

	if node.Type != TypeFile {
		return fmt.Errorf("not a file")
	}

	node.Content = content
	node.ModifiedAt = time.Now()

	return nil
}

func (fs *FileSystem) ListDirectory(path string) ([]string, error) {
	fs.Mu.RLock()
	defer fs.Mu.RUnlock()

	node, err := fs.getNode(path)
	if err != nil {
		return nil, err
	}

	node.Mu.RLock()
	defer node.Mu.RUnlock()

	if node.Type != TypeDirectory {
		return nil, fmt.Errorf("not a directory")
	}

	var names []string
	for name := range node.Children {
		names = append(names, name)
	}

	return names, nil
}

func (fs *FileSystem) DeleteFile(path string) error {
	fs.Mu.Lock()
	defer fs.Mu.Unlock()

	parts := strings.Split(strings.Trim(path, "/"), "/")
	current := fs.root

	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		if part == "" {
			continue
		}

		current.Mu.Lock()
		if child, exists := current.Children[part]; exists {
			current.Mu.Unlock()
			current = child
		} else {
			current.Mu.Unlock()
			return fmt.Errorf("path not found")
		}
	}

	filename := parts[len(parts)-1]
	current.Mu.Lock()
	defer current.Mu.Unlock()

	if _, exists := current.Children[filename]; !exists {
		return fmt.Errorf("file not found")
	}

	delete(current.Children, filename)
	return nil
}

func (fs *FileSystem) getNode(path string) (*FSNode, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	current := fs.root

	for _, part := range parts {
		if part == "" {
			continue
		}

		current.Mu.RLock()
		child, exists := current.Children[part]
		current.Mu.RUnlock()

		if !exists {
			return nil, fmt.Errorf("path not found: %s", path)
		}

		current = child
	}

	return current, nil
}

func (fs *FileSystem) Initialize() error {
	fs.CreateDirectory("/home")
	fs.CreateDirectory("/tmp")
	fs.CreateDirectory("/var")
	fs.CreateDirectory("/bin")
	return nil
}
