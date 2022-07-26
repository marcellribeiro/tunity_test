package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/marcellribeiro/tunity_test/pkg"
	"strconv"
)

func CollatzBash(argument string) {
	floatNumber, err := strconv.ParseFloat(argument, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	if floatNumber != 0 {
		result, err := pkg.GetCollatz(floatNumber)
		if err != nil {
			fmt.Println(err)
		}
		marshaledJson, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(marshaledJson))
	}
}
