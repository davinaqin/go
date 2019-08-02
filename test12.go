package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i,ch := range []byte(s){
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start{
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	//寻找最长不含有重复字符的子串
	//abcabcbb -> abc
	//bbbb -> b
	//pwwkew -> wke
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))

}
