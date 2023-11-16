package utility

import (
	"encoding/json"
	"errors"
	"net/http"
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
