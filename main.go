package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/yaml.v3"
)

// RequestBody 結構體用於解析輸入的 JSON
type RequestBody struct {
	YAML       string `json:"yaml,omitempty"`
	Properties string `json:"properties,omitempty"`
}

// yamlToJSON 將 YAML 字串轉換為 JSON 字串
func yamlToJSON(yamlStr string) (string, error) {
	var yamlInterface interface{}
	err := yaml.Unmarshal([]byte(yamlStr), &yamlInterface)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling YAML: %w", err)
	}

	jsonInterface := convertYAMLToJSONCompatible(yamlInterface)
	jsonBytes, err := json.Marshal(jsonInterface)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %w", err)
	}

	return string(jsonBytes), nil
}

func propertiesToJSON(propsStr string) (string, error) {
	propsMap := make(map[string]string)
	lines := strings.Split(propsStr, "\n")
	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue // 忽略空行與注釋行
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // 忽略不符合 key=value 格式的行
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		propsMap[key] = value
	}
	jsonBytes, err := json.Marshal(propsMap)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %w", err)
	}
	return string(jsonBytes), nil
}

func convertYAMLToJSONCompatible(yamlData interface{}) interface{} {
	switch value := yamlData.(type) {
	case map[interface{}]interface{}:
		jsonMap := make(map[string]interface{})
		for k, v := range value {
			jsonMap[fmt.Sprintf("%v", k)] = convertYAMLToJSONCompatible(v)
		}
		return jsonMap
	case []interface{}:
		for i, v := range value {
			value[i] = convertYAMLToJSONCompatible(v)
		}
	}
	return yamlData
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var reqBody RequestBody
	if err := json.Unmarshal(body, &reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var jsonResponse string
	if reqBody.YAML != "" {
		jsonResponse, err = yamlToJSON(reqBody.YAML)
	} else if reqBody.Properties != "" {
		jsonResponse, err = propertiesToJSON(reqBody.Properties)
	} else {
		http.Error(w, "No valid data provided", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}

func main() {
	http.HandleFunc("/tojson", postHandler)
	fmt.Println("Server is running on http://localhost:8080/tojson")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
