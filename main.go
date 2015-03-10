package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/haridas/sdiscover/util"
	"io/ioutil"
	"os"
)

// Command line arguments.
//
// $ sdiscover -config=<path>/config.json
//

var configOpt = flag.String("config", "",
	"Provide the Configuration files for the Sdiscover")

func main() {

	flag.Parse()

	if *configOpt != "" {
		file, err := ioutil.ReadFile(*configOpt)

		util.ExitOnError("Error while opening the file", err)

		var x SConf
		err = json.Unmarshal(file, &x)
		util.ExitOnError("Error while parsing the config file", err)

		fmt.Println(x)

	} else {
		fmt.Println("./sdiscover -config=config.json")
		os.Exit(1)
	}
}
