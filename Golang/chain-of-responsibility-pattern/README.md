# Chain of responsibility Pattern

[Credits to Matthias Bruns](https://medium.com/@MTrax/golang-chain-of-responsibility-pattern-1729895c9fd7)

The Chain of Responsibility pattern is a behavioral design pattern that allows an object to pass a request along a chain of handlers until one of the handlers can handle the request. This pattern is useful when there are multiple objects that can handle a request, but we do not know which object can handle the request until runtime.

The Chain of Responsibility pattern is a powerful pattern that allows us to handle requests in a flexible and extensible way. In Golang, we can implement this pattern using interfaces and structs, as shown in the example above.