package todo

import (
    "encoding/json"
    "os"
)

const fileName = "todos.json"

func SaveTodos(todos []Todo) error {
    data, err := json.MarshalIndent(todos, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(fileName, data, 0644)
}

func LoadTodos() ([]Todo, error) {
    var todos []Todo

    if _, err := os.Stat(fileName); os.IsNotExist(err) {
        empty := []byte("[]")
        os.WriteFile(fileName, empty, 0644)
    }

    data, err := os.ReadFile(fileName)
    if err != nil {
        return todos, err
    }

    err = json.Unmarshal(data, &todos)
    return todos, err
}
