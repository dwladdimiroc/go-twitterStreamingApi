package main

import (
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
	"os"
//	"fmt"
)

var configuration struct{
	ConsumerKey string `json:"ConsumerKey"`
	ConsumerSecret string `json:"ConsumerSecret"`
	AccessToken string `json:"AccessToken"`
	AccessTokenSecret string `json:"AccessTokenSecret"`
}

func printErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	configFile, err := os.Open("config.json")
	printErr(err)
		
	jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&configuration)
    printErr(err)
	
	anaconda.SetConsumerKey(configuration.ConsumerKey)
	anaconda.SetConsumerSecret(configuration.ConsumerSecret)
	client := anaconda.NewTwitterApi(configuration.AccessToken, configuration.AccessTokenSecret)

	stream := client.PublicStreamSample(nil)
	
	file, _ := os.Create("TerremotoChile-25Dic2016.json")
	defer file.Close()

	for {
		status := <-stream.C

		jsonTweet, err := json.Marshal(status)
		printErr(err)

		file.Write(jsonTweet)
		file.WriteString("\n")
		file.Sync()
	}
}
