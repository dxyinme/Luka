package FileServer

import (
	"os"
	"testing"
)

func TestFileClient_SendFile(t *testing.T) {
	fc := &FileClient{
		Host: "localhost:10505",
	}
	testFilename := "test.file"
	fp ,err := os.Create(testFilename)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10 * chunkSize + 7783; i ++ {
		_, err := fp.Write([]byte("x"))
		if err != nil {
			t.Fatal(err)
		}
	}
	fp.Close()
	err = fc.SendFile(testFilename, "testUser")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Remove(testFilename)
}