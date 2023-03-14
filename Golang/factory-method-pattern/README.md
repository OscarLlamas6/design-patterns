# Factory method pattern

[Credits to Matthias Bruns](https://medium.com/@MTrax/golang-factory-method-pattern-7c25c0c7cdbd)

The Factory Method pattern is one of the most commonly used creational patterns in software development. It is part of the famous “Gang of Four” design patterns mentioned in their book “Design Patterns: Elements of Reusable Object-Oriented Software”. In this article, we will discuss how to implement the Factory Method pattern in Golang.
Understanding the Factory Method Pattern

The Factory Method pattern is a creational pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. In simpler terms, it is a way of creating objects without specifying the exact class of object that will be created.

The Factory Method pattern is useful when we want to create objects without knowing their exact class. Instead of creating objects directly, we delegate the creation to a factory that can decide which class of object to create based on certain conditions. This makes our code more flexible and easier to maintain.
Implementing Factory Method Pattern in Golang

In Golang, we can implement the Factory Method pattern using an interface and a struct. The interface will define a method that will be implemented by the struct to create objects.

## Advantages of Using the Factory Method Pattern

The Factory Method pattern provides several advantages over other creational patterns. These include:

    1. Flexibility: The Factory Method pattern allows us to create objects without knowing their exact class. This makes our code more flexible and easier to maintain.

    2. Scalability: The Factory Method pattern can be easily scaled to support new types of objects. We can simply create a new subclass of the factory and implement the CreateProduct() method to create the new type of object.
    
    3. Decoupling: The Factory Method pattern decouples the creation of objects from their usage. This makes our code more modular and easier to test.

Conclusion

The Factory Method pattern is a useful pattern for creating objects without knowing their exact class. In Golang, we can implement this pattern using an interface and a struct. By using this pattern, we can write better and more flexible code. So, if you want to improve the scalability, maintainability and flexibility of your code, consider using the Factory Method pattern.