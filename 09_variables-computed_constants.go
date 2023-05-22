package main

// Constants must be known at compile time. Generaly it will be declared with a static value.
const myInt = 15

// Constants can be computed SO LONG AS the computation can happen at COMPILE TIME.
// That said, there is no way to declare a constant that can only be computed at runtime.
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName
