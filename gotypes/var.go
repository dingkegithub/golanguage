package main

import "fmt"

func main() {
	// var name type
	var name string
	var age int
	fmt.Println("name: [", name, "]")
	fmt.Println("age:  [", age, "]")

	// var var-1, var-2 type 
	var p, c string
	fmt.Println("province: [", p, "]")
	fmt.Println("city:     [", c, "]")

	// var name type = initvalue
	var nick string = "男"
	var class int = 5
	fmt.Println("nick:  [", nick, "]")
	fmt.Println("class: [", class, "]")

	// var var-1, var-2 type = val-1, val-2
	var province, city string = "青海省", "西宁市"
	fmt.Println("province: [", province, "]")
	fmt.Println("city:     [", city, "]")

	// var name = initvalue
	var country = "中国"
	var postCode = 11011
	fmt.Println("country:  [", country, "]")
	fmt.Println("postCode: [", postCode, "]")

	// var var-1, var-2 = val-1, val-2
	var a, b = "中国", 1
	fmt.Println("a:  [", a, "]")
	fmt.Println("b: [", b, "]")

	var (
		follor = 1
		zoo = "朝阳"
	)
	fmt.Println("follor: [", follor, "]")
	fmt.Println("zoo:    [", zoo, "]")

}
