package storage

import (
	"encoding/json"
	"os"

	"github.com/Revolyshn/todo-manager/internal/task"
)

type Storage struct {
	FilePath string
}

func NewStorage(filePath string) *Storage {
	return &Storage{
		FilePath: filePath,
	}
}

func (s *Storage) Save(tasks []task.Task) error {
	file, err := os.Create(s.FilePath) // Создает файл, если его нет
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func (s *Storage) Load() ([]task.Task, error) {
	// Проверяем, существует ли файл
	if _, err := os.Stat(s.FilePath); os.IsNotExist(err) {
		// Если файл не существует, возвращаем пустой список задач
		return []task.Task{}, nil
	}

	file, err := os.Open(s.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []task.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
