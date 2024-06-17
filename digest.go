package ethdigest

import (
	"encoding"
	"fmt"
	"math/big"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-opt"
)

// DigestLength is the length of the eth-digest measured in number of bytes.
const DigestLength = 32

// Digest holds the eth-digest.
//
// Note that type Digest is an optional-type.
//
// (The 'optional-type' is also known as an 'option-type' and 'maybe-type' in other programming languages and libraries.)
type Digest struct {
	optional opt.Optional[[DigestLength]byte]
}

var _ encoding.BinaryMarshaler = Digest{}
var _ encoding.BinaryUnmarshaler = &Digest{}
var _ encoding.TextMarshaler = Digest{}
var _ encoding.TextUnmarshaler = &Digest{}

// Nothing returns an empty digest.
//
// For example:
//
//	var digest ethdigest.Digest = ethdigest.Nothing()
//
// Nothing is is NOT the same thing as 0x0000000000000000000000000000000000000000
func Nothing() Digest {
	return Digest{}
}

// Something returns an digest whose value is equal to the contents of the 'value' variable.
//
// For example:
//
//	// 0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed
//	var digest ethdigest.Digest = ethdigest.Something( [32]byte{0x5a,0xae,0xb6,0x05,0x3f,0x3e,0x94,0xc9,0xb9,0xa0,0x9f,0x33,0x66,0x94,0x35,0xe7,0xef,0x1b,0xea,0xed,0x5a,0xae,0xb6,0x05,0x3f,0x3e,0x94,0xc9,0xb9,0xa0,0x9f,0x33} )
func Something(value [DigestLength]byte) Digest {
	return Digest{
		optional: opt.Something(value),
	}
}

// Parse returns the eth-digest represented by the hexadecimal-literal.
func Parse(text []byte) (Digest, error) {
	var digest Digest

	err := digest.UnmarshalText(text)
	if nil != err {
		return Nothing(), err
	}

	return digest, nil
}

// ParseString returns the eth-digest represented by the hexadecimal-literal.
func ParseString(text string) (Digest, error) {
	return Parse([]byte(text))
}

// ParseElsePanic is similar to the Parse func, except that it panic()s if there is an error.
func ParseElsePanic(text []byte) Digest {
	digest, err := Parse(text)
	if nil != err {
		panic(err)
	}

	return digest
}

// ParseStringElsePanic is similar to the ParseString func, except that it panic()s if there is an error.
func ParseStringElsePanic(text string) Digest {
	return ParseElsePanic([]byte(text))
}

// BigInt returns the eth-digest represented by the *big.Int.
func BigInt(bigint *big.Int) (Digest, error) {
	if nil == bigint {
		return Nothing(), errNilBigInt
	}

	if bigint.Cmp(minDigest) < 0 {
		return Nothing(), erorr.Errorf("ethdigest: digest-underflow — expected numerical value for digest to be between %s and %s but actually was %s", minDigest, maxDigest, bigint)
	}

	if bigint.Cmp(maxDigest) > 0 {
		return Nothing(), erorr.Errorf("ethdigest: digest-overflow — expected numerical value for digest to be between %s and %s but actually was %s", minDigest, maxDigest, bigint)
	}

	var digest [DigestLength]byte
	bigint.FillBytes(digest[:])

	return Something(digest), nil
}

// BigIntElsePanic is similar to the BigInt func, except that it panic()s if there is an error.
func BigIntElsePanic(bigint *big.Int) Digest {

	digest, err := BigInt(bigint)
	if nil != err {
		panic(err)
	}

	return digest
}

// Bytes returns the (decoded) bytes of the eth-digest.
//
// For example, if the hexadecimal-literal of an eth-addres is:
//
//	"0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed0f1e2d3c4b5a69788796a5b4"
//
// Then this Bytes method would return:
//
//	[]byte{0x5a, 0xae, 0xb6, 0x05, 0x3f, 0x3e, 0x94, 0xc9, 0xb9, 0xa0, 0x9f, 0x33, 0x66, 0x94, 0x35, 0xe7, 0xef, 0x1b, 0xea, 0xed, 0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4}
//
// It would NOT return []byte("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed0f1e2d3c4b5a69788796a5b4")
//
// I.e., it returns the "binary" representation (NOT the ASCII or UTF-8 "text" representation).
func (receiver Digest) Bytes() []byte {
	value, something := receiver.optional.Get()
	if !something {
		return nil
	}

	return value[:]
}

// If the receiver contains nothing, then BigInt returns nil.
// If the receiver contains something, then BitInt returns that something as a *big.Int.
//
// For example:
//
//	// 0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed0f1e2d3c4b5a69788796a5b4
//	var digest ethdigest.Something( [32]byte{0x5a, 0xae, 0xb6, 0x05, 0x3f, 0x3e, 0x94, 0xc9, 0xb9, 0xa0, 0x9f, 0x33, 0x66, 0x94, 0x35, 0xe7, 0xef, 0x1b, 0xea, 0xed, 0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4 )
//	
//	bigint := digest.BigInt()
func (receiver Digest) BigInt() *big.Int {
	value, something := receiver.optional.Get()
	if !something {
		return nil
	}

	return new(big.Int).SetBytes(value[:])
}

// If Digest contains nothing, then method Get returns false.
// If Digest contains something, then method Get return true and then [32]byte that represents the digest.
//
// For example:
//
//	var digest ethdigest.Digest = ethdigest.Nothing()
//	
//	// something == false
//	array, something := digest.Get()
//
// And, for example
//
//	// 0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed0f1e2d3c4b5a69788796a5b4
//	var digest ethdigest.Digest = ethdigest.Something( [32]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed,0x0f,0x1e,0x2d,0x3c,0x4b,0x5a,0x69,0x78,0x87,0x96,0xa5,0xb4} )
//	
//	// something == true
//	// array == [32]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed,0x0f,0x1e,0x2d,0x3c,0x4b,0x5a,0x69,0x78,0x87,0x96,0xa5,0xb4}
//	array, something := digest.Get()
func (receiver Digest) Get() ([DigestLength]byte, bool) {
	return receiver.optional.Get()
}

// If Digest contains nothing, then method GetElse returns the value of the 'alternative' parameter.
// If Digest contains something, then method GetElse returns the [32]byte that represents the digest.
//
// For example:
//
//	var alternative [32]byte{0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE}
//	
//	var digest ethdigest.Digest = ethdigest.Nothing()
//	
//	// array == alternative == [32]byte{0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE}
//	array := digest.GetElse()
//
// And, for example
//
//	var alternative [32]byte{0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE}
//	
//	// 0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed0f1e2d3c4b5a69788796a5b4
//	var digest ethdigest.Digest = ethdigest.Something( [32]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed,0x0f,0x1e,0x2d,0x3c,0x4b,0x5a,0x69,0x78,0x87,0x96,0xa5,0xb4} )
//	
//	// array == [32]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed,0x0f,0x1e,0x2d,0x3c,0x4b,0x5a,0x69,0x78,0x87,0x96,0xa5,0xb4}
//	array := digest.GetElse()
func (receiver Digest) GetElse(alternative [DigestLength]byte) [DigestLength]byte {
	return receiver.optional.GetElse(alternative)
}

// GoString returns the value of the receiver in Go syntax.
//
// This function is called when the "%#v" verb is used with the printing-functions from the Go built-in "fmt" package.
//
// For example:
//
//	var digest ethdigest = ethdigest.Something([32]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed,0x0f,0x1e,0x2d,0x3c,0x4b,0x5a,0x69,0x78,0x87,0x96,0xa5,0xb4})
//	
//	fmt.Printf("digest = %#v \n", digest)
func (receiver Digest) GoString() string {
	value, something := receiver.optional.Get()
	if !something {
		return "ethdigest.Nothing()"
	}

	return fmt.Sprintf("ethdigest.Something(%#v)", value)
}

// IsNothing return true if the receiver digest contains nothing.
// IsNothing return false if the receiver digest contains something.
func (receiver Digest) IsNothing() bool {
	return receiver.optional.IsNothing()
}

// IsSomething return true if the receiver digest contains something.
// IsSomething return false if the receiver digest contains nothing.
func (receiver Digest) IsSomething() bool {
	return receiver.optional.IsSomething()
}

// MarshalBinary returns the eth-digest in its binary form as a []byte.
func (receiver Digest) MarshalBinary() ([]byte, error) {
	value, something := receiver.optional.Get()
	if !something {
		return nil, errNothing
	}

	return value[:], nil
}

// MarshalText returns the eth-digest in its textual form as a []byte.
//
// (Note that this is different than the "binary" form of the eth-digest as a []byte.)
func (receiver Digest) MarshalText() ([]byte, error) {
	if receiver.IsNothing() {
		return nil, errNothing
	}

	return []byte(receiver.String()), nil
}

// MarshalText returns the eth-digest in its textual form as a string.
func (receiver Digest) String() string {
	array, something := receiver.optional.Get()
	if !something {
		return ""
	}

	return fmt.Sprintf("0x%32x", array[:])
}

// UnmarshalBinary sets the receiver to the eth-digest in its binary form as a []byte.
func (receiver *Digest) UnmarshalBinary(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	{
		const expected int = DigestLength
		var   actual   int = len(data)

		if expected != actual {
			return erorr.Errorf("ethdigest: the actual length of the data parameter (%d) is not what was expected (%d)", actual, expected)
		}
	}

	var digest [DigestLength]byte
	{
		copy(digest[:], data)
	}

	receiver.optional = opt.Something(digest)
	return nil
}

// UnmarshalText sets the receiver the eth-digest represented by the hexadecimal-literal.
func (receiver *Digest) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	var digest [DigestLength]byte

	err := unmarshalText(&digest, text)
	if nil != err {
		return err
	}

	receiver.optional = opt.Something(digest)
	return nil
}
