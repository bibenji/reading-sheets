MVC

CompositeView
= assemble primitives into more complex objects

View-Controller = example of the Strategy design pattern

MVC uses Factory, Decorator, Observer, Composer, Strategy

Abstract Factory = Provide an interface for creating families of related or dependent
                   objects without specifying their concrete classes.

Adapter = Convert the interface of a class into another interface clients expect.
          Adapter lets classes work together that couldn't otherwise because of incompat-
          ible interfaces.

Bridge = Decouple an abstraction from its implementation so that the two can vary
                   independently.

Builder = Separate the construction of a complex object from its representation so
          that the same construction process can create different representations.

Chain of Responsibility = Avoid coupling the sender of a request to its receiver by
                          giving more than one object a chance to handle the request. Chain the receiving
                          objects and pass the request along the chain until an object handles it.

Command = Encapsulate a request as an object, thereby letting you parameter-
          ize clients with different requests, queue or log requests, and support undoable
          operations.

Composite = Compose objects into tree structures to represent part-whole hierar-
            chies. Composite lets clients treat individual objects and compositions of objects
            uniformly.
            
Decorator = Attach additional responsibilities to an object dynamically. Decorators
            provide a flexible alternative to subclassing for extending functionality.

Facade = Provide a unified interface to a set of interfaces in a subsystem. Facade
         defines a higher-level interface that makes the subsystem easier to use.

Factory Method = Define an interface for creating an object, but let subclasses de-
                 cide which class to instantiate. Factory Method lets a class defer instantiation to
                 subclasses.
                 
Flyweight = Use sharing to support large numbers of fine-grained objects effi-
            ciently
            
Interpreter = Given a language, define a represention for its grammar along with
              an interpreter that uses the representation to interpret sentences in the language

Iterator = Provide a way to access the elements of an aggregate object sequentially
           without exposing its underlying representation.

Mediator = Define an object that encapsulates how a set of objects interact. Me-
           diator promotes loose coupling by keeping objects from referring to each other
           explicitly, and it lets you vary their interaction independently.

Memento = Without violating encapsulation, capture and externalize an object's
          internal state so that the object can be restored to this state later.

Observer = Define a one-to-many dependency between objects so that when one
           object changes state, all its dependents are notified and updated automatically.

Prototype = Specify the kinds of objects to create using a prototypical instance, and
            create new objects by copying this prototype.

Proxy = Provide a surrogate or placeholder for another object to control access to
        it.

Singleton = Ensure a class only has one instance, and provide a global point of
            access to it.

State = Allow an object to alter its behavior when its internal state changes. The
        object will appear to change its class.

Strategy = Define a family of algorithms, encapsulate each one, and make them
           interchangeable. Strategy lets the algorithm vary independently from clients that
           use it.

Template Method = Define the skeleton of an algorithm in an operation, deferring
                  some steps to subclasses. Template Method lets subclasses redefine certain steps
                  of an algorithm without changing the algorithm's structure.

Visitor = Represent an operation to be performed on the elements of an object
          structure. Visitor lets you define a new operation without changing the classes of
          the elements on which it operates.

Class & Creational : Factory Method
Object & Creational : Abstract Factory, Builder, Prototype, Singleton

Class & Structural : Adapter (class)
Object & Structural : Adapter (object), Bridge, Composite, Decorator, Facade, Flyweight, Proxy

Class & Behavioral : Interpreter, Template Method
Object & Behavioral : Chain of Responsibility, Command, Iterator, Mediator, Memento, Observer, State, Strategy, Visitor

Determining Object Granularity

P. 13

Principles :
- Program to an interface, not an implementation.
- Favor object composition over class inheritance.

Delegation

Delegation is an extreme example of object composition. It shows that you can always
replace inheritance with object composition as a mechanism for code reuse.

Parametrized Types
(or generics, or templates : =)

Acquaintance and aggregation
acquaintance is a weaker relationship than aggregation

Causes of redesign with design pattern solution :
- Creating an object by specifying a class explicitly.
- Dependence on specific operations.
- Dependence on hardware and software platform.
- Dependence on object representations or implementations.
- Algorithmic dependencies.
- Tight coupling.
- Extending functionality by subclassing.
- Inability to alter classes conveniently.

Application Programs

P. 25

Inheritance and composition

RULE : Favor object composition over class inheritance

...

Delegation...

In other words, instead of a Window being a Rectangle, it would have a Rectangle.
Window must now forward requests to its Rectangle instance explicitly, whereas before it would have inherited those operations.

Delegation is an extreme example of object composition.

Inheritance versus Parameterized Types

Parameterized types give us a third way (in addition to class inheritance and object composition) to compose behavior in object-oriented systems.

...

Acquaintance is a weaker relationship than aggregation and suggests much looser coupling between objects.

Aggregation relationships tend to be fewer and more permanent than acquaintance. Acquaintances, in contrast, are made and remade more frequently, sometimes existing only for the duration of an operation.

some common causes of redesign :
1. Creating an object by specifying a class explicitly.
2. Dependence on specific operations.
3. Dependence on hardware and software platform.
4. Dependence on object representations or implementations.
5. Algorithmic dependencies.
6. Tight coupling.
7. Extending functionality by subclassing.

Object composition in general and delegation in particular provide flexible alternatives to inheritance for combining behavior.

8. Inability to alter classes conveniently.

2. A Case Study: Design a Document Editor

P. 46

Embellishing the User Interface

P. 56

