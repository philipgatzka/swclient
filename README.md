*This is a work in progress*

## Example

```go
package main

import (
	"fmt"
	"github.com/philipgatzka/swclient"
	"github.com/philipgatzka/swclient/article"
)

func main() {
	// New swclient
	s, err := swclient.New("user", "key", "https://shop.ware/api")
	if err != nil {
		// handle
	}
	// The type of this struct determines the api endpoint we'll be requesting
	artcl := article.Article{}
	// "2" is the id of the article we want to get, all returned data will be unmarshaled into artcl
	err = s.Get("2", &artcl)
	if err != nil {
		// handle
	}
    
	fmt.Println(artcl.Name, artcl.DescriptionLong, artcl.MainDetail.Attribute.Attr3)
}
```

## TODO

 - More and better tests
 - Fix Digestclient and put it in a separate package (or import a better one from another project)
 - Implement more API resources (structs)