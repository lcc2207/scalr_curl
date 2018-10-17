package main

import "fmt"
import "io/ioutil"
import "log"
import "net/http"
import "os"
import "time"
import "crypto/hmac"
import "crypto/sha256"
import "encoding/base64"
import "strings"

const scalrSignatureVersion = "V1-HMAC-SHA256"

// process the request
func processrequest(scalr_url string, APIKeyID string, secret_key string, queryurl string, raction string) []byte{
  // time timestamp
  timestamp := time.Now().Format(time.RFC3339)
  // generate string
  stringToSign := strings.Join([]string{raction, timestamp, queryurl, "", ""}, "\n")
  // get sig
  key := []byte(secret_key)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(stringToSign))
  signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
  xsig := scalrSignatureVersion+" "+signature

  // build url
  url := strings.Join([]string{scalr_url, queryurl}, "")

  // create HTTP connection
  client := &http.Client{}
  req, err := http.NewRequest(raction, url, nil)
  // set Headers
  req.Header.Set("X-Scalr-Key-Id", APIKeyID )
  req.Header.Set("X-Scalr-Date", timestamp)
  req.Header.Set("X-Scalr-Signature", xsig)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-Scalr-Debug", "1")

  // execute request
  res, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
  return responseData
}

func main() {
  scalr_url := os.Getenv("SCALR_SERVER_URL")
  api_key := os.Getenv("SCALR_API_KEY_ID")
  secret_key := os.Getenv("SCALR_SECRET")
  queryurl := os.Getenv("QUERY_PATH")
  action := os.Getenv("METHOD")
  dowork := processrequest(scalr_url, api_key, secret_key, queryurl, action)
  fmt.Println(string(dowork))
}
