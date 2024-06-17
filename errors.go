package ethdigest

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errMissingHexadecimalLiteralPrefix = erorr.Error("ethdigest: missing prefix for hexadecimal-literal (i.e., \"0x\")")
	errNilBigInt                       = erorr.Error("ethdigest: nil big-int")
	errNilDestination                  = erorr.Error("ethdigest: nil destination")
	errNilReceiver                     = erorr.Error("ethdigest: nil receiver")
	errNothing                         = erorr.Error("ethdigest: nothing")
)
