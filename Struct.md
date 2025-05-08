You're moving like a true **Go warrior**! After maps, learning **structs** is the next solid step — it's the heart of writing real-world, modular Go programs (APIs, drivers, etc.).

Let’s go step-by-step and understand:

---

# 🧱 What is a `struct` in Golang?

A `struct` (short for “structure”) is a **composite data type** — it groups together variables (called **fields**) under one name.

> 📦 It’s like a box that holds different data types logically.

---

### ✅ Declaring a Struct

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

---

### ✅ Creating Struct Instances

```go
// 1. Named fields
p1 := Person{Name: "Adarsh", Age: 25, City: "Vasai"}

// 2. Positional (less readable)
p2 := Person{"Ram", 30, "Delhi"}

// 3. Zero value
var p3 Person
p3.Name = "Shyam"
```

---

### ✅ Accessing Fields

```go
fmt.Println(p1.Name)
p1.Age = 26
```

---

## ⚙️ Structs with Functions (Methods)

```go
func (p Person) greet() {
    fmt.Println("Hello,", p.Name)
}
```

> Add methods on structs like mini classes.

---

## 🧠 Pointers with Struct

To **modify struct fields** inside methods:

```go
func (p *Person) updateCity(city string) {
    p.City = city
}
```

Use pointer to **avoid copying** and allow mutation.

---

## 🔄 Anonymous Struct

```go
user := struct {
    Username string
    Active   bool
}{"adarsh", true}
```

Useful for temporary groupings.

---

## 🧱 Nested Structs

```go
type Address struct {
    City  string
    Pincode int
}

type Person struct {
    Name    string
    Address Address
}
```

Access like `person.Address.City`.

---

## ✨ Struct Tags (used in JSON, APIs)

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

Tags help map struct fields to JSON/XML/db fields.

---

## 🚀 Use Cases of Structs

| Use Case               | Example                       |
| ---------------------- | ----------------------------- |
| API Requests/Responses | JSON binding with `struct`    |
| Config & Settings      | `Config` struct for env vars  |
| Grouping data          | Group user/order/product info |
| Database mapping (ORM) | GORM/SQLX maps DB table rows  |
| Dependency injection   | Bundle services in a struct   |

---

## ❗ Things to Keep in Mind

| Concept             | Notes                                            |
| ------------------- | ------------------------------------------------ |
| Struct = value type | Copy happens by default                          |
| Use pointer `*`     | For performance & mutability                     |
| Export fields       | Capitalize fields to make them public/exportable |
| Tag rules           | Use backticks for JSON/XML (`json:"field"`)      |

---

## 🔁 Looping Through Slice of Structs

```go
people := []Person{
    {"A", 25, "City1"},
    {"B", 30, "City2"},
}

for _, person := range people {
    fmt.Println(person.Name)
}
```

---

## 🧪 Common Questions on Structs

1. Create and print student records
2. Compare two structs (deep equality)
3. Filter data from slice of structs
4. Implement basic methods like `greet()`
5. Parse JSON into struct (Go APIs)

---

## ✍️ Practice Task for You

Create a struct `Book` with fields:
`Title`, `Author`, `Price`, and method `discount(percent float64)` that returns price after discount.
