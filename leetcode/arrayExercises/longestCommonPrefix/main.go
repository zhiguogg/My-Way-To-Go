package main

import "strings"

// 序号14
func main()  {

}

// 首先这个前缀是公共的，我们可以从任意一个元素中找到它。 依次将基准元素和后面的元素进行比较 不断更
// 新基准元素，直到基准元素和所有元素都满足最长公共前缀的条件
func longestCommonPrefix(s []string ) string {
	if len(s) ==0{
		return ""
	}
	prefix := s[0]
	for i:=1;i<len(s);i++ {
		prefix = commonPrefix(prefix,s[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func commonPrefix(s1,s2 string) string {
	prefix := s1
	if strings.Index(s2,prefix) != 0 {
		if len(prefix) == 0{
			return ""
		}
		prefix = prefix[0:len(prefix)-1]
	}
	return prefix
}
