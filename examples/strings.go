package main

import (
	"fmt"
	"strings"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Strings operation")
	stringOperation()

	timeOperation()
}

func stringOperation() {
	str := "I'm a confident boy"
	//前后缀 strings.HasPrefix | strings.HasSuffix
	fmt.Printf("does str has perfix is %s? %t\n", "Im",strings.HasPrefix(str,"Im"))
	fmt.Printf("does str has surfix is %s? %t\n", "boy",strings.HasSuffix(str,"boy"))

	//字符串包含 strings.Contains
	fmt.Printf("does str contain \"a\"? %t\n",strings.Contains(str,"a"))

	//子串或字符的位置 strings.Index | strings.LastIndex
	// 有多个只显示第一个
	fmt.Printf("\"n\" first show in str at index %d\n",strings.Index(str,"n"))
	// 有多个显示最后一个
	fmt.Printf("\"o\" last show in str at index %d\n",strings.LastIndex(str,"o"))
	// 查找非ASCII码出现的第一个位置
	fmt.Printf("'i' first show in str at index %d\n",strings.IndexRune(str,rune('i')))

	//字符串替换 strings.Replace 后面那个数字是代码str中需要替换的子串个数，比如说，"ab ab ab ab ab",n为1则只替换第一个ab串
	fmt.Printf("str replace \"confident\" with \"handsome\": %s \n",strings.Replace(str,"confident","handsome",3))

	//统计字符出现的个数 strings.Count
	fmt.Printf("\"o\" show times %d\n",strings.Count(str,"o"))

	//重复字符串 strings.Repeat
	fmt.Printf("str repeat 2 times is %s\n",strings.Repeat(str,2))

	//修改大小写 strings.ToLower,strings.ToUpper
	fmt.Printf("str to low is %s\n str to uper is %s\n",strings.ToLower(str),strings.ToUpper(str))

	//修剪字符串 strings.TrimSpace | strings.Trim | strings.TrimLeft | strings.TrimRight
	str1 := " 'abcd efa' "
	//修剪开始和结尾的空格
	fmt.Printf("str1 cut space on begin or end:%s\n",strings.TrimSpace(str1))
	//将'a'剔除
	fmt.Printf("str1 cut 'a':%s\n",strings.Trim(str1,"a"))
	//剔除开头的字符串
	fmt.Printf("str1 cut ' 'a' on begin:%s\n",strings.TrimLeft(str1," 'a"))
	//剔除结尾的字符串
	fmt.Printf("str1 cut 'a' ' on end:%s\n",strings.TrimRight(str1,"a' "))

	//分割 strings.Fields | strings.Split
	// 会按空格进行分割
	fmt.Printf("str field is %v\n",strings.Fields(str))
	// 会按"o"进行分割
	fmt.Printf("str split by 'o' is %v\n",strings.Split(str,"o"))

	//拼接 strings.Join 注意slice的概念
	//将slice按拼接符组成一个新字符串
	aslice := [] string {"1","1","2","3"}
	fmt.Printf("\"1,1,2,3\" join a new str by '`' is：%s\n",strings.Join(aslice,"`"))

	//从字符串中读取内容 strings.Repeat
	// 将字符串读取2遍，相当于Repeat2次
	fmt.Printf("str read char at 7: %s\n",strings.Repeat(str, 2))

	//类型转换 strconv包
	orgin := "666"
	var a int
	var newstr string
	a,_ = strconv.Atoi(orgin)
	fmt.Printf("the a = %d\n",a)
	a = a + 5
	newstr = strconv.Itoa(a)
	fmt.Printf("the newstr = %s\n",newstr)
}

func timeOperation(){
	//获取当前时间
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d %02d %4d\n",t.Day(),t.Minute(),t.Year())

	//通用国际时间
	t = t.UTC()
	fmt.Println(t)

	//当前时间➕1周的时间
	var week time.Duration
	week = 60 * 60 * 24 * 7 * 1e9 //一周的纳秒表示时长
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)

	//时间格式化输出
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	//根据一个时间的标准化描述来打印当前时间
	fmt.Println(t.Format("2006-Jan-02 15:04"))
}


