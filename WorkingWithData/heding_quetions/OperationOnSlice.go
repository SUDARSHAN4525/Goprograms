package main
import "fmt"
func findIndex(slice []int,target int)int{
	for i,v:= range slice{
		if v == target{
			return i
		}
	}
	return -1
}
func main(){
	slice :=[]int{10,20,30,40,50}
	slice=append(slice, 60)
	fmt.Println("after appending 60:",slice)

	slice[2]=100
	fmt.Println("After modify index 2 to 100",slice)

	subSlice:=slice [1:4]
	fmt.Println("sub-slice from index 1 to 3",subSlice)

	slice = append(slice[:3],slice[4:]...)
	fmt.Println("after deleting element at index 3:",slice)

	index := findIndex(slice,100)
	if index != -1{
		fmt.Println("Index of 100",index)
	}else{
		fmt.Println("element not found")
	}
}