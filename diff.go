package jsondelta

import (
	"github.com/pkg/errors"

	"fmt"
)

func diff(leftStruct interface{}, rightStruct interface{}) (interface{}, error) {
	switch leftVal := leftStruct.(type) {

	case bool:
		rightVal, ok := rightStruct.(bool)
		if !ok || leftVal != rightVal {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}

	case float64:
		rightVal, ok := rightStruct.(float64)
		if !ok || leftVal != rightVal {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}

	case int:
		rightVal, ok := rightStruct.(int)
		if !ok || leftVal != rightVal {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}

	case nil:
		if leftVal != rightStruct {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}

	case string:
		rightVal, ok := rightStruct.(string)
		if !ok || leftVal != rightVal {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}

	case []interface{}:
		// TODO

	case map[string]interface{}:
		// TODO

	default:
		return nil, errors.New(fmt.Sprintf("Bad type on left side: %T", leftStruct))
	}

	return []interface{}{}, nil
}
