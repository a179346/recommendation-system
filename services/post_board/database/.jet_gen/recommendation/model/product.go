//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Product struct {
	ProductID   int32 `sql:"primary_key"`
	Title       string
	Price       float64
	Description string
	Category    string
}
