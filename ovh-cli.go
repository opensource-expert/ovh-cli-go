package main

import (
	//"encoding/json"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"context"
	"github.com/ovh/go-ovh/ovh"
	"io/ioutil"
	"os"
)

// Instantiate an OVH client and get the firstname of the currently logged-in user.
// Visit https://api.ovh.com/createToken/index.cgi?GET=/me to get your credentials.
func main() {

	// Initialize client, read credential in ovh.conf
	// See documentation for locations: https://github.com/ovh/go-ovh#configuration
	client, err := ovh.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	// check input on stdin
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	// mode of the stdin
	fmt.Fprintf(os.Stderr, "%v\n", fi.Mode())

	if fi.Size() > 0 {
		fmt.Fprintf(os.Stderr, "there is something to read\n")
	} else {
		fmt.Fprintf(os.Stderr, "stdin is empty\n")
	}

	method := os.Args[1]
	path := os.Args[2]

	fmt.Fprintf(os.Stderr, "%s\n", path)

	// call API, extracted from CallAPIWithContext()
	ctx := context.Background()
	req, err := client.NewRequest(method, path, nil, true)
	if err != nil {
		panic(err)
	}
	req = req.WithContext(ctx)
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Read all the response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	} else {
		// output JSON
		fmt.Println(string(body))
	}
}
