# Go Data Types

Go (Golang) is a statically typed language, meaning every variable has a defined type at compile time.

---

## Basic Data Types

### Numeric Types

#### Integer Types

| Type   | Size                  | Description                     |
| ------ | --------------------- | ------------------------------- |
| int    | platform (32/64 bits) | **Default integar**             |
| int8   | 8-bit                 | -128 to 127                     |
| int16  | 16-bit                | -32,768 to 32,767               |
| int32  | 32-bit                | -2³¹ to 2³¹-1                   |
| int64  | 64-bit                | Large integers                  |
| uint   | platform (32/64 bits) | **Unsigned integer**            |
| uint8  | 8-bit                 | 0 to 255 (byte alias)           |
| uint16 | 16-bit                | 0 to 65,535                     |
| uint32 | 32-bit                | 0 to 4,294,967,295              |
| uint64 | 64-bit                | 0 to 18,446,744,073,709,551,615 |

#### Floating Point Types

| Type    | Description      |
| ------- | ---------------- |
| float32 | Single precision |
| float64 | Double precision |

# Single vs Double Precision (Go)

Floating-point precision defines how many bits are used to store a number, which affects both **accuracy** and **range**.

---

## Single Precision (`float32`)

```go
var x float32 = 3.1415927
```

- Uses **32 bits (4 bytes)**
- Precision: ~**6–7 decimal digits**
- Structure:
  - 1 bit → sign
  - 8 bits → exponent
  - 23 bits → fraction (mantissa)

### Example

```go
var a float32 = 1.123456789
fmt.Println(a) // Output: 1.1234568 (rounded)
```

---

## Double Precision (`float64`)

```go
var x float64 = 3.141592653589793
```

- Uses **64 bits (8 bytes)**
- Precision: ~**15–16 decimal digits**
- Structure:
  - 1 bit → sign
  - 11 bits → exponent
  - 52 bits → fraction (mantissa)

### Example

```go
var b float64 = 1.123456789
fmt.Println(b) // Output: 1.123456789 (more accurate)
```

---

## Key Differences

| Feature       | float32 (Single)  | float64 (Double) |
| ------------- | ----------------- | ---------------- |
| Size          | 32 bits / 4 bytes | 64bits / 8 bytes |
| Precision     | ~6–7 digits       | ~15–16 digits    |
| Memory usage  | Lower             | Higher           |
| Accuracy      | Less              | More             |
| Default in Go | ❌                | ✅               |

---

## When to Use

### Use `float32` when:

- Memory usage matters (e.g., large datasets, graphics)
- Small precision loss is acceptable

### Use `float64` when:

- Accuracy is important (e.g., calculations, backend logic)
- General-purpose programming (recommended)

---

## Important Note

Floating-point numbers are not always exact:

```go
fmt.Println(0.1 + 0.2) // 0.30000000000000004
```

This happens due to binary representation limitations.

---

## Summary

- `float32` → smaller, faster, less precise
- `float64` → larger, slower, more precise
- Go uses `float64` by default for better accuracy

---

#### Complex Types

| Type       | Description              |
| ---------- | ------------------------ |
| complex64  | float32 real & imaginary |
| complex128 | float64 real & imaginary |

---

### Boolean

```go
var isActive bool = true
```

---

### String

```go
var name string = "Rohit"
```

- Immutable
- UTF-8 encoded

---

## Derived Types

### Array

```go
var arr [3]int = [3]int{1, 2, 3}
```

---

### Slice

```go
numbers := []int{1, 2, 3}
numbers = append(numbers, 4)
```

---

### Map

```go
m := map[string]int{
  "age": 25,
}
```

---

### Struct

```go
type User struct {
  Name string
  Age  int
}
```

---

### Pointer

```go
var x int = 10
var p *int = &x
```

---

### Function Type

```go
func add(a int, b int) int {
  return a + b
}
```

---

### Interface

```go
type Shape interface {
  Area() float64
}
```

---

## Special Types

### byte

```go
var b byte = 255
```

Alias for `uint8`

---

### rune

```go
var r rune = 'A'
```

Alias for `int32` (Unicode)

---

### nil

Represents zero value for:

- pointers
- maps
- slices
- interfaces
- channels
- functions

---

## Type Inference

```go
x := 10       // int
y := "hello"  // string
```

---

## Type Conversion

```go
var x int = 10
var y float64 = float64(x)
```

---

## Zero Values

| Type    | Zero Value |
| ------- | ---------- |
| int     | 0          |
| float   | 0.0        |
| bool    | false      |
| string  | ""         |
| pointer | nil        |

---

## Summary

- Go is statically and strongly typed
- Supports primitive and composite types
- Requires explicit type conversion
- Provides powerful built-in structures like slices and maps

---
