Quick Sort
===

- Quick sort uses divide and conquer to sort a list.
- Quick sort is faster than Selection Sort in most cases.
- Quick sort performs O(nlogn) in average and best case.
- Quick sort performs O(n<sup>2</sup>) in the worst.
- Quick sort can be recursive or iterative.
- The list is split into two parts.
    - This is based on a value know as **Pivot**.
    - The two parts are based on the values higher and lower than Pivot.
    - Higher part is moved right of the Pivot.
    - Lower part is moved left of the Pivot.
    - Pivot should be choosen from the middle of the list for better performance.
- Divide and Conquer
    - Base case is a empty list or a list with just one element.
    - List is broken down into smaller and smaller pieces until it meets the base case.