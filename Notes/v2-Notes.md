## Notes for Monkey-v2

### Parsers
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
- As described below, while parser generators can be used to solve problems, and just plugged into this project, the goal is to understand how parsers work which is why we are writing our own.
- In the Monkey language, everything other than let and return statements is an expressions. Essentially, expressions evaluated into some result while statements do not. So expression parsing such as operator precendence, function calls, is part of the challenge of writing a parser.


### Parser generators
- These are tools that, when given a language description, produces the parsers. Ex: ANTLR, yacc or bison.  
**Note: Read more about parsers**
- The output is basically code that can then be compiled/interpreted.
- Parsers are one of the most well understood parts of computer sciences, with a lot of time already invested in solving the problems
that can be thrown in the field. 
- So we have things like formats such as *CFG or Context Free Grammar* that describes the underlying grammar to create syntactically correct code. 
- The common forms of CFGs are BNF *(Backus-Naur Form)* or EBNF *(Extended Backus-Naur Form)*

### Parsing strategies
- Can either be top-down or bottom-up
- Top down starts with the root node and recursively builds the tree downwards. Bottom-up does the opposite
- Top-down parser examples: recursive decent parsing, early parsing, predictive parsing
    - Recursive decent, also called Pratt parsing, this is used in Monkey-v2. Recursive decent is basically about looking at the next
    bit of code and figuring out what to do.
    - This works well with keywords such as `let`, `if` and so on. Trickier with expressions
    - Pratt parsing solves this issue but mixing it with recursive descent - Read more [here (written in Java)](https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/)
    - Pratt parsing essentially associated a parsing function with a token type. If that token type is encountered, then the parsing function is
    executed and returns an AST node to represent it.
        - Each token type can have two parsing functions: for prefix or infix positions
- **Check: Parser error recovery**
- **Check: What is formal proof of correctness for parsers?**
- The idea is to get a minimal parser that works with Monkey, is extendible and a good starting point.
- What does a parser that work correctly mean? Accurate production of an AST that conveys the right information. This is where design comes in. Refer this [link](https://stackoverflow.com/questions/16066454/parsing-which-method-choose) and the first answer to find how design can impact the parser used.
- All of our parsing functions are going to follow this protocol: start with current token being the type of token you’re associated with and return with current token being the last token that’s part of your expression type. Never advance the tokens too far.

## AST used in this project
- The initial structure of the AST as put forth in [ast.go](/ast/ast.go) is as follows
    - Interface `Node` describes a node in the tree and has the `TokenLiteral()` method
    - There are `Statement` and `Expression` nodes, which are specialized interfaces, that implement `Node` and their own methods
    - We have the `Program` struct which is an array of `Statement` nodes
    - Initially, we have a `LetStatement` struct which describes the structure of a `let` statement, ie, the `let` keyword, identifier and expression. This implements `Statement` 
        - ***Question:* Must all the different types of statements be their own struct? Is this not tedious or is there a better way to do this, more efficient?**
        - From the book  

        ![Node tree](/Learning-Go/monkey-v2/assets/letStatementNodeTree.png)
- Excellent parser pseudocode
```
function parseProgram() {
    program = newProgramASTNode()
    
    advanceTokens()
    
    for (currentToken() != EOF_TOKEN) {
        statement = null

        if (currentToken() == LET_TOKEN) {
            statement = parseLetStatement()
        } else if (currentToken() == RETURN_TOKEN) {
            statement = parseReturnStatement()
        } else if (currentToken() == IF_TOKEN) {
            statement = parseIfStatement()
        }

        if (statement != null) {
            program.Statements.push(statement)
        }

        advanceTokens()
    }
    return program
}

function parseLetStatement() {
    advanceTokens()

    identifier = parseIdentifier()
    advanceTokens()
    if currentToken() != EQUAL_TOKEN {
        parseError("no equal sign!")
        return null 
    }
    advanceTokens()

    value = parseExpression()
    variableStatement = newVariableStatementASTNode()
    variableStatement.identifier = identifier
    variableStatement.value = value
    return variableStatement
}

function parseIdentifier() {
    identifier = newIdentifierASTNode()
    identifier.token = currentToken()
    return identifier
}   

function parseExpression() {
    if (currentToken() == INTEGER_TOKEN) {
        if (nextToken() == PLUS_TOKEN) {
            return parseOperatorExpression()
        } else if (nextToken() == SEMICOLON_TOKEN) {
            return parseIntegerLiteral()
        }
    } else if (currentToken() == LEFT_PAREN) {
        return parseGroupedExpression()
    }
    // [...]
}

function parseOperatorExpression() {
    operatorExpression = newOperatorExpression()    
    operatorExpression.left = parseIntegerLiteral()
    operatorExpression.operator = currentToken()
    operatorExpression.right = parseExpression()
    return operatorExpression()
}
// [...]
```


## General Go notes
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
- Functions can also have multiple return values
```
func A(x, y int) (string, string) {...}
```
- Consts do not use datatypes or the short hand operator `:=`
- Type convertion works like this: `datatype(value)`. Example: `float64(2)` converts int 2 into float64
- We also have type checking/assertion like this: `ident.(type)`. Example: `value.(string)` checks if value contains a string expression
- Interesting error message: `impossible type assertion: program.Statements[0].(*ast.Statement)
	*ast.Statement does not implement ast.Statement (type *ast.Statement is pointer to interface, not interface)`. What does this mean?