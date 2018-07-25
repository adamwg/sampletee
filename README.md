# sampletee

`sampletee` is a simple commandline utility for printing a sample of the output
of a command while recording the full output in one or more files. It operates
exactly like the standard Unix utility `tee`, except that instead of echoing all
lines of input to stdout, it echoes only every nth line. Invoking sampletee with
`-n 1` is equivalent to using regular `tee`.

## Usage

```
Usage: sampletee [-n n] [-a] [files...]
  -a    append to output files rather than overwriting
  -n int
        echo every nth line (default 10)
```

Note that, like `tee`, the output files are optional. Calling `sampletee`
without any output files will simply sample standard intput to standard output
according to the `-n` parameter.
