# jsondelta implementation for Go

See http://json-delta.readthedocs.io/en/latest/

## Current state

The `diff` implementation works and there are extensive tests, however:

* The algorithm for diffing lists is just a placeholder and as such extremely naive and highly suboptimal.
* There are some corner cases to work on, such as numeric types other than int and float64. It has to be decided if e.g. a float32 can equal a float64, or not, or if that can and should be made configurable. 
* Both performance and memory usage are likely be suboptimal. Especially the handling of custom types. 

There is no `patch` implementation yet.

A command line tool to use the library should be added.