// +build integration

package repository

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestConnectDatabase(t *testing.T) {
	if err := repo.db.Ping(); err != nil {
		t.Error(err)
	}
}

func TestSaveAndGet(t *testing.T) {

	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	want := "body_" + timestamp
	key := "title_" + timestamp
	err := repo.Save(key, want)
	if err != nil {
		t.Fatalf("cannot Save():%v", err)
	}

	data, err := repo.Get(key)
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 {
		t.Fatalf("none data of key:%v", key)
	}

	isExitst := false
	got := ""

	log.Printf("[TEST] TestSaveAndGet() want:%v, got:%v", want, data)
	for _, d := range data {
		got = d.Body
		if d.Body == want {
			isExitst = true
		}
	}
	if !isExitst {
		t.Fatalf("want %v, but %v:", want, got)
	}
}
