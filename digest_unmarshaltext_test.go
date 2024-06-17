package ethdigest_test

import (
	"testing"

	"github.com/reiver/go-ethdigest"
)

func TestDigest_UnmarshalText(t *testing.T) {
	tests := []struct{
		Text []byte
		Expected ethdigest.Digest
	}{
		{
			Text: []byte("0x0000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},



		{
			Text: []byte("0x1000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x10,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x2000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x20,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x3000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x30,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x4000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x40,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x5000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x50,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x6000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x60,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x7000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x70,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x8000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x80,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x9000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x90,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xA000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xA0,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xB000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xB0,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xC000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xC0,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xD000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xD0,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xE000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xE0,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xF000000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xF0,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x0100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x1100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x11,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x2100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x21,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x3100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x31,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x4100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x41,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x5100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x51,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x6100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x61,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x7100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x71,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x8100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x81,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x9100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x91,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xA100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xA1,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xB100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xB1,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xC100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xC1,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xD100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xD1,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xE100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xE1,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xF100000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xF1,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x0200000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x02,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x1200000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x12,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x2200000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x22,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x3200000000000000000000000000000000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x32,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},



		{
			Text: []byte("0x123456789ABCDEF2468ACE13579BDF0000000000000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x12,0x34,0x56,0x78,0x9A,0xBC,0xDE,0xF2,0x46,0x8A,0xCE,0x13,0x57,0x9B,0xDF,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},



		{
			Text: []byte("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359000000000000000000000000"),
			Expected: ethdigest.Something( [...]byte{0xfB,0x69,0x16,0x09,0x5c,0xa1,0xdf,0x60,0xbB,0x79,0xCe,0x92,0xcE,0x3E,0xa7,0x4c,0x37,0xc5,0xd3,0x59,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} ),
		},
		{
			Text: []byte("0x000000000000000000000000dbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"),
			Expected: ethdigest.Something( [...]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xdb,0xF0,0x3B,0x40,0x7c,0x01,0xE7,0xcD,0x3C,0xBe,0xa9,0x95,0x09,0xd9,0x3f,0x8D,0xDD,0xC8,0xC6,0xFB} ),
		},
		{
			Text: []byte("0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb0f1e2d3c4b5a69788796a5b4"),
			Expected: ethdigest.Something( [...]byte{0xD1,0x22,0x0A,0x0c,0xf4,0x7c,0x7B,0x9B,0xe7,0xA2,0xE6,0xBA,0x89,0xF4,0x29,0x76,0x2e,0x7b,0x9a,0xDb,0x0f,0x1e,0x2d,0x3c,0x4b,0x5a,0x69,0x78,0x87,0x96,0xa5,0xb4} ),
		},
	}

	for testNumber, test := range tests {

		var actual ethdigest.Digest
		err := actual.UnmarshalText(test.Text)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("TEXT: %q (%X)", test.Text, test.Text)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual unmarshaled value is not what wad expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("TEXT: %q (%X)", test.Text, test.Text)
				continue
			}
		}
	}
}

func TestDigest_UnmarshalText_error(t *testing.T) {
	tests := []struct{
		Text []byte
		ExpectedError string
	}{
		{
			Text: []byte(""),
			ExpectedError: "ethdigest: missing prefix for hexadecimal-literal (i.e., \"0x\")",
		},
		{
			Text: []byte("Hello world!"),
			ExpectedError: "ethdigest: missing prefix for hexadecimal-literal (i.e., \"0x\")",
		},
		{
			Text: []byte("pqrstuvwxyzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"),
			ExpectedError: "ethdigest: missing prefix for hexadecimal-literal (i.e., \"0x\")",
		},
		{
			Text: []byte("0xrstuvwxyzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"),
			ExpectedError: "ethdigest: byte number-0 (after \"0x\" prefix) of hexadecimal literal (114) ('r') is not a valid hexadecimal symbol",
		},



		{
			Text: []byte("0xW000000000000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-0 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0W00000000000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-1 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00W0000000000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-2 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000W000000000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-3 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000W00000000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-4 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000W0000000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-5 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000W000000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-6 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000W00000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-7 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000W0000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-8 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000W000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-9 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000W00000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-10 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000W0000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-11 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000W000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-12 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000W00000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-13 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000W0000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-14 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000W000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-15 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000W00000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-16 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000W0000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-17 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000W000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-18 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000W00000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-19 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000W0000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-20 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000W000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-21 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000W00000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-22 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000W0000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-23 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000W000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-24 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000W00000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-25 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000W0000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-26 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000W000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-27 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000W00000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-28 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000W0000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-29 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000W000000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-30 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000W00000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-31 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000W0000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-32 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000W000000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-33 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000000W00000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-34 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000000W0000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-35 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000000W000000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-36 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000000000W00000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-37 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000000000W0000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-38 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000000000W000000000000000000000000"),
			ExpectedError: "ethdigest: byte number-39 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000000000000W00000000000000000000000"),
			ExpectedError: "ethdigest: byte number-40 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000000000000W0000000000000000000000"),
			ExpectedError: "ethdigest: byte number-41 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000000000000W000000000000000000000"),
			ExpectedError: "ethdigest: byte number-42 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000000000000000W00000000000000000000"),
			ExpectedError: "ethdigest: byte number-43 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000000000000000W0000000000000000000"),
			ExpectedError: "ethdigest: byte number-44 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000000000000000W000000000000000000"),
			ExpectedError: "ethdigest: byte number-45 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000000000000000000W00000000000000000"),
			ExpectedError: "ethdigest: byte number-46 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000000000000000000W0000000000000000"),
			ExpectedError: "ethdigest: byte number-47 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000000000000000000W000000000000000"),
			ExpectedError: "ethdigest: byte number-48 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000000000000000000000W00000000000000"),
			ExpectedError: "ethdigest: byte number-49 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000000000000000000000W0000000000000"),
			ExpectedError: "ethdigest: byte number-50 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000000000000000000000W000000000000"),
			ExpectedError: "ethdigest: byte number-51 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000000000000000000000000W00000000000"),
			ExpectedError: "ethdigest: byte number-52 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000000000000000000000000W0000000000"),
			ExpectedError: "ethdigest: byte number-53 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000000000000000000000000W000000000"),
			ExpectedError: "ethdigest: byte number-54 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000000000000000000000000000W00000000"),
			ExpectedError: "ethdigest: byte number-55 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000000000000000000000000000W0000000"),
			ExpectedError: "ethdigest: byte number-56 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000000000000000000000000000W000000"),
			ExpectedError: "ethdigest: byte number-57 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000000000000000000000000000000W00000"),
			ExpectedError: "ethdigest: byte number-58 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000000000000000000000000000000W0000"),
			ExpectedError: "ethdigest: byte number-59 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000000000000000000000000000000W000"),
			ExpectedError: "ethdigest: byte number-60 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x0000000000000000000000000000000000000000000000000000000000000W00"),
			ExpectedError: "ethdigest: byte number-61 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x00000000000000000000000000000000000000000000000000000000000000W0"),
			ExpectedError: "ethdigest: byte number-62 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},
		{
			Text: []byte("0x000000000000000000000000000000000000000000000000000000000000000W"),
			ExpectedError: "ethdigest: byte number-63 (after \"0x\" prefix) of hexadecimal literal (87) ('W') is not a valid hexadecimal symbol",
		},



		// len -> 65
		{
			Text: []byte("0x000000000000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: the eth-digest is expected to be 66 bytes long, but was actually 65 bytes long",
		},

		// len -> 67
		{
			Text: []byte("0x00000000000000000000000000000000000000000000000000000000000000000"),
			ExpectedError: "ethdigest: the eth-digest is expected to be 66 bytes long, but was actually 67 bytes long",
		},
	}

	for testNumber, test := range tests {

		var digest ethdigest.Digest
		err := digest.UnmarshalText(test.Text)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("DIGEST: %#v", digest)
			t.Logf("TEXT: %q (%X)", test.Text, test.Text)
			continue
		}

		{
			expected := test.ExpectedError
			actual := err.Error()

			if expected != actual {
				t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("DIGEST: %#v", digest)
				t.Logf("TEXT: %q (%X)", test.Text, test.Text)
				continue
			}
		}
	}
}
