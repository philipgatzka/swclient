# Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/philipgatzka/swclient"
	"github.com/philipgatzka/swclient/article"
)

// error-handling omitted for better readability
func main() {
    // create a new api-client providing user, key, api-url and api-resource
	api, _ := swclient.New("user", "key", "https://shop.ware/api", "articles")
    // execute request
	response, _ := api.GetById(4)
    // unmarshal received data into article model
	artcl := article.Article{}
	_ = json.Unmarshal(response, &artcl)

	fmt.Println("inStock:", artcl.Data.MainDetail.InStock)
}
```