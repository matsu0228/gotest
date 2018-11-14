package repository

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var repo Database

// TestMain is called first. see additional: https://golang.org/pkg/testing/#hdr-Main
func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	// teardown()

	os.Exit(exitCode)
}

func setup() {
	fmt.Println("called setup()")

	if integrateFlag { // setup when execute integrat test
		var err error
		repo, err = NewDatabase("root", "mysql", "docker-mysql", "3306", "todo", "?parseTime=true&loc=Japan")
		if err != nil {
			log.Fatal(err)
		}
	}
}
