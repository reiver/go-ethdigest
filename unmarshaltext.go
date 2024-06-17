package ethdigest

import (
	"bytes"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-hexadeca"
)

// "0x"
var hexlitprefix [2]byte = [2]byte{'0', 'x'}

//unmarshalText unmarshals the ("0x" prefixed) hexadecimal-literal in 'text' into 'dst'.
//
// It, for example, turns this:
//
//	[]byte("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed0f1e2d3c4b5a69788796a5b4")
//
// Into:
//
//	[32]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed,0x0f,0x1e,0x2d,0x3c,0x4b,0x5a,0x69,0x78,0x87,0x96,0xa5,0xb4}
func unmarshalText(dst *[DigestLength]byte, text []byte) error {

	if nil == dst {
		return errNilDestination
	}

	var hex []byte
	{
		if !bytes.HasPrefix(text, hexlitprefix[:]) {
			return errMissingHexadecimalLiteralPrefix
		}

		hex = text[len(hexlitprefix):]
	}

	{
		var expected int = DigestLength*2 + len(hexlitprefix)
		var actual   int = len(text)

		if expected != actual {
			return erorr.Errorf("ethdigest: the eth-digest is expected to be %d bytes long, but was actually %d bytes long", expected, actual)
		}
	}

	{
		var limit int = len(hex) / 2

		for i:=0; i<limit; i++ {
			var indexToMostSignificant  int = i * 2
			var indexToLeastSignificant int = indexToMostSignificant + 1

			var mostSignificantHex  byte = hex[indexToMostSignificant]
			var leastSignificantHex byte = hex[indexToLeastSignificant]

			mostSignificant, ok := hexadeca.DecodeByte(mostSignificantHex)
			if !ok {
				return erorr.Errorf("ethdigest: byte number-%d (after \"0x\" prefix) of hexadecimal literal (%d) (%q) is not a valid hexadecimal symbol", indexToMostSignificant, mostSignificantHex, mostSignificantHex)
			}

			leastSignificant, ok := hexadeca.DecodeByte(leastSignificantHex)
			if !ok {
				return erorr.Errorf("ethdigest: byte number-%d (after \"0x\" prefix) of hexadecimal literal (%d) (%q) is not a valid hexadecimal symbol", indexToLeastSignificant, leastSignificantHex, leastSignificantHex)
			}

			var value byte = (mostSignificant << 4) | leastSignificant

			dst[i] = value
		}
	}

	return nil
}
