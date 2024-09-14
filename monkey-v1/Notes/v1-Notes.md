## Basic notes for Monkey v1
- The interpreter is in between a simple "toy compiler" and a much more complicated JIT Compiler
- This type of interpreter is called an AST interpreter, ie, AST = Abstract Syntax Tree
- Created a tree of tokens and walks though them, assigning whatever meaning, finally providing the output 
  - ASTs have the following major parts:
    - Lexer
    - Parser
    - AST itself, ie, the tree
    - Interal object system
    - Evaluator


### Lexers
- Changing the source code to a simpler form. 
- Instead of native machine language or bytecode, we will be changing it this way:  
  
        [[Source Code]] -> [[Tokens]] -> [[AST]]
- The buffer being used here doesn't require extra memory. Instead it uses pointers to scrub through the input
- Input is considered as a string for the sake of lowering complexity
    - **Exercise: Use *io.Reader* for file-based inputs with line numbers for better error logging**
- Our lexer's `readChar()` only supports ASCII. **Additional exercise: support for Unicode and/or UTF-8**


### General Golang notes
- Single chars must be enclosed in *single quotes(')*, double quotes will only work for strings
- `:=` is called the short assignment operator. It both declares and assigns a variable, and implicitly decides on the datatype as well
  - An example: `var i int =  5` can be shortened as `i := 5`, unless the declaration and assignment needs to be separate as shown  
  in the switch case of the `nextToken()` in `lexer.go`
- Function declarations and calls are surprisingly intuitive. For example
```
  func A(input string) int {
    // This function takes an input string and returns an integer value, ie, var output int = A("hello")
  } 

  func (input string) A() int {
    // This function takes no input but is instead called using a dot operator after a string and returns an int,  
    // ie, var output int = "hello".A()
  }
```
- Test methods but always start with a capital letter T and follow the format `func TestXxx(t *testing.T) {...}`. This must be done in a file called `something_test.go` where the method under test is `something.go`. Both must be in the same package.