package main

import "crud_with_TG/Golang-Poc-TG/external/services/tigergraph"

func main() {
	//	app.StartApplication()
	connString := tigergraph.TG{ConnectionString: "http://localhost:9000/graph"}
	//connString.GetAllTheVerticesOfAVertex("social", "person")
	//connString.GetVerticeOfAVertexByID("social", "person", "waqqar")
	type Person struct {
		Name   string `json:"name,omitempty"`
		Gender string `json:"gender,omitempty"`
		Age    int    `json:"age,omitempty"`
		State  string `json:"state,omitempty"`
	}
	//	p := Person{Name: "ranatunga", Gender: "female", Age: 32, State: "patna"}

	//	connString.UpsertSingleVertex("social", "person", "ranatunga", p)
	connString.UpsertSingleEdge("social", "friendship", "shubham", "Rahul")

}
