package main

import (
	"net/http"
	//"io/ioutil"
	//"log"
	"encoding/json"
	"fmt"
	)

func main(){
   //resp,_ := http.Get("https://jsonplaceholder.typicode.com/todos/1")
   resp,_ := http.Get("https://5aa6be85af352a001477f59c.mockapi.io/api/v1/blogs");
   var generic map[string] interface{}

   json.NewDecoder(resp.Body).Decode(&generic)

   //body,_ := ioutil.ReadAll(resp.Body)

   fmt.Println(generic)

}
