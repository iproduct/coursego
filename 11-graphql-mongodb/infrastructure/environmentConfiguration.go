package infrastructure

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"

	"gopkg.in/yaml.v2"
)

func (env *Environment) SetEnvironment() {
	_, filename, _, _ := runtime.Caller(1)
	env.path = path.Join(path.Dir(filename), "environment/Connection.yml")
	_, err := os.Stat(env.path)
	if err != nil {
		panic(err)
		return
	}
}

func (env *Environment) LoadConfig() {
	content, err := ioutil.ReadFile(env.path)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	err = yaml.Unmarshal([]byte(string(content)), env)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	if env.App.Debug == false {
		log.SetOutput(ioutil.Discard)
	}
	log.Println("Config load successfully!")
	return
}
