package main

import "fmt"

func main() {
	/*
	 *
	 *  Integers
	 *
	 */

	//int (int size depends on system architecture: usually 64-bit on modern 64-bit machines)
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
	var g float32 = 10.012345678910 //Stores approximate 32-bit floating-point value
	var h float64 = 20.012345678910 //Stores approximate 64-bit floating-point value

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
	 *  Array
	 *
	 */

	// Array declaration & initialization - Can't be declared with "const"
	var arr [2]int = [2]int{1, 2} //Fixed size

	//Short-hand array declaration & initialization
	arr2 := [2]int{3, 4}

	//Inferred size array: Let the compiler count the size of the array
	inferredArrSize := [...]int{1, 2, 3}

	// Specific Indices
	names := [3]string{2: "Amit", 0: "Rohit", 1: "Neha"} //OUTPUT: [Rohit Neha Amit]

	/*
	 *
	 *  *********************** Slice ***********************
	 *
	 */

	var slc []int = []int{1, 2, 3, 4, 5} //Dynamic size
	var makeArr = make([]int, 3, 5)      //type, size, capacity
	makeArr = append(makeArr, 4)
	fmt.Println(makeArr)

	// Why we use capacity, Underlying Array allocated: [0, 0, 0, 0, 0] (Size 5)
	// Your Slice Window only sees the first 3 elements.
	// Result: [0, 0, 0] (But it has 2 extra hidden slots reserved for future growth).
	// So appending more values to it will not create new array rather utilize the existing 1 unless the capacity exceeds

	/*
	 *  *********************** Map ***********************
	 *
	 *  A map has key-value pairs where keys are unique.
	 *	The key type and value type are defined when the map is created.
	 *	All keys must have the same type, and all values must have the same type.
	 *	Maps are implemented using hash tables, which provide efficient average-case lookups, insertions, and deletions.
	 *
	 */

	var m map[string]int = map[string]int{"a": 1, "b": 2}

	/*
	 *
	 *  *********************** Structs ***********************
	 *
	 *	Way to create dynamic style maps (Similar to JS Objects)
	 * 	But you can't add new fields to struct at runtime like you can do with JS objects, you have to define all the fields at compile time
	 *
	 */

	type Person struct {
		Name string
		Age  int
	}

	var p Person = Person{Name: "Amit", Age: 20}

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, arr, arr2, inferredArrSize, names, slc, m["a"], p.Name)

	/*
	 *
	 *  *********************** Pointers ***********************
	 *
	 *	A way to store memory address of a variable
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

	/*
	 *
	 *  Constants
	 *
	 */

	//Typed constants: We explicitly define the type of the constant, so it will always be of that type and cannot be used with other types without explicit conversion
	const TYPED_INT int = 10
	const TYPED_FLOAT float64 = 4.0
	// fmt.Println(TYPED_INT / TYPED_FLOAT) //OUTPUT: throws error:  TYPED_INT / TYPED_FLOAT (mismatched types int and float64)

	//Untyped constants: We do not explicitly type them, so the type will be inferred by the compiler based on the value assigned to it. This allows them to be used in a wider range of contexts without requiring explicit type conversion, as the compiler can automatically determine the appropriate type based on how the constant is used.
	const UNTYPED_CONST_INT = 10
	const UNTYPED_CONST_FLOAT = 4.0
	fmt.Println(UNTYPED_CONST_INT / UNTYPED_CONST_FLOAT) //OUTPUT: 2.5

	//Reassigning value to variable
	var test int = 10
	test = 20 //We can reassign value to variable

	fmt.Println(test) //OUTPUT: 20
}
