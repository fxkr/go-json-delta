package jsondelta

import (
	"github.com/pkg/errors"

	"fmt"
)

func diff(leftStruct interface{}, rightStruct interface{}) ([]interface{}, error) {
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

		results := []interface{}{}

		// Find removed keys
		for mapKey, _ := range leftVal {
			_, found := rightVal[mapKey]
			if !found {
				results = append(results, []interface{}{[]interface{}{mapKey}})
				continue
			}
		}

		// Find added keys
		for mapKey, rightMapVal := range rightVal {
			_, found := leftVal[mapKey]
			if !found {
				results = append(results, []interface{}{[]interface{}{mapKey}, rightMapVal})
				continue
			}
		}

		// Find updated keys
		for mapKey, leftMapVal := range leftVal {
			rightMapVal, found := rightVal[mapKey]
			if !found {
				continue // Added keys already handled above
			}

			// Recursive call
			subResults, err := diff(leftMapVal, rightMapVal)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprint("Error handling updated dict key %v", mapKey))
			}

			// Prefix sub-result keys with map key and add so modified sub-results to result
			for _, subResult := range subResults {
				subResult, ok := subResult.([]interface{})
				if !ok {
					panic(fmt.Sprintf("Bug: unexpected subresult %v of type %V in diff between %v and %v",
						subResult, subResult, leftMapVal, rightMapVal))
				}

				if len(subResult) == 0 {
					panic(fmt.Sprintf("Bug: unexpected empty subresult in diff between %v and %v",
						leftMapVal, rightMapVal))
				}
				subResultHead := subResult[0]

				subResultKey, ok := subResultHead.([]interface{})
				if !ok {
					panic(fmt.Sprintf("Bug: unexpected subresult key %v of type %V in diff between %v and %v",
						subResultHead, subResultHead, leftMapVal, rightMapVal))
				}

				subResultTail := subResult[1:]

				newSubResult := append(
					[]interface{}{append([]interface{}{mapKey}, subResultKey...)},
					subResultTail...,
				)

				results = append(results, newSubResult)
			}
		}

		return results, nil

	default:
		return nil, errors.New(fmt.Sprintf("Bad type on left side: %T", leftStruct))
	}
}
