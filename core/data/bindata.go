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

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486986386, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cocoonJobJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x54\x4d\x6f\xdb\x30\x0c\xbd\xf7\x57\x08\xc2\x8e\x43\xd1\x6c\xe8\x25\xc0\x0e\x99\xdd\x0d\x19\x92\xda\xc8\xc7\x69\x08\x0a\xd6\x66\x3d\x21\xb2\x64\xc8\x4c\xd6\xcc\xd0\x7f\x1f\x24\x3b\x8d\xbf\x12\x5f\xc2\xf0\x3d\xf3\x3d\xd2\x94\xaa\x3b\xc6\xf8\x2f\xfd\xca\xa7\xcc\x85\x8c\xf1\x15\x66\x42\x2b\x3e\x65\x3c\x93\xfa\x15\x24\xff\x5c\xe7\xe7\xa1\xcb\x25\x3a\xd1\x5a\xdd\x57\xd5\x3c\xb4\xf6\x0c\x3d\x43\x8e\x57\xc1\xcd\xa9\xf0\x60\x89\xe6\x28\x12\x3c\xa7\x63\x23\xb4\x11\x74\xe2\x53\xf6\xf8\xd0\xe4\x66\x52\xce\x28\x52\x89\xe3\xbf\x81\x2c\xb1\xc9\x87\x40\x90\xa0\x22\x34\x25\x9f\xb2\xdf\x3e\xc9\x18\x4f\x93\x09\xf7\xf1\xae\xe1\x05\x5a\x95\x64\x40\x28\x6a\xf3\xaa\xe6\x97\x31\xbe\xd8\x80\xc9\x90\x9c\x9d\x4f\x15\x10\x99\xfb\x3d\x1a\x85\xf2\x5e\x41\x8e\x67\xc7\xf5\x14\x2e\x4c\x29\xd4\xe1\xbd\x8d\x45\x05\x1a\x50\xa9\xc3\xbe\xf1\x26\x6d\x3b\x46\x36\x50\xee\x7f\x1a\x7d\x28\xae\xf8\xb8\x35\xb0\xa6\x93\x83\x72\xe2\x55\xe5\x23\x6b\x3b\x58\xbb\x4b\x75\x90\xb2\x05\x3a\xe5\xb6\x68\x57\x78\x4c\x9c\xa0\xdc\x0f\x1d\xd4\x73\x37\xe2\x88\xc6\x71\x53\x9d\xec\xd1\xf4\xf1\x40\xab\x37\x91\x7d\xac\x4e\x0b\x11\x39\x64\x5e\xa4\xaa\xe6\x2e\x1c\xd4\x66\x4e\x3f\xcf\x9b\x31\xca\x72\x08\x83\xc9\xfa\x8d\xd4\xcf\xae\x93\xb1\x3d\x4f\x4f\xea\x38\x66\x28\x88\x82\x28\x7a\x7e\x59\x3d\xc5\xd1\xcb\x76\xb5\x58\xd7\xde\x02\x3f\x82\x15\x16\x7a\xbb\x5a\x58\xcb\x6f\x56\x5e\xd7\x0b\xec\x4d\xed\x7a\xd8\x12\x09\x9c\x6c\xff\x9d\x85\xce\xae\x0f\x69\x09\xef\x3f\x84\xf4\x05\x27\x0f\x83\xfe\x1b\x74\x2d\xfe\xe1\xf2\xbb\xa7\xdc\x74\xb7\xc1\xbc\x90\x40\xe3\xf6\x66\x86\xc4\x1b\x24\x34\x0a\x86\xa2\x2c\x80\x92\x3f\x31\x9c\xa4\x86\xd4\xb7\xd1\x62\x5c\xe2\x5d\xfb\x7c\x60\xa9\x0f\xa6\x9e\x46\xbb\x31\x1e\xc4\xdb\x7a\x6f\xe3\xad\xed\x78\xe4\x4b\xcc\xb5\x39\xf9\x5e\xaa\xea\xfc\xa7\xc7\x09\x45\xb9\x6f\x18\x75\xd8\xc3\xe7\x51\xec\x3e\x5d\x67\x5a\xfc\x19\xe9\xaf\x36\xf5\xda\x5f\x96\xc3\x76\xdd\x12\x18\x8a\xb5\x14\xc9\xa9\xef\x78\xee\x2e\x96\x23\x48\x3e\x65\x5f\x1f\x2e\x4f\x47\x62\x46\x84\x79\x41\xc3\x4f\xc5\x43\x94\xe0\x4a\x7e\x79\x1c\x7f\x73\xa9\x53\x7f\x10\x52\xcf\x1b\x75\xf7\xb1\x3c\xa3\x77\xc9\xb6\x48\x81\xb0\x65\x9a\xaf\x09\xb2\xcc\x9f\xcb\xc9\xd0\xae\xdb\x9b\x18\x0c\x48\x89\xae\xa3\xc9\xdd\xb9\xa2\xbd\xb3\xff\x03\x00\x00\xff\xff\xbe\x2e\x71\x32\xed\x05\x00\x00")

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

	info := bindataFileInfo{name: "cocoon.job.json", size: 1517, mode: os.FileMode(420), modTime: time.Unix(1486986369, 0)}
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

