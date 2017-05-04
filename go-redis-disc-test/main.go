package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// struct to hold the values from the config.json.
type Config struct {
	Host  string   `json:"host"`
	Port  int      `json:"port"`
	Nodes []string `json:"nodes"`
}

// get the config.json file to read in the value to the struct.
func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func main() {
	fmt.Println("Starting the Application...")
	config, _ := LoadConfiguration("config.json")
	fmt.Println(config.Host)
	fmt.Println(len(config.Nodes))
	fmt.Println(config.Nodes)

	//fmt.Println(config.Nodes)
	//out := fmt.Sprintf("%s:%s", config.Database.Host, strconv.Itoa(config.Database.Port))
	//fmt.Println(out)

	/*client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Database.Host, strconv.Itoa(config.Database.Port)),
		Password: "", // no password set
		DB:       0,  // use default DB
	}) */
	/*
		client := redis.NewFailoverClient(&redis.FailoverOptions{
			SentinelAddrs: config.Nodes{":8100"},
			Password:      "", // no passwd set
			DB:            0,  // use default DB
		})

		fmt.Println("trying Redis...")
		pong, err := client.Ping().Result()
		fmt.Println(pong, err)
		// Output: PONG <nil>

		err1 := client.Set("key", "456", 0).Err()
		if err1 != nil {
			panic(err1)
		}
		val, err := client.Get("key").Result()
		if err == redis.Nil {
			fmt.Println("key does not exists")
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("key is", val)
		}

		val2, err2 := client.Get("key2").Result()
		if err2 == redis.Nil {
			fmt.Println("key2 does not exists")
		} else if err2 != nil {
			panic(err2)
		} else {
			fmt.Println("key2", val2)
		}
		// Output: key value
		// key2 does not exists
	*/
}
