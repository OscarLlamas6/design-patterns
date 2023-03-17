# Singleton Pattern

[Credits to Ramseykiang](https://medium.com/@ramseyjiang_22278/singleton-design-pattern-in-golang-6b140bdf8499)

The Singleton pattern is one of the most used design patterns out there and the easiest to grasp. It will provide a single instance of an object, and guarantee that there are no duplicates. At first, call to use the instance, if it’s created, and then reused between all the parts in the application that need to use the particular behaviour.

# Objectives

1. Create a single, shared value, of some particular type.
2. Create a restricted object creation of some type to a single unit along the entire

## Pros and Cons

### Pros of using the Singleton Pattern:

1. Provides a single point of access to a shared resource, ensuring that the resource is only instantiated once.
2. Promotes code reuse, as the shared resource can be accessed from anywhere in the code.
3. Improves code readability and organization, as the singleton pattern promotes the separation of concerns.
4. Enables global state management, as the singleton object can be used to store global state information.
5. Reduces memory footprint, as the singleton object is only instantiated once.

### Cons of using the Singleton Pattern:

1. Can introduce tight coupling between objects, as the singleton object may be accessed and used by multiple parts of the code.
2. It can make it more difficult to test code, as the singleton object may introduce a global state that can affect the behaviour of other parts of the code.
3. It can lead to race conditions if multiple threads try to access and modify the shared resource simultaneously.
4. It can introduce a global state that can make it more difficult to understand the code's behaviour.
4. It can make it more difficult to maintain and update the code, as the singleton object may be used by multiple parts of the code and may have multiple responsibilities.

It is important to carefully consider the pros and cons of using the Singleton Pattern and to choose the pattern that best fits the needs of your system. The Singleton Pattern is a powerful tool for managing shared resources, but it may not be the best choice for every situation.