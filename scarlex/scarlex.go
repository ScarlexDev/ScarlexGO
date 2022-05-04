package scarlex

import (
	base64 "encoding/base64"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
  }	

func MakeRequest(endpoint , username, password string)  string {
	
    c := http.Client{Timeout: time.Duration(2) * time.Second}
    req, err := http.NewRequest("GET", `https://scarlex.org/api/v1/` + endpoint, nil )
	req.Header.Add("Authorization", "Basic "+basicAuth(username, passowrd))
    if err != nil {
        fmt.Printf("error %s", err)
        return ""
    }
    req.Header.Add("Accept", `application/json`)
    resp, err := c.Do(req)
    if err != nil {
        fmt.Printf("error %s", err)
        return ""
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    //  fmt.Printf("Body : %s", body)
	return string(body)
}
