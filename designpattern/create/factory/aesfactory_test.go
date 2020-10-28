package factory

import "testing"

func TestAesFactory(t *testing.T) {
	srcString := "hello world"
	ecbString := "ECB:hello world"
	cbcString := "CBC:hello world"

	aesFactory := AesFactory{}

	aes := aesFactory.GetAes(AesTypeECB)
	encodeStr := aes.Encode(srcString)
	if encodeStr != ecbString {
		t.Log("ecb encoding failed: ", encodeStr, ecbString)
		t.FailNow()
	}

	decodeStr, err := aes.Decode(encodeStr)
	if err != nil {
		t.Log("decode error: ", err)
		t.FailNow()
	}

	if decodeStr != srcString {
		t.FailNow()
	}

	aes = aesFactory.GetAes(AesTypeCBC)
	encodeStr = aes.Encode(srcString)
	if encodeStr != cbcString {
		t.FailNow()
	}

	decodeStr, err = aes.Decode(encodeStr)
	if err != nil {
		t.Log("decode error: ", err)
		t.FailNow()
	}

	if decodeStr != srcString {
		t.FailNow()
	}
}
