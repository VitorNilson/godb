package model

import (
	"fmt"
	"strings"
	"github.com/vitornilson1998/godb/utils"
)

type Column struct{
	table Table
	ColumnName string
	ColumnType string
	rows []Row
}




func CreateColumns(columnDefinition string) []Column{
	
	columnDefinition = strings.ReplaceAll(columnDefinition, "(", "")
	columnDefinition = strings.ReplaceAll(columnDefinition, ")", "")
	
	columnsD := strings.Split(columnDefinition,",")
	columns := []Column{}



	for i:=0; i < len(columnsD); i++ {
		
		definitions := strings.Split(strings.Trim(columnsD[i], " "), " ")


		if(ValidateColumn(definitions)){
			
			columns = append(columns,Column{ColumnName:definitions[0], ColumnType:definitions[1]})

		}else{
			fmt.Println("Column definition invalid.")
		}
	}

	return columns
}

func ValidateColumn(columnDefinition []string) bool{
	
	var columnIsValid bool

	if(verifyIfHasNameAndType(columnDefinition)){
		
		if(!utils.VerifyIfContainsInvalidCharacters(columnDefinition[0])){
			
			if(!utils.VerifyIfContainsInvalidCharacters(columnDefinition[1])){
				
				if(verifyIfColumnTypeIsValid(columnDefinition[1])){
					columnIsValid = true
				}
			}
		}
	}

	return columnIsValid
}

func verifyIfColumnTypeIsValid(columnType string) bool{
	allowedTypes := []string{"string", "int", "float"}

	var isValid bool

	for i:=0; i < len(allowedTypes) ; i++ {
		
		if(strings.Contains(columnType, allowedTypes[i])){
			isValid = true
			break;
		}
	}

	return isValid

}

func verifyIfHasNameAndType(definitions []string) bool{
	var hasName bool 
	var hasType bool

	if(len(definitions) > 2){
		panic("There are more fields than acceptable")
	}

	if(definitions[0] != ""){
		hasName = true
	}

	if(definitions[1] != ""){
		hasType = true
	}

	return hasName && hasType
}

func ColumnToString(column []Column){
	for i:=0; i < len(column); i++{
		fmt.Println(column[i].ToString())
	}
}

func (column Column) ToString() string{
	return "ColumnName: " + column.ColumnName+"ColumnType: "+column.ColumnType
}