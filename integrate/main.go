package main

import (
	"log"
	"strconv"

	"github.com/matsu0228/gotest/integrate/repository"
)

// Repository is interface of datastore
type Repository interface {
	Get(title string) ([]repository.SomeData, error)
	Save(title, body string) error
}

func errorExit(err error) {
	log.Fatal(err)
}

// calcurate of keys data
func calcurate(repo Repository, keys []string) (int, error) {

	sum := 0

	// gather data
	var values []repository.SomeData
	for _, key := range keys {
		data, err := repo.Get(key)
		if err != nil {
			return 0, err
		}
		values = append(values, data...)
	}

	// calcurate of data
	for _, v := range values {
		i, err := strconv.Atoi(v.Body)
		if err != nil {
			return 0, err
		}
		log.Printf("[INFO] sum = sum:%v + i:%v", sum, i)
		sum = sum + i
	}
	return sum, nil
}

func main() {

	repo, err := repository.NewDatabase("root", "mysql", "docker-mysql", "3306", "todo", "?parseTime=true&loc=Japan")
	if err != nil {
		errorExit(err)
	}

	// save
	err = repo.Save("someTitle", "sample body")
	if err != nil {
		errorExit(err)
	}

	// get
	var data []repository.SomeData
	data, err = repo.Get("someTitle")
	if err != nil {
		errorExit(err)
	}
	log.Printf("[INFO] num:%v, result = %v", len(data), data)

	// save
	err = repo.Save("t1", "100")
	if err != nil {
		errorExit(err)
	}
	err = repo.Save("t2", "200")
	if err != nil {
		errorExit(err)
	}
	sum, err := calcurate(repo, []string{"t1", "t2"})
	if err != nil {
		errorExit(err)
	}
	log.Printf("[INFO] sum:%v", sum)

}
