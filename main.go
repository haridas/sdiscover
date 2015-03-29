package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/haridas/sdiscover/store"
	"github.com/haridas/sdiscover/util"
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
		file, err := ioutil.ReadFile(*configOpt)

		util.ExitOnError("Error while opening the file", err)

		var conf SConf
		err = json.Unmarshal(file, &conf)
		util.ExitOnError("Error while parsing the config file", err)

		fmt.Println("Configuration: ", conf, "\n")

		// Construct the data storage and initlize it.
		s3store := &store.S3Store{
			AccessKey:     conf.AwsConf.AccessKey,
			SecretKey:     conf.AwsConf.SecretKey,
			BucketName:    conf.AwsConf.S3BucketName,
			LockKeyName:   conf.LockFileName,
			MasterKeyName: conf.MasterFileName,
		}
		s3store.Init()

		// Create an Application Object with all the informations.
		sdiscover := &SDiscover{Store: s3store, Conf: conf}
		log.Fatal(sdiscover.Run())

	} else {
		fmt.Println("./sdiscover -config=config.json")
		os.Exit(1)
	}
}
