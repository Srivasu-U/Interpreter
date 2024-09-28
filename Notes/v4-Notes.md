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
