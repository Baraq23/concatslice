# concatslice
- this program takes cocatenates 2 or more slices int o one slice.

## Usage

```bash
go run .
```

_Output_

```bash
Slice1: [1 2 3 4]  Slice2: [5 6 7 8]
Concatenated:  [1 2 3 4 5 6 7 8]
```

## Notion
- Use of the '...' operator (the variadic operator).

#### Uses of Variadic Operator

* _Variadic Functions_: Define a function that takes a variable number of arguments using '...' in the parameter list.

_Example_

```bash
func printAll(strings ...string) {
    for _, s := range strings {
        fmt.Println(s)
    }


```

* _Passing Slices_: Expand a slice into individual arguments by appending ... to the slice when calling a variadic function.

```bash 
func main() {
    words := []string{"hello", "world", "from", "Go"}
    
    // Pass the slice as individual arguments

    printAll(words...)
}
```

* _Appending Slices_: Concatenate slices using append and the ... operator to expand the second slice.

_Example_

```bash
func main() {
    slice1 := []int{1, 2, 3}
    slice2 := []int{4, 5, 6}

    // Concatenate slice2 to slice1
    concatenated := append(slice1, slice2...)

    fmt.Println(concatenated)
}
```

```
Output: [1 2 3 4 5 6]
```


## Author
* [Barrack Kope Otieno](https://www.linkedin.com/in/barrack-kope-otieno-064a43244)