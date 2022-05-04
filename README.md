# Scarlex API Wrapper

- [Discord] : https://discord.gg/ZKzmmt4gvq
- [Site] : https://scarlex.org/api

# How To Use

- join our Discord for any support (link above).
- register your own profile on our website (link above).
- then login and you're ready to fetch endpoints and check your stats.

# Examples


```go
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
```

```json
// will return json response
{
  "code": 200,
  "message": "•••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••───────────────────────────────",
  "percentage": "69/100"
}
```
## Developers / Contributors
- **[A V I X I T Y#0001]( https://github.com/avixityyt )**
