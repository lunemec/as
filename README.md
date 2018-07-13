# As
`As` is a library to convert numeric types with overflow check in Go.


## Instalation

`go get -u github.com/lunemec/as`

## Example
```go
num, err := as.Int32(myBignumber)
// Checks at runtime if the value of *myBigNumber* overflows int32 or not.
if err != nil {
    return err
}
```

## Architecture support
`As` currently supports these architectures and correctly handles overflows for 64/32 bit numbers:
* 386
* amd64
* arm
* arm64

It may work for s390x or ppc64le, but I have no way to test that.
