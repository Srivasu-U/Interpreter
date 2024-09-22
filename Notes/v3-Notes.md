## Notes for Monkey v3

### Evaluation, the E in REPL
- Giving code meaning
- Set of rules that defines how a language is interpreter
- If we consider the following code, whether this returns `a` or `b` is dependent of if the interpreter being used considering 5 truthy or not. 
It is truthy in some languages but in some languages, we would have to use something like `5 != 0`
```
let num = 5;
if (num) {
    return a;
} else {
    return b;
}
```
- There are no clear distinctions between what an interpreter and a compiler is
    - The most commonly accepted idea is that compilers provide an executable artifacts while interpreters don't but this doesn't always hold true.
- Interpreters that traverse the AST and do what the node represents are called *tree-walking interpreters*
- There are also interpreters that traverse the AST but don't directly interpret it. Instead, the AST is converted in bytecode. 
    - Bytecode is an intermediate representation (IR) that is neither assembly language, nor machine code.
    - It is not executed by the OS or the CPU, but it is instead interpreted by a VM.
    - If this bytecode is not executed directly in the VM, but the VM instead converted this bytecode to native machine code right before execution, this is called a JIT interpreter/compiler.
        - JIT can also be the name for something that skips the bytecode step entirely. Just compile to native machine code and execute.


### Tree-Walking Interpreter        
- As usual, the choice of style is based on performance and usability needs. Tree-walking is the slowest but the most easy to build and extend. ***This is what is built as part of this project.***
- Based on the Lisp interpreter, inspiration from *The Structure and Interpretation of Computer Programs* - [click here](https://web.mit.edu/6.001/6.037/sicp.pdf)
- Pseudocode for our evaluator
```
function eval(astNode) {
    if (astNode is integerLiteral) {
        return astNode.integerValue
    } else if (astNode is booleanLiteral) {
        return astNode.booleanValue
    } else if (astNode is infixExpression) {
        leftEvaluated = eval(astNode.Left)
        rightEvaluated = eval(astNode.Right)

        if astNode.Operator == '+'{
            return leftEvaluated + rightEvaluate
        } else if astNode.Operator == '-' {
            return leftEvaluated - rightEvaluated
        }
    }
}
```
- What is the return type for eval? - Dependent on the internal object system of the interpreter (?)