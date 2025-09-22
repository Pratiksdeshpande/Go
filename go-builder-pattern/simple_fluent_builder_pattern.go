// Simple Fluent Builder Pattern Implementation in Go
//
// The Fluent Builder pattern is a variation of the Builder pattern that allows
// method chaining in any order to construct complex objects step by step.
// It's particularly useful when you need to create objects with many optional
// parameters or configurations.
//
// Key characteristics of this implementation:
// • Method chaining (fluent interface) for readable code
// • Runtime validation to ensure mandatory fields are set
// • Flexible order of method calls
// • Error handling for invalid states
// • Director pattern for common configurations

package main

import (
	"errors"
	"fmt"
)

func main() {
	demonstrateFluentBuilder()
}

// Pizza represents the complex object we want to build
// It contains various properties that can be set independently
type Pizza struct {
	Size      string // Size of the pizza (e.g., "Small", "Medium", "Large")
	Crust     string // Type of crust (e.g., "Thin", "Thick", "Stuffed")
	Cheese    bool   // Whether cheese is added
	Pepperoni bool   // Whether pepperoni is added
	Mushrooms bool   // Whether mushrooms are added
}

// PizzaBuilder defines the interface for building pizza objects
// Each method returns the builder itself to enable method chaining (fluent interface)
// This allows for readable and flexible object construction
type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder   // Sets the size of the pizza
	SetCrust(crust string) PizzaBuilder // Sets the crust type
	AddCheese() PizzaBuilder            // Adds cheese to the pizza
	AddPepperoni() PizzaBuilder         // Adds pepperoni to the pizza
	AddMushrooms() PizzaBuilder         // Adds mushrooms to the pizza
	Build() (Pizza, error)              // Finalizes and returns the constructed pizza with validation
}

// ConcretePizzaBuilder is the concrete implementation of the PizzaBuilder interface
// It maintains the state of the pizza being built and provides methods to configure it
type ConcretePizzaBuilder struct {
	pizza Pizza // The pizza object being constructed
}

// SetSize sets the size of the pizza and returns the builder for method chaining
func (p *ConcretePizzaBuilder) SetSize(size string) PizzaBuilder {
	p.pizza.Size = size
	return p
}

// SetCrust sets the crust type of the pizza and returns the builder for method chaining
func (p *ConcretePizzaBuilder) SetCrust(crust string) PizzaBuilder {
	p.pizza.Crust = crust
	return p
}

// AddCheese adds cheese to the pizza and returns the builder for method chaining
func (p *ConcretePizzaBuilder) AddCheese() PizzaBuilder {
	p.pizza.Cheese = true
	return p
}

// AddPepperoni adds pepperoni to the pizza and returns the builder for method chaining
func (p *ConcretePizzaBuilder) AddPepperoni() PizzaBuilder {
	p.pizza.Pepperoni = true
	return p
}

// AddMushrooms adds mushrooms to the pizza and returns the builder for method chaining
func (p *ConcretePizzaBuilder) AddMushrooms() PizzaBuilder {
	p.pizza.Mushrooms = true
	return p
}

// Build finalizes the construction and returns the completed pizza object
// Validates that mandatory fields (Size and Crust) are set before building
func (p *ConcretePizzaBuilder) Build() (Pizza, error) {
	// Validate mandatory field: Size
	if p.pizza.Size == "" {
		return Pizza{}, errors.New("pizza size is mandatory and cannot be empty")
	}

	// Validate mandatory field: Crust
	if p.pizza.Crust == "" {
		return Pizza{}, errors.New("pizza crust is mandatory and cannot be empty")
	}

	return p.pizza, nil
}

// PizzaDirector provides a high-level interface for constructing specific types of pizzas
// It encapsulates the logic for creating common pizza configurations
// This is optional in the Builder pattern but helps create predefined objects easily
type PizzaDirector struct{}

// CreateMargheritaPizza creates a classic Margherita pizza using the provided builder
// Margherita pizza: Large size, thin crust, with cheese
func (d *PizzaDirector) CreateMargheritaPizza(pizzaBuilder PizzaBuilder) (Pizza, error) {
	return pizzaBuilder.SetSize("Large").SetCrust("Thin").AddCheese().Build()
}

// CreateMushroomPizza creates a mushroom pizza using the provided builder
// Mushroom pizza: Large size, thin crust, with mushrooms
func (d *PizzaDirector) CreateMushroomPizza(pizzaBuilder PizzaBuilder) (Pizza, error) {
	return pizzaBuilder.SetSize("Large").SetCrust("Thin").AddMushrooms().Build()
}

// demonstrateFluentBuilder demonstrates the simple fluent builder pattern
func demonstrateFluentBuilder() {
	fmt.Println("=== SIMPLE FLUENT BUILDER PATTERN DEMONSTRATION ===")
	fmt.Println()

	// Create instances of the builder and director
	builder := &ConcretePizzaBuilder{}
	director := &PizzaDirector{}

	// Example 1: Using the Director to create predefined pizzas
	// The director encapsulates common pizza configurations
	fmt.Println("=== Predefined Pizzas (using Director) ===")

	margherita, err := director.CreateMargheritaPizza(builder)
	if err != nil {
		fmt.Printf("Error creating Margherita pizza: %v\n", err)
	} else {
		fmt.Printf("Margherita Pizza: Size=%s, Crust=%s, Cheese=%t, Pepperoni=%t, Mushrooms=%t\n",
			margherita.Size, margherita.Crust, margherita.Cheese, margherita.Pepperoni, margherita.Mushrooms)
	}

	mushroom, err := director.CreateMushroomPizza(builder)
	if err != nil {
		fmt.Printf("Error creating Mushroom pizza: %v\n", err)
	} else {
		fmt.Printf("Mushroom Pizza: Size=%s, Crust=%s, Cheese=%t, Pepperoni=%t, Mushrooms=%t\n",
			mushroom.Size, mushroom.Crust, mushroom.Cheese, mushroom.Pepperoni, mushroom.Mushrooms)
	}

	fmt.Println("\n=== Custom Pizza (using Builder directly) ===")

	// Example 2: Using the Builder directly for custom configurations
	// This demonstrates the flexibility of the Builder pattern
	// Method chaining (fluent interface) makes the code readable
	customPizza, err := builder.SetSize("Regular").SetCrust("Thick").AddCheese().AddPepperoni().AddMushrooms().Build()
	if err != nil {
		fmt.Printf("Error creating Custom pizza: %v\n", err)
	} else {
		fmt.Printf("Custom Pizza: Size=%s, Crust=%s, Cheese=%t, Pepperoni=%t, Mushrooms=%t\n",
			customPizza.Size, customPizza.Crust, customPizza.Cheese, customPizza.Pepperoni, customPizza.Mushrooms)
	}

	fmt.Println("\n=== Validation Examples ===")

	// Example 3: Demonstrate validation - missing size
	invalidBuilder1 := &ConcretePizzaBuilder{}
	_, err = invalidBuilder1.SetCrust("Thin").AddCheese().Build()
	if err != nil {
		fmt.Printf("Validation error (missing size): %v\n", err)
	}

	// Example 4: Demonstrate validation - missing crust
	invalidBuilder2 := &ConcretePizzaBuilder{}
	_, err = invalidBuilder2.SetSize("Large").AddCheese().Build()
	if err != nil {
		fmt.Printf("Validation error (missing crust): %v\n", err)
	}
}
