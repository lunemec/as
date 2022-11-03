package as_test

import (
	"fmt"

	"github.com/lunemec/as"
)

func Example() {
	for _, n := range []int{127, 128} {
		num, err := as.Int8(n)
		if err != nil {
			fmt.Printf("Input invalid: %d, err: %s\n", num, err)
		} else {
			fmt.Printf("Input valid: %d\n", num)
		}
	}
	// Output: Input valid: 127
	// Input invalid: -128, err: 128 (int) overflows int8
}

func ExampleT() {
	for _, n := range []int{127, 128} {
		num, err := as.T[int8](n)
		if err != nil {
			fmt.Printf("Input invalid: %d, err: %s\n", num, err)
		} else {
			fmt.Printf("Input valid: %d\n", num)
		}
	}
	// Output: Input valid: 127
	// Input invalid: -128, err: 128 (int) overflows int8
}

func ExampleSliceT() {
	out, err := as.SliceT[int, int8]([]int{127, 128})
	fmt.Printf("Output: %+v, error: %+v\n", out, err)
	// Output: Output: [127 0], error: 1 error occurred:
	// 	* at index [1]: 128 (int) overflows int8
}

func ExampleSliceTUnchecked() {
	out := as.SliceTUnchecked[int, int8]([]int{127, 128})
	fmt.Printf("Output: %+v\n", out)
	// Output: Output: [127 -128]
}
