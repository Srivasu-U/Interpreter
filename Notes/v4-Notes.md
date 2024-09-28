## Extending the interpreter

### Adding a new datatype
- Adding string support
- General series of steps
    - Add a new type of token
    - Modify the lexer
    - The token through the parser
    - Evaluate the parser
- The entire string is a single token, instead of each word/character. This makes it easier to move around the token.

### Built-in functions
- We create a new `Builtin` object as a wrapper, to wrap the functions and keep the representation in-line with our object system
- Built-in functions are kept in a separate environment
- `evalIdentifier()` must have a lookup as a fallback to check when a given identifier (function names are identifiers) does not have a binding in the current environment

### Arrays
- No restriction on the type of values to be held in an array. A single array can hold an int, a bool, a string and a function.
- The index operator has to have the highest precedence of all the operators since the actual value has to be retrieved.
- Just like function call expressions, we treat arrays/indexing of arrays like infix expressions instead of prefix
    - In function call, `(` is considered the operator, with function name being the left operand and params being the right operand
    - In indexing of array, `[` is considered the operator instead
- With arrays, the challenge is to be able to evaluate the actual array items as well as the array index.
    - For example, `[1, 2*2]` is about evaluating the second array item/array literal
    - `[1, 2, 3, 4][4/4]` is about evaluating the index to get value `1` and retrive the array element at index `1`
- When trying to retrieve a value with index out of bounds, Monkey returns a NULL instead of throwing an error.
- Built-in functions that operate on arrays
    - `len` - length (also works with strings)
    - `first` - returns first element of an array
    - `last` - return last element of array
    - `rest` - returns a new array containing all the elements of the array *except the first one*
        - The array is *newly allocated*, i.e., the original array is not modified.
    - `push` - returns a new array after appending a new element to the end of the original array
        - Once again, a *newly allocated* array is returned, and the original is unmodified
- `push` and `rest` indicate that arrays in Monkey are *immutable*.