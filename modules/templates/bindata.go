// Code generated by go-bindata.
// sources:
// ../../templates/add.html
// ../../templates/addsuccess.html
// ../../templates/base/base.html
// ../../templates/base/common.html
// ../../templates/base/common_noleft.html
// ../../templates/base/footer.html
// ../../templates/base/left_engines.html
// ../../templates/base/nav.html
// ../../templates/chgpass.html
// ../../templates/delsuccess.html
// ../../templates/login.html
// ../../templates/root.html
// DO NOT EDIT!

// +build bindata

package templates

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

var _TemplatesAddHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\xcd\x8a\xe4\x36\x10\x3e\xdb\x4f\xa1\x28\x03\xb6\xc9\xd8\xde\x65\x21\x84\x1e\x77\x43\x0e\x09\x39\x04\x12\xc8\x1e\x02\xcb\x30\xc8\x56\xb9\xad\x41\x96\x8c\x24\xf7\xee\xc4\xf8\x09\x72\xce\x73\xe4\x94\x57\x0a\x79\x8c\xa0\x1f\x77\xb7\xbd\x90\xec\xcf\x5c\x6c\x55\xd5\x57\x55\x9f\x4a\xa5\xd2\x34\x19\xe8\x07\x4e\x0c\x20\x5c\x13\x0d\xa5\xfd\x14\x9d\xe9\x39\x46\xc5\x3c\xc7\xef\xd9\x1b\xd9\xf7\x52\x5c\x21\xe2\x69\xa2\xd0\x32\x01\x08\x77\x40\x28\x9e\xe7\xb8\xd2\x8d\x62\x83\x41\xe6\x69\x80\x3d\x36\xf0\xce\x94\x8f\xe4\x44\xbc\x16\x1f\xe2\xe8\x26\xc5\x4c\x0c\xa3\x79\x33\x70\xd2\x40\x27\x39\x05\x75\x8f\xb3\x02\x48\xd3\xa5\xed\x28\x1a\xc3\xa4\x40\x69\x86\x26\x74\x93\x9a\x8e\xe9\xac\x20\xc6\xa8\x34\xd1\xec\x37\x48\x6e\x37\xca\xab\x20\x49\x56\x70\x10\x47\xd3\x65\x77\x68\xce\xee\xe2\x38\x3a\x47\x33\xa0\x4d\x9a\x4d\x71\x14\xdd\x14\xe4\x91\xbc\x4b\xed\x32\x1a\x15\xdf\x25\xa5\xb5\x25\xb7\x56\xa6\xc4\x90\x1d\x72\xa6\x28\x11\xa4\x87\x64\x87\x6e\x52\xfc\xa5\x5d\xe2\xac\xb0\x7b\x49\xb3\x5b\x6f\xee\xa4\x36\xc1\x6c\x97\x5b\xf3\xa8\x41\x5d\x45\x58\xc4\x2d\x8c\xd6\x57\x20\x2f\x6c\x21\x03\xd1\xfa\x2d\x0d\x10\x2f\xac\x21\xf3\x99\xfb\xeb\xa7\x01\x76\x28\x79\xd4\x52\xf8\x0d\xe9\xb1\x69\x40\xeb\x1d\x5a\x0a\x91\x36\x52\x18\x10\x26\x0b\xbb\x64\x2d\x5a\x54\x6f\x12\x50\x2a\xb9\x5f\x2c\x11\xe1\xa0\x4c\x8a\x5b\xc2\x38\xd0\x1d\xc2\x5f\x6d\x70\x0e\x35\x23\xe0\x1a\x82\x4b\xf0\x38\x27\xbd\x72\xd1\x86\x98\x51\x27\xf7\xd9\x9d\x63\xbc\x7c\xe6\x2c\x8e\xe6\xb8\x2a\x7d\x73\x1c\xe2\x69\x02\x41\xd7\x6d\x55\x4b\xfa\xe4\xda\x8a\xb2\x13\x6a\x38\xd1\x7a\x8f\x1b\xc9\xf3\x9e\xe6\x2f\x5f\x60\xeb\xc2\x5a\x54\xb4\x9c\xe8\xae\xf8\x4e\xa9\x6f\x29\xdd\x80\x6d\xa5\xf2\x06\x84\x01\x65\xbb\xaf\x1a\x16\x43\x7d\xcc\x41\x29\xe9\xb4\xd3\xb4\x0d\x11\x55\xe5\x70\x88\xab\x92\xb2\xd3\x85\x57\xd5\x4a\xd5\xa3\x1e\x4c\x27\xe9\x1e\xff\xfc\xd3\x2f\xaf\x31\x52\x92\xc3\x1e\x5b\x43\x08\xf4\xab\x56\xed\xf7\x52\xf5\x3f\x98\x9e\xbb\x40\x57\x64\x2c\x2c\x3f\x2a\x39\x0e\x1e\xcc\x5e\x7e\x23\x50\xf1\x23\x11\x47\x84\x69\xfd\xe0\xce\x7f\x9e\x2b\x77\x3f\x10\xa3\x7b\xec\x34\xc8\x7e\x97\xf5\x55\xc3\xef\xf1\x3a\x82\x73\x7b\xb8\xc4\xc1\xa5\xdd\xb0\xdf\xc2\xc7\xd0\xb0\x57\x17\x3b\xea\x1a\x38\x34\x9e\x09\x55\xec\x04\x6a\xe1\x12\xa4\x43\x1c\x45\xd3\xa4\x88\x38\x02\x2a\x68\xad\xad\x53\x54\xc9\xc1\x5d\xba\x13\xe1\x23\x58\x8e\xc5\x3c\xe3\x83\xfb\x55\xa5\xb7\x79\x3f\x5f\xd4\xa8\x2a\x7d\x9a\x4f\x21\x4b\x0c\xb1\x93\xe9\xc1\xdd\xc1\x55\xe5\x9c\x26\xb0\xf5\xeb\xff\xa8\xdc\x2a\xce\x43\xc7\x84\xf9\xe4\xf2\x2d\x91\x06\xa9\x36\x8c\x9c\x26\x30\xf2\xeb\x50\xa1\x57\xaf\x5e\x7c\xbd\xa1\xe7\x54\x9f\x95\x3f\x4c\x93\x15\x83\xa0\x5b\xce\xb0\xfe\xdf\x8e\xda\x44\x7b\x9e\xca\x9c\x87\xe1\x8a\xdb\x59\x1b\xd8\x5d\xe4\x0f\xe1\xb7\xa0\x9f\xe9\xec\xec\x98\x95\x8a\x6e\xce\xcf\x0f\xdf\xf0\xb2\x9d\x31\xcb\x89\x06\xeb\x87\xb0\x5d\x7c\x3f\x92\x6d\xe0\xe2\xf3\xeb\xb1\xee\xd9\xa5\x87\xd6\x99\x34\x39\xb9\x01\x70\x9e\x75\x46\xa0\xda\x88\x7c\x50\xac\x27\xea\x09\x23\x97\xf0\x8b\x3c\xaf\x08\xea\x14\xb4\x7b\x7c\x79\xa3\x77\xfe\xb1\xbc\x7b\xcf\xf9\x2d\x51\x82\x89\x23\x3e\xfc\xfd\xd7\xef\xff\xfc\xf9\x47\x55\x92\x43\x9e\x5f\x88\x57\xa5\x65\x7b\x9e\x99\x55\x50\x5e\x0f\xd0\x7f\x03\x00\x00\xff\xff\x62\x34\xe5\x37\x71\x08\x00\x00")

func TemplatesAddHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesAddHtml,
		"../../templates/add.html",
	)
}

func TemplatesAddHtml() (*asset, error) {
	bytes, err := TemplatesAddHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/add.html", size: 2161, mode: os.FileMode(420), modTime: time.Unix(1490778536, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesAddsuccessHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x4a\x4a\x2c\x4e\xd5\x07\x11\x7a\x19\x25\xb9\x39\x4a\x0a\x7a\xb5\xb5\x5c\x18\xf2\xc9\xf9\xb9\xb9\xf9\x79\x48\x2a\xb8\xaa\xab\x53\x52\xd3\x32\xf3\x52\x15\x94\x32\x52\x13\x53\x94\xc0\x9a\x52\xf3\x52\x50\xa5\x92\xf2\x53\x2a\x41\x52\x36\x29\x99\x65\x76\x5c\xd5\xd5\x99\x86\x16\x79\x0a\x7a\x3e\x89\x79\xe9\x0a\x4a\x89\x29\x29\xf1\x29\x49\xf1\xc5\xa5\xc9\xc9\xa9\xc5\xc5\x60\x55\xfa\x50\x65\x60\x73\x00\x01\x00\x00\xff\xff\xce\x6b\x05\x2f\xa4\x00\x00\x00")

func TemplatesAddsuccessHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesAddsuccessHtml,
		"../../templates/addsuccess.html",
	)
}

func TemplatesAddsuccessHtml() (*asset, error) {
	bytes, err := TemplatesAddsuccessHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/addsuccess.html", size: 164, mode: os.FileMode(420), modTime: time.Unix(1490778555, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesBaseBaseHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\xb1\x6e\x23\x21\x10\xad\xcd\x57\x70\xd4\xb7\x46\xa7\x6b\x81\xc6\x77\x6d\x92\x22\x4d\x4a\x0c\xb3\xd9\x71\x58\xd8\x30\xe3\xb5\x1c\xcb\xff\x1e\xe1\x75\x64\xbb\x1a\xcd\x7b\x8f\xc7\x7b\x1a\xf3\xeb\xdf\xf3\xe6\xf5\xed\xe5\xbf\x1c\x78\x4c\x4e\x98\x36\x64\xf2\xf9\xdd\xaa\xaf\xa1\xdb\x3c\xa9\x86\x81\x8f\x4e\x98\x11\xd8\xcb\xec\x47\xb0\x6a\x46\x38\x4c\xa5\xb2\x92\xa1\x64\x86\xcc\x56\x1d\x30\xf2\x60\x23\xcc\x18\xa0\xbb\x2c\xbf\x25\x66\x64\xf4\xa9\xa3\xe0\x13\xd8\x3f\xcd\x2b\x61\xfe\x90\x15\x92\x55\xc4\xc7\x04\x34\x00\xb0\x92\x43\x85\xde\x2a\x4d\xec\x19\x83\xde\x96\xc2\xc4\xd5\x4f\x7f\x75\x20\xba\xad\xeb\x11\xf3\x3a\x10\x35\x1f\x0a\x15\x27\x96\x54\xc3\xed\xdd\x8e\xf4\xee\x73\x0f\xf5\x78\x11\xee\x48\x39\xa3\x17\x9d\x13\xa7\x13\xc3\x38\x25\xcf\x20\x55\xeb\xa3\xe4\xfa\x7c\x16\x46\x5f\xbb\x6d\x4b\x3c\x3a\xf1\xa0\xca\x7e\x5e\x44\xc2\x44\x9c\x65\x48\x9e\xc8\xaa\xd6\xd7\x63\x86\xda\xf5\x69\x8f\x51\x39\xb1\xba\xa7\x6b\x39\x34\x68\x75\x6f\x94\xa0\xe7\xab\xd3\x23\xd1\x3e\x5d\x88\x95\xd1\x11\x67\x27\xae\xe3\x21\x47\x5f\x0a\x43\xfd\xc9\xbb\x04\x35\x7a\x39\xd7\x77\x00\x00\x00\xff\xff\xfd\x58\x70\x3d\xbf\x01\x00\x00")

func TemplatesBaseBaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesBaseBaseHtml,
		"../../templates/base/base.html",
	)
}

func TemplatesBaseBaseHtml() (*asset, error) {
	bytes, err := TemplatesBaseBaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/base/base.html", size: 447, mode: os.FileMode(420), modTime: time.Unix(1452220816, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesBaseCommonHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x4b\x2c\x53\xaa\xad\xe5\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x4a\x4a\x2c\x4e\xd5\xcf\x4b\x2c\xd3\xcb\x28\xc9\xcd\x51\x52\xd0\x03\x4b\xa7\xe6\xa5\xd4\xd6\x72\x71\x21\x34\xe6\xa4\xa6\x95\x60\xd5\x09\x92\x88\x4f\xcd\x4b\xcf\xcc\x4b\x2d\xc6\x6f\x44\x5a\x7e\x7e\x49\x6a\x11\x56\x43\x20\x52\x98\xda\x01\x01\x00\x00\xff\xff\xfe\xa6\x85\xb9\xb9\x00\x00\x00")

func TemplatesBaseCommonHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesBaseCommonHtml,
		"../../templates/base/common.html",
	)
}

func TemplatesBaseCommonHtml() (*asset, error) {
	bytes, err := TemplatesBaseCommonHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/base/common.html", size: 185, mode: os.FileMode(420), modTime: time.Unix(1490778494, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesBaseCommon_noleftHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x4b\x2c\x53\xaa\xad\xe5\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x4a\x4a\x2c\x4e\xd5\xcf\x4b\x2c\xd3\xcb\x28\xc9\xcd\x51\x52\xd0\x03\x4b\xa7\xe6\xa5\xd4\xd6\x72\x71\x21\x34\xe6\xa4\xa6\x95\x28\x61\x97\x4a\xcb\xcf\x2f\x49\x2d\xc2\x6a\x2c\x44\x0a\xd3\x64\x40\x00\x00\x00\xff\xff\xfd\xae\x3a\x68\x91\x00\x00\x00")

func TemplatesBaseCommon_noleftHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesBaseCommon_noleftHtml,
		"../../templates/base/common_noleft.html",
	)
}

func TemplatesBaseCommon_noleftHtml() (*asset, error) {
	bytes, err := TemplatesBaseCommon_noleftHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/base/common_noleft.html", size: 145, mode: os.FileMode(420), modTime: time.Unix(1490778570, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesBaseFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x4f\x4b\x33\x31\x10\x06\xf0\xf3\xdb\x4f\x31\xa4\xf0\x9e\xec\x46\x51\x50\xf6\xdf\xb5\x37\xf1\x20\xe2\x4d\x92\x4d\x9a\x86\x66\x67\xc2\x6c\xda\x6e\x29\xfd\xee\x92\x55\x7a\x59\x11\x79\x0e\xe1\x09\xe1\x97\x61\x6a\xe3\x0f\xd0\x05\x35\x0c\x8d\x60\x3a\x42\xb2\x63\x5a\x75\x16\x93\x65\x01\x43\x3a\x05\xdb\x88\x5e\xb1\xf3\xb8\x4a\x14\xcb\xfb\xdb\x38\x56\x5a\x75\x3b\xc7\xb4\x47\xb3\xea\x28\x10\x97\xcb\xcd\x53\x4e\xa5\x89\x8d\xe5\xe9\xe1\x5d\x1c\x61\xa0\xe0\x0d\x2c\xed\x63\x4e\x15\x95\x31\x1e\x5d\xf9\x90\x09\xd1\x2e\xf2\xcf\xed\xe2\xdf\x0b\x1d\x2d\x5b\x03\xfa\x04\xb5\x82\x2d\xdb\x4d\x23\xb6\x29\xc5\x52\x4a\x47\x41\xa1\x2b\x88\x9d\x80\xa4\xd8\xd9\xd4\x88\x0f\x1d\x14\xee\x44\xbb\xa6\x5a\xaa\x16\x0e\xe7\x73\xb1\xa6\x37\xcb\x97\xcb\x0d\xfc\x47\x3d\xc4\x6a\xce\xf8\xb4\xdd\xeb\xa2\xa3\x5e\x86\x3d\xe2\x49\x26\x85\x8e\xe6\xe4\x6b\xbe\xbe\xaa\x53\xfb\x1d\x1e\x89\xfb\xc2\xff\x20\xbd\x13\xf7\x57\x28\x97\xc9\xf9\x56\x40\xa1\xf9\xf3\xa8\x48\x46\xcf\xfd\x67\x32\xfa\xea\xe7\x32\xf9\x8b\x5a\x4e\x2b\xfd\x3a\x3e\x03\x00\x00\xff\xff\xbc\xc6\x16\x4e\xd9\x01\x00\x00")

func TemplatesBaseFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesBaseFooterHtml,
		"../../templates/base/footer.html",
	)
}

func TemplatesBaseFooterHtml() (*asset, error) {
	bytes, err := TemplatesBaseFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/base/footer.html", size: 473, mode: os.FileMode(420), modTime: time.Unix(1423189721, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesBaseLeft_enginesHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\xcf\x6a\xe3\x30\x10\xc6\xcf\xf6\x53\x08\xad\xd9\x5b\x6c\x76\xc9\x61\x59\x64\x97\x1e\x7a\x28\x94\xbe\x42\x91\x32\x8a\x33\xa0\xa8\xd4\x56\x92\xc2\xe0\x77\x2f\x92\xfc\x47\x36\x29\xbd\xd9\xa3\x99\x6f\xbe\x99\xdf\x08\xc0\x2b\x3b\x18\xd9\xf7\x35\x3f\xbc\x9b\xdd\x19\x76\x7f\x79\x93\x67\xe2\xb4\x6f\x88\xf0\xcf\x3f\xcb\xca\x17\x69\x5b\xc6\xa5\x31\x6f\xa0\xf8\x30\x88\xea\xb4\xf7\x19\x17\xd3\xe4\x19\x51\x61\xfc\xf3\xff\x3a\xe6\x0d\x43\x9e\x09\x83\x4d\x9e\x65\x44\x78\x64\x25\xf6\x8f\x00\x3e\x9a\x6d\xe4\x00\xa2\x5c\x78\xd1\xa6\xd7\xe1\x53\x48\x76\xea\xf4\xb1\xe6\x95\x04\x50\x7c\xeb\x61\x2a\x12\x95\x8c\x2d\xb4\x0d\xe2\xa2\x0a\x3d\x89\x8a\xc3\xa5\x7b\xb2\x2d\x5a\x1d\x2c\xcd\x7f\x3e\x89\xa8\x93\xb6\xd5\xac\xd4\x21\xd4\x2f\x5e\x83\xd5\x62\x95\x1c\x63\xfa\x23\x09\xb3\xf2\x55\x9e\xe3\x63\x46\xb4\xfc\x24\xf6\x13\xff\x57\xd4\xb7\x07\x84\x9a\xa8\x7c\x86\x61\xf0\xa3\x8c\x25\x5b\xf3\x77\xc7\xff\xa9\x7c\x19\x7d\x2a\xf9\xc5\x27\x90\xca\x59\xa6\x9c\xdd\xdd\x64\x67\xd1\xb6\xe1\xfb\xb3\x9f\x97\x19\x89\x71\x0d\xe8\xe6\x55\x2e\x8d\x41\x9b\xb4\xef\x56\x13\xfc\x0a\xbb\x6f\x24\x41\x1b\xed\xf4\x22\x3a\x51\x99\x29\xf9\x9b\x99\x8e\x2b\x85\xb3\xba\x29\x8f\x11\x21\xf0\xc3\xe9\x74\x0a\xa7\x42\xc0\xa9\x31\x30\xa2\x74\x52\x99\x48\x72\x21\x16\xf6\xc4\x8a\x31\x35\x20\x4e\x96\x17\x8f\x73\x0d\xcd\x60\x73\x6f\xf3\x85\x6f\xff\xdb\xa9\x7a\x2e\xdf\x60\x48\xd5\xec\x64\x75\x35\xac\xa8\x00\xaf\xcd\x57\x00\x00\x00\xff\xff\xdd\x8a\xbe\x82\x66\x03\x00\x00")

func TemplatesBaseLeft_enginesHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesBaseLeft_enginesHtml,
		"../../templates/base/left_engines.html",
	)
}

func TemplatesBaseLeft_enginesHtml() (*asset, error) {
	bytes, err := TemplatesBaseLeft_enginesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/base/left_engines.html", size: 870, mode: os.FileMode(420), modTime: time.Unix(1423225846, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesBaseNavHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\xbb\x8e\xdc\x3a\x0c\x86\xfb\x79\x0a\x1e\x6e\xad\x35\xb6\x3b\x85\xec\x62\x91\x26\xc0\x96\x01\x52\xd3\x16\x2d\x13\xd0\x48\x86\x24\x3b\x19\x18\xf3\xee\x81\xe3\xcb\x38\xce\xe6\xd2\xa4\x30\x2c\x52\xa4\xfe\x8f\x3f\xb5\x91\x11\x1a\x47\x29\x95\xd8\x04\x9f\x49\x3c\x47\xd5\xba\x41\x0c\x56\x97\xe3\x6d\x0c\x5f\xe6\x8c\xa7\x3d\xe3\x69\xac\x29\xc2\xf2\x53\x86\x5b\x1a\x5c\x46\x88\xc1\xf1\xf7\x4b\xb1\x94\x25\x78\xac\x2e\x00\xbf\xd7\x01\x00\xd0\xff\x29\x05\xaf\x91\xbc\x81\xf9\xcb\xc1\x5a\xc7\x60\x39\x83\x8d\x61\xe8\xd9\x40\x1b\x22\xd4\x9c\x33\x47\xb8\x86\x5a\x1c\x83\x91\xd4\x3b\xba\x81\x52\xeb\x1b\x07\x95\x15\xab\x63\x32\x1c\x57\x0d\x00\x5d\x0f\x39\x07\x0f\xf9\xd6\x73\x89\x4b\x80\xa7\x96\x45\x1a\xc1\x50\xa6\x35\x98\xa1\x9d\xa3\x3e\xed\x69\x8a\x96\x73\x89\x4f\x75\x52\xfc\x95\xae\xbd\x63\xb5\xb6\x6f\x95\xea\x65\x57\x05\xd0\xa9\x27\xbf\xe9\xa4\xa8\x82\x77\x37\xac\x3e\x2d\x43\x3e\xbc\xd2\xc5\x5c\xf7\x8b\x36\x69\x82\x57\x35\x45\xac\xfe\x65\x99\x2e\x16\x57\xf6\x98\x4e\xf6\xd4\xf3\x8e\x10\xba\xc8\x6d\x89\x4f\x58\x7d\x78\xfd\xcc\xb5\x2e\x68\x5d\x41\x61\x64\xfc\x79\x1b\x9b\x29\x70\x32\x09\x41\x4c\x89\x7f\x65\xa2\x1e\xdc\x81\x64\x7b\xc8\xd3\x78\xb4\xd9\xc9\x56\x43\x4d\x96\x91\xb1\xd2\xb4\xa2\x16\x58\x4d\x93\xbc\xfc\xef\xe1\xf9\x8d\xbc\x05\xec\xc2\x95\xf1\x7e\x9f\xd1\x75\xe1\xe4\xf1\xca\xee\xc4\xe0\xfe\x28\xbe\x1d\xa3\xd8\x2e\x43\x27\xc6\xb0\x57\xe9\xba\x33\x4d\x93\xb4\xf0\xfc\x31\xbd\x05\x2b\xfe\x7e\xbf\xec\x9c\x07\xb0\xa6\xb3\x3d\xa5\x74\xe6\x4b\x9c\xb3\x78\x9b\xde\x61\xfc\xb1\xdf\x05\x1b\x86\x7c\x6e\x5f\xb3\xef\x0d\x38\x4d\xec\xcd\x03\x66\x9f\xf3\xb1\xbc\xfd\xa4\x0b\x4f\x63\x75\x59\xe3\xe5\xf7\x2d\x00\x00\xff\xff\xeb\x69\x3b\x89\x35\x04\x00\x00")

func TemplatesBaseNavHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesBaseNavHtml,
		"../../templates/base/nav.html",
	)
}

func TemplatesBaseNavHtml() (*asset, error) {
	bytes, err := TemplatesBaseNavHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/base/nav.html", size: 1077, mode: os.FileMode(420), modTime: time.Unix(1423149144, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesChgpassHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x4f\x8b\xdb\x30\x10\xc5\xcf\xeb\x4f\x21\xe6\xd0\x5b\x6c\x7a\x2b\xd4\xf6\xad\xa5\x87\x42\x0b\xdb\x43\x6f\xcb\x58\x1a\xc5\x02\xfd\x31\xd2\x38\xdb\x60\xf2\xdd\x8b\x36\x8e\x93\x6c\xe2\x85\x5c\x0c\xd6\xd3\xbc\xf7\xd3\x43\x9a\x26\x26\x37\x58\x64\x12\xd0\x61\xa2\x2a\x7f\xca\x9e\x9d\x05\x51\x1e\x0e\xc5\x8d\x2e\x83\x73\xc1\xbf\xf8\x60\x49\xf3\xc5\xc6\x62\x9a\x14\x69\xe3\x49\x40\x4f\xa8\xe0\x6d\x96\xbc\xba\x96\xba\xa0\xf6\x59\xaa\x95\xd9\x09\x69\x31\xa5\x06\x98\xfe\xf1\x46\x92\x67\x8a\xd0\x16\xb5\x0e\xd1\x09\x47\xdc\x07\xd5\xc0\xef\x5f\xcf\x7f\xa0\x2d\xa6\xa9\xfc\x9b\xa2\xfe\x1e\xa2\xfb\xc1\xce\xbe\x79\x1b\x2d\x4a\x6d\x31\xf5\xa5\xec\xd1\x6f\xe9\x79\x94\x92\x52\x3a\x1c\x8a\xa7\x75\xf7\xa7\x7a\x38\x09\xdd\x76\x93\x8e\x23\x79\x7d\x9a\x56\xcd\xaa\x21\xcf\x55\xca\xec\xda\xe5\x48\x97\xf1\x81\x7b\x8a\xdf\x62\x0c\xf1\x91\x6c\xca\x03\x57\xc9\xef\x7c\xee\xc5\xd6\xf3\x9f\xf9\xfc\xc5\x8b\xf2\x27\xfa\xad\x80\x60\xd5\xcb\x80\x29\xbd\x86\x98\x5b\xaf\x8d\x1f\x46\x16\x1e\x1d\x35\x8b\x06\x82\xf7\x03\x35\xb0\xec\xbb\xe2\xb7\x6a\x4e\x3d\xe1\x95\x3d\xa6\x19\x70\x4e\xae\xda\x95\x89\x33\xfd\xb2\xb6\xc0\x1e\xd9\xef\x31\x7b\x7a\x5d\x65\x3e\x69\x37\xcc\x55\xfb\x91\xa5\xd4\x6e\xd5\xf2\xa4\x7d\x58\x83\xd4\xee\xc1\x1a\x2e\x26\xce\xd7\x67\x59\xbb\x5b\xc3\x8c\x75\xc4\x48\x63\xe7\x0c\x83\xd8\xa1\x1d\xa9\x81\xeb\xf3\xb8\xa0\x8c\xce\x6f\x05\x96\x4b\xc3\x5e\x74\xec\x37\x43\x34\x0e\xe3\x1e\xda\x4f\xbe\x4b\xc3\x57\x51\xd4\x28\xfa\x48\xba\x81\xea\x66\xaf\x22\x8d\xa3\x65\x68\xdf\x95\x85\x5e\x92\xcd\x35\x55\x78\x6e\xb5\xca\x4f\x6f\xf9\x9d\xf9\xff\x07\x00\x00\xff\xff\xbd\xb3\xc5\xc4\x21\x04\x00\x00")

func TemplatesChgpassHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesChgpassHtml,
		"../../templates/chgpass.html",
	)
}

func TemplatesChgpassHtml() (*asset, error) {
	bytes, err := TemplatesChgpassHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/chgpass.html", size: 1057, mode: os.FileMode(420), modTime: time.Unix(1490778585, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesDelsuccessHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x4a\x4a\x2c\x4e\xd5\x07\x11\x7a\x19\x25\xb9\x39\x4a\x0a\x7a\xb5\xb5\x5c\x18\xf2\xc9\xf9\xb9\xb9\xf9\x79\x48\x2a\xb8\xaa\xab\x53\x52\xd3\x32\xf3\x52\x15\x94\x32\x52\x13\x53\x94\xc0\x9a\x52\xf3\x52\x50\xa5\x92\xf2\x53\x2a\x41\x52\x36\x29\x99\x65\x76\x5c\xd5\xd5\x99\x86\x16\x79\x0a\x7a\x3e\x89\x79\xe9\x0a\x4a\x29\xa9\x39\xf1\x29\x49\xf1\xc5\xa5\xc9\xc9\xa9\xc5\xc5\x60\x55\xfa\x50\x65\x60\x73\x00\x01\x00\x00\xff\xff\x9e\x82\xc4\x99\xa4\x00\x00\x00")

func TemplatesDelsuccessHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesDelsuccessHtml,
		"../../templates/delsuccess.html",
	)
}

func TemplatesDelsuccessHtml() (*asset, error) {
	bytes, err := TemplatesDelsuccessHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/delsuccess.html", size: 164, mode: os.FileMode(420), modTime: time.Unix(1490778599, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesLoginHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x95\xcf\x8e\xdb\x2a\x14\xc6\xd7\x93\xa7\x40\x2c\xee\x0e\xfb\xce\x9d\x59\xdc\xaa\xd8\x52\x55\x4d\xd5\x45\xa5\x8e\xd4\x2e\xba\x1b\x61\xfb\x38\x46\xe5\x9f\x00\x67\x26\x42\x7e\xf7\x0a\x9b\x38\x4e\xc6\x69\x9b\xd6\x0b\x8c\xc3\x39\x1f\xe7\xfb\x41\x20\x04\x0f\xd2\x08\xe6\x01\xe1\x8a\x39\xc8\x63\x93\x75\x5e\x0a\x8c\xb2\x61\xd8\x6c\x42\x68\xa0\xe5\x0a\x10\xee\x80\x35\x78\x18\x36\x21\x80\x6a\x4e\x87\x14\xdb\x4d\x23\x67\x62\x8a\xed\x16\x5a\x2b\x89\x02\x5a\x7f\x41\xb3\xd2\xcd\x3e\x0e\xd1\x86\xef\x50\x2d\x98\x73\x05\xf6\xf0\xe2\x49\x0d\xca\x83\xc5\xe5\xc9\x48\xad\x05\x91\x0d\xb9\x47\xa9\xa3\xdb\xd6\x81\x27\xf7\x31\xac\xbb\x2d\x37\x21\xf0\xdb\xff\x15\xca\x3e\x31\xb5\x45\x58\xe8\x2d\x57\x4f\x9e\x7b\x01\xe3\x14\x79\x0c\xa1\x95\xcd\xcb\xcd\x4d\x08\xbc\x45\x59\x2b\x98\xeb\xb2\x77\xbd\xef\x1e\xac\xd5\x76\x18\x36\x68\x7c\x96\x73\x32\x01\xd6\xa3\xb1\x25\x0d\x53\xdb\xb1\x28\x94\x9e\x1b\x6a\xca\x10\x5e\xeb\xd0\xdc\x4c\x41\x37\x34\x6f\xf8\x6e\xea\x1f\xec\xd3\x56\x5b\x89\x24\xf8\x4e\x37\x05\x7e\xfc\xfc\xe5\x2b\x46\x56\x0b\x28\x70\x1c\xc0\x88\xd5\x9e\x6b\x55\xe0\x7c\x34\x80\xa7\x52\x93\xf0\xa1\xaa\x8e\x39\x02\xf1\x27\x9c\x54\xa3\xf9\xec\x9b\xb3\xed\x07\x6d\xe5\x47\x2f\xc5\x68\x66\x69\x24\x8a\x93\xad\xd5\xbd\x49\x06\xa8\x60\x15\x08\xd4\x6a\x5b\x60\xae\x4c\xef\x1f\x24\xe3\xe2\x0e\x2f\x69\x3b\x49\xee\x50\xad\x95\xb7\x5a\x90\x31\x1e\x97\x6b\x94\x15\x93\x11\x32\xcd\xc7\x98\xa4\x7f\xb6\x72\x4e\x92\x37\x33\x3b\x3a\xce\x88\xfc\xde\xc0\xb4\xe2\xf8\xa4\xce\x34\x25\x46\x51\xb8\xc0\xbd\x03\x8b\x11\x6f\xce\x0a\x35\x82\xd5\xd0\x69\xd1\x80\x2d\xf0\xa5\xb2\x9e\x3a\xae\xe2\xf6\xc3\x68\xc7\x44\x0f\x31\x30\x2d\x58\x54\x1d\x86\x03\x8d\xc3\x42\x1d\x3b\x71\xab\x1c\x9b\xab\x71\x3e\x32\xe7\x9e\xb5\x6d\xae\x25\x6a\x52\xde\x9f\xf3\x9c\x15\x12\xbe\xe3\xf7\x2a\xe3\x99\xeb\xa2\xe2\x9f\xa0\x3d\xa8\xcd\x60\x7f\x0d\xf0\x6a\x76\xef\x99\xf1\x75\xc7\xae\x24\x57\xa7\xac\xbf\xdc\x88\x13\xb4\xfa\xac\x84\x25\x33\x94\x64\x6e\x57\xe1\x5d\xe4\x71\x19\x41\x08\x59\x9a\x6f\xfa\xdf\xfe\x46\xca\xaa\xaf\x74\x22\xfe\x37\x17\xf8\xef\xd1\xe8\x32\xba\x83\xfa\x7b\xa5\x5f\x16\x47\x19\x5d\x00\x5b\x21\x73\xcc\x40\xa7\xcc\x2d\x48\x90\x15\xd8\xa7\xb8\x2d\xf0\x7c\x84\x46\x0f\x27\x8a\x8b\x83\xf0\x2a\x32\x57\xdb\xac\x7a\xef\xb5\x4a\x65\xbb\xbe\x92\xfc\x78\xb6\x54\x5e\xa1\xca\x2b\x62\x2c\x97\xcc\xee\x57\xcf\x32\x3c\x0c\xff\xa8\xca\x99\xb7\x88\x3a\xc3\xd4\x21\x75\x2b\xf6\xa6\xe3\xb5\x56\x68\xee\x91\x9a\xdb\x5a\x00\x61\xd6\xea\x67\x62\xf9\xb6\xf3\xb8\xa4\x79\xcc\x2a\x69\x3e\xd5\xb1\xee\x98\xe6\xd1\x65\x7c\xa7\xcf\xf1\xb5\x72\x3f\xb6\x5a\xc7\x6b\x70\xed\xda\x9d\x86\x5e\xdf\xbc\x3f\x02\x00\x00\xff\xff\x92\x19\x37\x96\xed\x07\x00\x00")

func TemplatesLoginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesLoginHtml,
		"../../templates/login.html",
	)
}

func TemplatesLoginHtml() (*asset, error) {
	bytes, err := TemplatesLoginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/login.html", size: 2029, mode: os.FileMode(420), modTime: time.Unix(1490669483, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesRootHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\x4d\x6f\xe3\x36\x13\x3e\x27\xbf\x82\xe0\xeb\x0d\xde\x02\xb5\x64\x6f\x51\x60\x9b\x95\xd2\x43\xb1\x87\x45\x83\x60\xbf\xd0\xeb\x82\x12\xc7\x36\x11\x8a\x54\x48\xca\x1f\x60\xf4\xdf\x8b\x21\x29\x59\xb2\x51\xf4\xba\x87\x5e\x6c\x6a\x66\xc8\x99\x79\x9e\x79\x48\xef\x1d\x34\xad\x64\x0e\x08\xad\x98\x85\x1c\x7f\xb2\x9d\x6b\x24\x25\x59\xdf\xdf\x5e\xf9\x6b\xdd\x34\x5a\x4d\x22\x6e\xbd\xe7\xb0\x11\x0a\x08\xdd\x01\xe3\x34\x6c\x02\xc5\xe7\xae\x4a\xf3\x13\xba\x0a\x2e\xf6\xa4\x96\xcc\xda\x92\xd6\x5a\x2e\x1b\xbe\xfc\x8d\x3e\xdc\xde\x4c\xed\x46\x1f\xd0\x74\x53\x6c\xb4\x69\x06\x23\xae\x97\x42\x49\xa1\x80\x12\xa3\x25\x44\x53\x08\x9c\xed\x6e\x3b\x29\x97\x46\x6c\x77\x8e\x38\x38\xba\x65\x0d\xca\x81\x89\x71\x37\x85\x50\x6d\xe7\x86\xd0\xca\x29\x52\x39\xb5\x6c\x8d\x68\x98\x39\x85\xb5\xdc\x52\xe2\x4e\x2d\x94\xd4\x76\x55\x23\x1c\x25\x7b\x26\x3b\x28\xa9\xf7\x62\xfd\x4e\x91\xec\x91\xa9\x2d\xa1\x70\x84\xba\x73\x40\xfb\x9e\xe6\xb1\x84\x9c\x8b\xfd\x58\x4c\xca\x86\x05\x30\x03\x8c\x28\xd6\xe0\x89\x2f\x92\xce\xfa\xa9\xb5\x72\x46\x4b\x6c\xe8\x60\x4b\xfa\x0b\xbd\x80\xe6\x57\x4a\x5a\xc9\x6a\xd8\x69\xc9\xc1\x94\xf4\x03\xb6\x42\xbe\x7e\x7e\xa4\xc4\xba\x13\x62\x70\x10\xdc\xed\xee\xdf\xad\xde\xbc\xa7\x0f\xde\x67\xf6\x45\xf6\x7d\x91\x0f\x79\x67\x4d\xc7\xae\x76\x82\x73\x50\x34\x55\x24\xf8\xa4\xbf\x4c\xf0\xeb\x76\x8a\x1c\x2b\x45\x86\xa2\xe5\xf6\xc6\xfb\x85\x63\x95\x04\x72\x5f\x92\x2c\xac\xfa\xfe\x9a\xc0\xa1\xc0\x86\x99\xad\x50\x4b\xa7\xdb\xfb\xb7\xab\xf6\xf8\x1e\x89\xf0\x5e\x6c\x48\x26\xec\x37\xdc\xfc\x97\x80\x03\x1e\xf0\x0f\x24\x26\x82\x15\x8b\x98\x12\x52\x74\x72\x8c\x62\x5b\xa1\x98\x13\x5a\xd1\xe4\x24\xc4\xfb\x45\xdd\x99\x4f\x6c\x1b\xeb\x4b\xeb\x90\x60\xf0\xbb\x2a\x96\x5e\xcd\xac\x82\x07\x2b\x62\x30\xb1\x4a\xd1\x08\x17\x1c\x61\x35\xf5\x19\xa6\xb6\x40\x16\xcf\x3f\x93\xc5\x1e\x23\xfe\xff\x25\x18\x32\xa7\x1d\x93\x98\xf4\xa7\x73\x74\x21\xc5\x43\xc1\xc8\xce\xc0\x26\x4c\xd2\x86\xc0\x0b\x59\x3c\x93\xc5\x58\xdf\xff\xbc\x07\x69\xa1\xef\xf3\xbd\x80\xc3\xef\x82\x97\xa1\xa6\xbe\xbf\x73\x55\x19\x8a\xee\xfb\x3b\xeb\x98\x71\xa5\xf7\x4d\x27\x9d\x08\xfb\x53\x55\x49\x71\x38\x03\x8c\x73\xf4\xac\x71\x0e\xd8\x43\x91\x4b\x31\xc1\x26\xea\x32\xe2\x98\x77\x32\x51\x9d\xc0\x1d\x39\x3f\xc7\x15\x91\xe9\x84\x77\xfc\x08\xbf\x4b\xeb\x8c\x68\x81\xa7\xaf\x4a\x1b\x0e\x06\x38\x8d\xdb\x23\xbd\x3b\xdd\x24\xe0\x0b\x67\xc6\x33\xae\x35\xe9\x38\x09\x63\x5c\xd2\xf5\x6a\xf5\x86\x92\x1d\x20\xf1\x25\x7d\xbb\x5a\x85\xe9\x14\x5b\x55\xd2\x46\x70\x2e\x01\x1b\x9c\xca\xf0\x00\xb2\xd6\x0d\xca\xb0\xc8\x1d\x4f\xed\x38\x93\x9a\x08\x78\xa6\xfc\x31\x97\xf7\x8b\xf6\xf9\x23\x3f\x06\x46\xc3\x2a\x06\xdc\x0c\x74\x66\xb5\x96\x5d\xa3\x6c\x32\x17\x8e\xa3\xac\x9e\x18\x76\x52\x54\x26\xc7\xaf\xaf\x9f\x1f\xbf\x9d\x5a\x18\xac\x43\xe2\x09\x6c\x09\x03\x61\x9f\x84\x24\x49\x2d\xaf\x44\xe9\x34\x3f\xc9\x2d\x41\x25\x67\xf6\xe9\xcf\x3f\x62\x5e\xf2\x8a\x93\xb1\x1e\xe2\x62\xc1\x58\xad\x50\x1c\x8e\xd7\xe1\xab\x49\xa4\xd8\x10\x0c\x7e\x4d\x79\x67\x09\x53\x27\x53\xec\x74\x0b\x26\x8a\x67\xda\xc4\xbc\x8d\x8b\x9e\xc6\xf5\x00\x32\xd6\x07\x6a\x2b\x14\x5c\xc8\x07\xa5\x83\x59\xd0\x86\xe9\x06\x6b\x42\xd9\x40\xad\x0d\xb7\x57\xec\x24\xf7\x1c\xfd\x1f\x07\x63\x2c\x68\x94\x71\xce\x41\x7e\x09\x7d\x24\xb5\x0e\x40\x9c\x35\x1b\x0e\x8a\x63\x72\x57\x6b\x19\xee\xdc\x90\x6d\x62\xdc\x33\x59\x7a\x1f\x53\x67\x64\x91\x86\xf2\x4e\x58\xd5\x35\x60\x44\x9d\x36\x0c\x43\xf7\xd1\x3e\x45\x7b\xdf\xd3\xcb\x77\x8c\x23\x7a\x26\x2c\x6d\x33\x4a\x25\x12\x41\x39\x48\x70\x51\x29\x78\x2d\x8c\x7c\xff\x3b\xc3\x49\x47\x53\xaa\x50\xb3\xb5\x96\xb6\x65\x0a\xef\x33\x44\x79\x50\xce\x2b\x5e\x40\x6b\xac\x6e\xaa\xe3\x74\x54\xbc\x18\x3e\xc4\xc7\xf3\x4c\xe3\x64\x2a\xd9\x66\x03\xb5\x03\xfe\x9d\x83\xad\x29\xc9\x86\xef\x61\x80\xce\xa2\xbe\xd8\xa8\xf4\x77\xce\x1c\xa3\x63\xe0\xa4\x93\xf9\xcd\x30\x75\xce\x6f\xbb\x3c\x30\x36\xbd\xc3\xfe\x7b\xa2\x7e\xf0\x27\x2a\x59\xd2\x5f\x32\xff\x1d\x00\x00\xff\xff\xc5\xcf\x4e\x36\xd6\x0a\x00\x00")

func TemplatesRootHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesRootHtml,
		"../../templates/root.html",
	)
}

func TemplatesRootHtml() (*asset, error) {
	bytes, err := TemplatesRootHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/root.html", size: 2774, mode: os.FileMode(420), modTime: time.Unix(1490778624, 0)}
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
	"../../templates/add.html":                TemplatesAddHtml,
	"../../templates/addsuccess.html":         TemplatesAddsuccessHtml,
	"../../templates/base/base.html":          TemplatesBaseBaseHtml,
	"../../templates/base/common.html":        TemplatesBaseCommonHtml,
	"../../templates/base/common_noleft.html": TemplatesBaseCommon_noleftHtml,
	"../../templates/base/footer.html":        TemplatesBaseFooterHtml,
	"../../templates/base/left_engines.html":  TemplatesBaseLeft_enginesHtml,
	"../../templates/base/nav.html":           TemplatesBaseNavHtml,
	"../../templates/chgpass.html":            TemplatesChgpassHtml,
	"../../templates/delsuccess.html":         TemplatesDelsuccessHtml,
	"../../templates/login.html":              TemplatesLoginHtml,
	"../../templates/root.html":               TemplatesRootHtml,
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"templates": &bintree{nil, map[string]*bintree{
				"add.html":        &bintree{TemplatesAddHtml, map[string]*bintree{}},
				"addsuccess.html": &bintree{TemplatesAddsuccessHtml, map[string]*bintree{}},
				"base": &bintree{nil, map[string]*bintree{
					"base.html":          &bintree{TemplatesBaseBaseHtml, map[string]*bintree{}},
					"common.html":        &bintree{TemplatesBaseCommonHtml, map[string]*bintree{}},
					"common_noleft.html": &bintree{TemplatesBaseCommon_noleftHtml, map[string]*bintree{}},
					"footer.html":        &bintree{TemplatesBaseFooterHtml, map[string]*bintree{}},
					"left_engines.html":  &bintree{TemplatesBaseLeft_enginesHtml, map[string]*bintree{}},
					"nav.html":           &bintree{TemplatesBaseNavHtml, map[string]*bintree{}},
				}},
				"chgpass.html":    &bintree{TemplatesChgpassHtml, map[string]*bintree{}},
				"delsuccess.html": &bintree{TemplatesDelsuccessHtml, map[string]*bintree{}},
				"login.html":      &bintree{TemplatesLoginHtml, map[string]*bintree{}},
				"root.html":       &bintree{TemplatesRootHtml, map[string]*bintree{}},
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