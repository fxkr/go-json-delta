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
			return newDiff(newUpdateStanza(rightStruct)), nil
		}
		return newDiff(), nil

	case float64:
		rightVal, ok := rightStruct.(float64)
		if !ok || leftVal != rightVal {
			return newDiff(newUpdateStanza(rightStruct)), nil
		}
		return newEmptyDiff(), nil

	case int:
		rightVal, ok := rightStruct.(int)
		if !ok || leftVal != rightVal {
			return newDiff(newUpdateStanza(rightStruct)), nil
		}
		return newEmptyDiff(), nil

	case nil:
		if leftVal != rightStruct {
			return newDiff(newUpdateStanza(rightStruct)), nil
		}
		return newEmptyDiff(), nil

	case string:
		rightVal, ok := rightStruct.(string)
		if !ok || leftVal != rightVal {
			return newDiff(newUpdateStanza(rightStruct)), nil
		}
		return newEmptyDiff(), nil

	case []interface{}:
		rightVal, ok := rightStruct.([]interface{})
		if !ok {
			return newDiff(newUpdateStanza(rightStruct)), nil
		}

		return diffArray(leftVal, rightVal)

	case map[string]interface{}:
		rightVal, ok := rightStruct.(map[string]interface{})
		if !ok {
			return newDiff(newUpdateStanza(rightStruct)), nil
		}

		return diffObject(leftVal, rightVal)

	default:
		return nil, errors.New(fmt.Sprintf("Bad type on left side: %T", leftStruct))
	}
}

func diffArray(leftVal []interface{}, rightVal []interface{}) ([]interface{}, error) {
	results := newEmptyDiff()

	// Update items
	for i := 0; i < len(leftVal) || i < len(rightVal); i++ {

		// For any extra items on the left, make 'remove' stanzas
		if i >= len(rightVal) {
			results = append(results, newArrayRemoveStanza(i))
			continue
		}

		// For any extra items on the right, make 'add' stanzas
		rightListVal := rightVal[i]
		if i >= len(leftVal) {
			results = append(results, newArrayAddStanza(i, rightListVal))
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
		results = appendWithPrefixIndex(results, i, subResults)
	}

	return results, nil
}

func diffObject(leftVal map[string]interface{}, rightVal map[string]interface{}) ([]interface{}, error) {
	results := newEmptyDiff()

	// Find removed keys
	for mapKey, _ := range leftVal {
		_, found := rightVal[mapKey]
		if !found {
			results = append(results, newObjectRemoveStanza(mapKey))
			continue
		}
	}

	// Find added keys
	for mapKey, rightMapVal := range rightVal {
		_, found := leftVal[mapKey]
		if !found {
			results = append(results, newObjectUpdateStanza(mapKey, rightMapVal))
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
		results = appendWithPrefixKey(results, mapKey, subResults)
	}

	return results, nil
}

func appendWithPrefixIndex(results []interface{}, index int, subResults []interface{}) []interface{} {
	return appendWithPrefix(results, index, subResults);
}

func appendWithPrefixKey(results []interface{}, mapKey string, subResults []interface{}) []interface{} {
	return appendWithPrefix(results, mapKey, subResults);
}

func appendWithPrefix(results []interface{}, prefix interface{}, subResults []interface{}) []interface{} {
	for _, subResult := range subResults {
		subResult, ok := subResult.([]interface{})
		if !ok {
			panic(fmt.Sprintf("Bug: unexpected subresult %v of type %V", subResult, subResult))
		}

		newSubResult := prependPrefixToStanza(prefix, subResult)

		results = append(results, newSubResult)
	}
	return results
}

func prependPrefixToStanza(prefix interface{}, subResult []interface{}) []interface{} {
	if len(subResult) == 0 {
		panic("Bug: unexpected empty subresult")
	}

	subResultHead := subResult[0]
	subResultKey, ok := subResultHead.([]interface{})
	if !ok {
		panic(fmt.Sprintf("Bug: unexpected subresult key %v of type %V", subResultHead, subResultHead))
	}

	subResultTail := subResult[1:]

	newSubResult := append(
		[]interface{}{append([]interface{}{prefix}, subResultKey...)},
		subResultTail...,
	)

	return newSubResult
}

func newDiff(stanzas ...interface{}) []interface{} {
	if stanzas == nil {
		stanzas = []interface{}{}
	}
	return stanzas
}

func newEmptyDiff() []interface{} {
	return newDiff()
}

func newAddStanza(newValue interface{}) interface{} {
	return []interface{}{[]interface{}{}, newValue}
}

func newArrayAddStanza(index int, newValue interface{}) []interface{} {
	return []interface{}{[]interface{}{index}, newValue}
}

func newObjectAddStanza(key string, newValue interface{}) []interface{} {
	return []interface{}{[]interface{}{key}, newValue}
}

func newUpdateStanza(newValue interface{}) interface{} {
	return newAddStanza(newValue)
}

func newArrayUpdateStanza(index int, newValue interface{}) []interface{} {
	return newArrayAddStanza(index, newValue)
}

func newObjectUpdateStanza(key string, newValue interface{}) []interface{} {
	return newObjectAddStanza(key, newValue)
}

func newRemoveStanza() interface{} {
	return []interface{}{[]interface{}{}};
}

func newArrayRemoveStanza(index int) interface{} {
	return []interface{}{[]interface{}{index}};
}

func newObjectRemoveStanza(key string) interface{} {
	return []interface{}{[]interface{}{key}};
}

