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
- Run the lexer by following these steps
  - `cd` into the `monkey-v1` directory
  - Run `go run ./main.go`

### Lexers
- Changing the source code to a simpler form. 
- Instead of native machine language or bytecode, we will be changing it this way:  
  
        [[Source Code]] -> [[Tokens]] -> [[AST]]
- The buffer being used here doesn't require extra memory. Instead it uses pointers to scrub through the input
- Input is considered as a string for the sake of lowering complexity
    - **Exercise: Use *io.Reader* for file-based inputs with line numbers for better error logging**
- Our lexer's `readChar()` only supports ASCII. **Additional exercise: support for Unicode and/or UTF-8**
- `readWhitespace(), eatWhitespace()` etc are methods used a lot in parsers or lexers to get around whatever character needs to be ignored,
such as whitespace, tab, newline and so on.

### REPL or Read Eval Print Loop
- This is basically the console that is present in Javascript or Python. 
- The REPL Reads the input, Evaluates it, Prints the result/output and does it all over again, aka, Loop.