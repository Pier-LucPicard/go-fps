package configUtil

import (
	"fmt"
	"os"
	"log"
	"encoding/json"
)

type Configuration struct {
	WINDOW_WIDTH int
	WINDOW_HEIGHT int
	NAME string
}


func  LoadConfig() (configuration Configuration){
	env := "config/dev.json"

	if os.Getenv("ENV") != "" {
		env = "config/"+ os.Getenv("ENV") + ".json"
	}
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

func main(){
	
}