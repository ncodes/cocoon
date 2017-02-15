// Code generated by go-bindata.
// sources:
// bindata.go
// cocoon.job.json
// DO NOT EDIT!

package data

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

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1487176452, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cocoonJobJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\xcf\x6f\xf2\x46\x10\xbd\xe7\xaf\x58\xad\xbe\x63\x03\xa4\xed\x77\x41\xea\xc1\xc5\x34\xa2\x02\x6c\xf1\xe3\x50\x55\x11\x9a\x2c\x13\x67\xc5\x7a\xd7\xda\x1d\xf8\x42\x2d\xff\xef\xd5\xda\xa6\x31\xb6\x21\xcd\x25\xcb\xbc\x37\x9e\xf7\xc6\x3b\xe3\xfc\x81\x31\xfe\xa7\x79\xe5\x63\xe6\x8f\x8c\xf1\x15\x26\xd2\x68\x3e\x66\x3c\x51\xe6\x15\x14\xff\xa9\x8a\xcf\x42\x1f\x13\x46\x18\xa3\x1f\xf3\x7c\x16\x16\xc5\x05\x5a\x42\x8a\x37\xc1\xcd\x39\x2b\x41\x87\xf6\x24\x05\x5e\xc2\xb1\x95\xc6\x4a\x3a\xf3\x31\xfb\x3e\xaa\x63\x81\x52\x01\x45\x5a\x78\xfe\x1b\x28\x87\x75\x3c\x04\x02\x81\x9a\xd0\x3a\x3e\x66\x7f\x97\x41\xc6\xf8\x5e\x3c\xf1\xf2\xfc\x52\xf3\x26\x46\x3b\xb2\x20\x35\x35\x79\x79\xfd\x9f\x31\x3e\xdf\x80\x4d\x90\xbc\x9c\x6f\x39\x10\xd9\xc1\x01\xad\x46\x35\xd0\x90\xe2\x45\x71\xd5\x85\x4f\xa6\x92\xfa\xf8\xd1\xc4\xa2\x0c\x2d\xe8\xbd\xc7\x7e\xe3\x75\xb8\xb8\x12\xb2\x01\x77\x78\xb6\xe6\x98\xdd\xd0\x71\xaf\x61\xb5\x93\xa3\xf6\xc5\xf3\xbc\x3c\x15\xc5\x15\xd6\x74\xa9\x8f\x4a\x35\x40\x5f\xb9\x59\xf4\xba\x70\x5f\x71\x02\x77\xe8\x2a\xa8\xfa\x6e\xe5\x09\xad\xe7\xee\x8d\x38\xa0\x6d\xe3\x13\xa3\xdf\x64\xf2\xdf\xd5\x69\x20\x32\x85\xa4\x2c\x92\xe7\x33\x7f\xec\x3c\x9b\xf9\xfa\x69\x5a\xb7\xf1\x15\xdc\x7b\x97\x00\x36\x29\xad\xf0\x6f\xf9\x32\x5a\x04\xe1\x6e\x31\xdd\x04\xbb\xf5\x64\x35\x8b\x37\xeb\x5d\x38\x5b\x15\xc3\x2b\x24\x9c\xc6\xf3\xe8\xaf\x9a\xb0\x5b\x06\x8b\x69\xc1\x5f\xae\x1e\x5a\xb4\x0c\x4c\xf5\xa9\x4f\xfd\x24\x9a\x44\xd1\x72\x77\xfb\xc6\x77\xb9\x93\x28\x9c\xee\xb6\xab\x79\xe5\x79\x52\xe6\x4c\xcc\x1e\xb7\xab\xf9\x57\x59\x9b\xe0\xb9\x9d\xb5\x81\xe4\xab\xac\x79\xb0\xec\xa4\xcd\x41\xfb\xbc\xbb\x96\xd7\xd5\x18\x96\x8d\x7d\x69\x61\x0b\x24\xe8\xeb\x47\xb7\xb1\xe5\x9d\xc0\xd3\xe3\x1e\x33\x65\xce\x8f\xc2\x68\x8d\x82\x8c\x7d\x1c\x0d\x46\x83\x5f\x07\x7d\x6f\xb3\xf1\xde\x7c\xf6\x50\x19\x01\x6a\xe8\x84\x95\x19\xb9\xfb\x9a\xe7\x26\xb9\x7d\xd5\x16\xf0\xf1\x87\x54\xa5\xa1\xa7\x51\xa7\x6a\x8d\xae\xe5\x3f\xb8\xf8\xbd\xa4\xdc\xad\xb4\xc1\x34\x53\x40\xfd\xed\x09\x2c\xc9\x37\x10\xd5\x72\x69\xeb\x60\x8c\x3f\x23\x11\xda\xb5\x39\xda\x72\x83\xf1\x77\xa2\xcc\x8d\x87\x43\x0b\x3f\x06\x89\xa4\xf7\xe3\xeb\xd1\xa1\x15\x46\x13\x6a\x1a\x08\x93\x0e\xb5\x30\x7b\x74\xc3\xea\x8a\x0d\x53\x70\x84\xf6\xd2\x92\xaf\xef\x76\xdb\x6c\xb9\xbc\x15\x90\x3c\x61\x88\xae\xde\x72\x29\x12\x0c\x9a\x33\xd3\x6a\x75\xdb\xe4\x0a\x5d\x69\xc0\xf5\x0e\x46\xbc\xad\x96\x52\xbc\x2d\x8a\x6e\xaf\x31\x35\xf6\x5c\xb6\x39\xcf\x2f\x3f\x7a\x78\xb3\x28\x5e\xf3\x31\xeb\xbe\xac\x25\xd2\x0f\x63\xab\xfd\x75\x7f\x70\x43\xe9\x32\x20\xf1\x1e\xc3\x59\x19\xf0\x3b\x24\x2f\x1a\x8c\xcf\xf3\x4b\x73\xa7\xdf\xb0\x76\xdb\xd6\xff\xb1\xe4\xb5\x1c\x6a\x46\x75\x6c\xe1\x7d\x76\x6f\x58\x2d\xae\xd5\x12\x58\x8a\x8d\x92\xe2\xdc\x56\x3c\xf3\x1f\xc3\x13\x28\x3e\x66\xbf\x8c\x3e\xff\xae\x4a\x04\x44\x98\x66\xd4\x1d\x0c\x1e\xa2\x02\xff\xc8\x9f\xbf\xf7\x67\x2e\xcc\x1e\xab\x09\xf7\xbc\x5e\x75\x97\x55\x51\xf4\x7e\xff\xb6\xd9\x1e\x08\x1b\xa2\xf9\x9a\x20\x49\xca\x6f\xc9\x53\x57\xae\x9f\xd2\x18\x2c\x28\x85\xde\xd1\xd3\xc3\xe5\x89\xc5\x43\xf1\x6f\x00\x00\x00\xff\xff\xc1\x60\x41\x49\xa1\x08\x00\x00")

func cocoonJobJsonBytes() ([]byte, error) {
	return bindataRead(
		_cocoonJobJson,
		"cocoon.job.json",
	)
}

func cocoonJobJson() (*asset, error) {
	bytes, err := cocoonJobJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cocoon.job.json", size: 2209, mode: os.FileMode(420), modTime: time.Unix(1487176375, 0)}
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
	"bindata.go": bindataGo,
	"cocoon.job.json": cocoonJobJson,
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
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"cocoon.job.json": &bintree{cocoonJobJson, map[string]*bintree{}},
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

