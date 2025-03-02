package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func handle_upload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	id := uuid.New().String()
	path := filepath.Join("./uploads", id)

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}
	defer os.Remove(path)

	var mods []mod

	moddata, _ := getModHashes(path)
	for _, data := range moddata {
		version_file := checkModrinth(data.Hash)
		modfile := mod{Hash: data.Hash, Filename: data.Filename, Expected_filename: version_file.Filename, Source_url: version_file.Url, Found: version_file.Found, Name: version_file.Name}
		mods = append(mods, modfile)
	}

	c.JSON(http.StatusOK, mods)

}
