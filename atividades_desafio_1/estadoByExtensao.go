package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

type Estado struct {
		Name 		string 	 `json:"name"`
		Km2 		string  `json:"km2"`
}

type EstadosBrasileiros struct {
	Brasil []Estado	
}

type byExtensao []Estado

func (s byExtensao) Len() int 		{	return len(s)				}
func (s byExtensao) Swap(i, j int)  {   s[i], s[j] = s[j], s[i]		}
func (s byExtensao) Less(i, j int) bool {
	k1,_ := strconv.ParseFloat(s[i].Km2, 32)
	k2,_ := strconv.ParseFloat(s[j].Km2, 32)
    return k1 > k2
}

func main() {

	res, _ := http.Get("https://raw.githubusercontent.com/JoaoPauloA/estadosBrasileirosJson/master/EstadosKm2.json")

	temp, _ := ioutil.ReadAll(res.Body)

	var estados EstadosBrasileiros

	err := json.Unmarshal([]byte(temp), &estados)

	if err != nil {
		fmt.Println("There was an error:", err)
	}

	fmt.Println(estados.Brasil,"\n")

	sort.Sort(byExtensao(estados.Brasil))

	fmt.Println(estados.Brasil)
}

