### Iterator Pattern

Suppose you have a collection, such as an array or list, and you need to provide a mechanism for accessing its elements. Without the iterator pattern, the client code needs to understand how the collection is structured, and different collections would require different methods to traverse them

## Iterator Pattern

Problem: How to access elements in a collection without exposing its internal representation

Solution: The Iterator Pattern provides a way to traverse a collection without revealing its underlying structure, offering a uniform interface for traversal

Structure:

-> Iterator: Interface for traversing a collection
-> Collection: Holds the elements and provides an iterator

# Iterator Pattern Benefits

1. Seperation of Concerns: The traversal logic is seperated from the collection itself, allowing you to change one without affecting the other

2. Uniform Interface: The same interface (Iterator) is used to traverse different types of collections, making the code more flexible

3. Simplified Client Code: The client doesn't need to know the underlying data structure, reducing coupling and making the code easier to maintain

4. Multiple Traversal Strategies: You can implement multiple types of iterators (e.g forward, backward, filtered) without changing the collection