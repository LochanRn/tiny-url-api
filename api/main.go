package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get the current path due to %v", err)
	}
	data, err := os.ReadFile(fmt.Sprintf("%s/tiny-url-server-swagger.json", pwd))
	if err != nil {
		log.Fatalf("Failed to read the file due to %v", err)
	}

	//For docs site
	GenerateSpec(data, pwd, "tiny-url-doc-apis.json", SpecGenerateOptions{
		IncludeDeprecated: true,
	})

	//For user sharing (postman)
	GenerateSpec(data, pwd, "tiny-url-apis.json", SpecGenerateOptions{
		IncludeDeprecated: false,
	})
}

type SpecGenerateOptions struct {
	IncludeDeprecated bool
}

func GenerateSpec(sourceSpec []byte, pwd, outputFileName string, options SpecGenerateOptions) {
	msg := make(map[string]interface{}, 0)
	err := json.Unmarshal(sourceSpec, &msg)
	if err != nil {
		log.Fatalf("Failed to decode the file data due to %v", err)
	}

	paths := msg["paths"].(map[string]interface{})
	for k, v := range paths {
		if !IsInternalPath(k) {
			value := v.(map[string]interface{})
			for k1, v1 := range value {
				if k1 != "parameters" {
					value1 := v1.(map[string]interface{})
					tags := value1["tags"].([]interface{})
					if IsPrivateTagExist(tags) {
						delete(value, k1)
					}
					if !options.IncludeDeprecated && IsDeprecated(value1["description"]) {
						delete(value, k1)
					}
				}
			}
			if IsApiBodyEmpty(value) {
				delete(paths, k)
			} else {
				paths[k] = value
			}
		} else {
			delete(paths, k)
		}
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", pwd, outputFileName))
	if err != nil {
		log.Fatalf("Failed to create a file due to %v", err)
	}
	defer f.Close()

	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to json marshal a file due to %v", err)
	}

	_, err = f.Write(jsonData)
	if err != nil {
		log.Fatalf("Failed to write to a file due to %v", err)
	}
}

func IsApiBodyEmpty(value map[string]interface{}) bool {
	if len(value) == 0 {
		return true
	}
	if len(value) == 1 && value["parameters"] != nil {
		return true
	}
	return false
}

func IsPrivateTagExist(params []interface{}) bool {
	tags := []string{"private", "system"}
	set := NewSetWithArray(tags)
	for _, v := range params {
		value := v.(string)
		if set.IsExists(value) {
			return true
		}
	}
	return false
}

func IsInternalPath(path string) bool {
	return strings.Contains(path, "/internal")
}

func IsDeprecated(description interface{}) bool {
	if description == nil {
		return false
	}
	desc := description.(string)
	return strings.Contains(desc, "Deprecated")
}
