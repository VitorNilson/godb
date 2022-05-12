package persistence

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/vitornilson1998/godb/model"
	"github.com/vitornilson1998/godb/utils"
)

const (
	tableExtension = ".godbtb"
	tablePath = "/godb-tables/"
)


var netWork bytes.Buffer

var encoder = gob.NewEncoder(&netWork)
var decoder = gob.NewDecoder(&netWork)


type TableOperations struct{}


func (t TableOperations) CreateTable(sql string, sqlCreateTable []string){

	tableName := sqlCreateTable[2]

	columnDefinitions := sql[strings.Index(sql, "("):]
	table := model.Table{TableName:tableName, Columns: model.CreateColumns(columnDefinitions) }

				
	if(utils.VerifyIfContainsInvalidCharacters(table.TableName)){
		panic("Invalid table name.")
	}
	
	// Create directory to store tables data
	if _, err := os.Stat(tablePath); errors.Is(err, os.ErrNotExist){
		err := os.Mkdir(tablePath, os.ModePerm)
		if err != nil{
			log.Fatal(err)
		}
	}

	// Create file and store
	f, err := os.Create(tablePath + table.TableName + tableExtension)
	
	if err != nil {
		log.Fatal(err)
	}
	error := encoder.Encode(table)

	if(error != nil){
		log.Fatalln("Encode Error", error)
	}

	defer f.Close()

	f.Write(netWork.Bytes())
	
}

func (t TableOperations) DropTable(tableName string){
	os.Remove(tablePath + tableName + tableExtension)
}