package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

var sess = connectAWS()

// use viper package to read .env file
// return the value of the key
func viperEnvVariable(key string) string {
	//Set config file as env
	viper.SetConfigFile(".env")

	//Read config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	//Get value
	value, ok := viper.Get(key).(string)

	if !ok {

		log.Fatalf("Invalid type assertion")
	}

	return value
}

/**
Initialize constant variables
NOTE: Cannot use const variable because const is only allowed for numeric types, strings and bools
*/
var (
	awsS3Region  = viperEnvVariable("AWS_S3_REGION")
	awsS3bucket  = viperEnvVariable("AWS_S3_BUCKET")
	awsAccessKey = viperEnvVariable("AWS_ACCESS_KEY")
	awsSecretKey = viperEnvVariable("AWS_SECRET_KEY")
	port         = viperEnvVariable("port")
)

/**
Initialize AWS session
*/
func connectAWS() *session.Session {

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, "")
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(awsS3Region),
			Credentials: creds,
		})

	if err != nil {
		panic(err)
	}
	return sess
}

/**
Upload Hanlder
*/
func uploadImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseMultipartForm(32 << 20)

	file, header, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	filename := header.Filename

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(awsS3bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Successfully uploaded, access file via this path %q\n", aws.StringValue(&result.Location))
	return
}

//Main go function
func main() {
	router := httprouter.New()

	router.POST("/uploadImage", uploadImage)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
