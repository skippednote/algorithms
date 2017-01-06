Breadth-first Search
===

- Breadth-first Search is a sorting algorithm.
  - It is used to find whether there is a path between to points.
  - It tells you the shortest distance between those two points.
- It takes O(Nodes + Edges) time to complete.
- Problems are modeled as Graphs in Breadth-first search.
  - Each graph can have multiple nodes.
  - These nodes are linked to each other by edges.
  - Nodes can be nested further but Breadth-first Search likes in the first
    relation.
  - If the item you are looking for isn't found in the first relation it move to second relation and so on.
  - Nodes that are linked to each other are called Undirected Graphs (<-->).
  - Nodes that are linked to other nodes but not back in turn are called Undirected Graphs (-->).
- Incase the same item you are looking for appears multiple times, it cause a
  infinite loop.
  - This can be caused if you are looking randomly in multiple relations
  - To handle this you will use queues which will look in the current relation
    first and then move to the next linked one.
- Queues work like real life queues, first in first out (FIFO).
  - You cannot access random data.
  - It has only two operations: enqueue and dequeue.
  - Stacks are last in first out (LIFO).
- Graphs can be implemented using Hash Tables.
  - Where the key is a node and the value are it's first relations.
  - Nodes that have no children and aren't Directed can have empty values.
