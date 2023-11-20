package utility

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

func MapReqBodyToJson(data interface{}, r *http.Request) (interface{}, error) {
	//type assertion here
	targetType, ok := data.(interface{})
	if !ok {
		return nil, errors.New("Invalid Type")
	}
	err := json.NewDecoder(r.Body).Decode(targetType)
	if err != nil {
		return err, nil
	}
	return targetType, nil
}

func copyFields(src, dest interface{}) {
	srcValue := reflect.ValueOf(src).Elem()
	destValue := reflect.ValueOf(dest).Elem()
	fmt.Println(srcValue, destValue)
	for i := 0; i < srcValue.NumField(); i++ {
		fieldName := srcValue.Type().Field(i).Name
		destField := destValue.FieldByName(fieldName)

		if destField.IsValid() && destField.CanSet() {
			destField.Set(srcValue.Field(i))
		}
	}
}
