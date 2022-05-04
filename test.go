package main


import "scarlex/scarlex"
import "fmt"
import "encoding/json"

func main() {
           Data := scarlex.MakeRequest("https://scarlex.org/api/v1/progressbar?max=100&value=69&size=100&style=3", "A V I X I T Y", "daddyscar")
            var result map[string]interface{}
            json.Unmarshal([]byte(Data), &result)
            fmt.Println(result["message"] , result["percentage"])
}