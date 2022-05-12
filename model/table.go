package model

import (
	"fmt"
)

type Table struct{
	TableName string
	Columns []Column
}


func (table Table) ToString(){
	fmt.Println("TableName:",table.TableName)
	ColumnToString(table.Columns)
}