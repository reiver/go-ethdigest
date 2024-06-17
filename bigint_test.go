package ethdigest_test

import (
	"testing"

	"fmt"
	"math/big"

	"github.com/reiver/go-ethdigest"
)

func TestBigInt(t *testing.T) {

	tests := []struct{
		BigInt *big.Int
		Expected ethdigest.Digest
	}{
		{
			BigInt: big.NewInt(0),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			BigInt: big.NewInt(1),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01} ),
		},
		{
			BigInt: big.NewInt(2),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x02} ),
		},
		{
			BigInt: big.NewInt(3),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x03} ),
		},
		{
			BigInt: big.NewInt(4),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x04} ),
		},
		{
			BigInt: big.NewInt(5),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x05} ),
		},
		{
			BigInt: big.NewInt(6),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x06} ),
		},
		{
			BigInt: big.NewInt(7),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x07} ),
		},
		{
			BigInt: big.NewInt(8),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x08} ),
		},
		{
			BigInt: big.NewInt(9),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x09} ),
		},
		{
			BigInt: big.NewInt(10),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0A} ),
		},
		{
			BigInt: big.NewInt(11),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0B} ),
		},
		{
			BigInt: big.NewInt(12),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0C} ),
		},
		{
			BigInt: big.NewInt(13),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0D} ),
		},
		{
			BigInt: big.NewInt(14),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0E} ),
		},
		{
			BigInt: big.NewInt(15),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0F} ),
		},
		{
			BigInt: big.NewInt(16),
			Expected: ethdigest.Something( [32]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x10} ),
		},



		{
			BigInt: func()*big.Int{
				var str string = "41016844021060681834146332200510513382994275193224758736949524404644867384756"

				var bigint big.Int
				_, ok := bigint.SetString(str, 10)
				if !ok {
					panic(fmt.Sprintf("could not load (decimal) %s into big-int", str))
				}

				return &bigint
			}(),
			Expected: ethdigest.Something( [32]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed,0x0f,0x1e,0x2d,0x3c,0x4b,0x5a,0x69,0x78,0x87,0x96,0xa5,0xb4} ),
		},



		{
			BigInt: func()*big.Int{
				var str string = "115792089237316195423570985008687907853269984665640564039457584007913129639935"

				var bigint big.Int
				_, ok := bigint.SetString(str, 10)
				if !ok {
					panic(fmt.Sprintf("could not load (decimal) %s into big-int", str))
				}

				return &bigint
			}(),
			Expected: ethdigest.Something( [32]byte{0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF} ),
		},
	}

	for testNumber, test := range tests {

		var actual ethdigest.Digest
		var err error

		actual, err = ethdigest.BigInt(test.BigInt)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("BIG-INT: %s", test.BigInt)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual digest is not what was expected", testNumber)
				t.Logf("EXPECTED: %s (%#v)", expected, expected.Bytes())
				t.Logf("ACTUAL:   %s (%#v)", actual,   actual.Bytes())
				t.Logf("BIG-INT: %s", test.BigInt)
				continue
			}
		}
	}
}


