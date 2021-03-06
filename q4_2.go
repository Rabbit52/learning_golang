/*
2. 建立一个程序统计字符串里的字符数量:
asSASA ddd dsjkdsjs dk
同时输出这个字符串的字节数。 提示: 看看 unicode/utf8 包。
3. 扩展上一个问题的程序,替换位置4开始的三个字符为“abc”。
*/

package main

/*
import "fmt"
import "unicode/utf8"
*/
// 请使用这种方法来引入多个包
import (
    "fmt"
    "unicode/utf8"
)

func main() {
    str := "asSASA ddd dsjkdsjs dk"
    fmt.Printf("String %s\nLength: %d, Runes: %d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
    // 不明 length 和 runes 的区别，求破
    fmt.Printf("%d\n", utf8.RuneCountInString(str))
    // 这个方法太扯了~
    str_1 := str[0:4] + "abc" + str[7:]
    fmt.Printf("%s\n", str_1)
}