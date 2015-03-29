package main

// Defines all the configuration types. We read the JSON configuration file
// and populate these values.

type AwsConf struct {
	AccessKey    string
	SecretKey    string
	S3BucketName string
}

type SConf struct {
	AwsConf        AwsConf `json:"aws"`
	LockFileName   string  // Lock file to serialize client access to the `masterFile`.
	MasterFileName string  // Name of the file On this we keep the master IPs of consul as JSON list.
	ConsulBinFile  string  // Full path to the consul binary.
	ConsulMode     string  // Running mode as agent or server.
}
