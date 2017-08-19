package jsondelta

import (
	"github.com/pkg/errors"

	"fmt"
)

func Diff(leftStruct interface{}, rightStruct interface{}) ([]interface{}, error) {
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

		results := []interface{}{}

		// Update items
		for i := 0; i < len(leftVal) || i < len(rightVal); i++ {

			// For any extra items on the right, make 'remove' stanzas
			if i >= len(rightVal) {
				results = append(results, []interface{}{[]interface{}{i}})
				continue
			}

			// For any extra items on the left, make 'add' stanzas
			rightListVal := rightVal[i]
			if i >= len(leftVal) {
				results = append(results, []interface{}{[]interface{}{i}, rightListVal})
				continue
			}

			// Compare items at same position
			leftListVal := leftVal[i];
			subResults, err := Diff(leftListVal, rightListVal)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprint("Error handling array item %d", i))
			}

			// For equal items at same position, do nothing
			if len(subResults) == 0 {
				continue
			}

			// For non-equal items, prefix their diffs stanzas with the array index and add them to the result
			for _, subResult := range subResults {
				subResult, ok := subResult.([]interface{})
				if !ok {
					panic(fmt.Sprintf("Bug: unexpected subresult %v of type %V in Diff between %v and %v",
						subResult, subResult, leftListVal, rightListVal))
				}

				if len(subResult) == 0 {
					panic(fmt.Sprintf("Bug: unexpected empty subresult in Diff between %v and %v",
						leftListVal, rightListVal))
				}
				subResultHead := subResult[0]

				subResultKey, ok := subResultHead.([]interface{})
				if !ok {
					panic(fmt.Sprintf("Bug: unexpected subresult key %v of type %V in Diff between %v and %v",
						subResultHead, subResultHead, leftListVal, rightListVal))
				}

				subResultTail := subResult[1:]

				newSubResult := append(
					[]interface{}{append([]interface{}{i}, subResultKey...)},
					subResultTail...,
				)

				results = append(results, newSubResult)
			}
		}

		return results, nil

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
			subResults, err := Diff(leftMapVal, rightMapVal)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprint("Error handling updated dict key %v", mapKey))
			}

			// Prefix sub-result keys with map key and add so modified sub-results to result
			for _, subResult := range subResults {
				subResult, ok := subResult.([]interface{})
				if !ok {
					panic(fmt.Sprintf("Bug: unexpected subresult %v of type %V in Diff between %v and %v",
						subResult, subResult, leftMapVal, rightMapVal))
				}

				if len(subResult) == 0 {
					panic(fmt.Sprintf("Bug: unexpected empty subresult in Diff between %v and %v",
						leftMapVal, rightMapVal))
				}
				subResultHead := subResult[0]

				subResultKey, ok := subResultHead.([]interface{})
				if !ok {
					panic(fmt.Sprintf("Bug: unexpected subresult key %v of type %V in Diff between %v and %v",
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
