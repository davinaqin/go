package main
//字符 和 字符串处理
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱golang！"
	fmt.Println(len(s))

	//UTF-8编码，每个中文占3个字节，采用可变长的编码
	//%X使用16进制的的数字表示每个得到的byte
	for _,ch := range []byte(s)  {
		fmt.Printf("%X ",ch)
	}
	fmt.Println()

	//ch是int32，实际上是一个rune,将每个中文转为Unico编码
	//中文3个字节，英文是1个字节
	for i,ch := range s{	//ch is rune
		fmt.Printf("(%d %X)",i,ch)
	}

	fmt.Println()

	bytes := []byte(s)
	for i,ch := range bytes{
		fmt.Printf(("%d %X\n"),i,ch)
	}

	//统计UTF-8的字符个数
	fmt.Println(utf8.RuneCountInString(s))

	fmt.Println(len(s))
}


