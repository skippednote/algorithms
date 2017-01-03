Recursion
===

- Recursive function is a function that calls itself until a base condition is met.
- Recursion has two cases: Base and Recurive.
    - Base case is when the function meets the condition and stops calling itself.
    - Recurive case is when the function hasn't met the condition and calls itself.
- Recursion doesn't necessarily provide any performance benefits.
    - It can be slower than `for` loop.
- Recursion make it simple to understand the problem being solved.

#### Stack
- Stack is a data structure to which we can write data and remove/read from.
    - It has two operations: push and pop.
- Stacks are last in first out.
- A call stack contains all the variables and function calls registered by the parent function.
    - All function calls go to call stack.
- Each of the variables and functions are removed from the call stack once the function that introduces them returns.
- Resurion can quickly lead to filling up the call stack which can lead to stackoverflow.
- Tail Call Recursion can solve this problem and is supported by languages like Clojure and Elixir which rely on recursion instead of loops.
