// GO is much faster than JS, Python, Ruby and PHP.
// This happens because GO is a compiled language
// Every interpreted programming language is slower, by it's nature.
//---------------------//

// Welcome to "Learn GO" from Boot.Dev
// For the next lessons, we'll write a "Textio" Web Server

// Since this program builds into an executable, we start using
package main

// Format function import
import "fmt"

// Every GO Program, starts with "main" function.
// It takes no inputs and doesn't return anything.

func main() {
	fmt.Println("Starting Textio Server") // "Print Line" function
}

// Compilation is the act of transforming human code writing in binaries or "Machine Code"

// GO enforces strong and static typing, meaning variables can only have a single type.
// One of the biggest benefits of strong typing is that errors can be caught at "compile time".
// Making it easier to catch bugs ahead of time, it's detected before the code even run.
