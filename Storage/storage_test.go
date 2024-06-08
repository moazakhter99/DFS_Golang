package storage_test

import (
	"bytes"
	"fmt"
	"testing"

	storage "github.com/moazakhter99/DFS_Golang/Storage"
)


func TestPathTransformerFunc(t *testing.T) {
	key := "Some Pictures"
	pathKey :=  storage.CASPathTransform(key)
	fmt.Println(pathKey)

}


func TestStore(t *testing.T) {
	opts := storage.StoreOpts{
		PathTransformerFunc: storage.CASPathTransform,
	}

	store := storage.NewStore(opts)

	data := bytes.NewReader([]byte("Some Images"))
	_ = store.WriteStream("mySpecialPicture", data)


}