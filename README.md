# Hello, World!

#### Install:
```
go get -u github.com/itv-go/go-requests
```

## How to Use

### Get Request

Example code:

```go
package main

import (
	"fmt"
	"github.com/itv-go/go-requests"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
    url := "https://api.example.com/users/1"
	
    var user User

    result, err := requests.Get(url, &user)
    if err != nil {
        fmt.Printf("Error fetching data: %s\n", err)
        return
    }

    fmt.Printf("Fetched User: %+v\n", result)
}
```

#### Example Output

Assume the API returns the following JSON:
```json
{
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com"
}
```

The program will output:
```
Fetched User: &{ID:1 Name:John Doe Email:john.doe@example.com}
```