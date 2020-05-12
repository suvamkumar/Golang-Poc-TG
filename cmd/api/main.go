package main

import (
	app "crud_with_TG/Golang-Poc-TG/cmd/api/myapp"
	"crud_with_TG/Golang-Poc-TG/external/services/tigergraph"
	"fmt"
)

func main() {
	connString := tigergraph.TG{ConnectionString: "http://localhost:9000/graph"}
	connString.SyncDataBaseWithGraph("social", "person", "friendship")
	fmt.Println("==========")
	app.StartApplication()

	//connString.GetAllTheVerticesOfAVertex("social", "person")
	//connString.GetVerticeOfAVertexByID("social", "person", "waqqar")
	// type Person struct {
	// 	Name   string `json:"name,omitempty"`
	// 	Gender string `json:"gender,omitempty"`
	// 	Age    int    `json:"age,omitempty"`
	// 	State  string `json:"state,omitempty"`
	// }
	// //	p := Person{Name: "ranatunga", Gender: "female", Age: 32, State: "patna"}

	//connString.UpsertSingleEdge("social", "friendship", "shubham", "Rahul")

}
