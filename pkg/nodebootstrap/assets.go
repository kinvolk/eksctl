// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/10-eksclt.al2.conf (999B)
// assets/bootstrap.al2.sh (843B)
// assets/bootstrap.flatcar.sh (2.252kB)
// assets/bootstrap.ubuntu.sh (1.946kB)
// assets/flatcar/kubelet.service (2.848kB)
// assets/kubelet.yaml (464B)

package nodebootstrap

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __10EkscltAl2Conf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xd1\x6a\xdb\x4a\x10\x7d\xd7\x57\x2c\x24\x0f\xf7\x82\x57\xe2\xe6\x5e\xee\x43\x40\x0f\xaa\xad\x04\x53\x55\x0e\x91\x43\x0b\x6d\x11\xe3\xdd\xb1\x33\x78\x35\x2b\x76\x57\x76\xd2\xe0\x7f\x2f\xb2\xac\xd6\xa1\xa1\xf4\x4d\xbb\x67\xe6\x9c\x33\xb3\x47\x17\x02\xb7\x5e\x05\x23\x7d\x8b\x8a\xd6\xa4\x84\x7f\xf6\x01\x1b\x2d\xb4\xb3\xad\x24\x16\x1d\x53\x10\x6b\xeb\xc4\xb6\x5b\xa1\xc1\x30\x39\x1e\xb2\x06\xbe\x59\x16\x05\x71\xf7\x24\xae\xc4\x5f\x59\x71\xf5\x77\x14\x7d\xae\xd0\xed\x48\xe1\xd7\xe8\x42\x14\x56\x81\x11\x0d\x06\xd0\x10\x40\xb4\xe0\xa0\xc1\x80\xce\x5f\x8b\xfb\xfc\x76\xbe\x28\x27\x22\xfb\x58\xd5\xb3\xfc\x26\x7b\x28\x96\xf5\x70\x17\xe5\xbc\x23\x67\xb9\x41\x0e\x37\x64\x30\x4d\x30\xa8\x64\xb0\x98\x8c\x5c\x31\xf2\x2e\xba\x10\xb7\xc6\xae\xc0\x08\x60\x2d\x7c\x80\x40\xea\x95\xc6\xb4\x78\xa8\x96\xf9\x7d\x3d\x2b\xab\x89\x28\x17\xb3\xbc\x2e\xb2\x77\x79\x31\x1e\x96\xd9\xbc\x5c\x56\xbf\x95\x3b\xcd\x7b\x52\x1b\xc6\x61\xcb\xf2\x0d\xb1\x23\xe5\xfc\x6e\x22\xe6\x65\xb5\xcc\xca\x69\x5e\xcf\x67\x7f\xc4\x6d\x7a\xd6\xa3\x42\x94\x3f\xa1\xaa\x02\xb8\x90\x9e\x7d\x26\x9d\x77\xc9\x8a\x78\x6c\x10\x5f\x22\x21\xa4\x64\xab\x51\x52\x9b\x5e\xbe\x9c\x94\x0f\xe7\x80\x81\x15\x1a\x3f\x82\xc3\xd8\x87\x09\x98\xf6\x11\xe2\x41\x3f\x26\x9b\x10\xfb\x00\xac\x50\x92\x4e\x2f\x5f\xce\x8c\x8f\x5c\x0d\x3c\xc9\xd6\xea\x9e\xe8\x43\xf6\xa9\xbe\x5b\xcc\xaa\x11\x72\xb8\x21\x1f\xd0\x1d\xf5\xd2\xe0\x3a\x3c\xbf\xdc\x53\x78\x94\x01\x88\xc3\x0f\x13\xc3\xba\xc7\x76\x30\xc6\xee\x65\xeb\x68\x47\x06\x37\xa8\x07\x86\x01\x53\xc6\x76\x5a\xb6\xce\xee\x48\xa3\x4b\x61\xef\x47\xc0\x72\xcf\x89\x4e\xba\x8e\x03\x35\x98\x6a\xab\xb6\xe8\xc6\xc9\x31\xec\xad\xdb\xca\xd6\x74\x1b\xe2\x54\x31\x8d\x7d\x4c\x72\x45\x2c\x35\xb9\x34\xb1\x6d\x48\x14\x53\xbf\xd2\x33\x58\x59\x5e\x0f\x78\xff\x44\x3d\xce\x18\x62\x7d\xaa\x68\xad\x96\xc4\x6b\x07\x67\x16\xa8\x81\x0d\xa6\x97\x2f\x7d\x82\xf3\xf7\x55\x9d\x4f\xef\xeb\x6c\x3a\x5d\x3c\x94\xcb\x43\xac\xb7\x2e\x46\xe5\xe2\x01\x7e\x1d\xf0\x43\x0c\xc7\x3f\x07\xf6\x3e\x56\xb6\xe9\xf3\x90\xb4\xd0\x79\x94\xd0\xe8\xff\xff\xbb\xfe\x37\xfe\xe7\x24\xdb\xbf\x78\x6f\x8c\x36\xbf\x24\x67\xb8\x8e\x9f\xa1\x31\x3f\x97\xf3\x56\x61\x1f\xb1\xbe\x2a\xfa\x1e\x00\x00\xff\xff\x5c\x3b\x51\x59\xe7\x03\x00\x00")

func _10EkscltAl2ConfBytes() ([]byte, error) {
	return bindataRead(
		__10EkscltAl2Conf,
		"10-eksclt.al2.conf",
	)
}

func _10EkscltAl2Conf() (*asset, error) {
	bytes, err := _10EkscltAl2ConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10-eksclt.al2.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x99, 0xe8, 0x6b, 0xb2, 0x26, 0x2e, 0xd0, 0xc, 0x90, 0xe7, 0x18, 0x1d, 0xc, 0xc2, 0x4c, 0x58, 0x6f, 0x79, 0x96, 0x47, 0x5b, 0xec, 0xa3, 0x4e, 0xb7, 0x5b, 0x79, 0x34, 0x47, 0x66, 0x3a, 0x36}}
	return a, nil
}

var _bootstrapAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\x5d\x6b\xdb\x30\x14\x7d\xd7\xaf\x38\x73\x4d\x69\x18\x8e\xd7\xd1\x0d\xda\x26\x83\xb2\x64\x90\x87\x25\x81\xf6\x61\x23\x64\x46\x91\x6f\x66\x51\x45\x32\xd2\x75\xd6\x12\xbc\xdf\x3e\x9c\xd9\xc5\x19\x7b\x1a\x7b\x92\xce\xc7\x3d\x5c\x1d\x74\xf6\x2a\xdd\x68\x9b\x6e\x64\x28\x84\x08\xc4\x48\x1c\xc8\x7b\x7a\xd2\xdc\xc1\x52\x97\xb4\x95\xda\x74\xd8\xba\xca\x06\x62\x21\xb6\x95\x55\xac\x9d\xc5\x77\xe2\x6c\x27\x9f\xb2\xd2\xe5\xe1\x62\x80\x83\x00\x7e\x14\xda\x10\x3c\xc9\x1c\xda\x06\x96\x56\x51\xc6\xcf\x25\xa1\xf1\xdc\x22\x77\x02\x00\xf4\x16\x58\xad\x10\xc5\x87\x13\x53\x1d\x61\x3c\x6e\xd8\xcb\x3a\xc2\x7a\x8d\xf3\xf3\xd6\xd5\x0c\x37\xe2\x4f\x7c\x5b\xbd\x49\xae\xd7\xaf\xe3\x46\xbe\x05\x17\x64\x8f\x81\x00\xa9\xc2\xa1\x75\xb6\x94\x27\xae\xfc\x6f\x7d\xab\x05\x90\x3b\x4b\x18\x21\x25\x56\x29\x3d\x06\xc5\x26\xed\xb6\x1f\xee\x64\x29\x6a\x21\xe6\x8b\xc9\x34\x9b\x2d\xc7\x51\x7c\xa1\x2a\x6f\x90\x24\x41\x1b\xb2\x8c\x82\xb9\xbc\x49\xd3\xcb\xf7\xd7\xc3\xb7\xef\xae\x86\xed\x99\x1a\xc9\x14\x38\xdd\x11\xcb\x24\x97\x2c\x53\xe3\x94\x34\x89\x2e\xf7\x57\x83\x48\xcc\xe6\xf7\x0f\x77\xf3\x8f\xd3\x6c\x36\xf9\xf7\xc4\xae\xa1\x44\xe7\xfd\xc8\x87\xaf\xcb\xe9\x7f\x08\x6d\x6a\x1f\x44\x42\x04\x57\x79\x45\x27\xdd\x3c\x56\x1b\x32\xc4\x43\xb2\x7b\x9c\x81\x0b\x1d\xa0\xa4\x85\xdb\x93\xf7\x3a\x27\x7c\xbe\xfb\x92\x2d\x17\x93\x7b\x21\x94\x64\x7c\xf8\xeb\xec\xb1\x8e\x63\xc2\x68\x34\x5d\x7c\x7a\xe9\x37\x3e\xb4\xb7\xfa\xa4\xa4\xf8\xd0\x43\xf5\x1f\x8f\xed\x89\x0d\xae\x45\xb7\xc0\x38\x3e\x74\xd7\x9b\x24\xbe\xe8\xff\xca\xe6\xf3\x9c\x4e\x45\x83\x5a\x34\x9b\x88\xf0\x1c\x98\x76\x8a\x0d\x72\x49\x3b\x67\x13\x4f\xc6\xc9\xbc\xc7\x93\x95\x1b\x43\x68\xdf\xd2\x13\x02\x4b\xcf\x2f\xfc\xaf\x00\x00\x00\xff\xff\x07\x7d\xb2\xb5\x4b\x03\x00\x00")

func bootstrapAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapAl2Sh,
		"bootstrap.al2.sh",
	)
}

func bootstrapAl2Sh() (*asset, error) {
	bytes, err := bootstrapAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.al2.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x63, 0xa0, 0x73, 0xd4, 0x3, 0x5, 0xfb, 0x57, 0xa7, 0x41, 0x53, 0xfb, 0x4b, 0x10, 0xfd, 0x3, 0x84, 0xe, 0xd5, 0xde, 0x6a, 0xea, 0x47, 0xee, 0xb7, 0x76, 0xb5, 0xc7, 0x89, 0x24, 0xc9, 0xd0}}
	return a, nil
}

var _bootstrapFlatcarSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x7f\x6f\xdb\x36\x10\xfd\x9f\x9f\xe2\xa6\x08\x45\x82\x82\x62\x52\x24\x19\x9a\xc6\x03\x82\xc4\xdb\x8c\x6d\x76\xd0\x74\xc3\x86\x20\x13\x68\xea\x6c\x11\xa1\x48\x81\xa4\x6c\x77\x86\xf6\xd9\x07\xca\x72\x26\xa5\xee\x92\x05\x5b\xff\x92\x78\xef\x7e\x3c\xde\x11\xf7\xf6\xbe\x62\x53\xa9\xd9\x94\xbb\x9c\x10\x87\x1e\xa8\x01\xb4\x16\x57\xd2\x6f\x8f\xa5\x2c\x71\xc6\xa5\xda\x9e\xb5\xa9\xb4\x43\x4f\xc8\xac\xd2\xc2\x4b\xa3\x61\x8e\x3e\x2d\xf8\x2a\x2d\x4d\xe6\xf6\x0f\x60\x4d\x00\x96\xb9\x54\x08\x16\x79\x06\x52\x3b\xcf\xb5\xc0\xd4\x7f\x2c\x11\x82\xcf\x3b\xc8\x0c\x01\x00\x90\x33\x80\xdb\x5b\x88\xe2\x75\xcf\xa9\x8e\x60\x30\x08\xd6\xa3\x3a\x82\xbb\x3b\x78\xf5\xaa\xf5\x0a\xc1\x01\xfc\x13\x7e\xbf\x3d\xa4\x6f\xef\x5e\xc7\x01\x7e\x07\x3e\x47\xdd\x24\x04\x40\x91\x1b\x68\x3d\x5b\x93\x45\x5f\xd9\x0d\x3e\x93\x04\x20\x33\x1a\xe1\x1c\x18\x7a\xc1\xf0\xde\x09\xaf\xd8\x96\x7d\x52\xf0\x92\xd4\x9d\xab\x35\xbc\x94\x4a\xef\xab\x29\x5a\x8d\x1e\xb7\x17\x6c\xea\x44\xa3\x0d\x2c\xf5\x1c\x7e\x78\xf0\x88\x08\x01\x10\x19\x01\x28\xee\x33\x69\x81\x96\xc0\x4c\xe9\x99\xd0\x32\x34\x9b\x00\x5c\x8e\x47\xe9\x2f\xc3\xf7\x37\xa3\xc9\x78\x10\x2d\x0e\x93\xd3\xe4\x30\x6a\xcd\xd7\x3f\xfe\xfc\xdd\x68\xdc\x43\xbf\x4e\x4e\x02\x7a\xf1\xfe\xf2\xfb\x41\xc4\x8b\xec\xf4\xb8\xa9\xb0\x9c\xa3\x87\xdc\xfb\xd2\x9d\x31\x36\x97\x3e\xaf\xa6\x89\x30\x05\x13\x46\x7b\x2e\x75\x43\x66\x69\xec\xbd\xd4\xf3\xa6\xb4\x45\x85\xdc\xa1\x63\x99\x59\x6a\x65\x78\xc6\xe2\x75\x87\x48\x1d\x9c\x68\xbc\x0e\x65\x6a\xda\x87\x12\x3f\xff\xe3\xcb\x57\x4c\x5c\xce\x4f\x8e\xde\x10\x80\xcd\x8f\xab\x0a\xa0\x02\x9e\x1b\xe4\xb9\x05\xba\x5a\xcc\x9e\x08\x00\x7a\xf9\x78\x3c\xb6\x78\x2a\xe6\x79\x1c\xfe\x65\xcb\x4a\x55\xcd\xa5\x76\x9f\x6d\x5b\xff\x69\x6c\xba\xd7\xc6\x3c\x22\xf3\xc8\xf3\x05\xe3\xfb\x5f\xb9\xfc\xc3\x60\x5f\x90\xa3\x37\xe7\x67\xc7\x7f\x7e\xec\xcf\x4f\xf1\x22\xc2\x9f\xac\x85\x4d\xf1\x66\x38\x74\xf2\x60\x62\x7c\xe9\xa8\xe4\x05\xe5\x55\xd8\x6e\x5e\x0a\xee\x8d\xdd\x35\xbe\xbf\x57\x13\x75\x72\xee\x76\x07\xee\x18\xe4\xe2\x30\x39\x4e\x0e\x77\xbb\xa7\x0d\x96\x2a\xa9\xab\x55\xda\xac\x9c\xb0\xd2\xf2\xc2\x64\xf0\x7a\xf5\x04\xc7\xb0\x3f\xc7\x93\xab\x61\x3a\xba\x1e\x44\xf1\xbe\xa8\xac\x02\x4a\x9d\x54\xa8\x37\xcf\xef\x8c\xb1\xa3\xd3\xb7\xc9\x9b\x93\xe3\xa4\xfd\x32\xc5\x3d\x3a\xcf\x0a\xf4\x9c\x66\xdc\x73\xa6\x8c\xe0\x8a\xca\x72\x71\x7c\x10\x91\xd1\xf8\xe6\xc3\xc5\xf8\x72\x98\x8e\xae\x5e\x9e\x71\xab\x30\x54\x66\xdd\x94\x1f\x7e\xbb\x1e\xfe\x07\x49\x83\x6c\x1d\x44\x84\x38\x53\x59\x81\x3d\x6d\x09\xf3\x51\xe8\x13\xd4\x0b\xd8\x03\x9f\x4b\x07\x82\x6b\x30\x0b\xb4\x56\x66\x08\x3f\x5d\xfc\x9a\x5e\x4f\xae\x6e\x08\x11\xdc\xc3\x37\x3b\x63\x9b\x76\x34\x19\xce\xcf\x87\x93\x6f\x1f\xfa\x1b\xaf\xdb\xbf\xba\xd7\xa4\x78\xdd\x39\xd5\x8f\x2e\xdb\x01\xc3\xb9\x26\x5b\x02\x83\x78\xbd\xfd\x3d\xa3\xf1\x7e\x57\xd5\x83\xf8\xf6\xa3\xa2\x83\x9a\x04\x26\x64\x0f\x5a\x11\xec\x28\x20\xf9\x54\x36\x09\xf1\xa6\x12\x79\xef\x76\x41\x88\x09\x71\x1f\x9d\xc7\x42\x78\x05\x19\xc7\xc2\x68\x6a\x31\xbc\xcf\x8e\x1d\x35\x9f\x2a\x84\xb6\x19\x1d\xc0\x79\x6e\xfd\x83\xfd\xaf\x00\x00\x00\xff\xff\x09\xd6\x2b\x1b\xcc\x08\x00\x00")

func bootstrapFlatcarShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapFlatcarSh,
		"bootstrap.flatcar.sh",
	)
}

func bootstrapFlatcarSh() (*asset, error) {
	bytes, err := bootstrapFlatcarShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.flatcar.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x87, 0x64, 0x68, 0x10, 0x38, 0xd7, 0x50, 0xac, 0x6d, 0x99, 0x75, 0xa7, 0xf6, 0xf9, 0x43, 0xe4, 0x18, 0x26, 0x5d, 0xa9, 0xd4, 0xed, 0x5, 0x37, 0x12, 0xf7, 0xf7, 0x48, 0x4d, 0xc8, 0xa9, 0x2f}}
	return a, nil
}

var _bootstrapUbuntuSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x61\x6f\x22\x37\x10\xfd\xee\x5f\x31\x25\xe8\x0a\x6a\xbc\xdb\x5c\xd3\x93\x2e\xc9\x56\xa5\x09\x57\x45\x4d\x21\xba\x70\x6a\xab\x28\x45\xc6\x1e\x58\x0b\xaf\xbd\xb2\x67\x21\x29\xa2\xbf\xbd\x32\xbb\x4b\x48\x74\xbd\x0f\x55\x3f\xb1\xf6\x7b\xf3\x76\xf6\xf9\x79\x38\xfa\x2a\x9d\x69\x9b\xce\x44\xc8\x19\x0b\x48\xc0\x1d\xa0\xf7\xf8\xa8\xa9\x5d\x96\xba\xc4\xb9\xd0\xa6\x5d\x5b\x57\xd9\x80\xc4\xd8\xbc\xb2\x92\xb4\xb3\xb0\x40\x9a\x16\xe2\x71\x5a\x3a\x15\x7a\x7d\xd8\x30\x80\x75\xae\x0d\x82\x47\xa1\x40\xdb\x40\xc2\x4a\x9c\xd2\x53\x89\x10\x39\xe7\xa0\x1c\x03\x00\xd0\x73\x80\xfb\x7b\xe8\x74\x37\x2f\x48\xdb\x0e\x64\x59\xdc\x3d\xd9\x76\xe0\xe1\x01\xde\xbc\x69\x58\xb1\x38\x82\x7f\xc3\x9f\xf7\xdf\xf2\xf7\x0f\xdf\x74\x23\x7c\x0e\x94\xa3\xdd\x09\x02\xa0\xcc\x1d\x34\xcc\xf3\x66\xcf\x23\x55\xbe\x26\xcc\x35\x03\x50\xce\x22\x5c\x40\x8a\x24\x53\x5c\x06\x49\x26\x6d\xdb\x4f\x0a\x51\xb2\x2d\x63\xa3\xf1\xd5\x70\x7a\x7d\x9b\x75\xba\x3d\x59\x79\x03\x9c\x07\x6d\xd0\x12\xe4\x44\xe5\x59\x9a\x9e\xbc\x7b\x9f\xbc\xfd\xfe\x34\x69\x7e\x53\x23\x08\x03\xa5\x05\x92\xe0\x4a\x90\x48\x8d\x93\xc2\x70\x5d\xae\x4e\xfb\x1d\x76\x3d\xba\x9b\x0c\x46\x97\xc3\xe9\xf5\xd5\x7f\x57\x6c\x2d\xe2\x5a\x1d\x4a\x4e\xfe\xb8\x1d\xfe\x0f\xa2\xd1\xf7\x7e\x87\xb1\xe0\x2a\x2f\xf1\x85\x37\xcb\x6a\x86\x06\x29\x41\xbb\x82\x23\xa0\x5c\x07\x90\xc2\x82\x5b\xa1\xf7\x5a\x21\xfc\x3a\xf8\x7d\x7a\x3b\xbe\xba\x63\x4c\x0a\x82\x1f\x3e\x5b\xbb\xb3\x63\xa7\x70\x71\x31\x1c\x7f\xd8\xfb\xdb\xdd\x34\x4f\xdb\x17\x26\x75\x37\x07\xab\xed\xab\x8f\x3d\x00\xe3\x7a\xcb\xda\x06\xb2\xee\xa6\x7d\x3c\xe3\xdd\xde\x61\x2c\x63\x7a\x5e\x56\x75\xfa\x5b\x16\x3b\x61\xc1\x8a\x12\x84\xd1\x22\x40\xd3\x2d\xc7\x65\x48\x9a\xe7\x76\xef\x35\x4d\x92\xd9\xd3\x24\x99\x76\xaf\xa6\x05\x72\xe5\xa1\x18\x0b\x4f\x81\xb0\x88\x3c\x8f\x01\x89\xc7\xfb\x84\x8a\xb1\x1e\x03\x38\x82\xc9\xf8\x6a\x7c\x16\x43\x1c\x10\x42\xee\x2a\xa3\x60\x86\x60\x9c\x5b\xa2\x02\x41\x80\x2b\xf4\x4f\x40\xba\xc0\x56\x14\x02\x09\x4f\x01\xaa\xf2\x78\xa7\xb0\xce\xb5\xcc\x41\x07\x58\xe7\x82\x60\x8d\xa0\x1c\x68\x0b\x83\x9b\xb7\xd0\xdb\x63\x33\x11\x50\x81\xb3\x50\x1a\xa1\x2d\xd4\x3d\xa9\x5a\x40\x58\x05\x05\x0a\x4b\x40\x2e\xbe\xbc\x74\x9e\xc4\xcc\x60\x5c\x16\x2e\x50\xcb\x06\xa5\x03\x79\x17\xfa\xc7\x30\xab\x08\x34\x7d\x1d\x76\xf5\xd6\x11\x48\x83\xc2\x43\xee\xd6\xb1\xc8\x38\xa1\x9a\x4f\x9a\x7b\x57\x3c\x37\x1e\xfd\x59\x6b\xca\x5d\x45\x90\x8b\x95\xb6\x8b\x9d\x00\x39\x90\x55\x20\x57\xe8\x80\xb1\xae\x26\x6a\x0a\x68\xe6\x0c\xe0\x0b\xb1\xdc\x47\xeb\xcb\xb4\x7f\x25\xc4\xcb\x10\xef\xc2\x8e\xc1\x00\xe6\x46\x2c\x42\xd6\xdb\x8d\x8b\x8e\x75\x0a\xb9\x2e\x0f\x72\xda\xa9\x81\x42\x3c\xf2\x18\xac\x83\xcc\xb5\xd0\xae\xc6\x88\x19\x9a\xd0\xd6\xdd\x0c\x7e\x1a\xde\xdc\x6d\x8f\x85\x29\x73\x91\xd4\x2f\x4e\xb4\x3b\xbc\xd3\xaf\x32\xdf\x68\x09\x63\xdc\x9a\x97\x5e\xaf\xb4\xc1\x05\xaa\x8c\x7c\x85\x0d\x56\x3a\xc5\xb5\x9d\x7b\xc1\xa5\xb3\x24\xb4\x45\xcf\x75\x21\x16\x98\x75\x37\x83\xdf\xee\xa6\xc3\x5f\xee\xa6\xc3\xcb\x8f\xd3\xc1\xe5\xe5\xf8\xd3\x68\xb2\x4d\xd4\xd2\x27\x28\x7d\x52\xc3\x57\xc3\x0f\x83\x4f\x37\x93\xe9\xc7\xe1\xcf\xd7\xe3\xd1\x36\x11\x85\xf8\xcb\x59\xb1\x0e\x89\x74\x45\x34\x27\x2d\x45\x15\x90\x8b\x42\xbd\x3b\x3d\xfb\x2e\x39\x69\x5e\x2b\x8d\xab\x14\x2f\xbd\x5b\x69\x85\x3e\x13\xeb\xd0\x02\x56\xf3\x99\xb6\x5c\x69\x9f\xa5\xae\xa4\x54\x5a\x1d\xff\x5d\x0e\x60\xe9\xec\xbc\xc6\xe3\x01\x44\xdc\x22\x25\xaa\x65\xec\x3f\xc3\x57\x36\xc6\x3d\x53\x4e\x2e\xd1\xb7\xbe\x22\xad\x9d\x5f\xf2\xd2\x54\x0b\x6d\x33\x69\x75\x03\x78\x5c\xe8\x40\xe8\x79\x74\xfe\xd0\xa1\x3d\x10\x03\xc7\xa3\x36\xed\x8f\x64\x32\xb8\x1e\x4d\xf6\x67\xb6\xbb\xbe\xce\xce\xf5\x22\x7b\x1d\x9e\x7a\x3b\x79\x12\x85\x79\xee\xf3\x73\xc4\x98\xb2\x96\xd5\x8f\x49\xaa\x67\xc1\xf3\x0c\x89\xa3\x20\x0e\xa2\x5d\xc2\xee\x7f\x7c\xd8\x76\x58\x9f\xb5\x13\x43\xf8\x17\x3c\xf6\x4f\x00\x00\x00\xff\xff\x8c\x16\xcc\x40\x9a\x07\x00\x00")

func bootstrapUbuntuShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapUbuntuSh,
		"bootstrap.ubuntu.sh",
	)
}

func bootstrapUbuntuSh() (*asset, error) {
	bytes, err := bootstrapUbuntuShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.ubuntu.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x99, 0x66, 0x92, 0x12, 0x80, 0x3a, 0x7d, 0x6f, 0xf5, 0xf9, 0x2c, 0x3d, 0xe8, 0x7d, 0xb3, 0x6b, 0xa4, 0x87, 0xc5, 0x2d, 0xf0, 0xec, 0x6e, 0x84, 0x93, 0x59, 0x1f, 0xc, 0x6b, 0x7f, 0x5f, 0x4b}}
	return a, nil
}

var _flatcarKubeletService = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\xdb\x6e\xdb\x38\x13\xbe\xd7\x53\x10\x4d\x2f\x4d\x31\x27\xa7\x6d\x00\x5d\xb8\x89\x9a\xdf\x88\xeb\x14\x96\xfd\x77\x81\x6e\x21\x50\xd4\xc4\x1e\x88\x22\x05\x92\x92\x9b\x2d\xf2\xee\x0b\x9d\x1a\x39\x75\xdc\xec\x9d\x25\x7e\x87\x19\xcd\x78\x38\x47\x64\xaa\x6c\x81\x06\x52\x72\x6f\x74\x7e\xe9\x1d\x91\x8d\x73\x85\xbd\x64\x6c\x8d\x6e\x53\x26\xbe\xd0\x39\xcb\x50\x55\x5a\x66\x4c\xea\x4c\xe7\xda\x61\x05\x34\x2b\x13\x30\x0a\x1c\x58\x96\x48\x9d\xb0\xf3\x77\xe3\xf3\x64\x9c\x9c\x9d\x8e\x2f\x92\xe3\x8b\x0f\xc9\xe9\xc5\xf8\x44\x1c\x9f\x89\xf1\xd9\x05\xbc\x7b\x7f\x7a\x76\x7e\x7a\x76\x7c\x3c\x4e\xdf\x33\xbe\xb5\xec\x5e\x72\x27\xb8\xa1\x12\x55\xf9\x83\x0d\xa4\xb6\xda\x64\x60\x2c\x13\xb2\xfb\xe9\x3f\xf0\x5c\xfa\x2e\x2f\xa4\xe7\x79\xdf\x56\x0a\xdd\x77\xef\x1a\xac\x30\x58\x38\xd4\x2a\xb8\x2d\x13\x90\xe0\x48\x85\x9c\xfc\xef\xa1\x00\x53\x8b\x79\x5f\xb9\x72\x36\x30\x85\xa0\xd6\x71\x97\xfa\x16\x4c\x85\x02\x3c\xef\x5b\xd4\xfe\xfa\xee\x1d\x91\x99\x16\x5c\x92\x1c\x1c\x4f\xb9\xe3\xa4\xe0\x86\xe7\xe0\xc0\xd8\x4b\xb2\x08\x6f\xa6\x77\xf3\x11\x99\x7c\x8d\xe2\xeb\xf0\xd3\x64\x35\x5b\xc6\xed\x3b\x2f\x54\x15\x1a\xad\x72\x50\xee\x13\x4a\x08\x18\x38\xc1\x20\xb3\xc2\x49\xd6\x6b\xf9\xa0\x2a\xef\x88\xdc\x48\x9d\x70\x49\xb8\x4a\x49\x1d\x06\x8a\x1d\x8f\xab\xd9\x2a\x5a\x86\x8b\xf8\x7a\x1e\x8d\xc8\xfc\xee\x3a\x8c\x67\x93\x8f\xe1\xac\x7f\x58\x4e\xa6\xf3\x65\x34\x22\xb7\xab\x8f\xe1\x2c\x5c\xc6\xff\x0f\x17\xd1\x9f\xfc\xb3\xf6\x63\x74\xf6\x6d\x7e\x4a\x2b\xba\xc7\xbd\xf1\x98\x7e\x19\x91\xe9\x3c\x5a\x4e\xe6\x57\x61\x3c\xbd\x7e\x95\xb6\xac\x55\x1b\x87\x21\x3c\x78\xd3\xc7\x39\xfd\x3c\xb9\x09\xe3\xd5\x62\x16\xa4\x5a\x64\x60\x2e\x19\xcb\xde\x5b\x7f\x2d\x8c\x8f\x9a\x6d\xfa\x12\xbd\x39\xc4\x5e\x4e\x6e\x82\xea\xc4\x3f\x39\xf7\x3f\xbc\x79\x66\xb3\xb8\x5d\xc6\x8b\xd5\x3c\x9e\x2c\x6e\xa2\x80\xd2\xb2\xc4\x94\xde\xa3\x04\x6a\x79\x05\x01\xab\xb8\x61\x82\x8b\x0d\xf4\xf1\xd2\x42\xa7\x7e\x8d\x22\x7f\x7b\x84\x50\x5a\x69\x59\xe6\x10\x18\xb0\x5a\x56\xa3\x0c\x55\x1a\x6c\xb4\x75\x23\xab\x4b\x23\xba\x8c\xdb\x43\x5f\x68\x75\xdf\xb1\x72\x5d\x2a\x47\x76\xb9\x8e\x9b\x35\xb8\x97\x18\x1d\xb6\xfd\x78\xf5\x7b\x5c\xbf\xe0\xd6\x42\xf6\x19\xed\x90\x87\x76\x3b\x94\x16\x4c\x74\xe1\x68\x82\x6a\x8f\x89\x2e\x1c\x4b\x50\xed\x73\xe8\x49\xbd\xf8\x2e\xb4\x03\x09\x85\x07\x13\x10\x0a\xf7\x69\x3f\xd1\x86\xa1\x3f\x81\xbb\xb8\xab\x66\x0c\x24\x54\x28\xdc\xa3\x5f\xd7\x53\x62\xf2\x92\xc7\x90\xdc\xbb\xfc\x4e\x79\xee\xc4\x25\x0a\x7d\xc8\xac\x01\x1c\xf4\x6b\x25\x7e\xb3\x1c\x12\x07\x75\x11\x0a\x0f\xd4\x46\x28\x3c\x54\x9f\x9e\x3c\xac\xd1\x2e\x65\x98\x9f\xde\x57\xa5\x26\x3e\xbd\x7e\x31\x23\xbd\xde\x4d\xe5\x17\xb4\x53\x46\x2b\x2c\xf2\x34\xdf\x23\x5d\x5a\xc3\x6c\x82\x8a\xf5\x98\x7d\x26\xbf\xf8\xbd\xcb\x4b\x2c\x54\x16\x44\x69\x80\xea\x66\xba\xdb\x00\x73\xbe\xae\x67\xc5\x0f\x10\x91\xe3\xc6\x7d\x31\x10\xd4\x99\xb3\x3c\x4b\xd1\x10\x5a\x90\xbe\x67\x0f\x62\xba\xce\x63\x0a\x9c\x9f\xfe\x51\xad\xfb\xba\x7f\x54\x1c\xdc\x5b\x39\x57\x78\x0f\xd6\xd9\xff\x42\x7a\x5d\x44\x83\x7e\x7e\x1d\xae\x69\xc2\x57\x41\xbb\x19\xc9\xda\x32\x15\xb2\x5c\xa3\x7a\x96\x01\x6d\x6a\x55\xd3\x4d\xe6\x88\xc9\xc9\x60\xe6\x1e\x1a\xb7\x4f\x2a\x6d\xb5\x9b\xc8\xb4\x01\x6d\x7f\x41\xb7\x86\x17\x05\x98\xae\xf4\x4a\xa7\x40\xb1\x08\xde\xfe\xec\xae\xa6\xc7\xe1\x81\xe4\x09\x48\xdb\x1f\xb6\x17\xe5\xe3\x88\xcb\x62\xc3\xfd\x76\x1a\xd6\xb7\x0b\x2a\xeb\xb8\x12\x40\x31\x0d\xde\xfe\x1c\xdc\x6c\xbd\x56\xce\x7f\xd4\x11\xd6\x42\x9f\x27\x7f\xc5\x5f\xee\xae\xa3\xfe\xc8\xc0\x1a\xad\x03\xd3\xf8\x05\xce\x94\x30\x7c\xb9\x45\xb7\xa1\x8e\x63\xbd\x56\x74\x41\xb4\x17\x74\x4f\xe7\x52\xea\x2d\x2d\x0c\x56\x28\x61\x0d\x69\xab\xd0\x9e\x09\xa9\xcb\x94\x16\x46\x57\x98\x82\x09\xf8\xd6\xf6\x07\x5a\xd5\x9a\x60\xa8\x29\x95\xc3\x1c\xba\x3b\xb3\xcf\x1c\x5c\xbd\x04\xd1\xb6\x30\xc1\xd3\x40\xab\x47\x42\x3d\x5b\x69\x8a\x26\xd8\x6d\xee\x0e\x51\xe8\x94\xa2\xba\x37\x7c\xe0\xd1\xfc\x97\x82\xb7\x3f\xeb\xa5\x26\xbc\x8d\xe2\xf0\x6a\x11\x4f\xae\xae\xee\x56\xf3\xe5\xa3\x9f\x66\xc6\x07\x61\xfc\xf6\x78\x77\xe7\x79\xf4\x79\xce\xff\xd1\x8a\x6f\x6d\xb3\x0d\x42\x66\x59\xc1\x4b\x0b\x94\xe7\xe9\xc5\xf9\xe5\x99\x7f\xd2\xd9\xd6\xa5\x6d\x87\xfe\x6f\xbb\x43\xfb\xba\x59\xe7\x9e\xb2\xdf\x07\xac\x97\x8c\x1a\xd5\x75\x90\x2e\x9e\xf5\xa0\x75\xba\x78\x75\x17\x2e\xc0\x36\x3d\xc8\xe5\x96\x3f\xd8\xfe\x31\x02\x11\x8c\xbd\x6f\xd3\xba\x5d\xa4\xfc\xde\xac\x8b\x90\x7e\x7c\x08\xf2\x52\x3a\xa4\xa5\x05\xe3\xb7\xa3\xca\xfb\x37\x00\x00\xff\xff\x20\x6b\x67\x51\x20\x0b\x00\x00")

func flatcarKubeletServiceBytes() ([]byte, error) {
	return bindataRead(
		_flatcarKubeletService,
		"flatcar/kubelet.service",
	)
}

func flatcarKubeletService() (*asset, error) {
	bytes, err := flatcarKubeletServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "flatcar/kubelet.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x81, 0x5f, 0xfd, 0x4a, 0xe6, 0x78, 0xaf, 0x52, 0x2a, 0xef, 0x76, 0xeb, 0xcd, 0xc7, 0xae, 0x47, 0x78, 0x1e, 0x54, 0x54, 0x3f, 0x8b, 0xcf, 0xaf, 0x15, 0xd8, 0x36, 0xf7, 0xe0, 0x5, 0x28, 0x8}}
	return a, nil
}

var _kubeletYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x4f\x4f\xc3\x30\x0c\xc5\xef\xf9\x14\xfe\x04\x6d\x07\x9a\x04\xb9\x8d\x4d\x70\x60\x27\x36\xe0\xec\xa6\xee\x16\x35\x8d\x27\xc7\x19\x7f\x3e\x3d\x5a\x5a\x90\x26\xa1\x9c\x9e\xde\xb3\xdf\x4f\xce\xe0\x63\x67\xe1\x39\xb7\x14\x48\xd7\x1c\x7b\x7f\xc8\x82\xea\x39\x1a\x3c\xf9\x37\x92\xe4\x39\x5a\x18\xa6\x40\xe5\x4a\xa2\x1a\xee\x52\xe5\xb9\x3e\x2f\x5a\x52\x5c\x18\x83\x5d\x27\x94\x92\x85\xa6\x2a\xcf\xb8\x90\x93\x92\x6c\x78\x44\x1f\x2d\xcc\xb2\x0a\xec\x30\x18\x83\x59\x8f\x14\xd5\xbb\x52\x64\x0d\x00\x46\x8e\x5f\x23\xe7\x74\x11\x00\x14\xb1\x0d\xd4\x59\xe8\x31\x24\x32\x00\x1f\xd4\x1e\x99\x87\xc9\x75\xe8\x8e\xb4\xdf\x6f\x2d\xdc\x8c\x4d\xba\x1e\x50\xc9\x97\xfc\xe7\xb2\xb9\x9f\xc3\xc1\x53\xd4\xf5\xea\xd1\x07\xb2\x50\x93\xba\x9a\x86\xe4\x34\xd4\x0e\x2b\x27\x3a\xd1\xb0\xf8\xef\x3f\x98\x91\x3b\xb2\xf0\x3e\x55\xfe\x5b\xbe\x9a\x47\xa8\x2b\x18\xcb\x5f\x8c\x62\xbe\x46\xbc\xb6\x6f\x9b\x64\x4c\x22\x39\x93\xec\xb7\xbb\x07\x66\x4d\x2a\x78\x9a\x61\x8d\x3b\x08\xe7\xd3\x46\xfc\x99\xc4\xc2\xa4\xfa\x64\x4c\x4f\xa8\x59\xe8\x09\x95\xca\x59\x5e\x58\x51\x69\xfe\xaa\x5d\x59\xb7\x26\x51\xdf\x5f\xee\x48\xf3\xb6\x9f\x00\x00\x00\xff\xff\x1f\x2f\xa9\x0f\xd0\x01\x00\x00")

func kubeletYamlBytes() ([]byte, error) {
	return bindataRead(
		_kubeletYaml,
		"kubelet.yaml",
	)
}

func kubeletYaml() (*asset, error) {
	bytes, err := kubeletYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubelet.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe1, 0x70, 0xd5, 0xbb, 0x81, 0xa2, 0xa6, 0x76, 0x99, 0x80, 0xe7, 0xe2, 0x47, 0xc5, 0xa0, 0xe0, 0xb4, 0xe1, 0x42, 0x2c, 0xb0, 0x60, 0xa0, 0xb0, 0x97, 0x53, 0xa7, 0x1a, 0x9, 0xc3, 0x1, 0x6d}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"10-eksclt.al2.conf":      _10EkscltAl2Conf,
	"bootstrap.al2.sh":        bootstrapAl2Sh,
	"bootstrap.flatcar.sh":    bootstrapFlatcarSh,
	"bootstrap.ubuntu.sh":     bootstrapUbuntuSh,
	"flatcar/kubelet.service": flatcarKubeletService,
	"kubelet.yaml":            kubeletYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"10-eksclt.al2.conf":   &bintree{_10EkscltAl2Conf, map[string]*bintree{}},
	"bootstrap.al2.sh":     &bintree{bootstrapAl2Sh, map[string]*bintree{}},
	"bootstrap.flatcar.sh": &bintree{bootstrapFlatcarSh, map[string]*bintree{}},
	"bootstrap.ubuntu.sh":  &bintree{bootstrapUbuntuSh, map[string]*bintree{}},
	"flatcar": &bintree{nil, map[string]*bintree{
		"kubelet.service": &bintree{flatcarKubeletService, map[string]*bintree{}},
	}},
	"kubelet.yaml": &bintree{kubeletYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
