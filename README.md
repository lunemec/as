# As

[![CI](https://github.com/lunemec/as/actions/workflows/ci.yml/badge.svg)](https://github.com/lunemec/as/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lunemec/as)](https://goreportcard.com/report/github.com/lunemec/as)
[![Maintainability](https://api.codeclimate.com/v1/badges/d0b5da039ba6172a1b3b/maintainability)](https://codeclimate.com/github/lunemec/as/maintainability)

`as` converts integer types and reports values that do not fit in the destination type. It requires Go 1.24 or newer.

## Installation

```sh
go get github.com/lunemec/as
```

## Usage

The destination type is explicit and the source type is inferred:

```go
for _, n := range []int{127, 128} {
	num, err := as.T[int8](n)
	if err != nil {
		fmt.Printf("Input invalid: %d, err: %s\n", num, err)
		continue
	}
	fmt.Printf("Input valid: %d\n", num)
}
// Output: Input valid: 127
// Input invalid: -128, err: 128 (int) overflows int8
```

Named helpers are generic wrappers around `T`:

```go
num, err := as.Int8(127)
```

All built-in integer types, `uintptr`, and user-defined types with an integer underlying type are supported.

## Slices

```go
out, err := as.SliceT[int, int8]([]int{127, 128})
fmt.Printf("Output: %+v, error: %+v\n", out, err)
// Output: Output: [127 0], error: 1 error occurred:
// 	* at index [1]: 128 (int) overflows int8
```

`SliceTUnchecked` also supports floating-point types and performs ordinary unchecked Go conversions.

Use the `Into` variants to reuse caller-owned storage without allocating on successful conversions:

```go
dst := make([]int8, 2)
err := as.SliceTInto(dst, []int{127, 128})
// dst is [127 0]; err identifies the overflow at index 1.

err = as.SliceTUncheckedInto(dst, []int{127, 128})
// dst is [127 -128]; err is nil.
```

`SliceTInto` and `SliceTUncheckedInto` return `io.ErrShortBuffer` before changing `dst` when it is too short. They fill only `dst[:len(src)]`, leaving any suffix unchanged. Checked conversion zeroes overflowing positions and returns every overflow in source-index order; building those errors may allocate. Source and destination slices must not overlap.

## Migrating dynamic inputs

Checked conversion functions accept integer values directly. Values stored in `any` must be type-asserted or handled with a type switch first, and pointers must be dereferenced:

```go
value := any(int64(42))
n, err := as.T[int8](value.(int64))

pointer := new(int64)
*pointer = 42
n, err = as.T[int8](*pointer)
```

Invalid source types are compile-time errors. `InvalidTypeError` and automatic pointer dereferencing are no longer part of the API.

## Performance and architecture support

Successful scalar and `Into` conversions perform no allocations. The allocating slice helpers allocate one result slice. Overflow checking still has a cost, so use the included benchmarks when evaluating a hot path.

The implementation has no architecture-specific files and follows the target's native widths for `int`, `uint`, and `uintptr`.

There are several [Go proposals](https://github.com/golang/go/issues/30613) for checked conversions. Libraries such as [overflow](https://github.com/JohnCGriffin/overflow) provide checked arithmetic operations in addition to conversion.
