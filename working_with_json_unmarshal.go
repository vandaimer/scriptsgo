package main

import "fmt"
import "net/http"
import "encoding/json"
import "io/ioutil"

type WObj struct {
	RESULTS []WRESULT `json:"results"`
}

type WRESULT struct {
	TYPES    []string  `json:"types"`
	GEOMETRY WGEOMETRY `json:"geometry"`
}

type WGEOMETRY struct {
	BOUNDS WBOUNDS `json:"bounds"`
}

type WBOUNDS struct {
	NORTHEAST WREGION `json:"northeast"`
	SOUTHWEST WREGION `json:"southwest"`
}

type WREGION struct {
	LAT float64 `json:"lat"`
	LNG float64 `json:"lng"`
}

func get_json() WObj {
	resp, err := http.Get("http://maps.googleapis.com/maps/api/geocode/json?address=92101")

	if err != nil {
		panic("ERRO")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var res WObj
	err = json.Unmarshal(body, &res)

	if err != nil {
		fmt.Println(err)
	}

	for _, each := range res.RESULTS {
		fmt.Println(each.GEOMETRY.BOUNDS.NORTHEAST.LAT)
		break
	}

	return res
}

func main() {
	fmt.Println("hello world")
	get_json()
}
