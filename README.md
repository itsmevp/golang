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
