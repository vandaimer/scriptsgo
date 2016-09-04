package main

import "fmt"
import "encoding/json"

type AUTOMOVEL struct {
	AVIAO AVIAO
	IATE  IATE
}

type AVIAO struct {
	RODAS       int
	PASSAGEIROS int
}

type IATE struct {
	PASSAGEIROS int
}

func get_json() AUTOMOVEL {

	x := AUTOMOVEL{AVIAO{RODAS: 2, PASSAGEIROS: 10}, IATE{PASSAGEIROS: 20}}

	_, err := json.Marshal(x)

	if err != nil {
		panic(err)
	}

	return x
}

func main() {
	fmt.Println("hello world")
	get_json()
}
