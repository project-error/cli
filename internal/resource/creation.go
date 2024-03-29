package resource

import (
	"log"
)

func CreateLuaResource(name string) {
	r, err := New(name, ".")
	if err != nil {
		log.Fatal(err)
	}

	r.AddFolder("client")
	r.AddFolder("server")

	r.AddFile("client/client.lua", []byte(""), ClientScript)
	r.AddFile("server/server.lua", []byte(""), ServerScript)

	r.CreateManifest()
};

func CreateJSResource(name string) {
	r, err := New(name, ".")
	if err != nil {
		log.Fatal(err)
	}

	r.AddFolder("client")
	r.AddFolder("server")

	r.AddFile("client/client.js", []byte(""), ClientScript)
	r.AddFile("server/client.js", []byte(""), ServerScript)

	r.CreateManifest()
}
