package repository_test

import (
	"bitsports/pkg/docker"
	"bitsports/testutil"
	"fmt"
	"testing"
)

var container *docker.Container

func TestMain(m *testing.M) {
	var err error
	container, err = testutil.StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer testutil.StopDB(container)

	m.Run()
}

