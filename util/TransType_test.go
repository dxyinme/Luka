package util

import (
	"log"
	"testing"
)

func TestTransInt16(t *testing.T) {
	for i := 0; i < (1 << 8) ; i ++ {
		int16i := int16(i)
		if int16i != ByteToInt16(Int16ToByte(int16i)) {
			t.Errorf("%d is trans error", int16i)
		} else {
			log.Printf("%d is verify OK", int16i)
		}
	}
}

func TestTranString(t *testing.T) {
	testStringList := []string{"sss", "31", "000000007000"}
	for i := range testStringList {
		nows := testStringList[i]
		if nows != ByteToString(StringToByteStaticLength(nows,32)) {
			t.Errorf("No. %d test case is error, src : %s\n", i, nows)
		}
	}
}