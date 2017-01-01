Binary Search
===

- Input is a sorted list of elements.
- If element being searched is present then it's position in the list is returned.
- If element being searched isn't present then null is returned.
- At most it takes log<sub>2</sub>n steps to find the element in the list.
- log<sub>10</sub>100 is like asking "How many 10s do we multiply to get 100?".
    - log<sub>10</sub>100 == 10<sup>2</sup>.
    - For BigO log is always log<sub>2</sub>, log<sub>2</sub>8 == 2<sup>3</sup>.
    - logrithm is the flip of exponentials.
- Binary search runs in logrithmic time.
    - Unlike Simple search where time is linear.
    - Simple search is O(n) whereas Binary search is O(log<sub>n</sub>)
- Binary search get faster as items in the list grow.
    - log<sub>2</sub>16 == 4.
    - log<sub>2</sub>1 billion == 29.897.
- BigO helps calculate how an algorithm performs as the list grows.
    - It doesn't tell you the seconds an operation completes in.
    - It lets you compare the number of operations it will take to complete.
    - It address the worst case, not O(1) but O(n).
