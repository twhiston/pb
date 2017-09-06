// Code generated by go-bindata.
// sources:
// tpl/block.xml.tpl
// tpl/schema.yml.tpl
// DO NOT EDIT!

package pb

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

var _tplBlockXmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xdf\x6b\xfa\x30\x10\x7f\xfe\xfa\x57\xdc\x37\xf8\x38\x8c\xbe\x8e\xb6\x50\xd6\x0d\x06\xb2\xc1\x7e\xbd\x88\x8c\xd8\x9e\x2e\xd8\x26\x52\x53\x47\x29\xf9\xdf\xc7\x25\x8d\xba\xa1\x93\xbd\x25\x77\x9f\x5f\xdc\x5d\xb4\x28\x75\xbe\x06\x25\x2a\x8c\x59\xd7\xc1\x88\x5e\x60\x2d\x4b\x06\x00\x00\x51\x2e\x0c\xae\x74\xdd\x26\xd4\x0b\x1f\xb0\x36\xe2\xfb\x8e\x07\x4a\xb5\xd4\x0e\x44\x0f\x07\x70\x15\xdf\x2c\x84\x11\x47\x1e\xce\xf3\x41\x54\x38\xd5\x9f\x58\x1f\xdc\xba\x0e\x6a\xa1\x56\x08\xc3\x35\xb6\x57\x30\xdc\x89\xb2\x41\xb8\x8e\x49\x74\xd3\x98\x2d\x58\xeb\x70\x4e\x73\x27\x6a\x29\x16\x25\xc2\x56\xe7\x6b\x34\x31\x93\x8a\x01\x79\xc6\xec\x59\xae\x94\x28\xc1\x91\x18\xf0\xe4\x02\x09\x0b\x69\xa8\x18\xb3\x09\x03\xe7\x19\xb3\xc9\x78\x1c\xd4\x32\x2c\x45\x0b\x46\x56\x08\x52\x41\xb5\xdd\x2b\x76\x1d\xa0\x2a\x42\xa8\xf3\xe1\x75\x63\x2e\xa7\xd7\x14\xd5\x1b\x7a\xea\x9f\x6d\x76\xa2\x3e\xe3\x51\x98\x76\xe3\x47\xdf\xc3\xad\xa5\x35\x50\x95\x86\x7f\xd8\xcc\xb7\x76\xb8\x84\xe3\x01\x72\xfe\xf8\x04\xf7\x77\xf0\x96\x4e\x5f\x6f\x21\xcd\xb2\x30\xaf\x31\x3b\x91\x35\xe2\xb4\xf8\xfe\x06\x96\x8d\xca\x8d\xd4\xaa\x77\x5b\xbe\xff\x76\x09\x8e\xf1\x7f\x76\x93\xa5\x2f\xe9\x2c\x28\x8f\xf6\x12\xd6\x0e\xfe\xcd\xe7\xbd\x30\x0f\xe5\xfe\xff\x81\xe5\xe6\x94\x0a\x29\x50\xef\x07\xdb\xc3\x23\xee\xa2\x24\x5f\x01\x00\x00\xff\xff\x2a\x60\x1b\x69\x10\x03\x00\x00")

func tplBlockXmlTplBytes() ([]byte, error) {
	return bindataRead(
		_tplBlockXmlTpl,
		"tpl/block.xml.tpl",
	)
}

func tplBlockXmlTpl() (*asset, error) {
	bytes, err := tplBlockXmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/block.xml.tpl", size: 784, mode: os.FileMode(420), modTime: time.Unix(1504557948, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplSchemaYmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\xb1\x6e\x83\x30\x10\x86\x77\x9e\xe2\x24\x32\x96\xb0\x23\x75\x69\x52\xb5\x5d\xd2\x2c\xcd\x6e\xe0\x82\x51\x88\x6d\xc1\x91\x08\x21\xde\xbd\xf2\xd9\xb8\x41\xad\x32\xa4\x4c\xf6\x7f\x9f\xcf\x9f\x74\x38\x7e\xf0\x8b\x62\xd8\x0b\x2a\xe4\x4b\xa3\x8b\x53\x06\xe3\x08\xeb\x9d\x38\x23\x4c\x53\x14\xf3\xee\xab\xc3\x56\xb9\x04\x12\x4e\xb6\x82\x70\xa7\xaf\x8c\x44\x31\xbc\xa1\xc2\x56\x10\x96\x90\x0f\x60\xf2\x0a\x15\x48\x22\xd3\x65\x69\x5a\xd5\x24\xfb\x7c\x5d\xe8\x73\x4a\x57\x59\x77\xa4\x55\xca\x44\xf4\xb0\xae\x55\x59\x6a\x16\x82\xb0\xd2\xed\xe0\xd2\x8d\xdf\xd9\x4a\xad\x8e\xda\xa5\x1f\xea\xa8\x6d\x32\x8e\x50\x1f\xed\xd6\xf4\xd4\x41\xc2\x8c\x5d\x66\xd1\x38\x26\xd0\x0a\x55\x21\xac\x4e\x38\x3c\xc1\xea\x22\x9a\x1e\x21\x7b\x0e\xf4\x34\x45\x00\x09\x84\xa6\x8e\x08\xad\x01\x00\x38\xb9\x2d\x1e\xb8\x89\xaf\x62\x59\x93\xc8\x9b\x05\xf0\xea\x33\x2f\x87\xaa\x64\xab\xe5\xd2\x2a\x7f\xf6\x14\x9c\xb5\x5b\xdf\x95\x9e\xf9\x7f\x59\xdf\x35\x3a\x88\xd6\xe9\x5c\x44\x7b\xdf\x85\x49\x2f\x52\xd2\x60\x16\x97\x6d\x6d\x30\xab\x84\xe9\xfa\xe2\x3c\xe3\xbf\x45\x24\x36\xc6\xcd\xf7\x1d\x1b\xc3\xff\x82\x2e\xd1\xa9\x5c\x6b\x92\xb0\xde\xe8\xd2\x9f\x4f\xd8\x79\x2f\x48\xba\xbb\x8c\x20\xe9\xce\xce\x99\x65\x6c\xe7\x1b\x9c\x1f\x85\xe3\xf3\x9f\xf7\x11\xd2\xdf\x56\xd3\xf4\x1d\x00\x00\xff\xff\x31\x74\xa8\xfd\x88\x03\x00\x00")

func tplSchemaYmlTplBytes() ([]byte, error) {
	return bindataRead(
		_tplSchemaYmlTpl,
		"tpl/schema.yml.tpl",
	)
}

func tplSchemaYmlTpl() (*asset, error) {
	bytes, err := tplSchemaYmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/schema.yml.tpl", size: 904, mode: os.FileMode(420), modTime: time.Unix(1504731659, 0)}
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
	"tpl/block.xml.tpl":  tplBlockXmlTpl,
	"tpl/schema.yml.tpl": tplSchemaYmlTpl,
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
	"tpl": &bintree{nil, map[string]*bintree{
		"block.xml.tpl":  &bintree{tplBlockXmlTpl, map[string]*bintree{}},
		"schema.yml.tpl": &bintree{tplSchemaYmlTpl, map[string]*bintree{}},
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
