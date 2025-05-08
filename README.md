# gorevision
Revising go concepts


# 🔰 **Data Structures – Broad Categories**

Data structures are mainly divided into **two types**:

---

### 📦 1. **Primitive Data Structures** (Basic building blocks)

These are **directly supported by the language**.

| Type      | Examples             |
| --------- | -------------------- |
| Integer   | `int`, `int64`       |
| Float     | `float32`, `float64` |
| Character | `rune`, `byte`       |
| Boolean   | `true`, `false`      |
| String    | `"hello"`            |

---

### 🏗️ 2. **Non-Primitive Data Structures**

These are **more complex** and built using primitive types.

They are further divided into:

---

## ✅ **I. Linear Data Structures**

Elements are stored **sequentially** (one after the other).

| DS Type           | Description                               | Example Use                    |
| ----------------- | ----------------------------------------- | ------------------------------ |
| **Array** / Slice | Fixed-size (Array), Dynamic (Slice in Go) | Store elements at fixed index  |
| **Linked List**   | Nodes connected using pointers            | Dynamic memory usage           |
| **Stack**         | LIFO – Last In First Out                  | Undo operation, function calls |
| **Queue**         | FIFO – First In First Out                 | Task scheduling, printer queue |

---

## 🌳 **II. Non-Linear Data Structures**

Data is not stored sequentially. Useful for complex relationships.

| DS Type   | Description                              | Example Use                      |
| --------- | ---------------------------------------- | -------------------------------- |
| **Tree**  | Hierarchical structure, nodes & children | File systems, XML parsing        |
| **Graph** | Nodes (vertices) connected with edges    | Social networks, Maps            |
| **Heap**  | Complete binary tree (Min/Max priority)  | Priority Queue, Heap Sort        |
| **Trie**  | Tree for strings/prefix matching         | Auto-complete, Dictionary search |

---

## 🎯 **III. Hash-Based Data Structures**

Quick access using keys.

| DS Type              | Description        | Example Use                        |
| -------------------- | ------------------ | ---------------------------------- |
| **Hash Table / Map** | Key-value storage  | Dictionary, Caching                |
| **Set**              | Unique values only | Remove duplicates, Membership test |

---

## 🧭 Summary Mindmap Style

```
Data Structures
│
├── Primitive: int, float, bool, char, string
│
└── Non-Primitive
    ├── Linear
    │   ├── Array / Slice
    │   ├── Linked List
    │   ├── Stack
    │   └── Queue
    │
    ├── Non-Linear
    │   ├── Tree
    │   ├── Graph
    │   ├── Heap
    │   └── Trie
    │
    └── Hash-Based
        ├── Map
        └── Set
```

---

# Searching

### 🔍 1. **Linear Search (Sequential Search)**

* **Concept:** Check each element one by one.
* **Use case:** Works on **unsorted** or **sorted** data.
* **Time Complexity:** `O(n)`

#### ✅ Example:

```go
func linearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}
```

---

### 🔎 2. **Binary Search (Divide and Conquer)**

* **Concept:** Repeatedly divide the **sorted array** and eliminate half.
* **Use case:** Requires **sorted array**.
* **Time Complexity:** `O(log n)`

#### ✅ Example:

```go
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
```

---

### 🎯 Bonus Category (Advanced Searching Types)

If we go a bit deeper, we can also include:

| Type                 | Example                      | Use case                         |
| -------------------- | ---------------------------- | -------------------------------- |
| Hash-based search    | Hash tables, maps            | Constant time lookups            |
| Search Trees         | Binary Search Tree (BST)     | Hierarchical data                |
| Trie Search          | Word prefixes, auto-complete | Fast string matching             |
| Interpolation Search | Uniformly distributed data   | Faster than binary in some cases |
| Exponential Search   | Very large sorted arrays     | Used with binary search          |

---

### 📌 Summary

| Type          | Works On | Time     | Notes                  |
| ------------- | -------- | -------- | ---------------------- |
| Linear Search | Any list | O(n)     | Simple, slow for large |
| Binary Search | Sorted   | O(log n) | Fast, needs sorting    |

---


# Sorting

## 🧮 All Major Sorting Algorithms (Grouped by Category)

### 🔹 1. **Simple / Comparison-based Sorts**

| Algorithm          | Best Case | Worst Case | Space | Stable |
| ------------------ | --------- | ---------- | ----- | ------ |
| **Bubble Sort**    | O(n)      | O(n²)      | O(1)  | ✅ Yes  |
| **Selection Sort** | O(n²)     | O(n²)      | O(1)  | ❌ No   |
| **Insertion Sort** | O(n)      | O(n²)      | O(1)  | ✅ Yes  |

📌 *Used for small datasets or educational purposes.*

---

### 🔹 2. **Efficient Comparison-based Sorts**

| Algorithm      | Best Case  | Worst Case | Space      | Stable | Type             |
| -------------- | ---------- | ---------- | ---------- | ------ | ---------------- |
| **Merge Sort** | O(n log n) | O(n log n) | O(n)       | ✅ Yes  | Divide & Conquer |
| **Quick Sort** | O(n log n) | O(n²)      | O(log n)\* | ❌ No   | Divide & Conquer |
| **Heap Sort**  | O(n log n) | O(n log n) | O(1)       | ❌ No   | Heap-based       |

📌 *Used in real-world applications, built-in libraries (like Go’s sort package uses QuickSort/HeapSort/MergeSort under the hood depending on data).*

---

### 🔹 3. **Non-Comparison-based Sorts (for integers / fixed-length data)**

| Algorithm         | Time     | Space    | Stable | Use-case                            |
| ----------------- | -------- | -------- | ------ | ----------------------------------- |
| **Counting Sort** | O(n + k) | O(k)     | ✅ Yes  | When range of elements is small     |
| **Radix Sort**    | O(nk)    | O(n + k) | ✅ Yes  | Works well with integers/strings    |
| **Bucket Sort**   | O(n + k) | O(n)     | ✅ Yes  | Uniformly distributed float/decimal |

📌 *These are super-fast but work only for specific kinds of data (e.g., integers, known ranges).*

---

## ✅ Which Ones Should You Focus On First?

| Level            | Sorts to Learn               |
| ---------------- | ---------------------------- |
| **Beginner**     | Bubble, Selection, Insertion |
| **Intermediate** | Merge, Quick, Heap           |
| **Advanced**     | Counting, Radix, Bucket      |

---

## 🚀 In Go (Golang), Built-in Sorting

You can use:

```go
import "sort"

arr := []int{5, 2, 9}
sort.Ints(arr) // Uses an optimized sorting algo
```
