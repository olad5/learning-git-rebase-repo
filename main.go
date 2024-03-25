package main

func main() {
	println("hello world")
	result := getResult(2)
	println("result is: ", result)
}

func getResult(num int) int {
	return num + num
}

func someFunc() {
	print("Hi there")
}
