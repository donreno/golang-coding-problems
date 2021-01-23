package recursionndynamic

import (
	"fmt"
)

func BooleanEvaluation(expression string, result bool) int {
	return booleanEvaluation(expression, result, make(map[string]int))
}

func booleanEvaluation(expression string, result bool, memo map[string]int) int {
	if len(expression) == 0 {
		return 0
	}

	if len(expression) == 1 {
		if stringToBool(expression) == result {
			return 1
		}

		return 0
	}

	expressionKey := fmt.Sprintf("%v%s", result, expression)
	if res, containsKey := memo[expressionKey]; containsKey {
		return res
	}

	ways := 0
	for i := 1; i < len(expression); i += 2 {
		operator := string(expression[i])
		left := string(expression[:i])
		right := string(expression[i+1:])

		leftTrue := booleanEvaluation(left, true, memo)
		leftFalse := booleanEvaluation(left, false, memo)
		rightTrue := booleanEvaluation(right, true, memo)
		rightFalse := booleanEvaluation(right, false, memo)
		total := (leftTrue + leftFalse) * (rightTrue + rightFalse)

		totalTrue := 0

		switch operator {
		case "^":
			totalTrue = leftTrue*rightFalse + leftFalse*rightTrue
		case "&":
			totalTrue = leftTrue * rightTrue
		case "|":
			totalTrue = leftTrue*rightTrue + leftFalse*rightTrue + leftTrue*rightFalse
		}

		if result {
			ways += totalTrue
		} else {
			ways += total - totalTrue
		}
	}

	memo[expressionKey] = ways

	return ways
}

func stringToBool(s string) bool {
	if s == "1" {
		return true
	}

	return false
}
