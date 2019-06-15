package Resources

import (
	"reflect"
	"encoding/json"
)

func StructToMapReflect(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StructToMapJson(v interface{}) (ret map[string]interface{}, err error) {
	//结构体转json
	jsonStr, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return 
	}
	//再将json转成map
	if err = json.Unmarshal([]byte(jsonStr), &ret); err != nil {
		return 
	} 
	return 
}