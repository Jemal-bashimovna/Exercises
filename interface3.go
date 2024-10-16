package main

import "fmt"

type DataBase interface {
	Save(data string)
	Load()
}

type MySQLDB struct {
	ConnectionString string
}

func (m MySQLDB) Save(data string) {
	fmt.Printf("\"%s\" saved in MySQL => %s\n", data, m.ConnectionString)
}

func (m MySQLDB) Load() {
	fmt.Println("Loading from MySQLDB => ", m.ConnectionString)
}

type PostgresDB struct {
	ConnectionString string
}

func (p PostgresDB) Save(data string) {
	fmt.Printf("\"%s\" saved in PostgreSQL = > %s\n", data, p.ConnectionString)
}
func (p PostgresDB) Load() {
	fmt.Println("Loading from PostgreSQL => ", p.ConnectionString)
}
func TestDataBase(d DataBase) {
	d.Save("fgfj")
	d.Load()
}

func main() {

	var data1 DataBase = MySQLDB{"String"}
	var data2 DataBase = PostgresDB{"String2"}

	TestDataBase(data1)
	TestDataBase(data2)
}
