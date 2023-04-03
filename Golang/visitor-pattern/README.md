# Visitor pattern

[Credits to Matthias Bruns](https://medium.com/@MTrax/golang-visitor-pattern-e3543f119987)

The Visitor Pattern is a behavioral design pattern that allows us to separate an algorithm from an object structure on which it operates. It is useful when we have a complex object structure and we want to perform some operations on it without changing the classes of the objects in the structure.

## How does it work?

The Visitor Pattern works by separating the algorithm from the object structure. To achieve this, we define an interface that represents the algorithm. This interface will have a method for each class in the object structure. Each method will take an instance of the class as an argument. The object structure will also have an interface that defines a method for accepting the visitor. This method will take an instance of the visitor interface as an argument. When the accept method is called, the visitor’s corresponding method for the class will be called.

The Visitor Pattern is a powerful design pattern that allows us to separate an algorithm from an object structure. It can be useful in many scenarios where we have a complex object structure and we want to perform some operations on it without changing the classes of the objects in the structure. In Golang, we can easily implement the Visitor Pattern using interfaces and methods.