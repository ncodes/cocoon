// Code generated by go-bindata.
// sources:
// data/bindata.go
// data/cocoon.job.json
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

var _dataBindataGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func dataBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_dataBindataGo,
		"data/bindata.go",
	)
}

func dataBindataGo() (*asset, error) {
	bytes, err := dataBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486976870, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataCocoonJobJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x54\x4f\x8f\xda\x3e\x10\xbd\xf3\x29\x2c\xeb\xa7\x3d\xfd\x84\xa0\xd5\x5e\x90\x7a\xa0\xc9\xb6\xa2\x02\x12\xf1\xe7\x54\xa1\xd5\x6c\x32\xa4\x16\x8e\x1d\x39\x03\xbb\x34\xca\x77\xaf\xec\x84\x25\x09\x01\x2e\x0c\xf3\x5e\xfc\xde\xbc\x0c\x2e\x06\x8c\xf1\x5f\xfa\x8d\x4f\x98\x2d\x19\xe3\x2b\x4c\x84\x56\x7c\xc2\x78\x22\xf5\x1b\x48\xfe\x7f\xd5\x9f\xf9\xb6\x17\xe9\x48\x6b\x35\x2c\x8a\x99\x5f\x96\x17\x68\x09\x29\xde\x05\x37\xe7\xcc\x81\x39\x9a\x93\x88\xf0\xd2\x0e\x8d\xd0\x46\xd0\x99\x4f\xd8\xf3\xa8\xee\x4d\xa5\x9c\x52\xa0\x22\xcb\xdf\x83\xcc\xb1\xee\xfb\x40\x10\xa1\x22\x34\x39\x9f\xb0\xdf\xae\xc9\x18\x8f\xa3\x31\x77\xf5\xae\xe6\x79\x5a\xe5\x64\x40\x28\x6a\xf2\x8a\xfa\x9b\x31\x3e\xdf\x80\x49\x90\xac\x9d\xff\x0a\x20\x32\xc3\x03\x1a\x85\x72\xa8\x20\xc5\x8b\xe3\x2a\x85\x2b\x53\x0a\x75\xfc\x68\x62\x41\x86\x06\x54\x6c\xb1\x6f\xbc\x6e\x97\x2d\x23\x1b\xc8\x0f\x3f\x8d\x3e\x66\x77\x7c\x3c\x0a\xac\x9e\xe4\xa8\xac\x78\x51\xb8\xaa\x2c\x5b\x58\x73\x4a\x75\x94\xb2\x01\x5a\xe5\xa6\x68\x5b\xb8\x4f\x9c\x20\x3f\xdc\x3a\xa8\x72\x37\xe2\x84\xc6\x72\x63\x1d\x1d\xd0\x74\x71\x4f\xab\xbd\x48\x3e\x57\xa7\x81\x88\x14\x12\x27\x52\x14\x33\x5b\xde\x9c\xcd\xac\x7e\x9a\xd6\x31\xca\x9c\x3d\x3d\xb1\xec\x3d\xe6\x2d\x52\xd9\xd1\x7b\x51\xa7\x3e\x31\x2f\xf0\x82\x60\xf9\xba\x7a\x09\x83\xd7\xed\x6a\xbe\xae\x74\x3d\x37\xde\x0a\x33\xbd\x5d\xcd\xcb\xf2\xf1\xc9\xeb\x6a\x39\x5d\x72\xbb\x0e\xb6\x40\x02\x2b\xdb\x7d\x66\xae\x93\xfb\x01\x2c\xe0\xe3\x87\x90\xee\xc0\xf1\xe8\x66\xf4\x1a\x5d\x8b\xbf\xb8\xf8\xee\x28\x0f\xdd\x6d\x30\xcd\x24\x50\xbf\xbd\xa9\x21\xb1\x87\x88\x7a\x41\x5f\xe4\x19\x50\xf4\x27\x84\xb3\xd4\x10\xbb\x31\x1a\x8c\x6b\xbd\x6b\xee\x3e\xe6\xfa\x68\xaa\x34\x9a\x83\x71\x2f\xdc\x56\x3b\x19\x6e\xcb\x96\x47\xbe\xc0\x54\x9b\xb3\x9b\xa5\x28\x2e\x3f\x3a\x1c\x5f\xe4\x87\x9a\x51\x95\x1d\x7c\x16\x84\xf6\xd5\xb5\xd2\xe2\x4b\xa4\x77\x6d\xaa\x95\xde\x0d\x7a\x02\xb2\x6e\x09\x0c\x85\x5a\x8a\xe8\xdc\x75\x3c\xb3\x97\xc6\x09\x24\x9f\xb0\xaf\xa3\xeb\xa7\x25\x31\x25\xc2\x34\xa3\xdb\x57\xc5\x7d\x94\x60\x8f\xfc\xf2\xdc\xff\xe4\x42\xc7\x6e\xc9\x63\xc7\xeb\x75\xf7\xb9\x3c\xbd\xf7\xc4\x36\x8b\x81\xb0\x61\x9a\xaf\x09\x92\xc4\xfd\xe7\xc6\xb7\x76\xed\xde\x84\x60\x40\x4a\xb4\x13\x8d\x07\x97\x13\xcb\x41\xf9\x2f\x00\x00\xff\xff\xcc\x71\x51\x26\xc9\x05\x00\x00")

func dataCocoonJobJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataCocoonJobJson,
		"data/cocoon.job.json",
	)
}

func dataCocoonJobJson() (*asset, error) {
	bytes, err := dataCocoonJobJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/cocoon.job.json", size: 1481, mode: os.FileMode(420), modTime: time.Unix(1486976539, 0)}
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
	"data/bindata.go": dataBindataGo,
	"data/cocoon.job.json": dataCocoonJobJson,
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
	"data": &bintree{nil, map[string]*bintree{
		"bindata.go": &bintree{dataBindataGo, map[string]*bintree{}},
		"cocoon.job.json": &bintree{dataCocoonJobJson, map[string]*bintree{}},
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

