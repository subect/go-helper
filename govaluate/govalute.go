package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func main() {
	audit := "(x0 && !x1) || (x2)"
	exp, err := govaluate.NewEvaluableExpression(audit)
	if err != nil {
		fmt.Println(err.Error())
	}
	params := make(map[string]interface{})
	params["x0"] = true
	params["x1"] = false
	params["x2"] = false
	value, err := exp.Evaluate(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	ok := value.(bool)
	fmt.Println(ok)
}
