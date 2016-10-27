# Example

```go
package main

import (
	"fmt"
	"github.com/philipgatzka/swclient"
	"github.com/philipgatzka/swclient/article"
)

func main() {
	// Get new swclient
	s, err := swclient.New("user", "key", "https://shop.ware/api")
	if err != nil {
		// handle
	}
	// The type of this struct determines the api resource which will be called
	artcl := article.Article{}
	// "2" is the id of the article we want to get, all returned data will be unmarshaled into artcl
	err = s.GetSingle("2", &artcl)
	if err != nil {
		// handle
	}
	// All fields of the "articles" resource are now set and accessible (except those that were null, 0, [], {} or "")
	fmt.Println(artcl.Name, artcl.DescriptionLong, artcl.MainDetail.Attribute.Attr3)
}
```
