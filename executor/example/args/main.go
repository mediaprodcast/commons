package main

import (
	"fmt"

	executor "github.com/mediaprodcast/commons/executor"
)

func main() {
	builder := executor.NewArgsBuilder()

	builder.With("command").
		With("-flag1").
		Withf("--option=%s", "value").
		With("arg1", "arg2").
		WithIf(true, "-flag2").
		WithIf(false, "-flag3") // Won't be added

	myMap := map[string]string{
		"key1": "value1",
		"key2": "", // Just the flag
		"key3": "value3",
	}

	builder.WithMap(myMap)

	mySlice := []string{"sliceArg1", "sliceArg2"}
	builder.AddSlice(mySlice)

	fmt.Println("String representation:", builder.String())
	fmt.Println("Slice representation:", builder.Build())
}
