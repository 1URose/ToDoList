package models

type Table struct {
	Name        string
	ListOfTasks []Task
}

func (t *Table) NewTable(name string) *Table {
	table := Table{
		Name:        name,
		ListOfTasks: make([]Task, 0),
	}
	return &table
}

func (t *Table) AddTask(task Task) {
	t.ListOfTasks = append(t.ListOfTasks, task)
}
