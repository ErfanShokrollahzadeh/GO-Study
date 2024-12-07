# GO-Study

## Introduction to GO Language

Go, also known as Golang, is an open-source programming language developed by Google. It is designed to be simple, efficient, and reliable, making it an excellent choice for building scalable and high-performance applications. Go is statically typed and compiled, which means that it catches errors at compile-time and produces fast executables.

## Installation Instructions

### Using Homebrew (macOS)

1. Open your terminal.
2. Install Homebrew if you haven't already by running the following command:
   ```sh
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   ```
3. Install Go using Homebrew:
   ```sh
   brew install go
   ```
4. Verify the installation:
   ```sh
   go version
   ```

### Using Terminal (Windows)

1. Download the Go installer from the official Go website: [https://golang.org/dl/](https://golang.org/dl/)
2. Run the installer and follow the on-screen instructions.
3. Open Command Prompt and verify the installation:
   ```sh
   go version
   ```

### Using Terminal (Linux)

1. Open your terminal.
2. Download the Go tarball from the official Go website:
   ```sh
   wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
   ```
3. Extract the tarball to `/usr/local`:
   ```sh
   sudo tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
   ```
4. Add Go to your PATH by adding the following line to your `~/.profile` or `~/.bashrc` file:
   ```sh
   export PATH=$PATH:/usr/local/go/bin
   ```
5. Apply the changes:
   ```sh
   source ~/.profile
   ```
6. Verify the installation:
   ```sh
   go version
   ```

## Step-by-Step Roadmap

1. **Introduction to Go**: Understand the basics of the Go programming language, including its syntax, data types, and control structures.
2. **Setting Up the Environment**: Install Go on your system and set up your development environment.
3. **Writing Your First Program**: Create a simple "Hello, World!" program to get started with Go.
4. **Exploring Go Packages**: Learn about Go packages and how to use them in your projects.
5. **Functions and Methods**: Understand how to define and use functions and methods in Go.
6. **Concurrency in Go**: Explore Go's concurrency model and learn how to write concurrent programs using goroutines and channels.
7. **Building and Testing**: Learn how to build and test your Go applications.
8. **Advanced Topics**: Dive into advanced topics such as error handling, reflection, and interfacing with other languages.
9. **Project Development**: Apply your knowledge by working on a real-world project using Go.

## Creating a Function in GO

In Go, you can create a function using the `func` keyword followed by the function name, parameters, and return type. Here's an example of a simple function that adds two numbers:

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 4)
    fmt.Println("Result:", result)
}
```

## Printing Output Using `fmt.Println`

The `fmt` package in Go provides various functions for formatted I/O. The `fmt.Println` function is used to print output to the console. Here's an example:

```go
package main

import "fmt"

func main() {
    message := "Hello, World!"
    fmt.Println(message)
}
```

In this example, the `fmt.Println` function prints the value of the `message` variable to the console.
