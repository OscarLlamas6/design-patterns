# Composite Pattern

[Credits to Matthias Bruns](https://medium.com/@MTrax/golang-composite-pattern-b611f30fb2f)

In object-oriented programming, the Composite pattern is a structural design pattern that allows you to compose objects into a tree structure and work with the tree as if it were a singular object. This pattern is part of the famous “Gang of Four” design patterns.
Understanding the Composite Pattern

The Composite pattern is useful when you need to work with objects with a hierarchical structure. It allows you to create a tree-like structure where each node in the tree can be either a single object or a group of objects. All the nodes in the tree can be treated as individual objects or as part of a larger group.

In the Composite pattern, there are two types of objects: leaf objects and composite objects. Leaf objects are the end nodes in the tree and do not have any children. Composite objects, on the other hand, can have one or more child nodes, which can be either leaf objects or other composite objects.