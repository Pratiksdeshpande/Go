package main

import "fmt"

// ============================================================================
// STAGED BUILDER PATTERN IMPLEMENTATION
// ============================================================================
// The Staged Builder pattern is a type-safe variation of the Builder pattern that
// enforces a specific sequence of operations through different interfaces at each stage.
// This provides compile-time guarantees that mandatory fields are set in the correct
// order and prevents the creation of invalid intermediate states.
// ============================================================================

func main() {
	demonstrateStagedBuilder()
}

// Car represents the complex product being built using the staged builder pattern
// This struct contains both mandatory fields (Make, Color) and optional features (HasGPS, IsElectric)
// The staged builder ensures mandatory fields are set before optional ones
type Car struct {
	Make       string // Mandatory: Car manufacturer (e.g., "Toyota", "Tesla", "Ferrari")
	Color      string // Mandatory: Car color (e.g., "Red", "Blue", "Yellow")
	HasGPS     bool   // Optional: Whether the car has GPS navigation system
	IsElectric bool   // Optional: Whether the car is electric powered
}

// MakeStage Stage 1: First mandatory step to set the car make
// This interface only allows setting the make and moving to the next stage
type MakeStage interface {
	SetMake(make string) ColorStage // Must set make first, returns next stage
}

// ColorStage Stage 2: Second mandatory step to set the car color
// This interface only allows setting the color and moving to the next stage
type ColorStage interface {
	SetColor(color string) OptionalStage // Must set color second, returns next stage
}

// OptionalStage Stage 3: Final stage for optional features and building
// This interface allows setting optional features and building the final car
type OptionalStage interface {
	WithGPS() OptionalStage      // Optional: Add GPS feature
	MakeElectric() OptionalStage // Optional: Make the car electric
	Build() Car                  // Build and return the final car object
}

// CarBuilder implements all stages of the staged builder pattern
// It maintains the car state and implements different interfaces for each stage
type CarBuilder struct {
	car Car // The car object being constructed through stages
}

// NewCarBuilder creates a new car builder and returns the first stage (MakeStage)
// This is the entry point for the staged builder pattern
func NewCarBuilder() MakeStage {
	return &CarBuilder{
		car: Car{}, // Initialize with empty car
	}
}

// SetMake : Stage 1 Implementation
// Sets the car make (mandatory field) and progresses to ColorStage
func (cb *CarBuilder) SetMake(make string) ColorStage {
	cb.car.Make = make
	return cb // Return self but typed as ColorStage interface
}

// SetColor : Stage 2 Implementation
// Sets the car color (mandatory field) and progresses to OptionalStage
func (cb *CarBuilder) SetColor(color string) OptionalStage {
	cb.car.Color = color
	return cb // Return self but typed as OptionalStage interface
}

// WithGPS : Stage 3 Implementation
// Adds GPS feature (optional) and remains in OptionalStage for method chaining
func (cb *CarBuilder) WithGPS() OptionalStage {
	cb.car.HasGPS = true
	return cb // Return self to allow method chaining of optional features
}

// MakeElectric : Stage 3 Implementation
// Makes the car electric (optional) and remains in OptionalStage for method chaining
func (cb *CarBuilder) MakeElectric() OptionalStage {
	cb.car.IsElectric = true
	return cb // Return self to allow method chaining of optional features
}

// Build : Stage 3 Implementation
// Finalizes construction and returns the completed car
// No validation needed here since mandatory fields are enforced by the staged interfaces
func (cb *CarBuilder) Build() Car {
	return cb.car
}

// Usage Examples:
//
// Basic car (mandatory fields only):
//   basicCar := NewCarBuilder().SetMake("Toyota").SetColor("Blue").Build()
//
// Luxury car (with all features):
//   luxuryCar := NewCarBuilder().SetMake("Tesla").SetColor("Red").WithGPS().MakeElectric().Build()
//
// Custom car (flexible optional features):
//   customCar := NewCarBuilder().SetMake("Ferrari").SetColor("Yellow").MakeElectric().Build()
//
// Compile-time safety examples (these would cause compile errors):
//   NewCarBuilder().SetColor("Red")           // Error: SetColor not available on MakeStage
//   NewCarBuilder().SetMake("Toyota").Build() // Error: Build not available on ColorStage
//   NewCarBuilder().WithGPS()                 // Error: WithGPS not available on MakeStage

// demonstrateStagedBuilder demonstrates the staged builder pattern with comprehensive examples
func demonstrateStagedBuilder() {
	fmt.Println("=== STAGED BUILDER PATTERN DEMONSTRATION ===")
	fmt.Println()

	// Example 1: Basic car with only mandatory fields
	// The staged builder enforces the order: Make → Color → Build
	fmt.Println("=== Basic Car (Mandatory fields only) ===")
	basicCar := NewCarBuilder().
		SetMake("Toyota"). // Stage 1: Must set make first
		SetColor("Blue").  // Stage 2: Must set color second
		Build()            // Stage 3: Build the car

	fmt.Printf("Basic Car: Make=%s, Color=%s, GPS=%t, Electric=%t\n",
		basicCar.Make, basicCar.Color, basicCar.HasGPS, basicCar.IsElectric)

	// Example 2: Luxury car with all optional features
	// Demonstrates method chaining in the optional stage
	fmt.Println("\n=== Luxury Car (With optional features) ===")
	luxuryCar := NewCarBuilder().
		SetMake("Tesla"). // Stage 1: Set make
		SetColor("Red").  // Stage 2: Set color
		WithGPS().        // Stage 3: Add optional GPS
		MakeElectric().   // Stage 3: Add optional electric feature
		Build()           // Stage 3: Build the final car

	fmt.Printf("Luxury Car: Make=%s, Color=%s, GPS=%t, Electric=%t\n",
		luxuryCar.Make, luxuryCar.Color, luxuryCar.HasGPS, luxuryCar.IsElectric)

	// Example 3: Different order of optional features
	// Shows flexibility in the optional stage while maintaining mandatory order
	fmt.Println("\n=== Sports Car (Different optional order) ===")
	sportsCar := NewCarBuilder().
		SetMake("Ferrari"). // Stage 1: Set make
		SetColor("Yellow"). // Stage 2: Set color
		MakeElectric().     // Stage 3: Make electric first
		Build()             // Stage 3: Build without GPS

	fmt.Printf("Sports Car: Make=%s, Color=%s, GPS=%t, Electric=%t\n",
		sportsCar.Make, sportsCar.Color, sportsCar.HasGPS, sportsCar.IsElectric)

	// Example 4: Economy car with only GPS
	fmt.Println("\n=== Economy Car (Single optional feature) ===")
	economyCar := NewCarBuilder().
		SetMake("Honda").  // Stage 1: Set make
		SetColor("White"). // Stage 2: Set color
		WithGPS().         // Stage 3: Add only GPS
		Build()            // Stage 3: Build the car

	fmt.Printf("Economy Car: Make=%s, Color=%s, GPS=%t, Electric=%t\n",
		economyCar.Make, economyCar.Color, economyCar.HasGPS, economyCar.IsElectric)
}
