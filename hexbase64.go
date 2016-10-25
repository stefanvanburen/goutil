package goutil

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexbytes []byte) ([]byte, error) {
	bytes := make([]byte, hex.DecodedLen(len(hexbytes)))
	_, err := hex.Decode(bytes, hexbytes)
	if err != nil {
		return nil, err
	}
	base64bytes := make([]byte, base64.StdEncoding.EncodedLen(len(bytes)))
	base64.StdEncoding.Encode(base64bytes, bytes)
	return base64bytes, nil
}

func HexStringToBase64String(hexstring string) (string, error) {
	decodedhex, err := hex.DecodeString(hexstring)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(decodedhex), nil

}
