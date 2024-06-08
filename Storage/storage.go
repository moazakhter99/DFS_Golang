package storage

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)


func CASPathTransform(key string) (PathKey){
	hash := sha1.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashString) / blocksize
	paths := make([]string, sliceLen)

	for i:= 0; i < sliceLen; i++ {
		from, to := i * blocksize, (i * blocksize) + blocksize
		paths[i] = hashString[from:to]
	}

	return PathKey{
		PathName: strings.Join(paths, "/"),
		OrignalFileName: hashString,
	} 

}


type PathTransformerFunc func (string) (PathKey)  

type PathKey struct {
	PathName string
	OrignalFileName string

}

func (path PathKey) Filename() string {
	return fmt.Sprintf("%s/%s", path.PathName, path.OrignalFileName)
}

type StoreOpts struct {
	PathTransformerFunc PathTransformerFunc  
}


type Store struct {
	StoreOpts
}


func NewStore(storeOpts StoreOpts) (store *Store) {
	store = &Store{
		StoreOpts: storeOpts,
	}
	return
}

func (store *Store) WriteStream(key string, reader io.Reader) (err error) {
	pathName := store.PathTransformerFunc(key)
	err = os.MkdirAll(pathName.PathName, os.ModePerm)
	if err != nil {
		return
	}

	pathKey := pathName.Filename()
	file, err := os.Create(pathKey )
	if err != nil {
		return
	}

	nBytes, err := io.Copy(file, reader)

	fmt.Printf("Written (%d) bytes to the disk : %s", nBytes, pathKey )

	return

}