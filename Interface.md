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
```go
// go-interface-example/database/database.go
package database

// Database interface defines the contract (POLYMORPHISM)
// This is similar to abstract classes in OOPs (Java/C++)
type Database interface {
	Close()
	SelectUser(userID int)
	CreateUser(name string)
}

// go-interface-example/database/mysql.go
package database

import "fmt"

type MySQL struct{}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (m *MySQL) Close() {
	fmt.Println("MySQL: Connection closed")
}

func (m *MySQL) SelectUser(userID int) {
	fmt.Printf("MySQL: Selected user with ID %d\n", userID)
}

func (m *MySQL) CreateUser(name string) {
	fmt.Printf("MySQL: Created user %s\n", name)
}

// go-interface-example/database/postgres.go
package database

import "fmt"

type Postgres struct{}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Close() {
	fmt.Println("Postgres: Connection closed")
}

func (p *Postgres) SelectUser(userID int) {
	fmt.Printf("Postgres: Selected user with ID %d\n", userID)
}

func (p *Postgres) CreateUser(name string) {
	fmt.Printf("Postgres: Created user %s\n", name)
}

// go-interface-example/database/mongo.go
package database

import "fmt"

type Mongo struct{}

func NewMongo() *Mongo {
	return &Mongo{}
}

func (m *Mongo) Close() {
	fmt.Println("MongoDB: Connection closed")
}

func (m *Mongo) SelectUser(userID int) {
	fmt.Printf("MongoDB: Selected user with ID %d\n", userID)
}

func (m *Mongo) CreateUser(name string) {
	fmt.Printf("MongoDB: Created user %s\n", name)
}

// go-interface-example/main.go
package main

import (
	"fmt"
	"golang_interface_example/database"
)

// MyApp uses a Database interface, not concrete type
// This supports dependency injection, loose coupling and easy testing
type MyApp struct {
	db database.Database
}

// NewApp initializes the application with any database that satisfies Database interface
func NewApp(db database.Database) *MyApp {
	return &MyApp{db: db}
}

// Run method demonstrates usage of the interface methods
func (a *MyApp) Run() {
	a.db.CreateUser("Adarsh")
	a.db.SelectUser(101)
	a.db.Close()
}

func main() {
	// Swap different DBs easily without changing MyApp
	fmt.Println("--- Running with MongoDB ---")
	db := database.NewMongo()
	// You can change above to NewMySQL() or NewPostgres() without any code change in MyApp

	app := NewApp(db)
	app.Run()
}

```