package FileServer

import (
	"github.com/golang/glog"
	"os"
	"sync"
	"time"
)

type fInfo struct {
	filepath string
	uploadTime time.Time
	locked bool
	lockMu sync.Mutex
	fp *os.File
}

func NewFInfo(from string, md5 string) *fInfo {
	ret := &fInfo{
		filepath:   from + "/" + md5,
		uploadTime: time.Now(),
		locked:     false,
		fp:         nil,
	}
	var err error

	if _, err := os.Stat(from); os.IsNotExist(err) {
		err = os.Mkdir(from, 0777)
		if err != nil {
			glog.Error(err)
			return nil
		}
		err = os.Chmod(from, 0777)
		if err != nil {
			glog.Error(err)
			return nil
		}
	}

	ret.fp, err = os.Create(ret.filepath)
	if err != nil {
		glog.Fatal(err)
		return nil
	}
	return ret
}

func (f *fInfo) Append(b []byte) error {
	f.lockMu.Lock()
	defer f.lockMu.Unlock()
	if f.locked {
		return nil
	}
	_, err := f.fp.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (f *fInfo) IsLocked() bool {
	f.lockMu.Lock()
	defer f.lockMu.Unlock()
	return f.locked
}

func (f *fInfo) Locked() error {
	f.lockMu.Lock()
	defer f.lockMu.Unlock()
	f.locked = true
	f.uploadTime = time.Now()
	return f.fp.Close()
}