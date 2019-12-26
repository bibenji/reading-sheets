# Clean Code

p. 40

p. 48

p. 62

FUNCTIONS SHOULD DO ONE THING. THEY SHOULD DO IT WELL. THEY SHOULD DO IT ONLY.

The Step-down Rule

p. 68

it does more than one thing<br />
Single Responsibility Principle (SRP)<br />
Open Close Principle (OCP)<br />
unlimited number of other functions that will have the same structure

Function Argument<br />
best : zero (= niladic)<br />
monadic, dyadic and triadic

make it an instance variable

monadic for event, or something transformed

flag arguments are ugly, instead use 2 functions

dyadic not good too, except for new Point(0, 0)<br />
event assertEquals(expect, actual) is not good

triads are evil

(include function inside sub class)

ok for argument lists like format functions (considered as dyadic functions)

assertEquals might be better written as assertExpectedEqualsActual(expected, actual)

Have No Side Effects

for function : DO ONE THING PRINCIPLE

Output Arguments : sometimes, you don’t know if argument is an output or an input<br />
output argument should be avoided<br />
prefer change state of the owning object

Command Query Separation : function should do something or answer something, not both

```
if (attributeExists(“username”)) {
	setAttribute(“username”, “unclebob”);
}
```

P. 77

Prefer Exceptions to Returning Error Codes<br />
Extract Try/catch Blocks<br />
Error Handling Is One Thing<br />
The Error.java Dependency Magnet -> Enum where all error codes are defined -> better is to think of errors as derivatives of the exception class

## 4 Comments

<i>"The proper use of comments is to compensate for our failure to express ourself in code."</i>

Comments Do Not Make Up for Bad Code

Explain Yourself in Code

Good Comments :
- Legal Comments
- Informative Comments
- Explanation of Intent (...)
- Clarification (in code that you cannot alter i.e.)
- Warning of Consequences
- TODO Comments
- Amplification
- Javadocs in Public APIs

Bad Comments :
- Mumbling
- Redundant Comments (not usefull to write comments to say the same thing as the property name i.e.)
- Misleading Comments
- Mandated Comments -> not everything should have a comment<br />
  i.e., not good:
  ```
  /**
  *
  * @param title The title of the CD
  * @param author The author of the CD
  * @param tracks The number of tracks on the CD
  * @param durationInMinutes The duration of the CD in minutes
  */
  public void addCD(String title, String author, int tracks, int durationInMinutes) {
    CD cd = new CD();
    cd.title = title;
    cd.author = author;
    cd.tracks = tracks;
    cd.duration = duration;
    cdList.add(cd);
  }
  ```
- Journal Comments
- Noise Comments
  ```
  /** The day of the month. */
  private int dayOfMonth;
  ```
- Scary Noise
- Don't use a comment when you can use a function or a variable
- Position Markers
- CLosing Brace Comments
- Attributions and Bylines
- Commented-Out Code (delete it!)
- HTML Comments
- Nonlocal Information
- Too Much Information
- Inobvious Connection (connection between a comment and the code it describes should be obvious)
- Function Headers
- Javadocs in Nonpublic Code

## 5 Formatting

P. 106

P. 112 (Instance variables)

## 6 Objects and Data Structures

P. 124

[...]

The Law Of Demeter

a method f of a class C should only call the methods of these:
- C
- An object created by f
- An object passed as an argument to f
- An object held in an instance variable of C

The method should not invoke methods on objects that are returned by any of the allowed functions. In other words, talk to friends, not to strangers.

= Train Wrecks

## 7 Error Handling

P. 134

Use different classes only if there are times when you want to catch one exception and allow the other one to pass through.

SPECIAL CASE PATTERN

Don't Return Null

If you are tempted to return null from a method, consider throwing an exception or returning a S PECIAL C ASE object instead.

If you are calling a null -returning method from a third-party API, consider wrapping that method with a method that either throws an exception or returns a special case object.

Don’t Pass Null

## 8 Boundaries

P. 144

[...]

learning tests...
help learn and help detect bugs if the third-party evolve

wrap third-party with custom code wrapper

## 9 Unit Tests

P. 152

TDD : Test Driven Developpment

[...]

BUILD-OPERATE-CHECK pattern

given-when-then convention

FIRST (Fast Independent Repeatable Self-validating Timely)

## 10 Classes

P. 166

The Single Responsibility Principle -> classes should have one responsibility, one reason to change

Cohesion

Classes should have a small number of instance variables<br />
Each of the methods of a class should manipulate one or more of those variables.

découper en classes qui regroupent des fonctions qui utilisent les mêmes variables (= un bon procédé)

OCP = Open-Closed Principle<br />
Classes should be open for extension but closed for modification.

Isolating from Change

DIP = Dependency Inversion Principle
-> our classes should depend upon abstractions, not on concrete details

## 11 Systems

P. 184

The seperation of concerns = one of the oldest and most important design techniques in our craft

The Single Responsibility Principle (déjà vu)

Separation of Main : assume everything has been build correctly

Factories

Abstract Factory pattern...

Dependency Injection (DI)

Inversion of Control (IoC)

JNDI = a "partial" implementation of DI where an object asks a directory server to provide a "service" matching a particular name

Scaling Up

Cross-Cutting Concernes

[...] AOP (aspect-oriented programming)

[...]

...never forget to use the simplest thing that can possibly work.

## 12 Emergence

P. 202





Further Reading :
---

- Literate Programming by Knuth

