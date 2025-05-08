You’re on fire! After Struct now **interface** and **methods** — These 2 concepts are pillar of **Go's OOP (Object-Oriented Programming)**

# Let’s go **step-by-step** and cover:

* What is interface
* What is Method
* Relation b/w Interface + struct
* Real-life examples and interview-style questions

---

## 🤔 What is an Interface in Golang?

> Interface defines **a set of method signatures** (just like a contract).
>
> Jo bhi struct **un sab methods ko implement karta hai**, wo struct **interface ko satisfy karta hai**.

### 🔑 Syntax:

```go
type Speaker interface {
    Speak() string
}
```

---

## 🧱 What is a Method in Golang?

> Method is just a function, **attached to a struct**.

### ✅ Example:

```go
type Person struct {
    Name string
}

// Method (value receiver)
func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}
```

Now `Person` has a method called `Speak()`.

---

## 📌 How Interface & Method Work Together

```go
type Speaker interface {
    Speak() string
}

type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hi! I am " + p.Name
}

// Function accepting any Speaker
func SaySomething(s Speaker) {
    fmt.Println(s.Speak())
}
```

### ✅ Usage:

```go
p := Person{Name: "Adarsh"}
SaySomething(p) // Person implements Speaker
```

---

## 🧠 Key Concepts

| Concept                           | Description                                                       |
| --------------------------------- | ----------------------------------------------------------------- |
| **Interface = set of methods**    | Struct must implement all of them                                 |
| **Implicit implementation**       | No `implements` keyword — Go matches methods automatically        |
| **Empty interface `interface{}`** | Can accept any type (used like `any`)                             |
| **Polymorphism**                  | Different structs can behave differently using the same interface |

---

## ✨ Real World Example: Shape Interface

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func printArea(s Shape) {
    fmt.Println("Area is:", s.Area())
}
```

### ✅ Usage:

```go
c := Circle{Radius: 5}
r := Rectangle{Width: 3, Height: 4}

printArea(c)
printArea(r)
```

> 📦 Both `Circle` and `Rectangle` satisfy `Shape` interface, because they have `Area()` method.

---

## 🔁 Interface with Slice of Structs

```go
shapes := []Shape{
    Circle{Radius: 2},
    Rectangle{Width: 3, Height: 5},
}

for _, s := range shapes {
    printArea(s)
}
```

---

## 🔍 Empty Interface: `interface{}`

```go
func printAnything(val interface{}) {
    fmt.Println(val)
}

printAnything(42)
printAnything("hello")
```

But real use is with type assertions or reflection (advanced).

---

## 💡 Interview Questions on Interface

| Question                                             | Concept |
| ---------------------------------------------------- | ------- |
| Difference between value and pointer receiver method |         |
| How does Go support polymorphism                     |         |
| Can a struct implement multiple interfaces           |         |
| What happens if you partially implement interface    |         |
| Explain empty interface and its use case             |         |

---

## 🛠 Practice Task

1. Define an interface `Animal` with method `Speak() string`
2. Create two structs: `Dog`, `Cat`
3. Implement `Speak()` method differently in both
4. Pass them to a function that accepts `Animal`

---
