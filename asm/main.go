//汇编测试
package main

func init() {
	println("main.init1")
}

func main() {
	i := 10
	p := &i
	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}
	//fmt.Println(i, arr, slice)
	t(i, p, arr, slice)
}
func t(i int, p *int, arr [3]int, slice []int) {

}

func init() {
	println("main.init2")
}
