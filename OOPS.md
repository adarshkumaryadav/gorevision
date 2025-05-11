
# 🏗️ Core OOP Concepts with Go

# 🔷 Full Guide to OOPs Concepts in Golang with Code Examples

---

## 🔰 What is OOP (Object-Oriented Programming)?

**OOP** is a programming paradigm that organizes software design around **objects**, which can contain **data (fields)** and **behavior (methods)**. It promotes **modularity**, **reusability**, and **scalability**.

---

## 🎯 Goals of OOP:

* Build real-world models using code
* Achieve modular and maintainable code
* Promote reuse via composition or inheritance
* Improve testability and extensibility

---

## 🧱 Four Main Pillars of OOP:

| Pillar               | Meaning                                                            |
| -------------------- | ------------------------------------------------------------------ |
| 1. **Encapsulation** | Wrap data and methods together, and restrict direct access to data |
| 2. **Abstraction**   | Show only essential features and hide unnecessary details          |
| 3. **Inheritance**   | Reuse functionality from existing entities (via composition in Go) |
| 4. **Polymorphism**  | Use a common interface to perform different behaviors              |

---
Before the four pillars (Encapsulation, Abstraction, Inheritance, Polymorphism), there are **two base concepts**:

---

## 🔹 1. **Class**

**📘 Definition**:
A **class** is a blueprint or template to define the structure (data + behavior) of an object.
In **Go**, we use the `struct` keyword instead of `class`.

### ✅ Go Equivalent:

```go
type Car struct {
    Brand string
    Speed int
}
```

---

## 🔹 2. **Object**

**📘 Definition**:
An **object** is an instance of a class (struct in Go). It holds actual data and can call methods defined for that struct.

### ✅ Creating Object in Go:

```go
func main() {
    myCar := Car{Brand: "Toyota", Speed: 180}
    fmt.Println("Brand:", myCar.Brand)
    fmt.Println("Speed:", myCar.Speed)
}
```

---

## ✅ Complete Example: Class and Object in Go

```go
package main

import "fmt"

// Define a 'class' using struct
type Car struct {
    Brand string
    Speed int
}

// Method on Car struct
func (c Car) Drive() {
    fmt.Println(c.Brand, "is driving at", c.Speed, "km/h")
}

func main() {
    // Create object (instance of struct)
    myCar := Car{Brand: "Honda", Speed: 120}

    // Accessing fields
    fmt.Println("Car Brand:", myCar.Brand)

    // Calling method
    myCar.Drive()
}
```

---

## 🔁 Recap of All OOP Concepts in Golang

| 🔢 | OOP Concept       | Description                               | Go Equivalent                 | Example Provided    |
| -- | ----------------- | ----------------------------------------- | ----------------------------- | ------------------- |
| 1  | **Class**         | Blueprint for object with data + behavior | `struct`                      | ✅ `type Car struct` |
| 2  | **Object**        | Instance of a class with real data        | Variable of struct type       | ✅ `myCar := Car{}`  |
| 3  | **Encapsulation** | Hide internal data, expose via methods    | Private fields + setters      | ✅ `Account` struct  |
| 4  | **Abstraction**   | Expose essential behavior using interface | Interface                     | ✅ `Shape` interface |
| 5  | **Inheritance**   | Reuse behavior from base to child         | Struct embedding              | ✅ `Animal -> Dog`   |
| 6  | **Polymorphism**  | One interface, multiple behaviors         | Interface with multiple impls | ✅ `Speaker`         |

---

### 🧠 So, OOP in Go follows this flow:

```
Class (struct)
     ↓
Object (instance of struct)
     ↓
Encapsulation → Abstraction → Inheritance → Polymorphism
```


---


# ✅ Let's Learn with Complete Code in Go

---

## 🧱 1. **Encapsulation**

**Goal**: Protect internal state and expose only necessary methods to access it.

### 🔧 Concept:

* Fields with lowercase names = private (package-level access)
* Methods (getters/setters) provide controlled access

### ✅ Code:

```go
package main

import "fmt"

type Account struct {
    balance int // unexported field
}

// Set method
func (a *Account) Deposit(amount int) {
    if amount > 0 {
        a.balance += amount
    }
}

// Get method
func (a *Account) GetBalance() int {
    return a.balance
}

func main() {
    acc := Account{}
    acc.Deposit(1000)
    fmt.Println("Balance:", acc.GetBalance())
}
```

---

## 🧱 2. **Abstraction**

**Goal**: Hide internal logic and show only what’s important.

### 🔧 Concept:

* Use `interface` to expose only necessary behavior
* Client code doesn’t know how it works internally

### ✅ Code:

```go
package main

import "fmt"

// Abstract Shape
type Shape interface {
    Area() float64
}

// Circle
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Rectangle
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    printArea(Circle{Radius: 7})
    printArea(Rectangle{Width: 5, Height: 4})
}
```

---

## 🧱 3. **Inheritance** (via Composition in Go)

**Goal**: Reuse existing behavior in a new entity.

### 🔧 Concept:

* No classical inheritance (`extends`)
* Achieved by **embedding** a struct inside another struct

### ✅ Code:

```go
package main

import "fmt"

// Base struct
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println(a.Name, "makes a sound")
}

// Derived struct
type Dog struct {
    Animal
    Breed string
}

func main() {
    d := Dog{
        Animal: Animal{Name: "Bruno"},
        Breed:  "Labrador",
    }

    d.Speak() // Inherited behavior
    fmt.Println("Breed:", d.Breed)
}
```

---

## 🧱 4. **Polymorphism**

**Goal**: Let different objects respond in their own way through a common interface.

### 🔧 Concept:

* Define interface
* Implement it with multiple structs

### ✅ Code:

```go
package main

import "fmt"

// Interface
type Speaker interface {
    Speak()
}

// Dog implements Speaker
type Dog struct {
    Name string
}

func (d Dog) Speak() {
    fmt.Println(d.Name, "says Woof!")
}

// Cat implements Speaker
type Cat struct {
    Name string
}

func (c Cat) Speak() {
    fmt.Println(c.Name, "says Meow!")
}

func makeAnimalSpeak(s Speaker) {
    s.Speak()
}

func main() {
    dog := Dog{Name: "Tommy"}
    cat := Cat{Name: "Kitty"}

    makeAnimalSpeak(dog)
    makeAnimalSpeak(cat)
}
```

---

# 🔁 Summary of Concepts and Go Support

| OOP Concept       | Purpose                                    | Go Feature               | Example Provided    |
| ----------------- | ------------------------------------------ | ------------------------ | ------------------- |
| **Encapsulation** | Hide internal state, expose via methods    | Private fields + methods | ✅ Bank Account      |
| **Abstraction**   | Show only what is needed, hide logic       | Interfaces               | ✅ Shape Interface   |
| **Inheritance**   | Reuse behavior from base to derived (is-a) | Composition (embedding)  | ✅ Dog embeds Animal |
| **Polymorphism**  | Same interface, different behavior         | Interfaces + struct      | ✅ Speaker interface |

---

## 🧠 Additional OOP Concepts (Beyond 4 Pillars)

| Concept               | Description                                                       |
| --------------------- | ----------------------------------------------------------------- |
| **Class**             | Go struct acts like a class                                       |
| **Object**            | Instance of struct                                                |
| **Method Overriding** | Not directly supported — can override by implementing same method |
| **Constructor**       | Go has no `constructor` keyword — we write custom functions       |
| **Access Modifiers**  | Public/Private based on uppercase/lowercase names in Go           |

---

## 🧰 Bonus: Constructor in Go

```go
type Student struct {
    Name string
    Age  int
}

// Constructor function
func NewStudent(name string, age int) Student {
    return Student{Name: name, Age: age}
}

func main() {
    s := NewStudent("Adarsh", 24)
    fmt.Println(s)
}
```

---
