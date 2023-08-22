package db_action

type Database struct {
	name  string
	table []*Table
}

func (d *Database) Table() []*Table {
	return d.table
}

func (d *Database) AddTable(name string) {
	t := NewTable(name)
	d.table = append(d.table, t)
}

func NewDatabase(name string) *Database {
	return &Database{name: name, table: []*Table{}}
}

func (d *Database) Name() string {
	return d.name
}

func (d *Database) SetName(name string) {
	d.name = name
}
