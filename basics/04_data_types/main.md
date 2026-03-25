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
