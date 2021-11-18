# innValidator
-  simple module for validation inn control number

#### Example of usage: 
```go
package main

import (
	"fmt"
	"github.com/SkurkovPavel/innValidator"
)

func example()  {

	inn := "1111111111"

	ok := innValidator.Validate(inn)
	
	if !ok {
		fmt.Printf("Inn is invalid:%v", inn )
		
	}
	
}
```
