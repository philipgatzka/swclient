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
	    s, _ := swclient.New("user", "key", "https://shop.ware/api", "articles")
        get(s)
}

func get(s swclient.Swclient) {
        // execute request
    	response, _ := s.GetById(4)
        
        // unmarshal received data into article model
    	artcl := article.Article{}
    	_ = json.Unmarshal(response, &artcl)
    
    	fmt.Println("inStock:", artcl.Data.MainDetail.InStock)
}

func put(s swclient.Swclient) {
        // define changes
        changeset := article.Data{
                Name: "Newname",
                MainDetail: article.Detail{
                        InStock: 99,
                        Attribute: &article.Attribute{
                                Attr4: "Update",
                        },
                },
        }
        // marshal into json
        b, _ := json.Marshal(changeset)
        
        // execute request
    	response, _ := s.PutById(4, b)
        
        // maybe inspect the returned data
    	artcl := article.Article{}
    	_ = json.Unmarshal(response, &artcl)
    	
        if !artcl.Success {
                // do something
        }
}
```