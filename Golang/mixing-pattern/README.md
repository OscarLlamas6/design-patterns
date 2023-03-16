# Mixing Pattern

[Credits to Ramseyjiang](https://medium.com/@ramseyjiang_22278/mixin-design-pattern-in-golang-with-unit-tests-672ceb94f348)

The mixin pattern is a structural design pattern that allows developers to dynamically add behaviour to an object without creating a new subclass or modifying the object itself. Mixins are like building blocks of functionality that can be added to an object, and they can be combined with other mixins to create a composite object with all the behaviour of its constituent parts. In the Mixin pattern, a struct that provides a specific behaviour is called a mixin, and a struct that receives that behaviour is called a receiver. The receiver struct can access the methods and properties of the mixin struct as its own.

## Objectives
The objectives of the mixin pattern are:

Allow developers to dynamically add behaviour to an object without modifying the object itself or creating a new subclass.
Encourage code reuse by breaking down functionality into smaller, more manageable pieces that can be combined in different ways.
Reduce the complexity of code by separating concerns into distinct modules or mixins.

## Pros and Cons

### The advantages of using the mixin pattern are:

1. It promotes code reuse and modularization.
2. It allows developers to dynamically add methods to an object without modifying the object or creating a new subclass.
3. It makes it easier to maintain and update code by separating concerns into distinct modules or mixins.
4. It encourages the use of composition over inheritance, which can lead to simpler, more flexible code.

### The disadvantages of using the mixin pattern are:

1. It can lead to code that is harder to understand and maintain if the number of mixins used is excessive.
2. It may increase the complexity of the codebase if not used judiciously.
3. It can lead to name collisions if the same method name is used in multiple mixins.

## Scenarios

The mixin pattern is useful in the following scenarios:

1. Implementing cross-cutting concerns such as logging, caching, or authentication across multiple classes.

2. Adding additional functionality to an existing class without modifying its source code.

3. Creating smaller, more focused classes that are easier to understand and maintain.

## How to implement

1. Define the interface that the mixin will implement. This interface should include all of the methods that the mixin will provide.
2. Define a mixin struct that implements the interface. This struct will provide the implementation code for all methods in the interface.
3. Define a struct that embeds the interface. This struct is the receiver struct, you want to mix in the behaviour to.
4. Use the implementation code from the mixin struct in the methods of the embedded interface. You can access the implementation code by using the field name of the embedded struct.
5. Create an instance of the mixin struct and initialize the embedded field in the target object.

## Conclusion

After learning the Mixin pattern, perhaps you feel it is similar to the decorator pattern. The main differences between the Mixin pattern and the Decorator pattern are as follows:

1. The Mixin pattern adds behaviour at compile-time, while the Decorator pattern adds behaviour at runtime.
2. In the Mixin pattern, the behaviour is explicitly inherited by the receiver struct, while in the Decorator pattern, the behaviour is dynamically wrapped around the original object.
3. The Mixin pattern is used for adding specific behaviours to a class, while the Decorator pattern is used for adding generic behaviours to a class.
4. The Mixin pattern enhances code reuse and flexibility by enabling a receiver struct to use the behaviour of one or more mixins. 
The Decorator pattern enhances code flexibility by enabling the dynamic addition and removal of behaviours at runtime. 

In summary, the Mixin pattern and the Decorator pattern are two different patterns that provide ways to add behaviours to objects. The Mixin pattern adds behaviour at compile-time and requires explicit inheritance, while the Decorator pattern adds behaviour at runtime and provides a flexible way to modify an object’s behaviour dynamically.