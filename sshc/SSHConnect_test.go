package sshc

import (
	"testing"
)

func TestSSHConnect(t *testing.T) {
	sess, err := SSHConnect("worker", "your password", "your host", 22)
	if err != nil {
		t.Fatal(err)
	}
	defer sess.Close()
	err = sess.Run("echo hahaha >> testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
}

