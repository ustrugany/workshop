## Cover 
- how to initialize module `run go mod init example.com/basics`
- download dependencies `go mod vendor`, try to inspect go.mod
- `package main` is *the application* and needs to implement `main()` method
- always handle errors
- import local package by name `import "example.com/basics/hello"`
- unwrap chain of errors and try to find error matching passed type
- error handling, when to wrap?
  - wrap an error to expose it to callers.
  - do not wrap an error when doing so would expose implementation details.
- we need to pass pointer to struct *FileProvider as pointer receiver implements the interface Provider
- defer is a language mechanism that puts your function call into a stack, be called before returning from function
- pointer method receiver fp would be self/this in different languages
- implicit interface implementation
- we define method Provide on struct HelloProvider
- it is pointer type of method receiver
- if function return error it should be last return argument
- convention is to return pointer from constructor New* type of methods
- fields and methods starting from capital letter will be exported from package, others are private
- interface
- each folder is a package
- no classes, functionality is grouped in packages
- data is grouped into structs
