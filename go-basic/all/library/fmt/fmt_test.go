package fmt

import (
	"fmt"
	"os"
	"testing"
)

/*
在 Go 语言中，fmt包使用类似于 C 的 printf() 和 scanf() 函数的函数来实现格式化 I/O。Go 语言格式中的fmt.Fprintf()函数根据格式说明符写入 w。
而且，这个函数是在 fmt 包下定义的。在这里，您需要导入“fmt”包才能使用这些功能。
*/
func TestFibList(t *testing.T) {
	// Declaring some const variables
	const name, dept = "GeeksforGeeks", "CS"

	// Calling Fprintf() function which returns
	// "n" as the number of bytes written and
	// "err" as any error ancountered
	n, err := fmt.Fprintf(os.Stdout, "%s is a %s portal.\n",
		name, dept)

	// Printing the number of bytes written
	fmt.Print(n, " bytes written.\n")

	// Printing if any error encountered
	fmt.Print(err)
}
