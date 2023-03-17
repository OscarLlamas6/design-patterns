# Decorator Pattern

[Credits to Ramseyjiang](https://medium.com/@ramseyjiang_22278/decorator-design-pattern-in-golang-with-unit-tests-652649fb93fb)

## Concept

The Decorator is a structural pattern that dynamically adds new behaviours to objects by placing them inside particular wrapper objects, called
decorators. It allows you to decorate an existing type with more functional features without touching it. When you think about extending legacy code without the risk of breaking something, you should think of the Decorator pattern first.

The “Wrapper” is the alternative nickname for the Decorator pattern that clearly expresses the main idea of the pattern. A wrapper is an object that can be linked with some target objects. The wrapper contains the same set of methods as the target and delegates it to all requests it receives. However, the wrapper may alter the result by doing something before or after passing the request to the target.

## Objectives

The objectives of the Decorator Pattern are the following:

1. Allow behaviour to be added to individual objects at runtime.
2. Provide an alternative to subclassing for extending functionality.
3. Improve the flexibility and reusability of code by separating concerns.

## Scenarios

The Decorator Pattern can be helpful in many scenarios, such as:

1. Add functionality to an object at runtime, such as a logging or caching layer.
2. Separate concerns in code, such as separating business logic from logging or caching logic.
3. Extend the functionality of an existing object without modifying its structure, such as adding new methods or properties.

## Pros and Cons

The advantages of the Decorator pattern are:

1. Allows behaviour to be added to objects at runtime.
2. Provides an alternative to subclassing for extending functionality.
3. Improves the flexibility and reusability of code by separating concerns.
4. Enables the creation of complex object compositions by adding and removing behaviours.
5. Provides a simple way to add functionality to an existing object without modifying its structure.

The disadvantages of the Decorator pattern are:

1. It can lead to a large number of small objects that can be hard to manage.
2. It can make the code more complex by adding layers of decorators.
3. It can be difficult to understand if not implemented correctly.

## How to implement

The following are the steps on how to implement the decorator pattern.

1. Define an interface that represents the object that will be decorated and any methods it should have.
2. Create a concrete implementation of the interface, representing the object that will be decorated.
3. Define a decorator interface that extends the Component interface, and any methods that it should have.
4. Create a concrete implementation of the decorator interface, representing the behaviour that will be added to the decorated object.
5. Create an instance of the object being decorated, and pass it to the decorator to add behaviour to it.

## Conclusion

After learning the Decorator pattern, perhaps you feel it is similar to the proxy pattern. What’s the difference?

Difference between decorator pattern and proxy pattern. The Decorator pattern is used to add additional responsibilities to an object dynamically. It provides a flexible alternative to subclassing for extending functionality. The decorator objects are wrapped around the original object and the original object is not modified. The decorator pattern is often used for implementing cross-cutting concerns, such as logging or security, in a clean and reusable way. On the other hand, the Proxy pattern provides a placeholder for another object to control access to it. The proxy object acts as an intermediary between the client and the real subject. This pattern is often used for creating objects on demand, remote communication, or performance optimization.

In summary, the Decorator pattern is focused on adding functionality to an object, while the Proxy pattern is focused on controlling access to an object.

The Decorator pattern is a useful design pattern for adding behaviour to an object in a flexible and dynamic way, without modifying the original code. The decorator type implements the same interface of the type it decorates, and stores an instance of that type in its members. This way, you can stack as many decorators as you want by simply keeping the old decorator in the field of the new one.