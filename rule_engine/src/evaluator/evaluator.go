package evaluator

import (
	"fmt"

	"github.com/Icheka/go-rules-engine/src/ast"
)

type Data map[string]interface{}
type Options struct {
	AllowUndefinedVars bool
}

var options *Options

func EvaluateConditional(conditional *ast.Conditional, identifier interface{}) bool {
	ok, err := EvaluateOperator(identifier, conditional.Value, conditional.Operator)
	if err != nil {
		panic(err)
	}
	return ok
}

func GetFactValue(condition *ast.Conditional, data Data) interface{} {
	value := data[condition.Fact]

	if value == nil {
		if options.AllowUndefinedVars {
			return false
		}
		// panic(fmt.Sprintf("value for identifier %s not found", condition.Fact))
		return nil
	}

	return value
}

func EvaluateAllCondition(conditions *[]ast.Conditional, data Data) bool {
	isFalse := false
	for _, condition := range *conditions {
		value := GetFactValue(&condition, data)
		if !EvaluateConditional(&condition, value) {
			isFalse = true
		}

		if isFalse {
			return false
		}
	}

	return true
}

func EvaluateAnyCondition(conditions *[]ast.Conditional, data Data) bool {
	for _, condition := range *conditions {
		value := GetFactValue(&condition, data)
		if EvaluateConditional(&condition, value) {
			return true
		}
	}

	return false
}

func EvaluateCondition(condition *[]ast.Conditional, kind string, data Data) bool {
	switch kind {
	case "all":
		return EvaluateAllCondition(condition, data)
	case "any":
		return EvaluateAnyCondition(condition, data)
	default:
		panic(fmt.Sprintf("condition type %s is invalid", kind))
	}
}

func EvaluateRule(rule *ast.Rule, data Data, opts *Options) bool {
	options = opts
	any, all := false, false
	if len(rule.Condition.Any) == 0 {
		any = true
	} else {
		any = EvaluateCondition(&rule.Condition.Any, "any", data)
	}
	if len(rule.Condition.All) == 0 {
		all = true
	} else {
		all = EvaluateCondition(&rule.Condition.All, "all", data)
	}

	fmt.Println("any", any)
	fmt.Println("all", all)
	return any && all

}
