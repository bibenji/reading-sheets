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

Getting Clean via Emergent Design

Simple Design Rules By Kent
1. Runs All the Tests
2. Contains no duplication
3. Expresses the intent of the programmer
4. Minimizes the number of classes and methods

Simple Design Rules 2-4: Refactoring

[...] Eliminate duplication, ensure expressiveness, and minimize the number of classes and methods.

No Duplication

the TEMPLATE METHOD pattern

Expressive
- choosing good names
- keeping functions and classes small
- using standard nomenclature
- well-written unit tests are also expressive

Minimal Classes and Methods

## 13 Concurrency

P. 208

...

servlets...

Challenges... P. 211

Single Responsibility Principle (again) = a given method/class/component should have a single reason to change

Corollary: Limit the Scope of Data

Take data encapsulation to heart; serverly limit the access of any data that may be shared.

Corollary: Use Copies of Data

Copy objects and treat them as read-only (if possible)

If not possible, copy objects, collect results from multiple threads in these copies and then merge the results in a single thread.

Corollary: Threads Should Be as Independent as Possible

des outils dans JAVA : ConcurrentHashMap, ReentrantLock, Semaphore, CountDownLatch... See java.util.concurrent, java.util.concurrent.atomic, java.util.concurrent.locks

Know Your Execution Models : Bound Resources, Mutual Exclusion, Starvation, Deadlock, Livelock

Producer-Consumer: the queue is a bound resource

Readers-Writers ...

Dinning Philosophers

Beware Dependencies Between Synchronized Methods :
- avoid using more than on method on a shared object
...

Keep Synchronized Sections Small

[...]

Writing Correct Shut-Down Code Is Hard ...

Testing Threaded Code ...

- Treat Spurious Failures as Candidate Threading Issues ...
- Get your nonthreaded code working first ...
- Make your threaded code pluggable ...
- Make your threaded code tunable ...
- Run with more threads than processors ...
- Run on different platforms ...
- Instrument your code to try and force failures ...

[...]

## 14 Successive Refinement

P. 224

[...]

remove the type-case...

[...]

## 15 JUnit Internals

P. 282

## 16 Refactoring SerialDate

P. 298

P. 308

## 17 Smells and Heuristics

P. 316

Comments

C1: Inappropriate Information
C2: Obsolete Comment
C3: Redundant Comment
C4: Poorly Written Comment
C5: Commented-Out Code

Environment

E1: Build Requires More Than One Step -> no, just one command
E2: Tests Require More Than One Step -> no, just one command to launch all

Functions

F1: Too Many Arguments -> should have a small number of arguments (more than three is not good)
F2: Output Arguments -> output arguments are counterintuitive
F3: Flag Arguments -> confusing, should be eliminated
F4: Dead Function ->

General

G1: Multiple Languages in One Source File -> no, only one language in one file
G2: Obvious Behavior Is Unimplemented
G3: Incorrect Behavior at the Boundaries -> look for every boundary condition (corner case, etc.)
G4: Overridden Safeties -> don't
G5: Duplication -> DRY principle (Don't Repeat Yourself), NEVER
G6: Code at Wrong Level of Abstraction -> higher level of abstraction in base class
G7: Base Classes Depending on Their Derivatives
G8: Too Much Information
G9: Dead Code
G10: Vertical Separation
G11: Inconsistency -> avoid
G12: Clutter -> should be removed (things not used, etc.)
G13: Artificial Coupling
G14: Feature Envy -> the methods of a class should be interested in the variables and functions of the class they belong to, and not the variables and functions of other classes.
G15: Selector Arguments -> Better to have many functions than to pass some code into a function to select the behavior
G16: Obscured Intent
G17: Misplaced Responsibility
G18: Inappropriate Static
G19: Use Explanatory Variables
G20: Function Names Should Say What They Do -> if you have to look at the implementation (or documentation) of the function to know what it does, then you should work to find a better name [...]
G21: Understand the Algorithm
G22 : Make Logical Dependencies Physical
G23: Prefer Polymorphism to If/Else or Switch/Case -> Just one switch to have polymorphic objects used by the rest of the system
G24: Follow Standard Conventions
G25: Replace Magic Numbers with Named Constants
G26: Be Precise
G27: Structure over Convention > naming conventions are good, but they are inferior to structures that force compliance. For example, switch/cases with nicely named enumerations are inferior to base classes with abstract methods.
G28: Encapsulate Conditionals
G29: Avoid Negative Conditionals
G30: Functions Should Do One Thing
G31: Hidden Temporal Couplings
G32: Don't Be Arbitrary
G33: Encapsulate Boundary Conditions
G34: Functions Should Descend Only One Level of Abstraction
G35: Keep Configurable Data at High Levels
G36: Avoid Transitive Navigation

Java

P. 338

J1: Avoid Long Import Lists By Using Wildcards
J2: Don't Inherit Constants
J3: Constants versus Enums

N1: Choose Descriptive Names
N2: Choose Names at the Appropriate Level of Abstraction
N3: Use Standard Nomenclature Where Possible
N4: Unambiguous Names
N5: Use Long Names For Long Scopes (i it's ok in a small for)
N6: Avoid Encodings
N7: Names Should Describe Side-Effects

Tests

T1: Insufficient Tests (= test everythings)
T2: Use a Coverage Tool!
T3: Don't Skip Trivial Tests
T4: An Ignored Test Is a Question about an Ambiguity
T5: Test Boundary Conditions
T6: Exhaustively Test Near Bugs
T7: Patterns of Failure Are Revealing
T8: Test Coverage Patterns Can Be Revealing
T9: Tests Should Be Fast

Appendix A! Concurrency II

Client/Server Example
Possible Paths of Execution

P. 352



P. 343



Further Reading :
---

- Literate Programming by Knuth

