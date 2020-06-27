package util

import (
	"log"
	"testing"
)

func TestJSON(t *testing.T) {
	var (
		err error
		dataByte = make([]byte,150)
	)
	type GT struct {
		Ui string
		Mi int
	}
	type test1 struct {
		Gg	int
		Gt	GT
	}

	var ans1 test1
	rep1 := test1{
		Gg:	123,
		Gt: GT{
			Ui 	:	"uiy",
			Mi	:	12,
		},
	}

	dataByte,err = IJson.Marshal(rep1)

	for len(dataByte)<150 {
		dataByte = append(dataByte, byte(0))
	}


	if err != nil {
		t.Fatalf("%v : marshal error , because %v", rep1, err)
	}

	err = IJson.Unmarshal(dataByte,&ans1)
	if err != nil {
		t.Fatalf("%v : s1 unmarshal error , because %v", string(dataByte), err)
	}
	log.Printf("%d,%s,%d\n",ans1.Gg,ans1.Gt.Ui,ans1.Gt.Mi)
}


func TestByte(t *testing.T) {
	type o struct {
		F int
		UC []byte
	}

	o1 := &o{
		F:  80,
		UC: []byte{1,1,1,1,1,0,0,1,89},
	}
	dataByte,err := IJson.Marshal(o1)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(dataByte)
	var ans1 o
	err = IJson.Unmarshal(dataByte,&ans1)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(ans1)
}