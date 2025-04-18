package field_test

import (
	"fmt"

	"github.com/wubin1989/gen/field"
)

func ExampleFunc() {
	expr := field.Func.UnixTimestamp()

	sql, vars := field.BuildToString(expr)
	fmt.Println(sql, vars)

	sql, vars = field.BuildToString(expr.Mul(100))
	fmt.Println(sql, vars)

	// Output:
	// UNIX_TIMESTAMP() []
	// (UNIX_TIMESTAMP())*? [100]
}
