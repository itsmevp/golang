# golang

**Formatting:**
  - Indentation: We use tabs for indentation and gofmt emits them by default. Use spaces only if you must.
  - Line length: Go has no line length limit.
  - Parentheses: Go needs fewer parentheses, control structures (if, for, switch) do not have parentheses in their syntax.

**Commentary:**
  - Block comments: "/* */"
  - Line comments: //
  - godoc is a program and a web server processes Go source files to extract documentation about the contents of the package.
  - Every package should have a package comment, a block comment preceding the package clause.
  - For multi-file packages, the package comment only needs to be present in one file, and any one will do.
  - Inside a package, any comment immediately preceding a top-level declaration serves as a doc comment for that declaration.
  - Every exported (capitalized) name in a program should have a doc comment.
  - `go doc -all [package name] | grep -i "name of the function"`
  - `godoc -http=:6060` to start an local http documentation server

**Naming:**
  - The visiblity of a name outside a package is determined by whether it's first character is upper case or not.
  - *Packaging:*
    - When a package is imported, the package name becomes an accessor for the contents.
    - By convention, packages are given lower case , single-word names; there should be no need for underscores or mixedCaps.
    - The function to make new instances of ring.Ring—which is the definition of a constructor in Go—would normally be called NewRing, 
      but since Ring is the only type exported by the package, and since the package is called ring, it's called just New, which clients of the package see as ring.New.
  - *Getters:*
    - Go doesn't provide automatic support for getters and setters.
    - If you have a field called owner (lower case, unexported), the getter method should be called Owner (upper case, exported), not GetOwner.
    - A setter function, if needed, will likely be called SetOwner.
  - *Interface Names:*
    - By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.
  - *MixedCaps:*
    - Finally, the convention in Go is to use MixedCaps or mixedCaps rather than underscores to write multiword names.

**Semicolon's:**
  - The rule is:
    - If the last token before a newline is an identifier (which includes words like int and float64), a basic literal such as a number or string constant, or one of the tokens
      `break continue fallthrough return ++ -- ) }` the lexer always inserts a semicolon after the token.
    - This could be summarized as, “if the newline comes after a token that could end a statement, insert a semicolon”.

**Control Structures:**
  - There is no do or while loop,
    only a slightly generalized for;
    switch is more flexible;
    if and switch accept an optional initialization statement like that of for;
    break and continue statements take an optional label to identify what to break or continue;
    and there are new control structures including a type switch and a multiway communications multiplexer, select.
  - `for init; condition; post { }` like a C for loop
  - `for condition { }` like a C while loop
  - `for { }` Like a C for(;;)
  - Go has no `,` operator
  - `++` and `--` are statements not expressions
