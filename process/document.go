package process

import (
	"crypto/sha1"
	"encoding/base64"
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
