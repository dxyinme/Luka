package FileServer

import (
	"log"
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
	md5, err := fc.SendFile(testFilename, "testUser")
	if err != nil {
		t.Fatal(err)
	}
	log.Println("md5: ", md5)
	//err = os.Remove(testFilename)
}

func TestFileClient_Download(t *testing.T) {
	fc := &FileClient{
		Host: "localhost:10505",
	}
	testFilename := "test.file.2"
	err := fc.Download(testFilename, "7bde4a2f4db0aff01c3632f1dc446465", "testUser")
	if err != nil {
		t.Fatal(err)
	}
}