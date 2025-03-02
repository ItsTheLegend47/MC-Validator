package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type mod struct {
	Hash              string
	Name              string
	Filename          string
	Expected_filename string
	Source_url        string
	Found             bool
}

type version_file struct {
	Found    bool
	Name     string
	Filename string
	Url      string
}

type modrinthresponse struct {
	GameVersions    []string    `json:"game_versions"`
	Loaders         []string    `json:"loaders"`
	ID              string      `json:"id"`
	ProjectID       string      `json:"project_id"`
	AuthorID        string      `json:"author_id"`
	Featured        bool        `json:"featured"`
	Name            string      `json:"name"`
	VersionNumber   string      `json:"version_number"`
	Changelog       string      `json:"changelog"`
	ChangelogURL    interface{} `json:"changelog_url"`
	DatePublished   string      `json:"date_published"`
	Downloads       int         `json:"downloads"`
	VersionType     string      `json:"version_type"`
	Status          string      `json:"status"`
	RequestedStatus interface{} `json:"requested_status"`
	Files           []struct {
		Hashes struct {
			Sha512 string `json:"sha512"`
			Sha1   string `json:"sha1"`
		} `json:"hashes"`
		URL      string      `json:"url"`
		Filename string      `json:"filename"`
		Primary  bool        `json:"primary"`
		Size     int         `json:"size"`
		FileType interface{} `json:"file_type"`
	} `json:"files"`
	Dependencies []struct {
		VersionID      interface{} `json:"version_id"`
		ProjectID      string      `json:"project_id"`
		FileName       interface{} `json:"file_name"`
		DependencyType string      `json:"dependency_type"`
	} `json:"dependencies"`
}

func checkModrinth(hash string) version_file {
	requestURL := fmt.Sprintf("https://api.modrinth.com/v2/version_file/%s", hash)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return version_file{Found: false, Name: "", Filename: "", Url: ""}
	}

	if res.StatusCode != 200 {
		return version_file{Found: false, Name: "", Filename: "", Url: ""}
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("could not read response body: %s\n", err)
		return version_file{Found: false, Name: "", Filename: "", Url: ""}
	}

	var data modrinthresponse

	err = json.Unmarshal(resBody, &data)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return version_file{Found: false, Name: "", Filename: "", Url: ""}
	}

	var name = data.Name
	var filename = data.Files[0].Filename
	var url = fmt.Sprintf("https://modrinth.com/mod/%s", data.ProjectID)

	return version_file{Found: true, Name: name, Filename: filename, Url: url}

}
