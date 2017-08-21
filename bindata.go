// Code generated by go-bindata.
// sources:
// assets/icon128.png
// assets/mac/icon.png
// DO NOT EDIT!

package timer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
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

var _assetsIcon128Png = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x80\x00\x00\x00\x80\x01\x03\x00\x00\x00\xf9\xf0\xf3\x88\x00\x00\x00\x06PLTE\xff\xff\xff\x00\x00\x00U\xc2\xd3~\x00\x00\x00\x01tRNS\x00@\xe6\xd8f\x00\x00\x00cIDATH\xc7c`\x18\x05\x04\x01\xfb\u007f(x@S\x81\a\xa84\xad\x04\xeaA\x96\xfe\x1b\x1e\x02\x8c \x81\x06\xda\n0\xc87\xf0\xff\xff\xff\x81\xf1\a\x035\x05\x18\xea\x18\xec\xff\x00\t\x1a\v0~``\xe0o\xa0\xb5@\x1d\x9c\xa0\xa1\x80<\xd8F\x940\xa5\xbe\x00\xbd2\xe1h\xce& @\x97\x10\xa3K\xb9>\n\xf0\x00\x00C\x00\xa5\x1b3\x96\xd2,\x00\x00\x00\x00IEND\xaeB`\x82")

func assetsIcon128PngBytes() ([]byte, error) {
	return _assetsIcon128Png, nil
}

func assetsIcon128Png() (*asset, error) {
	bytes, err := assetsIcon128PngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/icon128.png", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsMacIconPng = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x16\x00\x00\x00\x16\b\x04\x00\x00\x00n\xbd\xa4\xb0\x00\x00\x00LIDAT(\xcfc\xf8\xcf@<d\xa0\x8bb\\\x00\x87bl\xe6\xe1Q\x8ci*\xb5\x143`\xb8\xf6?\xae\xd0@U\x8e\xe1\x0f\xec\xde\xc1\xaa\x14[8\xff\xc7\x1d:x\x02\x8a\x90\xc9\xff\xff\xe3\xd5LQ\xa4\f\x9a\x18$\xda\x19$$\xd1\x01\xccV\x00\x15\x82A\xdb|\x1a\x87e\x00\x00\x00\x00IEND\xaeB`\x82")

func assetsMacIconPngBytes() ([]byte, error) {
	return _assetsMacIconPng, nil
}

func assetsMacIconPng() (*asset, error) {
	bytes, err := assetsMacIconPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/mac/icon.png", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"assets/icon128.png": assetsIcon128Png,
	"assets/mac/icon.png": assetsMacIconPng,
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
	"assets": &bintree{nil, map[string]*bintree{
		"icon128.png": &bintree{assetsIcon128Png, map[string]*bintree{}},
		"mac": &bintree{nil, map[string]*bintree{
			"icon.png": &bintree{assetsMacIconPng, map[string]*bintree{}},
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

