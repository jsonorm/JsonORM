package jsonorm

import (
	"encoding/json"
	"log"
)

type FullORM struct {
	defaultORM
	With []with
}

type with struct {
	defaultORM
	Relationship string
	Condition    string
}

type defaultORM struct {
	Table   string
	Where   []where
	OrWhere []where
	Select  string
	Offset  int
	Limit   int
	OrderBy string
	Handle  string
}
type where []interface{}

func Parse(jsonOrm string) error {
	var fullORM FullORM
	err := json.Unmarshal([]byte(jsonOrm), &fullORM)
	if err != nil {
		return err
	}
	log.Println(fullORM)

	return nil
}
