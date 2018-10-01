// Code generated by go-bindata.
// sources:
// assets/tuned/01-namespace.yaml
// assets/tuned/02-service-account.yaml
// assets/tuned/03-scc-tuned.yaml
// assets/tuned/04-cluster-role.yaml
// assets/tuned/05-cluster-role-binding.yaml
// assets/tuned/06-cm-tuned-profiles.yaml
// assets/tuned/07-cm-tuned-recommend.yaml
// assets/tuned/08-ds-tuned.yaml
// DO NOT EDIT!

package manifests

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

var _assetsTuned01NamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xb1\x0d\x02\x31\x0c\x05\xd0\x3e\x53\x78\x81\x03\xd1\x66\x08\x4a\x7a\x2b\xf9\x80\xc5\xe5\x3b\x8a\x7d\xcc\x8f\xa0\xa1\x7c\xd2\xd3\x69\x37\xac\x30\x67\x95\xf7\xa5\xbc\x8c\xbd\xca\x55\x07\x62\x6a\x43\x19\x48\xed\x9a\x5a\x8b\x88\x92\x9e\x9a\xe6\x8c\x2f\x45\x7c\x82\xf1\xb4\x7b\x9e\xcc\xcf\xf4\x8e\x2d\xb0\xa3\xa5\xaf\x2a\x45\x84\x3a\x50\xff\x69\x6b\xfb\x11\x89\xb5\xfd\x66\x1e\x34\x3e\xca\x27\x00\x00\xff\xff\x64\xcd\xee\x3d\x7f\x00\x00\x00")

func assetsTuned01NamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned01NamespaceYaml,
		"assets/tuned/01-namespace.yaml",
	)
}

func assetsTuned01NamespaceYaml() (*asset, error) {
	bytes, err := assetsTuned01NamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/01-namespace.yaml", size: 127, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned02ServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc7\xb1\x0d\x02\x31\x0c\x05\xd0\x3e\x53\x78\x81\x14\xb4\xee\x98\x01\x89\xde\x72\x3e\x60\xc1\xfd\x44\x89\x73\xf3\xd3\x5c\xf9\x6c\xc4\x13\x73\x45\xa7\xca\x79\x2b\xdf\x60\x53\x79\x60\x9e\xe1\xb8\xbb\xf7\xcd\x2c\x07\xd2\x9a\xa5\x69\x11\xa1\x1d\x50\xc9\x4d\xb4\x4b\x6b\x98\x43\xa5\x0f\x70\x7d\xe2\x95\xd5\x7f\x7b\x25\x66\x65\x6f\xa8\xb9\x19\x7c\x97\x7f\x00\x00\x00\xff\xff\x65\xfe\xcc\xca\x67\x00\x00\x00")

func assetsTuned02ServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned02ServiceAccountYaml,
		"assets/tuned/02-service-account.yaml",
	)
}

func assetsTuned02ServiceAccountYaml() (*asset, error) {
	bytes, err := assetsTuned02ServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/02-service-account.yaml", size: 103, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned03SccTunedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xbf\x6a\x23\x41\x0c\xc6\xfb\x79\x8a\xe9\x0c\x07\x5b\x5c\x3b\x9d\xf1\x71\x4e\x11\x12\xe3\x10\xf7\xca\xae\x76\x23\x3c\x2b\x0d\x92\xc6\xf6\xbe\x7d\x98\xe0\x10\x43\x4c\x4a\xe9\xfb\xc3\x8f\x0f\x0a\x1d\x50\x8d\x84\x53\x3c\xfd\x0d\x47\xe2\x21\xc5\x17\xec\xab\x92\x2f\x1b\x61\xc7\x8b\x6f\x84\xcd\x15\x88\xdd\xc2\x8c\x0e\x03\x38\xa4\x10\x23\x30\x8b\x83\x93\xb0\xb5\x93\x61\xc6\x14\xa5\x20\xdb\x3b\x8d\xde\xf5\xb9\x9a\xa3\x76\x2c\x03\x76\x5e\x99\x78\x0a\x90\xb3\x9c\x1f\xc4\xfc\x1f\xe9\x41\x72\x9d\x71\x97\xeb\x44\x9c\xa2\x6b\xc5\x6f\xf9\x09\xfd\x2c\x7a\xbc\x7d\xef\x94\x4e\x94\x71\xc2\xa1\x61\x01\x31\xea\xad\x8c\xc3\x06\x0a\xbc\x51\x26\x27\xb4\x14\xba\xb8\xfa\xb3\x0a\xa3\x6d\x55\x6a\x69\x78\xbe\x14\x4c\x71\x5f\x79\x6d\x6b\x5e\xc2\xd4\xfe\x9f\x3e\x5b\xcc\x71\x4e\x0d\xd3\x82\x22\x0c\xcf\x9c\x97\xbd\x88\xff\xa7\x8c\x57\x31\x8e\x90\x0d\x83\xb6\xf4\xab\xa1\xde\x29\x34\x7c\x24\xae\x97\xeb\x66\x77\x0d\x7d\x2f\x73\xd9\xa9\x8c\xad\xf8\x0b\xd1\x6a\x29\x19\x67\x64\x87\xbc\xbd\x42\xfd\xc8\x56\x43\xbd\x85\x35\xd4\x13\xf5\x08\x7d\x2f\x95\x3d\xfd\x3a\x7a\xf2\xca\x38\x84\x8f\x00\x00\x00\xff\xff\x15\x71\x4f\x5c\xea\x01\x00\x00")

func assetsTuned03SccTunedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned03SccTunedYaml,
		"assets/tuned/03-scc-tuned.yaml",
	)
}

func assetsTuned03SccTunedYaml() (*asset, error) {
	bytes, err := assetsTuned03SccTunedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/03-scc-tuned.yaml", size: 490, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned04ClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8c\xb1\x6a\x03\x41\x0c\x05\x7b\x7d\x85\x70\xbf\x17\xd2\x85\x6d\x53\xa4\x4f\x91\x5e\xde\x15\xb6\xf0\x5a\x3a\x24\xed\x05\xf2\xf5\xe1\x7c\x57\xbd\xc7\xc0\xcc\x43\xb4\x57\xfc\x1c\x33\x92\xfd\xdb\x06\x03\xad\xf2\xc3\x1e\x62\x5a\xd1\xaf\xd4\x16\x9a\x79\x37\x97\x3f\x4a\x31\x5d\x1e\x1f\xb1\x88\xbd\x6d\xef\xf0\xe4\xa4\x4e\x49\x15\x10\x95\x9e\x5c\xb1\x1d\x99\xa2\xd6\xb9\xe4\x54\xd1\x5b\xcd\xa9\xdc\xc1\xe7\xe0\xa8\x50\x90\x56\xf9\x72\x9b\x6b\xec\x56\xc1\xcb\x05\x10\x9d\xc3\xa6\x37\x3e\xd9\x6e\x07\x20\x6e\xec\xd7\x13\xdd\x38\x5f\x3b\x24\x8e\xf3\x4b\xd9\xee\xf0\x1f\x00\x00\xff\xff\x04\x60\x32\xd4\xbd\x00\x00\x00")

func assetsTuned04ClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned04ClusterRoleYaml,
		"assets/tuned/04-cluster-role.yaml",
	)
}

func assetsTuned04ClusterRoleYaml() (*asset, error) {
	bytes, err := assetsTuned04ClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/04-cluster-role.yaml", size: 189, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned05ClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xcf\xb1\x4e\xc3\x40\x0c\xc6\xf1\xfd\x9e\xe2\x5e\x20\x41\x6c\xc8\x1b\xb0\x33\x14\x89\xdd\xbd\xfb\xda\x9a\x26\x76\x74\xf6\x55\x82\xa7\x47\x51\x2a\x16\x10\xa2\xfb\xdf\xbf\x4f\xe6\x45\xde\xd0\x5c\x4c\x29\xb7\x3d\x97\x91\x7b\x9c\xac\xc9\x27\x87\x98\x8e\xe7\x07\x1f\xc5\xee\x2e\xf7\xe9\x2c\x5a\x29\x3f\x4f\xdd\x03\x6d\x67\x13\x9e\x44\xab\xe8\x31\xcd\x08\xae\x1c\x4c\x29\x67\xe5\x19\x94\xcb\x16\x0d\x6a\x15\x43\x74\x15\x3d\x52\x74\x45\x4d\xcd\x26\xec\x70\x58\xd3\x1f\xde\x3f\xce\xb7\xc2\x17\x2e\xa0\x6c\x0b\xd4\x4f\x72\x88\xe1\x97\x83\xe4\x7d\xff\x8e\x12\x4e\x69\xb8\x4e\xbd\xa2\x5d\xa4\xe0\xb1\x14\xeb\x1a\xdf\x6b\x37\xcb\xdd\xd1\x5e\xd6\x78\xa5\xfd\xc3\x03\x33\xf9\x66\xf3\x66\xd3\x9f\xc0\xf5\x97\xaf\x00\x00\x00\xff\xff\xc1\xcb\xc9\x87\x79\x01\x00\x00")

func assetsTuned05ClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned05ClusterRoleBindingYaml,
		"assets/tuned/05-cluster-role-binding.yaml",
	)
}

func assetsTuned05ClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsTuned05ClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/05-cluster-role-binding.yaml", size: 377, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned06CmTunedProfilesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\x38\xb4\x7b\x48\xb0\x4a\x8b\x9c\xd8\xc9\x0c\x68\x2f\x59\x80\x3e\xac\x48\xb1\x74\x7b\x19\x06\x81\xa5\x4e\xd2\xc1\xe4\x91\x20\x8f\x76\xdc\x6d\xdf\x7d\x90\x22\xcb\x49\x16\x03\xe9\xea\x27\x9b\xbc\xfb\xfd\xb9\x3f\xb4\xf2\xf4\x3b\x86\x48\x8e\x57\xb0\x29\x66\x6b\xe2\x7a\x05\xd7\x8e\x1b\x6a\x3f\x28\x3f\xb3\x28\xaa\x56\xa2\x56\x33\x00\x56\x16\x57\x20\x89\xb1\xce\x7c\x70\x0d\x19\x8c\xe3\x71\xf4\x4a\xe3\x0a\x9c\x47\x8e\x1d\x35\x92\x69\x93\xa2\x60\xc8\xd8\xd5\x98\x49\x62\xe2\x76\xb6\xc7\x79\x8a\x90\x0d\xc7\xf0\xf7\x0c\x00\x0e\x00\xfb\x03\x80\x3f\xac\x22\xfe\x73\xfc\x11\x93\xb5\x2a\xec\xca\x5b\x2f\x64\xe9\x0b\x42\xdc\x45\x41\x1b\x21\x24\xee\x39\xe0\xd6\x23\xdf\xf5\x00\x70\xe2\x55\x40\x16\x18\x79\x4e\x47\x04\x62\x6d\x52\x8d\xe5\x77\x7f\x35\xab\x0d\x05\xa9\x74\x87\x7a\x3d\x7c\x4d\xca\x64\x6d\xc2\x28\x2b\xe9\x82\x4b\x6d\xe7\x93\x64\x1e\x43\xe3\x82\x55\xac\xf1\x9f\xd9\x5e\x52\x44\x43\x9c\xee\xf7\xaa\xd4\x46\x57\x5a\xe9\x0e\x2b\xe9\x02\xc6\xce\x99\xba\x5c\x2e\x16\xe7\xcb\x29\x81\x51\xf6\xc1\xdc\x54\xda\x31\x4b\x50\x7a\x5d\x75\x2a\x76\x91\xbe\x60\x59\x9c\x17\x67\x97\xf3\x03\xc1\x2e\x6a\x31\x53\x0a\x4a\x4e\x7e\x73\x91\x93\xaf\x1a\x17\xb6\x2a\xd4\x65\x31\xde\xad\x31\x30\x9a\xdc\x53\x5d\x59\x75\x5f\xfe\x34\x02\x1d\x12\x19\xa5\x21\x23\x18\xf2\x27\xcc\x7d\x70\x71\x76\x71\xb5\xb8\x5c\x8e\xc1\x4d\xcc\x89\x9d\x50\xb3\xcb\xad\xba\xaf\x52\xc4\x50\x6d\x95\xe8\x0e\xe3\xe8\xe6\x99\x1a\x46\x6a\xbb\xbc\xc6\x46\x25\x23\x79\xab\x47\xf7\x45\x79\x55\xfc\x38\x7f\x65\xf0\xbc\x3c\x9f\x5f\x2e\xaf\x5e\x19\x7d\xfe\x92\x90\xe5\xd7\x08\x39\x1a\xfc\xa2\x90\xa3\xd1\x8f\x85\x3c\x1a\x7a\xc7\x12\x9c\xc9\xbc\x51\x8c\xdf\x3e\xc1\x23\x1c\x0c\x70\xcf\xc6\x77\x22\x3d\x32\x32\x6f\x61\xdd\xaf\x19\x3c\x9c\x42\x44\x11\xe2\x36\xbe\x03\xab\xee\x7b\xde\x9e\x89\x7e\x70\x70\x98\xf4\x7d\xe2\x04\xf0\x81\x98\xac\x32\xe0\x03\xa2\xf5\x42\x8e\xa1\x0d\x8a\x93\x51\x81\x64\x07\x8d\x0b\x70\xfd\xf1\xb7\xec\xb3\x4b\x5c\x83\xa8\xb8\x8e\xab\x29\xf7\x64\xac\xd8\x0a\x0a\xb0\x11\xf5\x5b\x80\x93\x02\xbe\x07\x32\xae\x3d\x61\xed\x53\x3c\x3d\x7d\x07\x89\x49\xe2\x0a\x58\xb1\x8b\xa8\x1d\xd7\xf1\xf4\xe9\x54\x47\xdd\x61\x5d\x59\xe2\xea\x11\x73\xc5\xb1\x2c\xce\x1e\x3e\x13\xe1\xa7\x0e\x41\x9c\x28\x03\x42\x16\x41\x3a\x84\x21\x39\x19\x0c\xb0\x25\x63\xfa\x6a\x46\xaa\x31\x80\x02\x4b\x6d\x50\x82\x75\xff\x36\x68\x8c\x71\x42\x79\x33\x6c\x31\x74\x4e\xde\x80\xea\x5d\x75\x29\x82\xc1\x18\xc1\xd0\x1a\xcd\x0e\xc4\xc1\x67\x84\x80\xd9\x1e\xe2\xe0\xf8\xa1\x8f\x30\x1a\x07\x8a\xb0\x18\x24\xbe\x03\xca\x31\x87\xb3\x7c\x01\xf6\x98\xbf\x1e\x8b\x1c\x57\xda\x45\xe9\xed\x2d\x9e\xb9\xbb\xbb\x7e\x7f\xf3\x73\x75\xfb\xe9\xfd\xcd\xaf\xb0\x55\x6b\xcc\x92\x7f\xdc\x8b\xfc\x3f\xbd\xfb\xf8\x72\xcf\xb6\x1d\xf2\x43\xab\x06\x18\x48\x3e\x07\xf8\xc5\x6d\x31\x0c\x15\xdb\x28\x93\xfa\x32\x4e\x30\x64\x7d\x70\x1b\x9c\x38\x8d\x12\x64\xbd\x1b\x6b\xb3\x9f\x9c\x61\x14\xf6\x57\x3a\x90\x90\xee\xfb\xd0\xd3\xe4\x2f\xf9\xed\xd1\x92\x7f\xde\xd2\x8b\x47\x9e\x0f\x2b\xd5\xff\x7f\x7c\xfb\x26\xf5\x28\xf1\x2b\x37\x68\x7a\x8b\x44\xfb\xaa\x51\x51\xfa\x84\xf2\xfc\x55\x8f\xe5\xcb\x3e\x32\x8c\xff\xd7\xca\xcd\x1d\x38\x7e\xad\xa1\x81\xec\x88\xab\x8d\x1d\xf4\x5a\xe5\x2b\xed\x12\x4b\x39\x5f\xce\x8b\x8b\x8b\xd9\xbf\x01\x00\x00\xff\xff\xa3\x21\x66\x41\x07\x08\x00\x00")

func assetsTuned06CmTunedProfilesYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned06CmTunedProfilesYaml,
		"assets/tuned/06-cm-tuned-profiles.yaml",
	)
}

func assetsTuned06CmTunedProfilesYaml() (*asset, error) {
	bytes, err := assetsTuned06CmTunedProfilesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/06-cm-tuned-profiles.yaml", size: 2055, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned07CmTunedRecommendYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xd0\x3f\x4b\x04\x31\x10\x05\xf0\x3e\x9f\x62\xc0\x4e\xf6\x0f\xb6\x81\xab\xb4\xd5\x6b\xc4\x46\x2c\x66\xb3\xb3\xbb\xe1\xb2\x93\x90\x99\x1c\x16\x7e\x78\xd9\xdc\x89\x87\x68\x21\x5c\x15\x78\x24\xbf\xf7\x08\x26\xff\x42\x59\x7c\x64\x0b\xc7\x3b\x73\xf0\x3c\x5a\xb8\x8f\x3c\xf9\xf9\x11\x93\x59\x49\x71\x44\x45\x6b\x00\x18\x57\xb2\xa0\x85\x69\x6c\x33\xb9\xb8\xae\xc4\xe3\x39\x97\x84\x8e\x2c\xc4\x44\x2c\x8b\x9f\xb4\x75\xa1\x88\x52\x6e\x39\x8e\xd4\x6a\x61\xcf\xb3\xf9\x82\x4e\x44\x74\xe9\x9b\xb1\xf0\x61\x00\x00\x5e\x2f\x84\xc8\x9a\x63\x68\x53\x40\xa6\x66\xc5\x8d\x7b\xab\x97\xfa\x23\xe6\x3e\xf8\xa1\xaf\x50\xbf\x41\xb5\x26\xe0\x40\x41\x3a\x37\xcd\xbb\xee\xb6\x26\x39\x06\xea\x0e\x65\xa0\xcc\xa4\x24\x9d\x8f\xfd\x09\xda\x69\x2e\x64\xaa\x76\x03\x4f\x11\xa4\xb8\x05\xea\x7b\xa0\x77\x2f\x2a\x0d\x3c\xef\x1f\xf6\x3f\x37\x55\x94\xa4\x39\x9f\x57\x98\x43\x01\x45\xbd\x13\xc2\xec\x96\x8b\x55\x7f\x7e\xc4\x06\x5d\xa1\xd7\xf3\x94\xf1\xf7\xbe\x7f\x36\x98\xcf\x00\x00\x00\xff\xff\xfe\xb3\x5a\x52\x43\x02\x00\x00")

func assetsTuned07CmTunedRecommendYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned07CmTunedRecommendYaml,
		"assets/tuned/07-cm-tuned-recommend.yaml",
	)
}

func assetsTuned07CmTunedRecommendYaml() (*asset, error) {
	bytes, err := assetsTuned07CmTunedRecommendYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/07-cm-tuned-recommend.yaml", size: 579, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned08DsTunedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xcb\x6e\x1b\x31\x0c\xbc\xef\x57\xf0\x07\xd6\x9b\xa2\xe8\x65\x6f\x41\x1e\xbd\xd4\x8e\xd1\x14\xbd\x06\xb4\x44\xdb\xaa\xf5\x82\x44\x6d\x6a\x14\xfd\xf7\x42\xd9\x87\xb5\x89\xdd\x14\xdd\x93\x2d\x0e\x47\x43\x72\x44\xf4\xea\x3b\x85\xa8\x9c\x6d\x01\xbd\x8f\x4d\xf7\xa1\x3a\x28\x2b\x5b\xb8\x45\x32\xce\x3e\x12\x57\x86\x18\x25\x32\xb6\x15\x80\x45\x43\x2d\x70\xb2\x24\x87\x7f\xd1\xa3\xa0\x16\x9c\x27\x1b\xf7\x6a\xcb\xb5\xd0\x29\x32\x85\xda\x3a\x49\x35\x27\xab\xec\xae\x02\xd0\xb8\x21\x1d\x33\x05\x14\x58\xf4\x7e\x64\x8b\x9e\x44\x0e\x47\xd2\x24\xd8\x85\x1e\x6a\x90\xc5\xfe\x4b\x91\x7b\x21\x1b\x80\xc9\x78\x8d\x4c\x43\x5e\xa1\x39\x7f\x7a\x46\x71\x91\x04\x60\x94\xf1\xf2\x9b\x42\xa7\x04\x5d\x0b\xe1\x92\xe5\xd5\xac\xf4\xfc\x61\x62\x67\x72\xe8\x71\x06\xfc\xe6\x0e\x64\x5b\xe0\x90\x68\x00\x0a\x67\x19\x95\xa5\x30\x29\xa8\x41\x38\x63\xd0\xca\x93\xa4\x1a\x9a\x0e\x43\xa3\xd5\xa6\x79\xb9\xa4\xd9\x28\xdb\x84\x64\x27\x80\x32\xb8\xa3\x16\xa4\x13\x07\x0a\x0b\xe5\x9a\x1f\x86\xac\xc0\x43\x0f\xaf\xbd\x13\x73\xe8\x3a\x69\xbd\x76\x5a\x89\x63\x0b\xd7\xfa\x19\x8f\x71\x8a\xdb\x37\xb5\x00\x04\x8a\x2e\x05\x41\xb1\x85\x5f\xbf\xa7\xd3\x48\x22\x05\xc5\xc7\x1b\x67\x99\x7e\xf2\x49\x2e\x80\x0f\xaa\x53\x9a\x76\x24\x67\xc5\xe6\x51\x04\xa3\x2c\xb2\x72\x76\x49\x31\x66\x29\xc8\xfb\x16\x1a\x49\x5d\x53\x04\x6b\xed\x76\x7f\x4b\x1a\xb4\xdf\x2b\x7d\xe2\xee\x9c\x4e\x86\x96\xb9\xcf\xb1\xec\x5d\x5f\x11\xb1\xa8\xfb\x6e\x04\xca\x0d\x26\x2b\x0b\xc1\x2f\xb3\x1a\xa4\x10\x8b\xa1\xcd\x13\x72\x21\xdf\xf0\x75\x18\x6a\xad\x36\x03\xa7\x0f\x6e\xab\x34\xc5\x3a\x3b\xeb\x02\xef\x7c\x84\xe7\x33\x46\xf6\x58\x4c\x64\xce\x52\x46\xc8\x76\x65\xd7\xc7\xe4\x87\x9b\xf5\xd3\xea\xe1\xf6\xee\x69\x75\xbd\xbc\x2b\xe2\x00\x1d\xea\x44\xf7\xc1\x99\x76\x76\x0c\xb0\x55\xa4\xe5\x57\xda\xbe\x3e\x1f\x22\xfd\xdd\xf9\x01\x2c\xf2\xdb\xcd\x7e\xaf\xca\xa6\x17\xe6\x7d\xab\x7f\xef\x62\x2f\xbe\x34\xc8\xeb\x62\xb2\xeb\xed\x56\xed\x96\xe8\x4b\x9c\x62\x32\x71\x5e\xe3\x81\x8e\x83\x3f\x6b\x27\xfc\xd9\x69\x8e\xfc\x9f\xae\xea\xe9\x35\x2f\x32\x7f\x81\x29\x7c\x7e\x96\xc3\xf9\xec\x37\xd4\xaf\x0c\xfc\x9e\x99\xfe\xab\x8e\x4b\xee\x19\xeb\x98\xa3\x16\x47\x34\xfa\x42\x25\x23\xe6\x5f\x0b\x79\xdf\xc5\xd2\xc6\xf1\xb5\xdd\xf4\xeb\xfb\x5e\x85\xc8\xd5\x69\xb6\x2b\xe2\x67\x17\x0e\x33\xfe\x40\x91\x31\xf0\xd9\x1d\x13\xc5\x9e\x64\xd2\x14\xfa\xb5\x29\x69\x8b\x49\x73\x3d\x1d\x4f\x2b\x76\xbe\x5e\x4e\xab\xa7\xd8\x07\x9f\x03\x0a\x5a\x53\x50\x4e\x3e\x92\x70\x56\xc6\x16\x3e\x5e\x55\x7f\x02\x00\x00\xff\xff\xeb\x13\xcc\xa0\xbc\x06\x00\x00")

func assetsTuned08DsTunedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned08DsTunedYaml,
		"assets/tuned/08-ds-tuned.yaml",
	)
}

func assetsTuned08DsTunedYaml() (*asset, error) {
	bytes, err := assetsTuned08DsTunedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/08-ds-tuned.yaml", size: 1724, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"assets/tuned/01-namespace.yaml": assetsTuned01NamespaceYaml,
	"assets/tuned/02-service-account.yaml": assetsTuned02ServiceAccountYaml,
	"assets/tuned/03-scc-tuned.yaml": assetsTuned03SccTunedYaml,
	"assets/tuned/04-cluster-role.yaml": assetsTuned04ClusterRoleYaml,
	"assets/tuned/05-cluster-role-binding.yaml": assetsTuned05ClusterRoleBindingYaml,
	"assets/tuned/06-cm-tuned-profiles.yaml": assetsTuned06CmTunedProfilesYaml,
	"assets/tuned/07-cm-tuned-recommend.yaml": assetsTuned07CmTunedRecommendYaml,
	"assets/tuned/08-ds-tuned.yaml": assetsTuned08DsTunedYaml,
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
		"tuned": &bintree{nil, map[string]*bintree{
			"01-namespace.yaml": &bintree{assetsTuned01NamespaceYaml, map[string]*bintree{}},
			"02-service-account.yaml": &bintree{assetsTuned02ServiceAccountYaml, map[string]*bintree{}},
			"03-scc-tuned.yaml": &bintree{assetsTuned03SccTunedYaml, map[string]*bintree{}},
			"04-cluster-role.yaml": &bintree{assetsTuned04ClusterRoleYaml, map[string]*bintree{}},
			"05-cluster-role-binding.yaml": &bintree{assetsTuned05ClusterRoleBindingYaml, map[string]*bintree{}},
			"06-cm-tuned-profiles.yaml": &bintree{assetsTuned06CmTunedProfilesYaml, map[string]*bintree{}},
			"07-cm-tuned-recommend.yaml": &bintree{assetsTuned07CmTunedRecommendYaml, map[string]*bintree{}},
			"08-ds-tuned.yaml": &bintree{assetsTuned08DsTunedYaml, map[string]*bintree{}},
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

