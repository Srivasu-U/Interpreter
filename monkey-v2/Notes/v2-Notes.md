## Notes for Monkey-v2
- Adding a parser to the previously created lexer in v1
- So, if the lexer tokenizes our input, a parser is used to turn the input into a data structure so it is easier to work with
- In our case, the data structure we will be using is the AST, or the Abstract Syntax Tree. Abstract because things like whitespace are abstracted
out and have no meaning in our programming language. Or how things such as `{}` and `()` guide in the tree structure instead.
- ASTs can differ on the meaning of symbols and how a language should perform. This is up to the designer
- Example of a parser in JavaScript:
``` 
> var input = 'if (3 * 5 > 10) { return "hello"; } else { return "goodbye"; }';
> var tokens = MagicLexer.parse(input);
> MagicParser.parse(tokens);
{
    type: "if-statement",
    condition: {
        type: "operator-expression",
        operator: ">",
        left: {
            type: "operator-expression",
            operator: "*",
            left: { type: "integer-literal", value: 3 },
            right: { type: "integer-literal", value: 5 }
        },
        right: { type: "integer-literal", value: 10 }
    },
    consequence: {
        type: "return-statement",
        returnValue: { type: "string-literal", value: "hello" }
    },
    alternative: {
        type: "return-statement",
    returnValue: { type: "string-literal", value: "goodbye" }
    }
}
```
- Our Monkey-v1 basically takes the tokens from the lexer as input to create an AST. 
- It also analyzes the input to assert the expected structure. Which is why parsing is also called syntactic analysis.
- As described below, while parser generators can be used to solve problems, and just plugged into this project, the goal is to understand how parsers work
which is why we are writing our own.

### Parser generators
- These are tools that, when given a language description, produces the parsers. Ex: ANTLR, yacc or bison.  
**Note: Read more about parsers**
- The output is basically code that can then be compiled/interpreted.
- Parsers are one of the most well understood parts of computer sciences, with a lot of time already invested in solving the problems
that can be thrown in the field. 
- So we have things like formats such as *CFG or Context Free Grammar* that describes the underlying grammar to create syntactically correct code. 
- The common forms of CFGs are BNF *(Backus-Naur Form)* or EBNF *(Extended Backus-Naur Form)*