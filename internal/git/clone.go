package cli_git

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func CloneTemplateRepo(path string) {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:        "",
		Progress:   os.Stdout,
		NoCheckout: true,
	})

	if err != nil {
		log.Fatal(err)
	}
}
