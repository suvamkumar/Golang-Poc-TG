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
func (tg TG) UpsertSingleVertex(graphName string, verticesName string, id string, postBodyData interface{}) ([][]byte, *errors.RestErr) {
	reqData := `{ "vertices":{"` + verticesName + `":{"` + id + `":{`
	m := make(map[string]interface{})
	jsonBytes, _ := json.Marshal(postBodyData)
	json.Unmarshal(jsonBytes, &m)
	fmt.Println(m)
	for k, v := range m {
		if reflect.TypeOf(v).String() == "string" {
			reqData = reqData + `"` + k + `": { "value" :"` + v.(string) + `"},`
		}
		if reflect.TypeOf(v).String() == "float64" {
			ageBytes, _ := json.Marshal(v)
			reqData = reqData + `"` + k + `": { "value" :` + string(ageBytes) + "},"
		}
	}
	reqData = reqData[:len(reqData)-1] + "}}}}"
	fmt.Println(reqData)
	response, err := http.Post(tg.ConnectionString+"/"+graphName, "application/json", bytes.NewBuffer([]byte(reqData)))
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	b, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(b))
	fmt.Println("===============")
	fmt.Println(err)
	return nil, nil
}

//UpsertSingleEdge ...
func (tg TG) UpsertSingleEdge(graphName string, edgeName string, fromID string, toID string) ([][]byte, *errors.RestErr) {

	reqData := `{"edges":{"person":{"` + fromID + `":{"` + edgeName + `":{"person":{"` + toID + `":{"connect_day":{"value":"` + date_utils.GetNowDBFormat() + `"}}}}}}}}`

	response, err := http.Post(tg.ConnectionString+"/"+graphName, "application/json", bytes.NewBuffer([]byte(reqData)))
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	b, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(b))
	fmt.Println("===============")
	fmt.Println(err)

	return nil, nil
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
