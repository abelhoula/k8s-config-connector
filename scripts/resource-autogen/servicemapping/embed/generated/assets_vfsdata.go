// Code generated by vfsgen; DO NOT EDIT.

//go:build !integration
// +build !integration

package generated

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/bigquery.yaml": &vfsgen۰CompressedFileInfo{
			name:             "bigquery.yaml",
			modTime:          time.Time{},
			uncompressedSize: 5123,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4b\x73\xda\xc8\x13\xbf\xeb\x53\x74\x99\xcb\xff\x5f\x65\x43\xec\xa3\xf6\x44\x70\x1e\xda\x75\xf0\x2e\x90\x4d\xe5\x44\xb5\x46\x8d\xd4\xce\x68\x66\x32\x33\x82\xa8\x58\x7f\xf7\x2d\xbd\x30\xb2\x71\xec\x7d\xa4\x96\x83\xb9\x58\xee\xe7\xaf\xbb\x7f\x3d\x0c\x1a\xc0\x44\x9b\xd2\x72\x9a\x79\xb8\x78\x75\x71\x01\xef\xb4\x4e\x25\xc1\xd5\xd5\x24\x18\x04\x03\xb8\x62\x41\xca\x51\x02\x85\x4a\xc8\x82\xcf\x08\xc6\x06\x45\x46\x9d\xe6\x14\x7e\x27\xeb\x58\x2b\xb8\x18\xbe\x82\xff\x55\x06\x27\xad\xea\xe4\xff\x3f\x05\x03\x28\x75\x01\x39\x96\xa0\xb4\x87\xc2\x11\xf8\x8c\x1d\xac\x58\x12\xd0\x37\x41\xc6\x03\x2b\x10\x3a\x37\x92\x51\x09\x82\x0d\xfb\xac\x4e\xd3\x06\x19\x06\x03\xf8\xdc\x86\xd0\xb1\x47\x56\x80\x20\xb4\x29\x41\xaf\xf6\xed\x00\x7d\x0d\xb8\xfe\x64\xde\x9b\x70\x34\xda\x6c\x36\x43\xac\xd1\x0e\xb5\x4d\x47\xb2\xb1\x74\xa3\xab\x68\xf2\x66\x3a\x7f\x73\x76\x31\x7c\x55\xfb\x7c\x54\x92\x9c\x03\x4b\x5f\x0b\xb6\x94\x40\x5c\x02\x1a\x23\x59\x60\x2c\x09\x24\x6e\x40\x5b\xc0\xd4\x12\x25\xe0\x75\x05\x78\x63\xd9\xb3\x4a\x4f\xc1\xe9\x95\xdf\xa0\xa5\x60\x00\x09\x3b\x6f\x39\x2e\x7c\xaf\x5b\x1d\x3c\x76\x3d\x03\xad\x00\x15\x9c\x8c\xe7\x10\xcd\x4f\xe0\xf5\x78\x1e\xcd\x4f\x83\x01\x7c\x8a\x16\xef\xaf\x3f\x2e\xe0\xd3\x78\x36\x1b\x4f\x17\xd1\x9b\x39\x5c\xcf\x60\x72\x3d\xbd\x8c\x16\xd1\xf5\x74\x0e\xd7\x6f\x61\x3c\xfd\x0c\xbf\x44\xd3\xcb\x53\x20\xf6\x19\x59\xa0\x6f\xc6\x56\xf8\xb5\x05\xae\xfa\x48\x49\xd5\xb4\x39\x51\x0f\xc0\x4a\x37\x80\x9c\x21\xc1\x2b\x16\x20\x51\xa5\x05\xa6\x04\xa9\x5e\x93\x55\xac\x52\x30\x64\x73\x76\xd5\x34\x1d\xa0\x4a\x82\x01\x48\xce\xd9\xa3\xaf\x25\x0f\x8a\x1a\x06\x01\x1a\x6e\xe7\x1f\x82\xd0\x96\x86\x42\xd9\x7c\x28\xa4\x2e\x92\x61\x5a\x53\x69\x28\x74\x3e\x5a\x9f\xa3\x34\x19\x9e\x07\x5f\x58\x25\x21\xcc\xc9\xae\x59\xd0\x07\x34\x86\x55\x1a\xe4\xe4\x31\x41\x8f\x61\x00\xa0\x30\xa7\x10\x62\x4e\xbf\x16\x64\xcb\xc3\xe1\x5a\x33\x67\x50\x50\x08\x95\xc9\x99\x2b\x9d\xa7\x3c\xa8\x8a\xbb\x8b\xf2\x9a\xd3\xdf\xaa\x28\x01\xc0\xba\x03\xb9\x3e\x8f\xc9\xe3\x79\x00\xe0\x1a\x10\xef\xb5\xf3\xd3\x7e\xd2\x26\x13\x1a\x76\x6d\x36\x4b\x4e\x17\x56\x90\xab\x62\x03\x9c\xb5\xf1\x1b\xbb\x65\xe7\xb7\xac\x6a\x70\xe4\x83\x86\x84\x4d\xa9\x1d\x86\xcb\x9e\x0e\x0b\xaf\xdf\x91\x22\x8b\x9e\x92\x10\xbc\x2d\xa8\xd5\x70\xb2\xa0\xdc\x48\xf4\x14\xc2\x89\xb1\xfa\x86\x84\x77\xa3\xed\xb6\x7d\xbc\xbd\x1d\xb5\x59\x2a\x61\xfb\xb8\xe4\xe4\xf6\xf6\xe4\x41\x80\x09\xaa\xd7\xf4\xd1\x51\xb2\xd0\x1f\xd0\x8b\x6c\xd6\x56\xd1\x54\xbb\x42\xe9\xba\xa4\x5d\x7d\xe3\x35\xb2\xac\x48\x1f\xa9\xb1\x73\xe4\x23\xb5\x26\xe5\xb5\x2d\xfb\xe6\xdd\xbc\xda\xf9\x85\xad\xb8\x6b\xfb\x1d\xaa\x9d\x42\x62\x4c\xd2\x85\xed\xdf\x7b\x59\xa3\xcb\xbb\x08\x1e\x6d\x4a\xfe\x2d\x93\x4c\x0e\x04\xca\x98\x2c\x5a\x91\xb1\x40\x39\xa3\x15\x59\x52\xbb\xa1\x34\x83\xf1\xa5\xa1\x10\xda\x66\xed\xe4\x00\x5f\xa8\xdc\x89\x67\xb4\xba\x87\xe0\x91\x58\xab\x16\xc7\x73\xc3\x55\x9f\x84\x9c\xb0\x6c\x7c\x4d\xb6\x3f\xce\xf6\x34\x00\x8b\x8c\x3a\x27\xf0\x19\xfa\xe6\x2c\xec\x40\x40\x4c\x52\xab\xd4\x81\xd7\xc3\x3d\xb7\x74\xfd\x25\xec\x45\x69\x78\xf5\xeb\x03\x4c\x07\x49\xbe\x17\xc7\xea\xc2\x84\xbb\x6c\x39\x2a\x4c\xc9\x3e\xba\x60\x4f\xb2\x7c\x89\x42\x90\x73\xdf\x23\xfb\x78\xdf\xe2\x07\x50\x7e\xb4\xdd\x5a\x2d\xa9\x7e\x28\x1c\xd9\x65\x5c\x2e\x29\x47\x96\xb5\xa4\x2e\xb8\x2f\x4a\x74\x8e\xac\xea\xc7\xfa\x18\x44\xb9\xac\xad\x6a\x09\x63\xbe\xcc\x29\x8f\xc9\xd6\xff\xae\x99\x36\x8d\x53\x93\xb1\x4d\x57\x78\x56\x74\x4c\xdb\xd6\x42\x7a\xee\x4e\xf5\xcd\x5f\x16\xea\x08\x16\xea\x46\xc7\x07\xb7\xe8\xe7\x9d\xfc\x9f\xec\xce\x8d\x8e\x2b\xc1\x8d\x8e\x9b\x9d\x91\x5a\xd4\xdf\xe7\xa3\xed\xb6\x7b\x3c\x26\x3a\x37\x38\x9f\xcb\xe6\x9e\xf5\x0b\x99\x8f\x80\xcc\xbe\x62\xc0\x41\x3a\x2f\xf6\x34\xdf\x21\x34\xe6\x13\xad\x56\xbc\x47\x0c\xa3\x25\x8b\x72\xfa\x78\xba\x65\x75\x74\x37\x56\xf7\x9c\x3e\xd4\xe7\xf9\x53\xae\xcd\xa9\xbf\x73\xb5\xdd\xbc\x9b\x01\xef\xf5\xa9\x29\xb9\x75\x4c\xf6\x14\x0d\xb1\x2a\xf5\x4e\xe8\x0a\x63\xb4\xf5\x6e\xa2\x55\xc2\xf5\x05\xfa\x5f\xfa\xd2\xab\xd3\x57\xd2\x0e\xc7\x31\xad\xef\x83\xde\xfc\xbd\x9b\xdf\xbd\x30\x2f\x9b\x7d\x04\x9b\xdd\xbf\x39\xf4\x77\x7b\xd6\xd3\xfd\x88\xab\x5e\x9b\xdd\xdd\x5d\xc2\x8e\x8c\xf9\x77\xa8\xfe\xe2\x55\xec\x85\xe6\xff\x09\xcd\x9b\x26\xd4\x95\xb6\x44\xeb\x57\xba\x6b\xcf\x81\x9f\xb3\x4f\xf7\x21\xba\xec\xde\x4f\xb5\xee\x20\xb4\xf2\xc8\xf5\x6b\x96\xa6\x2f\xcd\xf4\x9f\xd1\x8b\xc3\xaf\x10\x9e\xdf\x93\x27\x5e\xa6\xfc\x19\x00\x00\xff\xff\x9e\x22\x79\x0a\x03\x14\x00\x00"),
		},
		"/datacatalog.yaml": &vfsgen۰CompressedFileInfo{
			name:             "datacatalog.yaml",
			modTime:          time.Time{},
			uncompressedSize: 5707,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4f\x6f\xe2\xc8\x13\xbd\xfb\x53\x94\xe0\xf2\xfb\x49\x09\x24\x39\x7a\x4f\x0c\xc9\x64\xac\x4d\xc8\x08\xc8\x8e\xe6\x84\x8a\x76\x61\xf7\xc6\xee\xee\xed\x2e\x43\x10\xcb\x77\x5f\xf9\x0f\x60\x02\x21\x64\x32\xda\x5d\xad\xf0\x29\xe9\xae\x7e\x55\xf5\xba\xde\x6b\x89\x26\x74\xb5\x99\x5b\x19\xc5\x0c\x57\x17\x57\x57\x70\xab\x75\x94\x10\xdc\xdd\x75\xbd\xa6\xd7\x84\x3b\x29\x48\x39\x0a\x21\x53\x21\x59\xe0\x98\xa0\x63\x50\xc4\xb4\xda\x39\x83\xdf\xc8\x3a\xa9\x15\x5c\xb5\x2e\xe0\x7f\x79\x40\xa3\xda\x6a\xfc\xff\x17\xaf\x09\x73\x9d\x41\x8a\x73\x50\x9a\x21\x73\x04\x1c\x4b\x07\x13\x99\x10\xd0\xb3\x20\xc3\x20\x15\x08\x9d\x9a\x44\xa2\x12\x04\x33\xc9\x71\x91\xa6\x02\x69\x79\x4d\xf8\x5e\x41\xe8\x31\xa3\x54\x80\x20\xb4\x99\x83\x9e\xd4\xe3\x00\xb9\x28\xb8\xf8\x62\x66\xe3\xb7\xdb\xb3\xd9\xac\x85\x45\xb5\x2d\x6d\xa3\x76\x52\x46\xba\xf6\x5d\xd0\xbd\xe9\x0d\x6e\xce\xaf\x5a\x17\xc5\x99\x47\x95\x90\x73\x60\xe9\x8f\x4c\x5a\x0a\x61\x3c\x07\x34\x26\x91\x02\xc7\x09\x41\x82\x33\xd0\x16\x30\xb2\x44\x21\xb0\xce\x0b\x9e\x59\xc9\x52\x45\x67\xe0\xf4\x84\x67\x68\xc9\x6b\x42\x28\x1d\x5b\x39\xce\x78\x8b\xad\x55\x79\xd2\x6d\x05\x68\x05\xa8\xa0\xd1\x19\x40\x30\x68\xc0\xa7\xce\x20\x18\x9c\x79\x4d\xf8\x16\x0c\xbf\x3c\x3c\x0e\xe1\x5b\xa7\xdf\xef\xf4\x86\xc1\xcd\x00\x1e\xfa\xd0\x7d\xe8\x5d\x07\xc3\xe0\xa1\x37\x80\x87\xcf\xd0\xe9\x7d\x87\x5f\x83\xde\xf5\x19\x90\xe4\x98\x2c\xd0\xb3\xb1\x79\xfd\xda\x82\xcc\x79\xa4\x30\x27\x6d\x40\xb4\x55\xc0\x44\x97\x05\x39\x43\x42\x4e\xa4\x80\x04\x55\x94\x61\x44\x10\xe9\x29\x59\x25\x55\x04\x86\x6c\x2a\x5d\x7e\x9b\x0e\x50\x85\x5e\x13\x12\x99\x4a\x46\x2e\x56\x76\x9a\x6a\x79\x1e\x1a\x59\xdd\xbf\x0f\x42\x5b\x6a\x09\x65\xd3\x96\x48\x74\x16\xb6\xa2\x62\x94\x5a\x42\xa7\xed\xe9\x25\x26\x26\xc6\x4b\xef\x49\xaa\xd0\x87\x01\xd9\xa9\x14\x74\x8f\xc6\x48\x15\x79\x29\x31\x86\xc8\xe8\x7b\x00\x0a\x53\xf2\x21\xff\x4f\x20\x63\xa2\xa3\xfd\x88\x55\xa4\x33\x28\xc8\x87\x3c\xe4\xdc\xcd\x1d\x53\xea\xe5\xfd\x6d\x80\xae\x91\xb1\x5b\x02\x79\x00\xd3\x55\xa9\xd3\xcb\x31\x31\x5e\x7a\x00\xae\x2c\xe5\x8b\x76\xdc\xdb\x49\x5d\xe6\x43\x23\x5d\x95\xd3\x92\xd3\x99\x15\xe4\xf2\x0c\x00\xe7\x55\x96\x32\x6e\x94\x1f\x1d\x55\x67\x47\xa4\xd8\xce\x47\x91\xd5\x99\xf1\xca\xa9\x2c\x7b\xaf\x55\x74\x93\x87\xdc\xd6\x22\x30\x63\x7d\x4b\x8a\x2c\x32\x85\x3e\xb0\xcd\xa8\xda\x91\x98\x76\xb5\x9a\xc8\xc8\xaf\x16\x00\x8c\x4e\xa4\x98\xf7\x8e\x29\x60\x24\x31\x1d\x95\xf1\x2f\x8e\xdf\x53\x3a\x26\x7b\x3c\x48\x5a\xc4\xaf\x41\x2c\x4d\xc8\x92\x12\xf4\x59\x52\x12\x6e\x6a\x5b\xd1\xbf\x4b\x42\xfe\xf1\xdc\x90\x0f\x32\x5c\x2f\xb9\xcc\x18\x6d\xd9\x75\xb5\x0a\x65\x31\x6c\x3e\x4c\x30\x71\xeb\xee\xc3\x21\xa5\x26\x41\x26\x1f\x1a\x8b\x45\x8e\xbd\x5c\x36\x76\x36\xbb\xa8\x3e\xd1\xa3\xa3\x70\xa8\xef\x91\x45\xdc\xaf\x6e\xab\x6c\xaf\x0e\xb8\xba\xc7\xce\x14\x65\x92\xab\x3c\x50\x1d\xe7\x88\x03\x35\x25\xc5\xda\xce\xb7\xc3\xf3\x21\x21\xbb\xbe\x99\xe0\xba\x6c\xb7\x68\xf2\x05\x62\x70\xbd\x61\x81\xd1\x46\xc4\xbb\xb1\x00\x53\x4c\x32\xaa\xf5\x64\xac\xfe\x9d\x04\xbb\xf6\x62\x51\xfd\xb9\x5c\xb6\x13\x2d\x4a\xe5\xb5\x17\x0b\x4b\x91\xd4\x6a\xb9\x6c\xd3\x7a\x66\xf2\xe5\x02\x67\x43\x45\x2c\xc9\xa2\x15\xb1\x14\x98\xf4\x57\x57\xe3\x36\x05\x9d\x57\xd4\x57\x39\x6a\x57\xf2\x44\xf3\xf5\x72\x9f\x26\x2f\x9a\x7a\x05\x6b\x52\xb5\x76\x2c\x5c\xfe\x85\xe4\x84\x95\x86\x0b\x19\xfe\x79\x5e\xdb\x01\x18\xc6\xb4\x3a\x04\x1c\x23\x97\x6f\xc5\xaa\x08\x18\x53\xa2\x55\xe4\x80\x75\xab\x76\x2c\x9a\x3e\xf9\x5b\x28\xa5\xcc\xbe\xee\xd4\xb4\x57\xfe\x35\x9c\x9c\x53\x7f\x9d\x2d\x45\x85\x11\xd9\x57\xdd\xe7\x08\xf1\x1f\x94\xfd\xdb\x8a\xff\xaf\xcf\xfc\x62\x51\x73\x87\x6a\xb2\x25\xed\x4c\xf5\x21\x9a\x19\xa3\x11\x57\x88\xaf\xb1\x3d\xc4\x68\xb8\x1d\xf2\x53\x5d\xb6\x5e\xc2\x8f\xdb\xec\x0e\xca\x3b\x7d\x76\x0f\x11\x70\x32\xda\x0f\x19\x2d\x6f\xe6\xe6\xe4\xb4\xc5\xf7\xaf\x74\x5a\xc6\xe8\x80\xf2\x4f\x2e\xbb\x58\x18\xb4\xa4\xb8\x9c\xe8\x77\xba\xeb\xb3\x56\x3a\x7d\xf5\x1d\x1b\x6e\xef\xff\x64\x5b\x2d\xb1\x3f\x62\xa9\x35\x84\x77\xdb\xe9\x56\x67\x70\xb2\xd2\x0f\x5a\x69\x41\xe7\x9e\xc7\xfd\x64\xa4\x35\x69\xfd\xb3\x46\x5a\xaa\xea\x90\x9f\x7e\x2d\x22\x8e\x72\xd5\xf7\x0b\x7e\x93\xfe\xc7\x25\xff\x02\xe3\x9d\xa2\xdf\x21\x00\x4e\xb2\xdf\xfb\xa0\xac\xdc\x71\xb9\x6c\x9b\xd5\x48\xec\x51\xf6\x61\x25\x96\x72\x2b\x9e\xa6\xf5\x5c\x6d\xcb\x6e\xa3\xd5\x22\x6a\x77\x40\x0f\x48\xe9\xc0\xd4\x1e\xaf\xab\x63\x7e\x84\x7a\x9b\xbf\xaa\xd5\x15\x69\xaf\xf4\xb8\xe7\xc5\x79\xbb\xb5\xe1\xee\xa1\xbf\xb3\x33\xa8\x6e\xa6\x92\xfe\x5f\x01\x00\x00\xff\xff\xe8\x37\xf0\xea\x4b\x16\x00\x00"),
		},
		"/tags.yaml": &vfsgen۰CompressedFileInfo{
			name:             "tags.yaml",
			modTime:          time.Time{},
			uncompressedSize: 3058,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\x4d\x6f\xe2\x48\x10\xbd\xfb\x57\x94\xe0\xb2\x2b\x25\x66\xc2\xd1\x7b\x22\x24\x33\x63\x4d\x02\x2b\x20\x19\xcd\x29\x2a\xda\x65\xd3\x1b\xbb\xbb\xb7\xbb\x0c\x63\x8d\xe6\xbf\xaf\xdc\x36\x60\x08\x64\x3e\x34\xd2\x6a\xa5\xf5\x81\x43\xd7\xab\x8f\x57\xf5\xaa\xe8\xc3\x58\x9b\xca\xca\x6c\xc5\x30\x7c\x33\x1c\xc2\x3b\xad\xb3\x9c\xe0\xee\x6e\x1c\xf4\x83\x3e\xdc\x49\x41\xca\x51\x02\xa5\x4a\xc8\x02\xaf\x08\x46\x06\xc5\x8a\xb6\x96\x0b\x78\x24\xeb\xa4\x56\x30\x0c\xdf\xc0\x6f\x35\xa0\xd7\x9a\x7a\xbf\xff\x11\xf4\xa1\xd2\x25\x14\x58\x81\xd2\x0c\xa5\x23\xe0\x95\x74\x90\xca\x9c\x80\x3e\x0b\x32\x0c\x52\x81\xd0\x85\xc9\x25\x2a\x41\xb0\x91\xbc\xf2\x69\xda\x20\x61\xd0\x87\x4f\x6d\x08\xbd\x64\x94\x0a\x10\x84\x36\x15\xe8\xb4\x8b\x03\x64\x5f\xb0\xff\x56\xcc\x26\x1a\x0c\x36\x9b\x4d\x88\xbe\xda\x50\xdb\x6c\x90\x37\x48\x37\xb8\x8b\xc7\xb7\x93\xf9\xed\xe5\x30\x7c\xe3\x7d\x1e\x54\x4e\xce\x81\xa5\xbf\x4b\x69\x29\x81\x65\x05\x68\x4c\x2e\x05\x2e\x73\x82\x1c\x37\xa0\x2d\x60\x66\x89\x12\x60\x5d\x17\xbc\xb1\x92\xa5\xca\x2e\xc0\xe9\x94\x37\x68\x29\xe8\x43\x22\x1d\x5b\xb9\x2c\xf9\xa0\x5b\xdb\xf2\xa4\x3b\x00\x68\x05\xa8\xa0\x37\x9a\x43\x3c\xef\xc1\xf5\x68\x1e\xcf\x2f\x82\x3e\x7c\x8c\x17\xef\xa7\x0f\x0b\xf8\x38\x9a\xcd\x46\x93\x45\x7c\x3b\x87\xe9\x0c\xc6\xd3\xc9\x4d\xbc\x88\xa7\x93\x39\x4c\xdf\xc2\x68\xf2\x09\x3e\xc4\x93\x9b\x0b\x20\xc9\x2b\xb2\x40\x9f\x8d\xad\xeb\xd7\x16\x64\xdd\x47\x4a\xea\xa6\xcd\x89\x0e\x0a\x48\x75\x53\x90\x33\x24\x64\x2a\x05\xe4\xa8\xb2\x12\x33\x82\x4c\xaf\xc9\x2a\xa9\x32\x30\x64\x0b\xe9\xea\x69\x3a\x40\x95\x04\x7d\xc8\x65\x21\x19\xd9\xbf\xbc\x20\x15\x06\x01\x1a\xd9\xce\x3f\x02\xa1\x2d\x85\x42\xd9\x22\x14\xb9\x2e\x93\x30\xf3\x52\x0a\x85\x2e\x06\xeb\x2b\xcc\xcd\x0a\xaf\x82\x67\xa9\x92\x08\xe6\x64\xd7\x52\xd0\x3d\x1a\x23\x55\x16\x14\xc4\x98\x20\x63\x14\x00\x28\x2c\x28\x02\xc6\xcc\x9d\x0e\xd5\x42\x9c\x41\x41\x11\xd4\x90\x4b\x57\x39\xa6\x22\xa8\x89\xed\x23\x2c\x30\x73\x01\xc0\x7a\x5b\xdc\xfa\x6a\x49\x8c\x57\x01\x80\x6b\x92\xbf\xd7\x8e\x27\xfb\x64\x4d\x06\x34\xd2\xb5\x59\x2c\x39\x5d\x5a\x41\xae\x8e\x09\x70\xd9\xc6\x6d\x70\x4f\xb5\x4f\xfd\xf3\xf4\x4c\x55\xd0\x88\xae\xa1\x56\xe7\x5d\x60\xf6\x61\xf7\x8c\x25\xeb\x77\xa4\xc8\x22\x53\x12\x01\xdb\x92\x5a\x8b\xc4\x62\xac\x55\x2a\xb3\xa8\x7d\x00\x30\x3a\x97\xa2\x9a\x9c\xcb\xf4\x24\xb1\x78\x6a\x30\x47\x2e\xf7\x54\x2c\xc9\xbe\xee\x58\x78\xcc\xce\xd1\x52\x4a\x96\x94\xa0\xb7\x92\xf2\x64\x5f\x43\x67\x06\x1d\x76\xf5\xc7\x95\xa1\xc8\x5b\x77\x8f\xae\x34\x46\x5b\x76\x63\xad\x12\xe9\x65\x12\x41\x8a\xb9\xdb\x71\x4c\x16\x54\x98\x1c\x99\x22\xe8\xb1\xef\x8b\x1b\x7c\xf9\x52\xc7\xf8\xfa\xb5\xf7\x02\x34\x46\x75\x4d\x0f\x8e\x92\x85\xbe\x47\x16\xab\x59\x3b\x84\x86\x58\x37\xf0\x76\x3c\xa3\x35\xca\xbc\xde\xd3\x58\x8d\x9c\x23\x8e\xd5\x9a\x14\x6b\x5b\x1d\xc2\xeb\xa1\x93\xdd\xcd\x21\xbe\x69\x48\x77\xc9\x6c\x23\xc6\x37\xfb\x5e\x30\xda\x8c\xf8\x18\x7b\x4e\x0b\x6b\xcc\x77\xc3\x3d\x50\xc3\x63\xc7\xf0\xcb\xf4\xe0\xb3\xfd\x9c\x22\xf6\xae\x3f\xa1\x89\x2e\xcb\x5f\xa4\x0a\xdf\x9f\xff\xbe\x2e\xf6\xd8\xd9\xb6\x8f\x6e\xef\x73\x09\xcf\x54\x45\x60\xd0\x92\xe2\x19\xa5\xdd\x16\xa6\x6d\xa8\xc6\xd8\xb1\x64\xeb\xe7\xee\x14\xce\x5e\x99\xe6\x3b\x71\xed\x3a\xa1\xac\x2e\xcd\x37\x4e\xeb\xb7\xf8\xf9\x24\xf5\xb0\x4e\x6e\xb5\xb7\xec\xc7\xe7\xb5\xe8\x09\x75\x64\x7e\x6e\x77\x96\x52\x25\xf5\x7f\xc1\x09\x96\xd7\x07\xa6\x57\xf6\xe7\x48\x55\xad\xdf\xff\xba\xfa\x7e\x5d\xfd\x69\xf5\x5f\x24\xf8\x87\x45\xb5\xad\xaf\x40\x85\x19\xd9\x1f\xd4\x57\x79\x70\x83\x5e\x2a\x6c\x30\xf0\xa1\x8e\x93\x1c\xfe\x65\x0f\x4c\x53\xfb\xf7\x0a\xb1\xd3\xba\xed\x05\x3a\xd3\xbc\x53\x47\xef\xf5\xbd\x7c\x3c\x42\xff\x7b\x9b\xb9\xbb\xac\xbb\x96\xfc\x13\x00\x00\xff\xff\xa2\x71\x11\x39\xf2\x0b\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/bigquery.yaml"].(os.FileInfo),
		fs["/datacatalog.yaml"].(os.FileInfo),
		fs["/tags.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
