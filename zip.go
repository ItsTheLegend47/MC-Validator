package main

import (
	"archive/zip"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"path/filepath"
)

type modfile struct {
	Filename string
	Hash     string
}

func getModHashes(source string) ([]modfile, error) {

	var mods []modfile

	// Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// Iterate over .jar files and get hashes
	for _, file := range reader.File {
		if !file.FileInfo().IsDir() {

			name := file.FileInfo().Name()
			hash, _ := hashFile(file)

			if filepath.Ext(file.FileInfo().Name()) == ".jar" {
				mod := modfile{Filename: name, Hash: hex.EncodeToString(hash)}
				mods = append(mods, mod)
			}
		}
	}

	return mods, nil
}

func hashFile(f *zip.File) ([]byte, error) {

	// Ignore directorys
	if f.FileInfo().IsDir() {
		return nil, nil
	}

	// Unzip content and hash it
	zippedFile, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer zippedFile.Close()

	hash := sha512.New()

	if _, err := io.Copy(hash, zippedFile); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

//func main() {
//    err := unzipSource("testFolder.zip", "")
//    if err != nil {
//        log.Fatal(err)
//    }
//}
