package main

func main() {

	arr := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		println(&arr[i])
	}
	println("----")
	slice := arr[1:3]

	for i := 0; i < len(slice); i++ {
		println(&slice[i])
	}

	slice = append(slice, 6)

	for i := 0; i < len(slice); i++ {
		println(&slice[i])
	}

	println("----")
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}
