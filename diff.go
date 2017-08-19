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
		return []interface{}{}, nil

	case float64:
		rightVal, ok := rightStruct.(float64)
		if !ok || leftVal != rightVal {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}
		return []interface{}{}, nil

	case int:
		rightVal, ok := rightStruct.(int)
		if !ok || leftVal != rightVal {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}
		return []interface{}{}, nil

	case nil:
		if leftVal != rightStruct {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}
		return []interface{}{}, nil

	case string:
		rightVal, ok := rightStruct.(string)
		if !ok || leftVal != rightVal {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}
		return []interface{}{}, nil

	case []interface{}:
		rightVal, ok := rightStruct.([]interface{})
		if !ok {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}

		// TODO
		_ = rightVal

		return []interface{}{}, nil

	case map[string]interface{}:
		rightVal, ok := rightStruct.(map[string]interface{})
		if !ok {
			return []interface{}{[]interface{}{[]interface{}{}, rightStruct}}, nil
		}

		// TODO
		_ = rightVal

		return []interface{}{}, nil

	default:
		return nil, errors.New(fmt.Sprintf("Bad type on left side: %T", leftStruct))
	}
}
