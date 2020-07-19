package util


func ByteToInt16(b []byte) int16 {
	return int16(b[0]) << 8 | int16(b[1])
}

func Int16ToByte(v int16) []byte {
	return []byte { uint8(v >> 8) , uint8(v - (v >> 8)) }
}

// 将s转换成长度为l的[]byte
// nil for len(s) + 1 > l
func StringToByteStaticLength(s string, l int) []byte {
	if len(s) + 1 > l {
		return nil
	}
	var res = make([]byte, l)
	for i := range s {
		res[i] = s[i]
	}
	res[len(s)] = uint8(0)
	return res
}

func ByteToString(b []byte) string {
	ed := 0
	lb := len(b)
	for ;ed < lb && b[ed] != uint8(0); {
		ed ++
	}
	return string(b[:ed])
}