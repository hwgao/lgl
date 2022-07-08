# lgl

`lgl` is a simple logger supporting log level.

## Features
- Work out of the box.
- Only global functions are provided. Just `import` the package, then you can use it.
- Supported log levels: `Trace`, `Debug`, `Warning`, `Error`.
- By default show the file name and line number.
- By default log level is `Error`. Call `SetLogLevel()` in `main()` to enable more log.
- Call `Set()` in `main()` to save log to a file or other location.
