# Abstract factory pattern

[Credits to Matthias Bruns](https://medium.com/@MTrax/golang-abstract-factory-pattern-ec9132b589ab)

The “Gang of Four” design patterns are a set of commonly used software engineering principles that provide solutions to recurring problems in software design. One of these patterns is the Abstract Factory Pattern, which provides an interface for creating families of related objects without specifying their concrete classes. In this article, we will explore how to implement the Abstract Factory pattern in Golang, with examples.
What is the Abstract Factory Pattern?

The Abstract Factory pattern is a creational pattern that provides an interface for creating families of related objects without specifying their concrete classes. In other words, the Abstract Factory pattern defines an abstract class or interface for creating related objects without specifying their exact class types. This allows us to create objects based on a family of related objects, rather than creating individual objects.
Why use the Abstract Factory Pattern?

The Abstract Factory pattern provides a way to create related objects without specifying their concrete classes. This is useful in situations where we want to create objects based on a certain theme or context. For example, if we are building a GUI application, we might want to create a set of related objects, such as windows, buttons, and menus, that all have the same look and feel. By using the Abstract Factory pattern, we can create these objects without specifying their concrete classes, which makes it easy to change the theme or context of our application.
Implementing the Abstract Factory Pattern in Golang

To implement the Abstract Factory pattern in Golang, we first define an interface that defines methods for creating related objects. This interface is called the abstract factory. We then define concrete implementations of this interface, which are called concrete factories. The concrete factories are responsible for creating the actual objects.