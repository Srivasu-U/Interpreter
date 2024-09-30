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
    - In function call, `(` is considered the operator, with function name being the left operand and params being the right operand, i.e, this is an *infix* operation
    - In *indexing of array and hashes both*, `[` is considered the operator instead, and this is an *infix* operation. (The right operand is the expression that resolves to an index and the left operand can be an array/hash or an identifier that resolves to an array/hash)
        - Both hashes and arrays hence use the same prefix parsing method in our code, and it is registered with the `[` operator
        - In the parsing on an *array literal*, `[` is considered as the operator again, but this is a *prefix* operation
    - In hashes, `{` is the operator and this is *prefix*
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
- This is now possible in Monkey. The line `iter(rest(arr), push(accumulated, f(first(arr))));` executes inside out, just for future reference.
```
let map = fn(arr, f) {
        let iter = fn(arr, accumulated) {
                if (len(arr) == 0 ) {
                        accumulated
                } else {
                iter(rest(arr), push(accumulated, f(first(arr))));
                }
        };
        iter(arr, []);
};
let a = [1, 2, 3, 4];
let double = fn(x) { x * 2 };
map(a, double);
```

### Hash maps
- Keys can be of any type in Monkey: string, int or bool. It can even be an expression that resolves into one of the keys.
- This uses Go's `map` as the underlying data structure, naturally.
    - But since we can have keys of any datatype, we will have to modify certain behaviours.
- The retrieval of values from a hash map is an interesting problem
    - If we define the `Hash` object in our system as shown below, retrieval is convoluted. 
    ```
    type Hash struct {
        Pairs map[Object]Object
    }
    ```
    - Trying to retrieve with the following code, while the first line evaluated to `*object.String` with `.Value` as "
    name" mapped to `*object.String` with `.Value` as "Monkey".
    ```
    let hash = {"name": "Monkey"};
    hash["name"]
    ```
    - But when trying to access the value with the second "name" evaluates to a new `*object.String` and the comparison between the first and the second is false because both of these are pointers to different memory locations
        - While the individual `.Value`s can be compared, this increased the time from O(1) to O(n), which is not what we want with hashes
        ```
        name1 := &object.String{Value: "name"}
        monkey := &object.String{Value: "Monkey"}
        pairs := map[object.Object]object.Object{}
        pairs[name1] = monkey
        fmt.Printf("pairs[name1]=%+v\n", pairs[name1])
        // => pairs[name1]=&{Value:Monkey}
        name2 := &object.String{Value: "name"}
        fmt.Printf("pairs[name2]=%+v\n", pairs[name2])
        // => pairs[name2]=<nil>
        fmt.Printf("(name1 == name2)=%t\n", name1 == name2)
        // => (name1 == name2)=false
        ```
    - Hence we need a way to generate hashes for objects that are easy to compare between different pointers having the same value.
    - This is solved by using a Hash object (struct) with just an integer value
    ```
    type Hash struct {
        Type ObjectType
        value uint64
    }
    ```