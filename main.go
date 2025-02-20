package main

func main() {
	todos := Todos{}

	Storage := NewStorage[Todos]("todos.json")
	Storage.Load(&todos)
	CommandFlags := NewCmdFlags()
	CommandFlags.Execute(&todos)
	Storage.Save(todos)
}
