package main

import "fmt"

func main() {
	/*
	 *
	 *  Integers
	 *
	 */

	//int (System default integer type)
	var a int = 10
	var b int = 20

	//int8, int16, int32, int64
	var c int8 = 10  //takes 8bit memory
	var d int16 = 20 //takes 16bit memory
	var e int32 = 30 //takes 32bit memory
	var f int64 = 40 //takes 64bit memory

	/*
	 *
	 *  Floats
	 *
	 */
	var g float32 = 10.012345678910 //Output will be 10.012345 - Takes 32bit
	var h float64 = 20.012345678910 //Output will be 20.01234567891 - Takes 64bit

	/*
	 *
	 *  Booleans
	 *
	 */
	var i bool = true
	var j bool = false

	/*
	 *
	 *  Strings
	 *
	 */
	var k string = "String value"

	/*
	 *
	 *  Constants
	 *
	 */
	const l int = 10 //can't be modified once declared

	/*
	 *
	 *  Array
	 *
	 */

	var arr [2]int = [2]int{1, 2} //Fixed size

	/*
	 *
	 *  Slice
	 *
	 */

	var slc []int = []int{1, 2, 3, 4, 5} //Dynamic size

	/*
	 *
	 *  Maps
	 *
	 */

	var m map[string]int = map[string]int{"a": 1, "b": 2}

	/*
	 *
	 *  Structs: Way to create dynamic style maps (Similar to JS Objects)
	 *
	 */

	type Person struct {
		Name string
		Age  int
	}

	var p Person = Person{Name: "Amit", Age: 20}

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, arr, slc, m["a"], p.Name)

	/*
	 *
	 *  Pointers: A way to store memory address of a variable
	 *
	 */
	var salary int = 10_000      //Memory allocated
	var salaryPtr *int = &salary //Memory address
	/*
		Dereferencing a pointer means retrieving the value of the variable from Memory location it points to
		and assign to a new variable which will be stored in the new memory location
	*/
	var derefSalary int = *salaryPtr

	fmt.Println(salaryPtr, &derefSalary) //0xe3fab5f40f0 0xe3fab5f40f8 - Memory addresses are different

}
