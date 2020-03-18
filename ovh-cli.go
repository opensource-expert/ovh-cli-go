package main

import (
	"encoding/json"
	"fmt"
	"github.com/ovh/go-ovh/ovh"
)

// Instantiate an OVH client and get the firstname of the currently logged-in user.
// Visit https://api.ovh.com/createToken/index.cgi?GET=/me to get your credentials.
func main() {
	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	var result map[string]interface{}
	client.Get("/me", &result)

	str, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(str))
	}
}
