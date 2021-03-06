// Code generated by go-bindata.
// sources:
// cmd/genus/schema/plan_schema.json
// cmd/genus/schema/schema.go
// DO NOT EDIT!

package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cmdGenusSchemaPlan_schemaJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\x4d\x4f\xe3\x3c\x10\xbe\xf7\x57\x8c\xf2\x72\x0c\x6f\x77\xa5\x3d\x71\x83\xa5\x20\x24\xc4\x56\x50\x69\x0f\x2b\x0e\x6e\x32\x6d\x0c\x89\x9d\xb5\x9d\x42\xb5\xe2\xbf\xef\xd8\xf9\x68\x93\x38\xc0\x42\xca\x01\x81\xc7\x33\xcf\xe3\xf9\x78\x26\x7f\x26\x00\xc1\x91\x8e\x12\xcc\x58\x70\x02\x41\x62\x4c\x7e\x32\x9d\x3e\x68\x29\x8e\xcb\xd3\xff\xa5\x5a\x4f\x63\xc5\x56\xe6\xf8\xcb\xb7\x69\x79\xf6\x5f\x10\x5a\x3f\xb3\xcd\xd1\x3a\xc9\xe5\x03\x46\xa6\x3c\xcb\x95\xcc\x51\x19\x8e\x9a\x2c\x36\x3a\x9d\xdd\x15\xab\x15\x7f\x6e\xfe\xdf\xf3\xd4\x46\x71\xb1\x76\x9e\xee\x3c\xe3\xe2\x1a\xc5\xda\x24\x64\xfc\xda\x9c\xc6\xa8\x23\xc5\x73\xc3\xa5\xb0\x4e\x65\x38\x90\x2b\x30\x98\xe5\x29\x33\x84\xe5\xae\xbe\x94\x1e\xc1\xa2\x3a\x3e\xe7\x6a\x34\xd0\x3a\x26\xc4\x5c\xd1\x63\xa5\xda\xb6\x31\xcf\x98\x1e\x15\xcf\xc6\xdb\x61\xc1\x4a\x2a\x58\xa3\x40\x45\x14\x62\x88\x64\x8c\x21\xc4\xb8\x62\x45\x6a\x9c\xed\x68\xfe\xf3\xbc\x4f\x68\xce\xa2\x47\xb6\xc6\x71\x49\xad\x25\xe4\x65\x5c\x87\x9c\xca\x88\xa5\xc0\xb3\x5c\x2a\xa3\x81\x8b\x0e\xcd\x36\xa9\xbb\x47\x9e\xcf\x9e\xb9\x36\xda\xc7\x69\x29\x65\x8a\x4c\x04\x83\x85\x47\x03\x46\x82\x51\x05\x02\x5f\xc1\x56\x16\xf0\xc4\x84\x3b\xd3\x14\xb8\x46\xa6\xdb\x20\x45\x44\xf4\x78\x8a\x80\x25\x5c\x8f\xc6\x85\x54\x19\x33\x07\xa1\x21\x6d\x62\x28\x38\x25\xf8\x3d\x44\xf8\xf3\x55\x99\xbd\x43\x90\xa1\x49\xb1\x34\xaa\xfa\xb4\xd1\xe7\x29\x13\x57\x34\x43\x5e\x60\xa6\x14\xdb\xee\x60\x79\xe7\x9e\x7f\xfc\x2b\x8b\x47\x04\x2a\x4b\x4f\x0a\x5e\x6f\xcb\xd7\x9a\x73\x20\x1f\x43\xe2\x50\xfe\xbc\xec\x7b\x0f\x08\xc5\xd8\x94\x86\xa4\xc3\xc7\xa8\x2f\x23\x63\xb3\xf9\xb0\xb0\x0c\xd1\xed\x8b\xcc\x41\x28\x7f\x44\x76\x7c\x94\xbd\x12\xd4\x61\xdc\x1d\x3a\x3f\xad\x4f\xcb\xd1\x10\xbd\x9e\x34\x1d\x86\xde\xdb\x32\x35\x48\xd0\x27\x59\x07\x21\xe9\x93\x2f\x1f\xab\x0f\x36\x61\x87\x4d\x15\xc5\xaa\xc7\x7b\xdb\xc9\x6a\xe8\xa2\xc4\xf8\x27\x64\x14\x45\x46\xd6\x5f\xc1\xdd\xd5\xcd\xe5\xf5\x6c\xf1\xe3\x26\x08\x21\xb8\x9d\xcd\x67\xa7\x8b\xd3\xb3\xeb\x59\x70\xff\xae\xc4\xed\x1c\x6c\xfa\x48\x66\x04\x70\x4d\x35\xc4\xa8\xb0\xdc\x33\x1a\x63\x9e\x53\x55\x0d\xcf\x50\xc3\x13\x37\x09\xe0\x06\x69\xee\x35\x71\xa2\x73\x2b\xeb\x6e\x7a\xf8\x06\x05\xc4\xcc\xb0\x10\xa4\x49\x50\x3d\x71\x9a\x39\x4d\x20\xdc\xe1\x34\x2c\x67\xc3\xa9\xb8\x45\x52\x39\x8a\x33\x4e\x2d\xea\x68\xcd\xd4\x5b\x49\x4f\xb0\x53\x18\x28\x44\x8c\xca\x19\x6a\xe9\x1c\xe4\x77\x41\xed\x2d\x58\xf6\x59\x62\x75\x18\x4b\x88\xfa\xb5\x2f\x9f\x26\xa1\x12\x44\x54\x89\x25\x02\x83\x4b\xd9\x6c\x22\x60\x51\x84\xb9\x9b\x37\x4b\x18\xb3\x25\xc6\xe4\x75\x4e\x59\x77\x85\x78\x7b\x51\xdd\x10\xee\x2b\x33\xd7\xde\xd9\x95\xb1\xbf\xb9\xfd\xaf\x6e\x99\x5b\xf0\xc3\x1b\xcd\xe6\x41\x87\x50\xe8\x82\xa5\x29\x35\x95\xa4\xb4\xa8\x5d\xe1\xa8\xdb\x5c\xcb\x35\x09\xd0\xe5\x76\x56\x98\xc9\x0d\xc6\xc3\xef\x7d\x53\x5d\x7a\xdf\x1c\x3e\x9a\xdf\x25\xc1\x22\xc1\xd6\x6b\xc2\x2e\x0e\x57\x32\x3b\x28\x61\x53\xa3\xa5\x96\x69\xb1\x77\xad\xbb\x5e\x42\x78\xc4\xad\x9d\x2b\x96\x72\xa6\xeb\x46\x2c\x8d\x21\x30\x11\xc3\x86\xa5\x56\xc0\xb4\x33\xb0\xc8\x50\x3a\xdc\xf3\x87\x5f\x68\x8b\x3e\x76\x21\xab\xa4\xb4\x0b\x39\xf1\xfd\xbd\xc7\x25\x50\xf8\xbb\xa0\xcf\x81\xd8\xa9\x51\x23\x67\xf7\x93\x7d\x1f\xfb\xdb\xf9\xb4\x6e\xdf\x4f\x5e\x26\x7f\x03\x00\x00\xff\xff\x0b\x8c\x8a\xc8\x41\x0e\x00\x00")

func cmdGenusSchemaPlan_schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_cmdGenusSchemaPlan_schemaJson,
		"cmd/genus/schema/plan_schema.json",
	)
}

func cmdGenusSchemaPlan_schemaJson() (*asset, error) {
	bytes, err := cmdGenusSchemaPlan_schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/genus/schema/plan_schema.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdGenusSchemaSchemaGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x96\xdb\x8f\xdb\xc6\x15\xc6\x9f\x97\x7f\x05\xb3\x40\x0a\x2d\xe0\x4a\xbc\x5f\x0c\xf8\x25\xb6\x5b\xb8\x40\x9d\xa2\xf1\x5b\x55\x18\x43\x72\xb8\x61\x2b\x89\xaa\x28\xa5\x5a\x07\xf9\xdf\x7b\x7e\x73\x8e\x9c\x4d\x5b\x37\x2f\x5d\x80\x2b\x72\x38\xe7\xf6\x9d\xef\x7c\xc3\xcd\x26\x7e\x3d\x0f\x3e\x7e\xf4\x07\x7f\x72\x67\x3f\xc4\xdd\x53\xfc\x38\xff\xb6\x9b\x0e\x83\x3b\xbb\x75\x24\x1b\x96\xf9\x72\xea\xfd\xf2\x92\xfb\x7e\x3f\x6c\x64\xef\x65\xd9\x2c\xfd\xf7\x7e\xef\x36\xc7\x9d\x3b\x7c\xd4\xfb\xf5\xdf\x96\xf9\xf0\x5f\x37\xd9\xfb\xc7\x99\xb7\x6f\xbe\x8d\xdf\x7f\xfb\x21\x7e\xfb\xe6\xdd\x87\xaf\xa2\xe8\xe8\xfa\xbf\xbb\x47\x1f\xeb\x96\x28\x9a\xf6\xc7\xf9\x74\x8e\x57\xd1\xdd\x7d\xf7\x74\xf6\xcb\xbd\xdc\xf4\xf3\xfe\x78\xf2\xcb\xb2\x79\xfc\x34\x1d\x59\x18\xf7\x67\x7e\xa6\x59\xff\x6f\xa6\xf9\x72\x9e\x76\x3c\xcc\xc1\xe0\xe8\xce\xdf\x6f\xc6\x69\xe7\xb9\x61\x61\x39\x9f\xa6\xc3\x63\x78\x77\x9e\xf6\xfe\x3e\x7a\x88\xa2\xf1\x72\xe8\x63\xab\xf3\xcf\xde\x0d\x2b\x6e\xe2\xbf\xfc\x95\xb0\x2f\xe2\x83\xdb\x4b\x52\xc1\xec\x21\x5e\xdd\x56\xfd\xe9\x34\x9f\x1e\xe2\x1f\xa3\xbb\xc7\x4f\xe1\x29\x7e\xf9\x2a\x26\xab\xf5\x7b\xff\x4f\x9c\xf8\xd3\x2a\xa4\xcd\xf3\x37\x97\x71\x94\x67\xdc\x3e\x3c\x44\x77\xd3\x18\x0c\xbe\x7a\x15\x1f\xa6\x1d\x2e\xee\x4e\xfe\x7c\x39\x1d\x78\x7c\x11\x4b\x49\xeb\xb7\x78\x1f\x57\xf7\x38\x8a\xbf\xfe\xc7\xcb\xf8\xeb\x1f\xee\x35\x93\x10\x4b\x7c\xfc\x14\x45\x77\x3f\xb8\x53\xdc\x5d\xc6\x58\xe3\x68\x90\xe8\xee\xa3\xa6\xf3\x2a\x9e\xe6\xf5\xeb\xf9\xf8\xb4\xfa\x8d\xec\x79\x21\xb9\x89\x55\xbf\x7b\x7b\xcb\x74\xfd\x7a\x37\x2f\x7e\x25\xe5\xff\x9f\xf2\xc1\x8d\xfa\xff\x82\x23\xd9\xa8\x79\xdb\xa2\xa4\xb5\xfe\x86\xd4\x57\x0f\x2f\xd8\x11\xc9\xbb\xf3\xd3\xd1\xc7\x6e\x59\xfc\x19\xc8\x2f\xfd\x19\x2f\xa1\x3e\xeb\x87\x84\x39\x8c\x73\x1c\xcf\xcb\xfa\x77\xd2\xd6\x77\xf2\xf0\xd9\xce\x5a\x78\x5b\x7f\xe6\x21\xf4\x50\xfe\xb4\x8d\xd1\xdd\x32\x7d\x0a\xcf\xd3\xe1\x5c\x15\xd1\xdd\x1e\xe2\xc7\x9f\x9d\xfe\x51\x1e\xc3\xe2\x07\x61\x48\x0c\x4d\xd6\xdc\x11\x27\x50\x65\x35\x4e\xff\x1e\xeb\x21\x7e\x2f\x21\x56\x0f\x16\x81\x98\x56\xe5\x38\xad\x89\x2e\xc6\x5f\xb6\xfd\x4e\xd2\x11\xdb\x90\xcd\x2f\x4d\x49\xf4\x7f\x9a\x92\xab\x98\x3e\xcb\xfc\x97\x0e\x28\xed\xd7\x1c\x50\x9c\xf8\xf8\x5c\xe8\x7f\x78\xb0\xea\xbf\xec\xe4\xdd\xf2\x66\x3a\x89\x8b\x6e\x9e\x77\xcf\xad\xdd\x6e\xf9\x95\xca\x9f\x16\x2d\xdc\x9f\x46\xd7\xfb\x1f\x7f\x7a\x66\x6d\x94\x80\xe5\x1f\x45\x48\x7e\x8f\x8e\x7c\x17\xc4\xe1\x4f\x3f\x4b\xcd\x1f\x44\x69\x84\xeb\x4a\x8e\xd5\xfd\xf6\x9a\x8e\xdb\x6b\xd3\x6d\xaf\x49\x23\x57\x62\x57\xbb\xbd\x56\x5e\xd6\x6d\x6d\x94\x3d\x5d\xb1\xbd\x96\xd5\xf6\x5a\x0c\x72\xc9\xb3\xcf\xb7\xd7\xbc\x17\x7b\x79\xdf\xc9\xde\xb1\x96\xf7\x72\x35\xb2\x36\x66\xdb\x6b\x2d\x57\x22\xf7\x95\xec\xad\x65\xdd\x95\xb2\x5f\x6c\xeb\x54\xf6\xe4\xfa\x9c\x89\x6d\x26\x7e\x7b\xf3\x5d\xca\x73\x25\xb1\x13\xb1\xc9\xc8\xc9\x6b\x1e\xb9\xf8\xaa\x06\xf5\xd7\xc8\xfb\x56\xee\xbb\x52\x7f\x8b\x4c\xef\x3d\xbf\xe4\x25\xd7\x20\x79\x8f\xd4\x20\xbf\x2d\x79\xca\x6f\x9f\xe8\xde\xde\x89\x1f\x72\x90\xab\x97\xbc\x72\x79\xdf\x5b\x3d\xd8\xd4\xb2\x37\x93\x5c\xea\x51\x7f\xa9\xbf\x97\xbd\x2d\x36\x92\x4b\x8a\x0f\xc9\xa3\x94\x7d\x35\x58\xc9\x73\x21\xef\x2a\x7e\x65\xbd\xb6\x7c\xc9\x2d\xf7\x9a\x43\xd6\xaa\x6d\x2f\x35\x0d\xb9\xe2\x49\xfd\xa5\x53\x2c\x2b\x72\x28\x15\x03\x2f\xd7\xd8\xe8\xde\xae\x56\x3c\x6a\xb9\x46\xf0\x19\x15\x6f\xec\x72\x7a\x82\x9d\xc4\x19\x24\x7e\x2e\x6b\xbd\xec\xf3\xe2\x27\x91\xf5\x42\xf6\xbb\x4a\x7b\x84\xaf\xb6\xb4\xbc\x65\x6f\x6a\xf9\xb4\x62\x93\xc9\x5a\x5e\xa9\x3d\x39\x0f\x72\xa5\xb2\xd7\x75\xfa\x0b\x0e\xf4\x60\x94\x6b\x90\xfb\x51\x7c\x0f\x85\xfa\xa7\x97\x9d\x5c\xb5\xdc\xb7\xf2\xde\x7b\x8d\x07\x96\xf4\x23\x75\x5a\xd7\xe0\xb4\xcf\x95\x5c\x63\xaf\xcf\x2d\x79\x91\x23\xf8\x58\xfd\x35\xbf\xc6\x8d\x8e\x7e\x81\x5f\xa9\xfd\x6b\x13\xe5\x44\x2e\xbf\x6d\xa3\x75\x82\x2b\xfd\x6b\x0a\x7d\x76\x5e\xb9\xd8\xca\x73\xea\xb5\x6f\x2e\xd3\xda\xe8\x85\x07\x4f\xee\xe5\xfd\x90\xe8\x3a\x3d\x0e\xfc\xeb\x15\xc7\xca\xe2\x93\x63\x07\x2e\xa9\xf2\x83\x98\xf8\x04\x13\x9e\xc9\x7d\x90\x7c\xaa\x54\xe3\x14\xf8\x72\xca\x0b\x07\xfe\xe4\x2f\x7b\xd3\x42\xb9\xc1\x2c\x51\x3f\xf3\x91\xa5\x1a\xb3\x6b\x8c\x37\xa5\xe6\x05\x26\x7e\x50\xce\x80\x37\x58\x77\xbd\x71\x04\x4e\x78\xad\xa7\xa8\x35\xb7\xde\x7a\x50\x08\x0e\x4e\xec\x32\xe6\xa0\x50\xcc\xa8\xa7\xa9\x0d\x67\xa7\xf1\xa9\x8b\x39\x80\x37\xbc\x27\x5f\x72\x65\x16\x1a\x9b\x37\xf8\x04\x1f\x9c\xf8\xec\x3a\x8d\x45\x6f\x03\x06\xe0\xdb\xea\x7b\xf0\xe9\x8d\x9f\xf4\x81\x98\x8d\xd3\x19\x48\x0a\x9d\xf1\xa6\x54\xfb\x1b\x1f\x9b\x4c\xf9\x97\x94\x3a\x37\xf0\x19\xec\xe0\x7d\x0a\xef\x13\xc5\x85\x7e\xe6\x36\x2b\xa9\xe1\x74\xe3\x33\x98\x67\x86\x6d\x01\x9f\x0a\x9d\x59\x62\x37\x89\x62\x10\x6a\xb7\x1e\x91\x43\x59\x28\xe7\xe1\x4a\x82\x16\xa5\xda\x03\x34\x05\xfc\xb3\x46\x39\x86\x2d\xb3\x87\x4e\x05\xdf\x8d\xf6\xa0\x44\x6b\xc0\x83\xf5\x5c\xb9\xe8\x0c\x77\xf8\x0f\x9f\x9c\xc4\x1f\x9d\xea\x24\x3c\x82\x6b\xe0\x92\xe6\xaa\x8b\xec\xc1\xb6\x84\xd3\x89\xd6\x42\x6d\xd4\xe3\x6d\x8d\x39\xf7\x86\x37\x73\xc7\xac\xc0\xf9\xc6\xfa\x00\xae\xf4\xac\xe8\x94\x13\x89\x33\x7f\xa3\x69\x4b\xad\x33\xcd\x1c\x81\x11\xda\x5e\xd8\x2c\xa2\xad\x70\x0a\x3e\x79\x74\xaa\x53\x9d\x45\xf7\xc3\x8c\x36\x3a\x37\x0d\x38\x17\x3a\x87\xf4\x37\xcc\xe7\xa8\x33\xca\x6c\x74\xa6\x93\xcc\x27\x57\x62\x73\x03\x6f\xc9\x13\x8d\x41\x0f\xc1\x94\x1a\xc0\x11\x7d\xa8\x7b\xf5\x0f\x86\x68\x28\x1a\xc2\xec\x0c\x9c\x21\x83\xcd\xe8\xa8\x3c\x2e\x6e\xf5\xa4\x76\xbe\x0c\x8a\x09\x18\xb2\x97\x1c\x07\xaf\xd8\xa1\xb5\x39\x5c\x68\x74\xa6\x33\xce\xa1\x56\xfb\x4c\x8f\x7d\xab\x67\x00\x35\x80\x87\xb3\xba\xe1\x30\xf1\x99\x91\xc2\x34\x0f\x1c\x9c\xf1\x1a\x7c\xf0\x8f\xa6\x32\x9b\xbe\xd1\xd9\xeb\x8d\x17\xc4\x61\x9e\xe9\x61\xe0\x22\x39\x55\x3f\x73\x1a\x1d\x1d\x4a\x9d\x27\xb4\x69\x64\xc6\x39\x1b\x53\xd5\xc7\x70\xb6\x34\x9a\x23\xfc\x46\x7b\xd9\x5b\x81\x79\xad\x58\x0e\x86\xb7\xef\xb4\xa6\xda\xce\x62\x74\x8e\x3c\xc0\x3c\xed\xcd\xaf\x53\x0c\xaa\x4a\x79\xc2\x5e\xea\x02\x1b\xf6\xd2\x47\xfa\x05\xde\xcc\x01\xda\x4c\x3f\xc1\xb5\x34\x8e\x27\x83\xce\x39\x5c\xa1\xff\x70\x37\xaf\xf5\x7b\xc0\xd3\xeb\x4a\xe7\x3d\xe8\x7c\xa9\xda\x03\x26\x81\x6f\x9d\xe6\xc9\x59\x51\x3b\xc5\x88\xfd\x89\xc5\x87\x17\xc4\xa7\x9f\xf0\x8b\xd9\x21\xce\xed\x9b\x00\xde\xe6\xad\xe6\x9c\x59\x9e\xde\x66\x8b\x33\x0a\x5c\xc2\xd9\xd0\x9a\x1e\x30\xd7\x99\xda\x30\xfb\xd4\xc4\x79\x91\x81\xb7\x33\xbd\x1c\x94\x67\x65\xa7\xf1\xe0\x15\x39\x50\x6f\xda\x28\x47\x0b\xd3\x64\x34\x32\xbd\xf5\xba\x55\x5c\xd1\x52\xbe\x5d\x0a\xe3\x37\xbe\xd1\xf7\x70\x16\xb7\xaa\x1d\x75\xa9\xbe\x98\x49\x34\x36\x68\x40\xa9\x7c\x6b\xed\x1b\x02\xad\x80\xf3\xc4\x42\x8b\xf8\x86\xc9\x4c\x03\xd1\x49\x6c\x87\x56\xfb\x98\x65\xaa\x05\xe8\x24\x3a\x88\xde\xb6\x9d\xf6\x81\x7d\xc4\x82\x13\xd8\xa3\x6f\x21\xc7\x52\x6b\x0f\x1c\xa9\x55\xf3\xd0\x44\x38\x5c\x1a\x17\xc1\x9d\xbe\x85\x33\x34\x57\xad\x08\x33\xd4\xea\xb7\x01\x3a\xc6\xf7\x13\xf1\x46\xd9\xef\xec\xdc\x47\xf7\xc0\x6f\xbc\x71\xa2\x33\x4e\x57\x8a\x0b\xf3\x43\x1f\x99\x0f\xf4\xdd\xdd\xce\xe5\x52\xfd\xd2\xeb\xce\xf8\x57\x0e\xda\x53\xee\x99\xad\xd2\xbe\x4d\xd8\x0b\xc7\x98\x5f\x66\x05\xbd\xe0\x4c\xe0\xcc\x65\xfe\x02\x2f\x2b\x8d\xcf\xf7\x5d\x3d\x18\x67\x07\xe5\x1a\xdf\x25\xf4\x2b\xb7\x6f\x1a\x9e\x03\x3e\xa9\xe2\x19\xe6\x11\xee\x92\x3b\x7a\xef\xf5\xac\x0b\xba\x9c\xe9\x59\x03\x17\xd0\x03\xfc\x3a\xfb\x06\x68\x8d\xef\xe1\x1b\xa7\x56\x6d\xe2\x4c\x44\x73\xff\x15\x00\x00\xff\xff\xf1\x51\x92\xf7\x00\x10\x00\x00")

func cmdGenusSchemaSchemaGoBytes() ([]byte, error) {
	return bindataRead(
		_cmdGenusSchemaSchemaGo,
		"cmd/genus/schema/schema.go",
	)
}

func cmdGenusSchemaSchemaGo() (*asset, error) {
	bytes, err := cmdGenusSchemaSchemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/genus/schema/schema.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"cmd/genus/schema/plan_schema.json": cmdGenusSchemaPlan_schemaJson,
	"cmd/genus/schema/schema.go": cmdGenusSchemaSchemaGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"cmd": &bintree{nil, map[string]*bintree{
		"genus": &bintree{nil, map[string]*bintree{
			"schema": &bintree{nil, map[string]*bintree{
				"plan_schema.json": &bintree{cmdGenusSchemaPlan_schemaJson, map[string]*bintree{}},
				"schema.go": &bintree{cmdGenusSchemaSchemaGo, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

