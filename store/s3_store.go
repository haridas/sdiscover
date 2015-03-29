package store

import (
	"fmt"
	"github.com/haridas/sdiscover/util"
	"gopkg.in/amz.v3/aws"
	"gopkg.in/amz.v3/s3"
	"os"
)

const (
	s3LockPath = "/lock"
)

// S3 struct. With Store specific configurations.
type S3Store struct {
	Region        string
	AccessKey     string
	SecretKey     string
	BucketName    string
	LockKeyName   string
	MasterKeyName string

	// S3 Store specific constants.
	Conn   *s3.S3
	Bucket *s3.Bucket
}

/*
* Initilize the Store.
 */
func (store *S3Store) Init() {
	AWSAuth := aws.Auth{
		AccessKey: store.AccessKey,
		SecretKey: store.SecretKey,
	}

	region := aws.USEast

	// Just to remember the type at this point of time.
	var err error
	store.Conn = s3.New(AWSAuth, region)
	store.Bucket, err = store.Conn.Bucket(store.BucketName)

	util.ExitOnError("Error while connecting to S3", err)

	// Create the bucket for this operation.
	err = store.Bucket.PutBucket(s3.Private)
	util.ExitOnError("Error while creating the Bucket: ", err)
}

/*
*
* Acquire a lock before doing any modifications to the store.
* We create one empty object under the path s3LockPath to notify the caller that
* we have exclusive access to the bucket.
*
* If it failed to create lock, it returns the error back to the application.
 */
func (store *S3Store) AcquireLock() error {
	return store.Bucket.Put(s3LockPath, []byte(""), "application/octet-stream",
		s3.Private)
}

/* Remove the lock object created in the store or return error if it failes.
 */
func (store *S3Store) ReleaseLock() error {
	return store.Bucket.Del(s3LockPath)
}

/*
* Add any Object into the given path of the s3 Bcuket
* returns error if it fails to do so.
 */
func (store *S3Store) AddObject(path string, obj []byte) error {
	return store.Bucket.Put(path, obj, "application/octet-stream", s3.Private)
}

/*
* Add the given file into the s3 bucket on the given path.
*
* returns error if it fails to do so.
 */
func (store *S3Store) AddFile(path string, file *os.File, contentType string) error {
	fileInof, _ := file.Stat()
	return store.Bucket.PutReader(path, file, fileInof.Size(), contentType,
		s3.Private)
}

/*
* Read the given object stored on the S3 Bucket.
 */
func (store *S3Store) ReadAsObject(path string) ([]byte, error) {
	return store.Bucket.Get(path)
}

/*
* Read the content into the File object. The caller has to manage the file
* object. returns error if reading failes.
 */
func (store *S3Store) ReadAsFile(path string, file *os.File) error {
	data, err := store.ReadAsObject(path)
	if err != nil {
		return err
	}

	_, err = file.Write(data)

	if err != nil {
		return err
	}
	return nil
}
