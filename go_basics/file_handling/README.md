- Working with low level api's
  - In go the file mode value is represented by `os.FileMode` type, which is derived from `uint32`
  - `os` file package provides a `File.Chmod()` function to change the permission of a file, which takes numeric file mode value of type `os.FileMode`
  - **File Open-time Flags**
    - When we open a file using some open() function call, we can pass an extra argument which is called an open-time flag that determines how a file should be opened and what operations can be performed on it.
      ![open_time_flags](https://miro.medium.com/max/512/1*LIWYx3qdT81h9uvWPbWZwQ.png)
    - `func OpenFile(name string, flag int, perm FileMode) (*File, error)` 
    - The os.OpenFile function opens an existing file on the system or creates one with the perm file mode if not already exists. The flag argument is the open-time flags.
