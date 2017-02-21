package uf

import (
	"reflect"
	"strings"
)

func copystruct(aa interface{}, bb interface{}) { //aa,bb为2个struct指针，将后一个结构体相同字段赋值给前一个
	atype := reflect.TypeOf(aa)
	btype := reflect.TypeOf(bb)
	if !(atype.Kind() == reflect.Ptr && btype.Kind() == reflect.Ptr) {
		return
	}

	at := atype.Elem()
	bt := btype.Elem()
	if at.Name()[0] < 65 || at.Name()[0] > 90 || bt.Name()[0] < 65 || bt.Name()[0] > 90 {
		return
	}
	ai := at.NumField()
	bi := bt.NumField()
	for jj := 0; jj < ai; jj++ {
		for kk := 0; kk < bi; kk++ {
			af := at.Field(jj)
			bf := bt.Field(kk)
			if af.Name == bf.Name && af.Type == bf.Type && string(af.Name[0]) == strings.ToUpper(string(af.Name[0])) && string(af.Name[0]) != "_" {
				reflect.ValueOf(aa).Elem().Field(jj).Set(reflect.ValueOf(bb).Elem().Field(kk))

			}
		}
	}

}
