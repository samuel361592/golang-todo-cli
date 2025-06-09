package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "todo-cli/todo"
)

func main() {
    todos, err := todo.LoadTodos()
    if err != nil {
        fmt.Println("讀取失敗:", err)
        return
    }

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("\n=== Todo CLI ===")
        fmt.Println("1. 查看待辦事項")
        fmt.Println("2. 新增待辦事項")
        fmt.Println("3. 標記完成")
        fmt.Println("4. 刪除事項")
        fmt.Println("5. 離開程式")
        fmt.Print("請輸入選項：")

        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        switch input {
        case "1":
            showTodos(todos)
        case "2":
            fmt.Print("輸入新的事項：")
            title, _ := reader.ReadString('\n')
            title = strings.TrimSpace(title)
            newTodo := todo.NewTodo(title)
            todos = append(todos, newTodo)
            todo.SaveTodos(todos)
        case "3":
            showTodos(todos)
            fmt.Print("輸入要標記完成的編號：")
            numStr, _ := reader.ReadString('\n')
            numStr = strings.TrimSpace(numStr)
            index, _ := strconv.Atoi(numStr)
            if index >= 1 && index <= len(todos) {
                todos[index-1].Completed = true
                todo.SaveTodos(todos)
                fmt.Println("已標記完成")
            } else {
                fmt.Println("無效的編號")
            }
        case "4":
            showTodos(todos)
            fmt.Print("輸入要刪除的編號：")
            numStr, _ := reader.ReadString('\n')
            numStr = strings.TrimSpace(numStr)
            index, _ := strconv.Atoi(numStr)
            if index >= 1 && index <= len(todos) {
                todos = append(todos[:index-1], todos[index:]...)
                todo.SaveTodos(todos)
                fmt.Println("已刪除")
            } else {
                fmt.Println("無效的編號")
            }
        case "5":
            fmt.Println("Bye!")
            return
        default:
            fmt.Println("請輸入 1~5 選項")
        }
    }
}

func showTodos(todos []todo.Todo) {
    fmt.Println("\n待辦清單：")
    if len(todos) == 0 {
        fmt.Println("（目前沒有待辦事項）")
        return
    }
    for i, t := range todos {
        status := " "
        if t.Completed {
            status = "✔"
        }
        fmt.Printf("%d. [%s] %s\n", i+1, status, t.Title)
    }
}
