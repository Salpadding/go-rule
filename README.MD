# GO Rule engine support only number type

## example

```go
package main

import "fmt"

func main(){
	ctx := NewContext(map[string]interface{}{
		"c0": -9,
		"c1": 0,
		"c2": 0,
		"c3": 0,
		"c4": 0,
		"c5": 0,
		"c6": 0,
		"c7": 0,
	})
	ruler, err := ParseRule(`($or c0<=-9  
			($atLeast 2 c0>=80  c1>=8 c2>=8 c3>=8 c4>=8 c5>=8 c6>=8 c7>=8) 
			($atLeast 2 c0<=-8  c1<=-8 c2<=-8 c3<=-8 c4<=-8 c5<=-8 c6<=-8 c7<=-8)   			  
        )`)
    fmt.Println(ruler.Rule(ctx))    
}
```

