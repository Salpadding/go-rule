# GO Rule engine support only number type

## example

```go
package main

import(
	rule "github.com/Salpadding/go-rule"
	"fmt"
)


func main(){
	evaultor := rule.NewEvaluator()
	// let assign variable rs1rule1 a rule
	evaultor.Eval(`(let rs1rule1 'c0 <= -9')`);
	// (or rule1 rule2) return composite rule
	evaultor.Eval(`(let rs1rule1 
			(or rs1rule1 (atLeast 2 'c0>=8' 'c1>=8' 'c2>=8' 'c3>=8' 'c4>=8' 'c5>=8' 'c6>=8' 'c7>=8'))
		)`)
	// (atLeast 2 rule1 rule2 rule3) success when at least 2 of rule1 rule2 success	
	evaultor.Eval(`(let rs1rule1
			(or rs1rule1 
				(atLeast 2 'c0<=-8' 'c1<=-8' 'c2<=-8' 'c3<=-8' 'c4<=-8' 'c5<=-8' 'c6<=-8' 'c7<=-8')
			)
		)`)	
	evaultor.Eval(`(
		let rs1rule1 (and rs1rule1 'isSportOrDrunk=1' )
	)`)		
	rs1rule1, _ := evaultor.Eval("((rs1rule1))")	
	ctx := rule.NewContext(map[string]interface{}{
		"c0": 10,
		"c1": 0,
		"c2": 0,
		"c3": 0,
		"c4": 0,
		"c5": 0,
		"c6": 0,
		"c7": 0,
		"isSportOrDrunk": 1,
	})
	results, _ := evaultor.Expect(ctx)
	fmt.Printf("%v", results["rs1rule1"].Ok)
}
```

