package Group

import (
	"log"
	"testing"
)

func TestImpl(t *testing.T) {
	var (
		test 	Group
	)
	test = New("test", "test")
	_ = test.Join("r1")
	_ = test.Join("r2")
	_ = test.Join("r3")
	log.Println(test)
}
