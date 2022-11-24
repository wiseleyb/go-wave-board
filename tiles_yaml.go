package main

import (
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v3"
)

type TilesYaml struct {
	m map[string]Tile
}

func yamlToTiles(yamlFileName string) Tiles {
	yfile, err := ioutil.ReadFile(yamlFileName)

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[interface{}]interface{})

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	tiles := reflect.ValueOf(data["tiles"])
	tt := Tiles{}
	var nullTile Tile
	for _, k := range tiles.MapKeys() {
		v := tiles.MapIndex(k)
		vs := v.Interface().(string)
		ks := k.Interface().(string)
		//fmt.Println(vs, ks)
		t := newTile(stringToArr(vs), ks)
		if ks == "null" {
			nullTile = t
		} else {
			tt.tiles = append(tt.tiles, t)
		}
	}
	//fmt.Println(tt)
	return newTiles(tt.tiles, nullTile)
}
