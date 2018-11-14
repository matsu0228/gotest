package main

import (
	"strings"
	"testing"

	"github.com/matsu0228/gotest/integrate/repository"
)

func TestCalcurate(t *testing.T) {
	trepo := repoMock{}

	want := 300

	got, err := calcurate(trepo, []string{"t_100", "t_200"})
	if err != nil {
		t.Fatalf("calcurate() fails err:%v", err)
	}
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

// repoMock is struct of Repository
type repoMock struct{}

func (r repoMock) Get(title string) ([]repository.SomeData, error) {
	// testing local rule
	ary := strings.Split(title, "_")
	if len(ary) == 0 {
		return []repository.SomeData{}, nil
	}
	return []repository.SomeData{
		{Body: ary[1]},
	}, nil
}

func (r repoMock) Save(title, body string) error {
	return nil
}
