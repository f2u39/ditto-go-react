package encode

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var (
	ShiftJIS  = EncName("ShiftJIS")
	EUCJP     = EncName("EUCJP")
	ISO2022JP = EncName("ISO2022JP")
)

type EncName string

func Encode(encname EncName, b []byte) ([]byte, error) {
	enc, err := encode(encname)
	if err != nil {
		return nil, err
	}
	r := bytes.NewBuffer(b)
	encoded, err := ioutil.ReadAll(transform.NewReader(r, enc.NewEncoder()))
	return encoded, err
}

func Decode(encname EncName, b []byte) ([]byte, error) {
	enc, err := encode(encname)
	if err != nil {
		return nil, err
	}
	r := bytes.NewBuffer(b)
	decoded, err := ioutil.ReadAll(transform.NewReader(r, enc.NewDecoder()))
	return decoded, err
}

func encode(encname EncName) (enc encoding.Encoding, err error) {
	switch encname {
	case ShiftJIS:
		enc = japanese.ShiftJIS
	case EUCJP:
		enc = japanese.EUCJP
	case ISO2022JP:
		enc = japanese.ISO2022JP
	default:
		err = fmt.Errorf("Unknown Encoding %s", encname)
	}
	return
}
