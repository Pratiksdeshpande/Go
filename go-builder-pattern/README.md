# 🏗️ Go Builder Pattern

## 📑 Table of Contents

- [🤔 What is the Builder Pattern?](#-what-is-the-builder-pattern)
- [✅ When to Use the Builder Pattern?](#-when-to-use-the-builder-pattern)
- [❌ When NOT to Use the Builder Pattern?](#-when-not-to-use-the-builder-pattern)
- [🔧 How does the Builder pattern solve the telescoping constructor problem?](#-how-does-the-builder-pattern-solve-the-telescoping-constructor-problem)
- [🚀 Builder vs Functional Options](#-in-go-we-often-use-functional-options-for-optional-parameters-why-might-you-choose-builder-instead)
- [🎭 Role of the Director in the Builder pattern?](#-role-of-the-director-in-the-builder-pattern)
- [🔒 How to make a builder thread-safe in Go? Should you?](#-how-to-make-a-builder-thread-safe-in-go-should-you)
- [💻 Implementation Examples](#-implementation-examples)
- [🚀 Quick Start](#-quick-start)
- [📚 Further Reading](#-further-reading)

## 🤔 What is the Builder Pattern?

The Builder pattern separates the construction of a complex object from its representation so you can build different representations using the same construction process.

### 🧩 Structure (Conceptual Components)

- **Product** — the final object (e.g., Pizza)
- **Builder** — interface or type that exposes methods to set parts/attributes
- **Concrete Builder** — concrete implementation of builder methods; stores interim state
- **Director (optional)** — orchestrates a particular sequence of steps to create a preset configuration (e.g., BuildMargherita())

*Note: You don't always need a Director — in Go it's rare unless you have many preset recipes.*

## ✅ When to Use the Builder Pattern?

Use Builder when one or more of the following are true:

- The object has many optional parameters and plain constructors cause confusion or long parameter lists
- Construction is multi-step or requires validation across fields (e.g., field A must exist when B is set)
- You want a fluent, readable API: `NewBuilder().SetA(x).SetB(y).Build()`
- You need staged construction (enforce ordering and required steps at compile time)
- You build nested complex objects (e.g., House that contains Rooms, each with its own builder)
- You want to reuse building logic (Directors) to produce well-known object variants

## ❌ When NOT to Use the Builder Pattern?

- If you only have a couple of optional parameters, functional options are usually simpler and more idiomatic in Go
- If the object is small and construction is trivial — Builder adds unnecessary complexity
- If you only need immutability/defaults and no multi-step construction — use simple constructors or functional options

## 🔧 How does the Builder pattern solve the telescoping constructor problem?

In many languages (Java, C++), if a class has many optional parameters, you either:
1. Create multiple constructors with different parameter combinations → explosion of constructors
2. Or make a single constructor with many parameters (booleans, strings, ints, etc.) → hard to read, error-prone

**🛠️ How Builder fixes it:**
The Builder pattern replaces this messy constructor with self-documenting, chainable methods.

The Builder pattern avoids long, ambiguous constructors by:
- Making required parameters explicit (e.g., Size, Crust)
- Turning optional parameters into fluent methods
- Using Build() to validate before creating the final object

## 🚀 In Go, we often use functional options for optional parameters. Why might you choose Builder instead?

Functional options is a common Go idiom. Instead of a huge constructor, you define option functions that modify the object during construction.

**📊 Example:**
```go
type Pizza struct {
    Size      string
    Crust     string
    Cheese    bool
    Pepperoni bool
    Mushrooms bool
}

type Option func(*Pizza)

func WithCheese() Option {
    return func(p *Pizza) { p.Cheese = true }
}

func NewPizza(size, crust string, opts ...Option) *Pizza {
    p := &Pizza{Size: size, Crust: crust}
    for _, opt := range opts {
        opt(p)
    }
    return p
}

// Usage:
p := NewPizza("Large", "Thin", WithCheese())
```

**⚖️ Builder vs Functional Options:**

Both solve the telescoping constructor problem, but they shine in different contexts:

**✅ When Builder is better:**
- **🔄 Multi-step construction** - If building the object requires multiple phases or dependent fields. Example: Pizza builder validating that "Stuffed crust" can't be used with "Small" size
- **🛡️ Staged construction / compile-time enforcement** - Builders can enforce mandatory steps (size, crust) before optional toppings. Functional options don't enforce order — you could accidentally skip required fields
- **🏢 Complex, nested objects** - When building objects with sub-objects or hierarchies (e.g., HouseBuilder → RoomBuilder → WindowBuilder)
- **📖 Readability for very large objects** - Chainable builder calls are more expressive than dozens of functional options
- **✔️ Validation at the end** - With Builder, validation can happen in Build() where you check all constraints. Functional options apply immediately, making cross-field validation trickier

**✅ When Functional Options is better:**
- **⚡ Simplicity** — The idiom is common, concise, and very Go-like
- **🔢 Few optional fields** — If you only have 2–3 optional fields, options are lightweight
- **🔄 Flexibility** — You can easily pass options around (they're just functions)
- **🔄 Stateless** — Options don't require holding intermediate state like a builder does
- **Flexibility** — You can easily pass options around (they're just functions)
- **Stateless** — Options don't require holding intermediate state like a builder does

## 🎭 Role of the Director in the Builder pattern?

A Director encapsulates and orchestrates the sequence of builder calls to construct specific, pre-defined variants (recipes) of the product — *it tells the builder how to assemble a particular configuration so callers don't repeat the assembly steps*.

**🎯 Responsibilities of the Director:**

1. **🔄 Orchestrate steps** - Call builder methods in a specific order to produce a known configuration (e.g., BuildMargherita, BuildPepperoni)
2. **♻️ Reuse construction logic** - When many callers need the same preconfigured product, Director centralizes the recipe
3. **🎯 Keep clients simple** - Clients either use the Director for common presets or the builder directly for custom constructions
4. **🔌 Decouple recipe from product** - Director knows what combination of setter calls produces a variant; the builder knows how to apply each step

**🤷 Do you always need a Director in Go?**

No — you don't always need it. In Go, Director is optional and often unnecessary. Use it only when it adds clear value.

**❌ When you likely don't need a Director:**
- The API is small and readable: direct chaining is clear and concise
- Clients require high customization: they prefer calling builder methods themselves
- You already provide simple factory/helper functions for common variants

**✅ When a Director makes sense:**
- You have many callers that need the same complex recipes (avoid duplicated builder sequences)
- Recipes are complex (many steps or conditional logic) and benefit from centralization
- You want to expose a small surface API for common variants while still letting power users use the builder directly

## 🔒 How to make a builder thread-safe in Go? Should you?

**⚠️ Problem:** Builders hold mutable state (pointer receiver). If multiple goroutines share the same builder instance, data races can occur. This can lead to inconsistent or corrupted internal state.

**🛡️ Ways to make a builder thread-safe:**

**Option A — 🔐 Use a mutex in the builder**
- **✅ Pros:** Safe for concurrent modification
- **❌ Cons:** Adds complexity, small performance overhead, rarely needed

**Option B — 🔄 Make the builder immutable**
Each method returns a new copy of the builder instead of modifying the same instance.
- **✅ Pros:** No mutex needed, safer for concurrency
- **❌ Cons:** Can be memory-intensive for large objects, less common in Go

**Option C — 🎯 Avoid sharing builders** *(Most idiomatic Go solution)*
- Each goroutine uses its own builder instance
- Builders are cheap and small structs — creating one per goroutine is simple and safe
- Only share the final product after it's built

## 💻 Implementation Examples

This repository includes two different Builder pattern implementations:

### 🔗 Simple Fluent Builder (`simple_fluent_builder_pattern.go`)
- **Method chaining** for readable code
- **Runtime validation** to ensure mandatory fields are set
- **Flexible order** of method calls
- **Error handling** for invalid states
- **Director pattern** for common configurations

### 🏗️ Staged Builder (`staged_builder_pattern.go`)
- **Type-safe construction** through different interfaces at each stage
- **Compile-time guarantees** that mandatory fields are set in correct order
- **Prevents invalid intermediate states**
- **Progressive interface exposure** as you complete each stage

## 🚀 Quick Start

Ready to see the Builder patterns in action? Run these examples to explore both implementations:

```bash
# Run the Simple Fluent Builder example
# This demonstrates flexible method chaining with runtime validation
go run simple_fluent_builder_pattern.go

# Run the Staged Builder example  
# This shows compile-time enforcement of construction steps
go run staged_builder_pattern.go
```

**What you'll see:**
- 🍕 **Pizza Builder**: Fluent API for building pizzas with various toppings and validation
- 🚗 **Car Builder**: Staged construction ensuring mandatory fields (make, color) before optional features
- 🎭 **Director Pattern**: Pre-configured recipes for common object variants
- ✅ **Validation Examples**: How builders handle invalid states and missing required fields

**Try customizing the examples:**
- Modify the pizza toppings in `simple_fluent_builder_pattern.go`
- Add new car features in `staged_builder_pattern.go`
- Create your own Director recipes for common configurations

## 📚 Further Reading

### 🏗️ Design Patterns & Architecture
- **[Gang of Four - Design Patterns](https://en.wikipedia.org/wiki/Design_Patterns)** - Original Builder pattern specification
- **[Martin Fowler - FluentInterface](https://martinfowler.com/bliki/FluentInterface.html)** - Understanding fluent APIs and method chaining
- **[Clean Code - Object Creation](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship/dp/0132350884)** - Robert Martin's principles for clean object construction

### 🚀 Advanced Go Patterns
- **[Go Patterns Repository](https://github.com/tmrts/go-patterns)** - Comprehensive collection of Go design patterns
- **[Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)** - Production Go best practices
- **[Go Proverbs](https://go-proverbs.github.io/)** - Rob Pike's wisdom on Go philosophy

### 🔧 Practical Applications
- **[Kubernetes Client-Go](https://github.com/kubernetes/client-go)** - Real-world builder patterns in Kubernetes API
- **[Go Kit](https://gokit.io/)** - Microservice toolkit showcasing advanced Go patterns
- **[Testify](https://github.com/stretchr/testify)** - Popular testing library using builder-like patterns

### 📺 Videos & Talks
- **[GopherCon Talks](https://www.youtube.com/c/GopherCon)** - Annual conference with Go pattern discussions
- **[Effective Go Talks](https://talks.golang.org/)** - Official Go team presentations
- **[Go Time Podcast](https://changelog.com/gotime)** - Regular discussions on Go patterns and best practices
