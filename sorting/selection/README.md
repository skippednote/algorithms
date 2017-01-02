Selection Sort
===

#### Memory
- Memory can be thought of as a huge set drawers.
    - Each drawer stores some data.
    - Each drawer has an address(feØffeeb) which can be used to access it.
    - Multiple items can be stored in memory using Arrays and Lists.

#### Arrays
- Arrays are stored contiguously (next to each other) in memory.
- If the adjacent slot in memory is not empty then the whole array is moved to a location where all the items can sit next to each other.
- Adding data into arrays is slow.
    - You need to move the whole array if there isn't space in the adjacent slot.
- To avoid moving the array from one location to another you can hold up memory slots during initialization.
    - It can lead to wastage of memory as some slots might not end up being used.
    - You might need more than 10 slots and then you still have to move to a different location in memory.
- Arrays are fast at reading data.
    - Since all the items in the array are placed next to each other, it is easy to calculate the index of any item in the array.
    - If the first item in an array is at address 101 and the list is 10 items longs, then the last item would be at the address 111 (101 + 10).
- Arrays start from index 0 not 1.
    - Name of the array is a pointer and array[n] refers to a memory location n-element away from starting element.
    - array[0] is same as the memory reference to the array.
    - Array indexs are used as offsets.
    - array[0] is 0 elements away from the first item.
- Arrays aren't the best at adding data to the middle.
    - You might end up moving the whole array which is slow (depending on the size of the array).
    - Array insertions can fail if there is no memory left.
- Arrays are slow at removing data.
    - You might end up moving the whole array which is slow.
- Arrays are used more because of their fast random access property.

#### Linked Lists
- Data can be added anywhere in memory with Linked Lists.
    - You don't need to move the list around as you add more data.
- The data in Linked List store the content and the address to the next item in the list.
    - The first item has the address of the second, the second has the address of third and so on.
- Adding data into Linked Lists is fast.
    - You just store the data anywhere in memory and store the address in the last item of the list.
- Linked Lists are slow at reading data.
    - The first item has the address of second, second has the address of third and so on until the last item.
    - If you have to get the data from the last item in the list, then you will have to traverse the whole list.
    - With Linked List we don't know the address of the items beforehand.
- Linked Lists are better at adding data to the middle.
    - You will need to update the address of the item so you can link it to the new item.
    - Linked List insertions can fail if there is no memory left.
    - Every subsequent item needs to be updated to update the memory address of the next item in the list.
- Linked Lists are better at deleting data.
    - If you want to delete the n<sup>th</sup> item, you just need to update the memory address of n+1 to point to n-1 instead of n.
        - You should keep track of first and last item of Linked List for fast deletion.
        - It can be slow if you want to delete the middle element as you first need to traverse to it.
- Linked Lists aren't used as much as arrays because of their sequential access property.
    - Linked List is however, fast at sequential access.

|  | Arrays | List |
| - | - | - |
| Reading | O(1) | O(n) |
| Inserting | O(n) | O(1) |
| Deletion | O(n) | O(1) |

#### Selection Sort
- It takes O(n<sup>2</sup>) to sort a list using Selection Sort.
- Even though the number of items to look through is decresed with every iteration we still use n x n every time.
- It is not the fastest sorting algorithm.
- List is divided into two parts sorted and unsorted.
    - Sorted part sits on the left side.
    - Unsorted part sits on the right side.
- You find the minimum element in the list and move it to the first place.
    - The value at the first place and the minimum value are swapped everytime.
    - Then we increment the iterator, repeating for n-1.
    - Then find the minimum element in the remaining list and move it to the second place and so on.