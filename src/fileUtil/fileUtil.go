package fileUtil

import (
	"fmt"
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
)

type Shader struct {
	VERTEX string
	FRAGMENT string
}

type Configuration struct {
	WINDOW_WIDTH int
	WINDOW_HEIGHT int
	NAME string
	SHADER Shader
	FRAME_CAP int
}


func  LoadConfig() (configuration Configuration){
	env := "config/dev.json"

	if os.Getenv("ENV") != "" {
		env = "config/"+ os.Getenv("ENV") + ".json"
	}

	fmt.Println("Loading config:", env)

	file, err := os.Open(env)
	if err != nil {
		log.Fatalln("failed to open config:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return 
}

func LoadShader(fileName string) (shader string){

	fmt.Println("Loading shader:", fileName)
	bytes, err:= ioutil.ReadFile(fileName)

	shader = string(bytes) + "\x00"
	if err != nil {
		log.Fatalln("failed to open config:", err)
	}
	return 
}