package resource

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Resource struct {
	Name             string
	Path             string
	CompletePath     string
	ClientScripts    []string
	ServerScripts    []string
	MapClientScripts map[string]string
	MapServerScripts map[string]string
}

var ClientScript = "client"
var ServerScript = "server"

func New(name string, path string) (*Resource, error) {
	completePath := fmt.Sprintf("%v/%v", path, name)

	rs := &Resource{
		Name:             name,
		Path:             path,
		CompletePath:     completePath,
		ClientScripts:    []string{},
		ServerScripts:    []string{},
		MapClientScripts: map[string]string{},
		MapServerScripts: map[string]string{},
	}

	err := os.Mkdir(rs.CompletePath, 0777)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (r *Resource) AddFile(name string, data []byte, script string) error {
	filePath := fmt.Sprintf("%v/%v", r.CompletePath, name)
	path := filepath.Base(filePath)

	err := ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		return err
	}

	if script == ClientScript {
		r.MapClientScripts[name] = path
	}

	if script == ServerScript {
		r.MapServerScripts[name] = path
	}

	return nil
}

func (r *Resource) AddFolder(name string) error {
	dirPath := fmt.Sprintf("%v/%v", r.CompletePath, name)
	err := os.Mkdir(dirPath, 0777)

	if err != nil {
		return err
	}

	return nil
}

func (r *Resource) CreateManifest() {
	filePath := fmt.Sprintf("%v/%v", r.CompletePath, "fxmanifest.lua")
	var clientScripts = []string{}
	var serverScripts = []string{}

	var cs string = ""
	var ss string = ""

	if len(r.MapClientScripts) > 0 {
		for key := range r.MapClientScripts {
			script := fmt.Sprintf("%q", key)
			s := append(clientScripts, script)
			clientScripts = s
		}

		clientScriptsJoined := strings.Join(clientScripts, ", ")
		cs = fmt.Sprintf("client_scripts {\n%v\n}", clientScriptsJoined)
	}

	if len(r.MapServerScripts) > 0 {
		for key := range r.MapServerScripts {
			script := fmt.Sprintf("%q", key)
			s := append(serverScripts, script)
			serverScripts = s
		}

		serverScriptsJoined := strings.Join(serverScripts, ", ")
		ss = fmt.Sprintf("server_scripts {\n%v\n}", serverScriptsJoined)
	}

	manifestString := fmt.Sprintf("%v\n\n%v", cs, ss)

	err := ioutil.WriteFile(filePath, []byte(manifestString), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
