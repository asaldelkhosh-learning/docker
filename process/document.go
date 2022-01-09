package process

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/hamed-yousefi/gowl"
)

type Document struct {
	content string
	hash    string
}

func (d *Document) Start() error {
	hash := sha1.New()
	hash.Write(nil)
	d.hash = base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return nil
}

func (d *Document) Name() string {
	return "hashing-process"
}

func (d *Document) PID() gowl.PID {
	return "p-1"
}

func (d *Document) Hash() string {
	return d.hash
}
