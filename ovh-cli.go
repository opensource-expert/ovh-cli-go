// vim: set ts=4 sw=4 sts=4 et:
//
// ovh-cli is a command line tool for manipulating OVH api
// See: https://eu.api.ovh.com/console/
//
package main

import (
	"context"
	//"encoding/json"
	"fmt"
	// because I miss some OVH internal yet, to pass raw JSON string directly
	"github.com/acrazing/cheapjson"
	// data dumper
	//"github.com/davecgh/go-spew/spew"
	"github.com/docopt/docopt-go"
	"github.com/ovh/go-ovh/ovh"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var copyleft = `
Copyleft  (Ɔ)   2020 Sylvain Viart.
Issues:   https://github.com/opensource-expert/ovh-cli/issues
License   GPLv3 <https://www.gnu.org/licenses/gpl-3.0.txt>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
`

var Usage string = `Shell interface for OVH API

Usage:
  ovh-cli [--debug] METHOD URL_API [JSON_INPUT]

Options:
  -h, --help         This help message.
  -v, --version      Show program version and Licence.
  --debug            Output extra debug information on stderr.

Arguments:
  METHOD             GET | PUT | POST | DELETE to be passed to the API.
  URL_API            OVH URL of the API without '/1.0/' prefix.
  JSON_INPUT         Parameter for the API, it can be tansmitted via
                     command line argument. Or it can be passed from
                     stdin. If both are passed, then stdin will take
                     precedance.

Examples:
  ovh-cli  GET /me
  ovh-cli  GET /cloud/project/
  ovh-cli  GET /cloud/project/$project_id
  echo '{ "description" : "change-project-name-here"}' | ovh-cli  PUT /cloud/project/$project_id
  # othe form as argument JSON_INPUT
  ovh-cli  PUT /cloud/project/$project_id '{ "description" : "change-project-name-here"}'
`

// vars defined at compile time
var (
	// defined by https://github.com/ahmetb/govvv
	BuildDate string
	GitCommit string

	// contents of ./VERSION file
	Version string

	// defined by Makefile
	GoBuildVersion string
	ByUser         string

	// will be built by main()
	Ovh_cli_Version string
)

var Debug bool = false

// ======================================================================
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

func Debug_print(f string, args ...interface{}) {
	if Debug {
		log.Printf(f, args...)
	}
}

// ====================================================================== main

func main() {
	// prepare our command line parser
	parser := &docopt.Parser{
		OptionsFirst: true,
	}

	// build the Version string
	// Global are filed by
	Ovh_cli_Version = fmt.Sprintf("ovh-cli %s commit %s built at %s by %s\nbuilt from: %s\n%s",
		Version,
		GitCommit,
		BuildDate,
		ByUser,
		GoBuildVersion,
		strings.TrimSpace(copyleft))

	arguments, err := parser.ParseArgs(Usage, nil, Ovh_cli_Version)
	if err != nil {
		panic(err)
	}

	// ARGUMENTS
	Debug = arguments["--debug"].(bool)
	method := arguments["METHOD"].(string)
	path := arguments["URL_API"].(string)

	var bytes []byte
	json_input, err := arguments.String("JSON_INPUT")
	if err == nil {
		if len(json_input) > 0 {
			Debug_print("JSON_INPUT copied to []bytes: '%s'\n", json_input)
			bytes = []byte(json_input)
		}
	} else {
		Debug_print("JSON_INPUT can't be fetched: %v\n", err)
	}

	//======================================================================

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

	if fi.Mode()&os.ModeNamedPipe != 0 {
		// there is something to read
		bytes, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	}

	Debug_print("stdin bytes '%v' %d '%s'\n", bytes, len(bytes), string(bytes))
	Debug_print("Call: %s %s stdin: %v\n", method, path, bytes != nil)
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
	Debug_print("%v\n", req)
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if Debug {
		var s string
		// extrac response Header
		for k, v := range response.Header {
			s += fmt.Sprintf("%s : %s\n", k, v)
		}
		Debug_print("response: %d \n %s\n", response.StatusCode, s)
	}

	// Read all the response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	} else {
		// output response
		fmt.Println(string(body))
		if response.StatusCode != 200 {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
}
