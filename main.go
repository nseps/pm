package main

import (
	"fmt"
	"reflect"
	"./pm"
)

func main() {

	type OO struct {
		a string
		b int
	}

	i := 5
	s := "ksksk"
	o := pm.NewProject("kokos")

	m := map[string]interface{} {

		"int": i,
		"str": s,
		"obj": o,
	}

	fmt.Println(reflect.TypeOf(m["obj"]))

	b := pm.BoolField{}


}
