# Adapter Pattern

The Adapter Pattern, one of the Gang of Four design patterns, is used to enable the interaction between two incompatible interfaces by creating a wrapper object that can translate requests between both interfaces. In Golang, this pattern is quite useful when you have two interfaces, but their functionalities do not match, and you need to make them work together.
Implementation

[Credits to Matthias Bruns](https://medium.com/@MTrax/golang-adapter-pattern-904598fe615d)

---

## The Adapter Pattern consists of four components:

    1. Target Interface: the interface that the client is expecting to interact with.
    
    2. Adapter: the object that implements the Target Interface and wraps the Adaptee.
    
    3. Adaptee: the interface that needs to be adapted to be used by the client.
    
    4. Client: the component that uses the Target Interface.

