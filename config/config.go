package config

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

//NewYaml ...
func NewYaml(file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		return err
	}

	yaml.NewEncoder(f).Encode(Config{})

	return nil
}

//LoadYaml ...
func LoadYaml(file string, updateCallback func(Config)) (cancel func(), err error) {
	initStat, err := os.Stat(file)
	if err != nil {
		return
	}

	config, err := readFile(file)
	if err != nil {
		return
	}
	updateCallback(config)

	ticker := time.NewTicker(time.Second)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if stat, err := os.Stat(file); err == nil && stat.ModTime() != initStat.ModTime() {
					initStat = stat
					config, err := readFile(file)
					if err != nil {
						log.Print(err)
						log.Printf("error on reading config.yaml")
						continue
					}
					updateCallback(config)
				}
			}
		}
	}()

	return
}

func readFile(path string) (config Config, err error) {
	log.Printf("reading %s", path)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(f, &config)
	return
}
