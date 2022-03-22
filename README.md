# Interpreter for the [Monkey programming language.](https://interpreterbook.com)

Based on Thorsten Ball's book [Writing An Interpreter In Go](https://interpreterbook.com)

# Usage
Run the REPL:
```shell
$ go run main.go
Hello xiovv!
>> 
```

## Getting started
### Variables
Variables are defined using the `let` keyword. They may be integers, strings, arrays or hashes.
```shell
let a = 5;
let b = 10;
let c = a + b; 
let hello = "hello, ";
let world = "world";

>> puts(a, b, c, hello, world, hello + world)
5
10
15
hello, 
world
hello, world
```

### Arithmetic operations
Monkey supports all the basic arithmetic operations.
```shell
let a = 10;
let b = 5;
>> puts(a + b, a - b, a * b, a / b);
15
5
50
2
```

## Builtin data structures
Monkey contains two builtin data structures: an `array` and a `hash`.

### Arrays
Arrays in monkey can hold multiple types and even expressions.
```shell
let a = [1, 2, 5 + 5, "hello", true];
>> puts(c);
[1, 2, 10, hello, true]
```
Adding items to an array can be achieved via the `push` function:
```shell
push(a, "new item");
```
Fetching the first item of the array can be achieved via the `first` function:
```shell
first(a);
```
Fetching the last item of the array can be achieved via the `first` function:
```shell
last(a);
```
Fetching every element of the array except the first one can be achieved via the `last` function:
```shell
rest(a);
```

### Hashes
A hash is a key/value store. Known as a map or dictionary in other languages. Keys can only be of type `int`, `string` and `boolean`.
```shell
let a = { "language": "monkey", true: 1, 5: "five" };
>> puts(a);
{language: monkey, true: 1, 5: five}
>> puts(a["language"], a[true], a[5])
monkey
1
five
```

## Functions
Monkey uses the `fn` keyword to define a function which is assigned to a variable.
```shell
let add = fn(a, b) { return a + b; };
>> puts(add(1, 2));
3
```
Functions can be treated as normal variables:
```shell
let addTwo = fn(a, b, f) { return 2 + f(a, b); };
>> puts(addTwo(1, 2, add));
5
```

### If-else statements
Monkey supports classic if-else statements.
```shell
let a = 10;
let b = 5;
if (a > b) {
  puts("a is greater than b");
} else {
  puts("a is less than b");
}

a is greater than b
```