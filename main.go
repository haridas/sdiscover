package main

import (
	"encoding/json"
	"flag"
	"fmt"
	//	"github.com/haridas/sdiscover/util"
	"io/ioutil"
	"log"
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
		file, err := ioutil.ReadFile("conf.json")

		//err_msg := fmt.Printf("Error while opening the config file: %v", *configOpt)
		//util.ExitOnError("Error while opening the file", err)
		//fmt.Printf("read 100 bytes %d, data: %q", count, data[:100])
		//err_msg := fmt.Printf("Error while parsing the config file: %v", *configOpt)

		var x SConf
		json.Unmarshal(file, &x)
		//fmt.Printf("File data: %s\n", string(file))
		//util.ExitOnError("Error while parsing the config file", err)
		fmt.Println(x)

		if err != nil {
			fmt.Println("Error while parsing the file")
			log.Fatal(err)
		}

		//file, e := ioutil.ReadFile("test.json")
		//if e != nil {
		//	fmt.Printf("File error: %v\n", e)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s\n", string(file))

		////m := new(Dispatch)
		////var m interface{}
		////var jsontype jsonobject
		//var jsontype AwsConf
		//json.Unmarshal(file, &jsontype)
		////fmt.Printf("Results: %v\n", jsontype)
		//fmt.Println(jsontype)

	} else {
		fmt.Println("./sdiscover -config=config.json")
		os.Exit(1)
	}
}
