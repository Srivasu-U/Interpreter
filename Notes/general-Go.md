## General Golang notes
- Valuable links: [Go By Example](https://gobyexample.com/), [Tour of Go](https://go.dev/tour/list) and [Effective Go](https://go.dev/doc/effective_go)
- Single chars must be enclosed in *single quotes(')*, double quotes will only work for strings
- `:=` is called the short assignment operator. It both declares and assigns a variable, and implicitly decides on the datatype as well
  - An example: `var i int =  5` can be shortened as `i := 5`, unless the declaration and assignment needs to be separate as shown  
  in the switch case of the `nextToken()` in `lexer.go`
- Function declarations and calls are surprisingly intuitive. For example
```
  func A(input string) int {
    // This function takes an input string and returns an integer value, ie, var output int = A("hello"). 
    // This is called a function
  } 

  func (input string) A() int {
    // This function takes no input but is instead called using a dot operator after a string and returns an int,  
    // ie, var output int = "hello".A()
    // Conventionally, this style is called a method
  }
```
- Test methods but always start with a capital letter T and follow the format `func TestXxx(t *testing.T) {...}`.  
This must be done in a file called `something_test.go` where the method under test is `something.go`. Both must be in the same package.
- Imports from other files must always start from the actual directory of the `go.mod`. For example, `import "Learning-Go/monkey-v1/token"` is the correct way,  
not `import "monkey=v1/token"` unless the *monkey-v1* directory has a `go.mod`. 
- **Pointers** are really useful when utilizing variable of large sizes. Often a good practice to always just pass references for structs,  for example, instead of an entire struct, a reference to the struct can be passed instead
- Note: [Useful article](https://medium.com/@mathieu.durand/how-to-use-golang-interface-vs-java-1fc8b281c101) to the differences in Interfaces between Java and Golang
    - Another [article](https://gobyexample.com/interfaces) showing a clear example of the usage of interfaces at a basic level
- Interfaces are the one thing in Golang from OOP. It does not have classes. Interfaces can be declared using 
```
type X interface {
    methodX() return_type
}
```
- Inheritance or implementation of interfaces can also work like this 
```
type Y interface {
    X       // Here, Y implements methodX() from X as well. Y **must** provide its own implementation of methodX()
    methodY() return_type
}
```
- Interfaces can somehow be struct elements?, ie, structs implement interfaces. I am not sure how
- Interfaces can be parameters to functions
- In Golang, only methods or values starting with a capital letter can be exported. For example, in the `lexer.go` file
`New()` and `NewToken()` are exportable methods while `isDigit()` and others are not.
- Functions can also have multiple return values. Probably one of the best things about golang
```
func A(x, y int) (string, string) {...}
```
- Consts do not use datatypes or the short hand operator `:=`
- Type convertion works like this: `datatype(value)`. Example: `float64(2)` converts int 2 into float64
- We also have type checking/assertion like this: `ident.(type)`. Example: `value.(string)` checks if value contains a string expression
- Interesting error message: `impossible type assertion: program.Statements[0].(*ast.Statement)
	*ast.Statement does not implement ast.Statement (type *ast.Statement is pointer to interface, not interface)`. What does this mean?
- String formatting reference: [click here](https://gobyexample.com/string-formatting)