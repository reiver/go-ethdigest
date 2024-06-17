# go-ethdigest

Package **ethdigest** provides tools for working with **eth-digests**, for the Go programming language.

The **eth-digest** (also called an **ethereum-digest**, **evm-digest**, and  mistakenly also called an **ethereum-hash** or **evm-hash**) is a 32-byte digest — and is commonly used by **EVM based networks**, as well as other places.

`ethdigest.Digest` is meant to be a **replacement** for `go-ethereum/common.Hash` from the **official Ethereum golang package**.

## Examples

Here is an example of loading an` ethdigest.Digest` from a hexadecimal-literal stored in a Go `string`:

```golang
digest, err := ethdigest.ParseString("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed0f1e2d3c4b5a69788796a5b4")
```

Here is an example of loading an` ethdigest.Digest` from a hexadecimal-literal stored in a Go `[]byte`:

```golang
var bytes []byte = []byte("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed0f1e2d3c4b5a69788796a5b4")

digest, err := ethdigest.Parse(bytes)
```

Here is an example of loading an` ethdigest.Digest` from a Go `[20]byte`:

```golang
digest := ethdigest.Something( [32]{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed,0x0f,0x1e,0x2d,0x3c,0x4b,0x5a,0x69,0x78,0x87,0x96,0xa5,0xb4} )
```

Here is an example of loading an` ethdigest.Digest` from a Go `*big.Int`:

```golang
var bigint *big.Int = // ...

digest := ethdigest.BigInt(bigint)
```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-ethdigest

[![GoDoc](https://godoc.org/github.com/reiver/go-ethdigest?status.svg)](https://godoc.org/github.com/reiver/go-ethdigest)

## Import

To import package **ethdigest** use `import` code like the follownig:
```
import "github.com/reiver/go-ethdigest"
```

## Installation

To install package **ethdigest** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-ethdigest
```

## Author

Package **ethdigest** was written by [Charles Iliya Krempeaux](http://reiver.link)
