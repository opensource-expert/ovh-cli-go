package main

import (
	"context"
	//"encoding/json"
	// because I miss OVH internal to pass JSON directly
	"fmt"
	"github.com/acrazing/cheapjson"
	//"github.com/davecgh/go-spew/spew"
	"github.com/ovh/go-ovh/ovh"
	"io/ioutil"
	"os"
)

/*
// for OVH imported code
import (
	"bytes"
	"crypto/sha1"
	"net/http"
	"strconv"
)

type MyClient ovh.Client

// From OVH code, NewRequest returns a new HTTP request
func (c *MyClient) OurNewRequest(method, path string, reqBody interface{}, needAuth bool) (*http.Request, error) {
	var body []byte
	var err error

	// our reqBody is alreay a JSON string
	//if reqBody != nil {
	//	body, err = json.Marshal(reqBody)
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	target := fmt.Sprintf("%s%s", c.endpoint, path)
	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Inject headers
	if body != nil {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("X-Ovh-Application", c.AppKey)
	req.Header.Add("Accept", "application/json")

	// Inject signature. Some methods do not need authentication, especially /time,
	// /auth and some /order methods are actually broken if authenticated.
	if needAuth {
		timeDelta, err := c.TimeDelta()
		if err != nil {
			return nil, err
		}

		timestamp := getLocalTime().Add(-timeDelta).Unix()

		req.Header.Add("X-Ovh-Timestamp", strconv.FormatInt(timestamp, 10))
		req.Header.Add("X-Ovh-Consumer", c.ConsumerKey)

		h := sha1.New()
		h.Write([]byte(fmt.Sprintf("%s+%s+%s+%s%s+%s+%d",
			c.AppSecret,
			c.ConsumerKey,
			method,
			c.endpoint,
			path,
			body,
			timestamp,
		)))
		req.Header.Add("X-Ovh-Signature", fmt.Sprintf("$1$%x", h.Sum(nil)))
	}

	// Send the request with requested timeout
	c.Client.Timeout = c.Timeout

	return req, nil
}
*/

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

	var bytes []byte

	if fi.Mode()&os.ModeNamedPipe != 0 {
		// there is something to read
		bytes, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	}

	// type PutProjectIn struct {
	// 	Description string
	// }
	// var m PutProjectIn
	// err = json.Unmarshal(bytes, &m)
	// spew.Dump(m)

	// ARGUMENTS
	method := os.Args[1]
	path := os.Args[2]
	if len(os.Args) > 3 {
		bytes = []byte(os.Args[3])
	}

	fmt.Fprintf(os.Stderr, "stdin bytes '%v' %d '%s'\n", bytes, len(bytes), string(bytes))
	fmt.Fprintf(os.Stderr, "Call: %s %s stdin: %v\n", method, path, bytes != nil)
	var data interface{}
	if len(bytes) > 0 {
		json_parsed_value, err := cheapjson.Unmarshal(bytes)
		if err != nil {
			panic(err)
		}
		data = json_parsed_value.Value()
	}

	// call API, extracted from CallAPIWithContext()
	ctx := context.Background()
	req, err := client.NewRequest(method, path, data, true)
	if err != nil {
		panic(err)
	}
	req = req.WithContext(ctx)
	fmt.Fprintf(os.Stderr, "%v\n", req)
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Read all the response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	fmt.Fprintf(os.Stderr, "%v\n", response)

	if err != nil {
		panic(err)
	} else {
		// output JSON
		fmt.Println(string(body))
	}
}
