package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/consul/api"
)

type statusPage struct {
	ID          string
	Name        string
	Group_id    string
	Description string
}

var statusID = "teststatus"
var ChannelID = "test-chan2"

func addApps() {
       //  Will read in a json of statuspage components and put them in consul
	filePath := "./components.json"
	fmt.Printf("// reading file %s\n", filePath)
	file, err1 := ioutil.ReadFile(filePath)
	if err1 != nil {
		fmt.Printf("// error while reading file %s\n", filePath)
		fmt.Printf("File error: %v\n", err1)
		os.Exit(1)
	}

	fmt.Println("// defining array of struct StatusPage")
	var components []statusPage

	err2 := json.Unmarshal(file, &components)
	if err2 != nil {
		fmt.Println("error:", err2)
		os.Exit(1)
	}

	fmt.Println("// loop over array of structs of StatusPage")
	for k := range components {
		fmt.Printf("The component '%s' is: '%s'\n", components[k].ID, components[k].Name)
		fmt.Printf("The group is '%s' and description: %s\n", components[k].Group_id, components[k].Description)
	}
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	for k := range components {
		fmt.Printf("Inserting component '%s' is: '%s'\n", components[k].ID, components[k].Name)
		//	p := &api.KVPair{Key: "sre/REDIS_MAXCLIENTS", Value: []byte("1000")}
		// Txt := fmt.Sprintf("Description: %s ", slackout.Submission["description"])
		//Keyid := fmt.Println("sre/", components[k].ID)
		//   keyname := components[k].Name
		p := &api.KVPair{Key: ("ComponentId/" + components[k].ID), Value: []byte(components[k].Name)}

		_, err = kv.Put(p, nil)
		if err != nil {
			panic(err)
		}
	}
}
func subscribeUsers() {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "component/testcomp/testuser", Value: []byte("testname")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("sre/REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}

// func addToChannel() {
// 	//JoinChannel will create the channel if it's not already there in our authenticated Slack
// 	// Get a new client
// 	client, err := api.NewClient(api.DefaultConfig())
// 	if err != nil {
// 		panic(err)
// 	}
// 	kv := client.KV()
// 	pair, _, err := kv.Get("component/testcomp/", nil)
// 	channel, err := *slack.Client.InviteUserToChannel(ChannelID, pair)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//channel, err := *slack.Client.InviteUserToChannel(channelID, pair.Value string)
// 	// if err != nil {
// 	// 	fmt.Println("Channel Created ", channel)
// 	// 	return code, err.Error()
// 	// }
// }
func main() {
	addApps()
	subscribeUsers()
}
