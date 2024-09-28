## The Monkey Interpreted Language
- This repo consists of different incremental versions of the Monkey programming language, as implemented from *Writing an Interpreter in Go* by Thorsten Ball.
- Each subdirectory has its own version of the source code for the interpreter
    - `monkey-v1` mainly focuses on the lexer/tokenizer.
    - `monkey-v2` deals with writing a correct parser and AST for the language.
    - `monkey-v3` looks into the evaluation of the parsed AST.
    - `monkey-v4` additionally supports strings, arrays, hash-maps and certain built-in functions.
        - `len` for arrays and strings
        - `first`, `last`, `rest` and `push` for arrays (Read `Notes/v4-Notes.md`) to get more information about these methods
    - `Notes` consists of general notes regarding Golang, as well as notes related to what I understood as I went through each version of Monkey

### Execution of code
- Each subdirectory can be executed by cloning the repo and running
```
> cd monkey-<version-number>/src
> go run main.go
```
- The test cases can be executed by running
```
> go test ./lexer
> go test ./parser
> go test ./ast
> go test ./evaluator
```
