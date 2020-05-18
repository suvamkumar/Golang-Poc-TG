package tigergraph

import (
	"bytes"

	"crud_with_TG/Golang-Poc-TG/internal/utils/date_utils"
	"crud_with_TG/Golang-Poc-TG/internal/utils/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

//TG ...
type TG struct {
	//ConnectionString ...
	ConnectionString string
}

//GetAllTheVerticesOfAVertex ...
func (tg TG) GetAllTheVerticesOfAVertex(graphName string, vertexName string) ([][]byte, *errors.RestErr) {
	response, err := http.Get(tg.ConnectionString + "/" + graphName + "/vertices/" + vertexName)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	result := getJSONResult(response)
	return result, nil

}

//GetVerticeOfAVertexByID ...
func (tg TG) GetVerticeOfAVertexByID(graphName string, vertexName string, id string) ([][]byte, *errors.RestErr) {
	response, err := http.Get(tg.ConnectionString + "/" + graphName + "/vertices/" + vertexName + "/" + id)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	result := getJSONResult(response)
	return result, nil
}

//UpsertSingleVertex ...
func (tg TG) UpsertSingleVertex(graphName string, verticesName string, id string, postBodyData interface{}) map[string]interface{} {
	reqData := `{ "vertices":{"` + verticesName + `":{`
	reqData = reqData + createJSONBodyForVertices(id, postBodyData) + "}}}"
	response, err := http.Post(tg.ConnectionString+"/"+graphName, "application/json", bytes.NewBuffer([]byte(reqData)))
	if err != nil {
		return map[string]interface{}{"message": "could not make the post request to tiger graph"}
	}
	b, err := ioutil.ReadAll(response.Body)
	return createResponse(b)
}

//UpsertManyEdges ...
func (tg TG) UpsertManyEdges(graphName string, edgeName string, postData interface{}) map[string]interface{} {
	// reqData := `{"edges":{"person":{"` + fromID + `":{"` + edgeName + `":{"person":{"` + toID + `":{"connect_day":{"value":"` + date_utils.GetNowDBFormat() + `"}}}}}}}}`
	reqData := `{"edges":{"person":{`
	d := reflect.ValueOf(postData)
	for i := 0; i < d.Len(); i++ {
		m := make(map[string]interface{})
		bytesData, _ := json.Marshal(d.Index(i).Interface())
		json.Unmarshal(bytesData, &m)
		if i < d.Len()-1 {
			reqData = reqData + createJSONBodyForEdges(m["from"].(string), m["to"].(string), edgeName, "") + ","

		} else {
			reqData = reqData + createJSONBodyForEdges(m["from"].(string), m["to"].(string), edgeName, "")

		}
	}
	reqData = reqData + "}}}"
	// reqData := `{"edges":{"person":{` + createJSONBodyForEdges(fromID, toID, edgeName, "") + `}}}`
	response, err := http.Post(tg.ConnectionString+"/"+graphName, "application/json", bytes.NewBuffer([]byte(reqData)))
	if err != nil {
		return map[string]interface{}{"message": "could not make the post request to tiger graph"}
	}
	b, err := ioutil.ReadAll(response.Body)
	return createResponse(b)
}

//UpsertMultipleVertex ...
func (tg TG) UpsertMultipleVertex(graphName string, verticesName string, postBodyData interface{}) map[string]interface{} {
	reqData := `{ "vertices":{"` + verticesName + `":{`
	dataValue := reflect.ValueOf(postBodyData)
	for i := 0; i < dataValue.Len(); i++ {
		m := make(map[string]interface{})
		dataBytes, _ := json.Marshal(dataValue.Index(i).Interface())
		json.Unmarshal(dataBytes, &m)
		var id string
		for k, v := range m {
			fmt.Println(k, v)
			if strings.EqualFold(k, "name") {
				id = v.(string)
			}
		}
		if i < dataValue.Len()-1 {
			reqData = reqData + createJSONBodyForVertices(id, m) + ","
		} else {
			reqData = reqData + createJSONBodyForVertices(id, m)
		}
	}
	reqData = reqData + "}}}"
	response, err := http.Post(tg.ConnectionString+"/"+graphName, "application/json", bytes.NewBuffer([]byte(reqData)))
	if err != nil {
		return map[string]interface{}{"message": "could not make the post request to tiger graph"}

	}
	b, err := ioutil.ReadAll(response.Body)
	return createResponse(b)
}

//UpsertSingleEdge ...
func (tg TG) UpsertSingleEdge(graphName string, edgeName string, fromID string, toID string) map[string]interface{} {
	// reqData := `{"edges":{"person":{"` + fromID + `":{"` + edgeName + `":{"person":{"` + toID + `":{"connect_day":{"value":"` + date_utils.GetNowDBFormat() + `"}}}}}}}}`
	reqData := `{"edges":{"person":{` + createJSONBodyForEdges(fromID, toID, edgeName, "") + `}}}`
	response, err := http.Post(tg.ConnectionString+"/"+graphName, "application/json", bytes.NewBuffer([]byte(reqData)))
	if err != nil {
		return map[string]interface{}{"message": "could not make the post request to tiger graph"}
	}
	b, err := ioutil.ReadAll(response.Body)
	return createResponse(b)
}

//SyncDataBaseWithGraph ...
func (tg TG) SyncDataBaseWithGraph(graphName string, verticesName string, edgeName string, person interface{}, frienship interface{}) (bool, *errors.RestErr) {

	reqData := `{ "vertices":{"` + verticesName + `":{`
	dataValue := reflect.ValueOf(person)
	var jsonBody = ""
	for i := 0; i < dataValue.Len(); i++ {
		m := make(map[string]interface{})
		dataBytes, _ := json.Marshal(dataValue.Index(i))
		json.Unmarshal(dataBytes, &m)
		var id string
		for k, v := range m {
			fmt.Println(k, v)
			if strings.EqualFold(k, "name") {
				id = v.(string)
			}
		}
		if i < dataValue.Len()-1 {
			reqData = jsonBody + createJSONBodyForVertices(id, person) + ","
		} else {
			reqData = jsonBody + createJSONBodyForVertices(id, person)
		}
	}
	reqData = reqData + jsonBody + "}},"

	jsonBody = `"edges":{"person":{`
	d := reflect.ValueOf(frienship)
	for i := 0; i < d.Len(); i++ {
		m := make(map[string]interface{})
		bytesData, _ := json.Marshal(d.Index(i))
		json.Unmarshal(bytesData, &m)
		if i <= d.Len()-1 {
			reqData = jsonBody + createJSONBodyForEdges(m["from"].(string), m["to"].(string), edgeName, "") + ","

		} else {
			reqData = jsonBody + createJSONBodyForEdges(m["from"].(string), m["to"].(string), edgeName, "")

		}
	}
	reqData = reqData + jsonBody + "}}}"
	response, err := http.Post(tg.ConnectionString+"/"+graphName, "application/json", bytes.NewBuffer([]byte(reqData)))
	if err != nil {
		return false, errors.NewInternalServerError(err.Error())
	}
	b, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(b), err)
	return true, nil
}

func getJSONResult(response *http.Response) [][]byte {
	m := make(map[string]interface{})
	r, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(r, &m)
	d := reflect.ValueOf(m["results"])
	s := make([]interface{}, d.Len())
	var result [][]byte
	for key := range m {
		if key == "results" {
			for i := 0; i < reflect.ValueOf(m["results"]).Len(); i++ {
				s[i] = d.Index(i).Interface()
				iter := reflect.ValueOf(s[i]).MapRange()
				for iter.Next() {
					k := iter.Key()
					v := iter.Value()
					if reflect.Value(k).Interface() == "attributes" {
						jsonBytes, _ := json.Marshal(reflect.Value(v).Interface())
						result = append(result, jsonBytes)
					}
				}
			}
		}
	}
	return result
}

func createJSONBodyForVertices(id string, postBodyData interface{}) string {
	m := make(map[string]interface{})
	jsonBytes, _ := json.Marshal(postBodyData)
	json.Unmarshal(jsonBytes, &m)
	reqData := `"` + id + `":{`
	for k, v := range m {
		if reflect.TypeOf(v).String() == "string" && k != "_id" {
			reqData = reqData + `"` + k + `": { "value" :"` + v.(string) + `"},`
		}
		if reflect.TypeOf(v).String() == "float64" && k != "_id" {
			ageBytes, _ := json.Marshal(v)
			reqData = reqData + `"` + k + `": { "value" :` + string(ageBytes) + "},"
		}
	}
	return reqData[:len(reqData)-1] + "}"
}

func createJSONBodyForEdges(fromID string, toID string, edgeName string, connect_day string) string {
	if connect_day == "" {
		return `"` + fromID + `":{"` + edgeName + `":{"person":{"` + toID + `":{"connect_day":{"value":"` + date_utils.GetNowDBFormat() + `"}}}}}`
	}
	return `"` + fromID + `":{"` + edgeName + `":{"person":{"` + toID + `":{"connect_day":{"value":"` + connect_day + `"}}}}}`

}

func createResponse(response []byte) map[string]interface{} {
	m := make(map[string]interface{})
	json.Unmarshal(response, &m)
	return m

}

// func checkErrorMessage(response []byte) (message string, results string, err bool) {
// 	m := make(map[string]interface{})
// 	json.Unmarshal(response, &m)
// 	fmt.Println(m)
// 	iter := reflect.ValueOf(m).MapRange()
// 	for iter.Next() {
// 		if iter.Key().Interface() == "error" {
// 			err = iter.Value().Interface().(bool)
// 			fmt.Println("===========")
// 			fmt.Println(err)
// 			fmt.Println("===========")
// 		}
// 		// k := iter.Key()
// 		// v := iter.Value()
// 	}
// 	return "", "", false
// 	// {"version":{"edition":"developer","api":"v2","schema":0},
// 	// "error":false,"message":"","results":[{"accepted_vertices":1,"accepted_edges":0}],
// 	// "code":"REST-0001"}

// 	{
// 		"error": false,
// 		"message": "",
// 		"version": {
// 		  "schema": 0,
// 		  "edition": "developer",
// 		  "api": "v2"
// 		},
// 		"results": {
// 		  "deleted_vertices": 0,
// 		  "v_type": "person"
// 		}
// 	  }
//}
