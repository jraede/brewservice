package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const CACHE_LOCATION = "/usr/local/Cellar/.brewservicecache"

func loadCacheOrFail() Cache {
	cache, err := LoadCache()
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Services have not been parsed yet. Did you run `brewservice update`?")
			os.Exit(0)
		} else {
			fatal(err)
		}
	}

	return cache
}

type Cache map[string]Service

func (c Cache) Save() error {
	file, err := os.OpenFile(CACHE_LOCATION, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(c)
}

func LoadCache() (cache Cache, err error) {
	var file *os.File
	file, err = os.Open(CACHE_LOCATION)
	if err != nil {
		return
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cache)
	return cache, err
}

func readdir(dirname string) ([]string, error) {
	dir, err := os.Open(dirname)

	if err != nil {
		return nil, err
	}

	matched := []string{}

	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.Mode().IsDir() {
			children, err := readdir(dirname + file.Name() + "/")
			if err != nil {
				return nil, err
			}
			matched = append(matched, children...)

		} else if file.Mode().IsRegular() {

			if strings.HasPrefix(file.Name(), "homebrew") && strings.HasSuffix(file.Name(), ".plist") {
				matched = append(matched, dirname+file.Name())
			}
		}
	}

	return matched, nil
}

func update(dirname string) (Cache, error) {

	matched, err := readdir(dirname)

	services := Cache{}
	if err != nil {
		return services, err
	}

	for _, m := range matched {
		spl := strings.Split(m, "/")
		services[spl[4]] = Service(m)
	}

	return services, nil
}
