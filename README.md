# GLI Command Line Helper

The purpose of this package is to provide a simple struct based framework for cli argument parsing and command routing

if you need any more than that take a look at one of the other many command line frameworks available for go.

## Limitations
 - supported types int +variants, uint +variants, float32/64, string, Slices of the previous tyoes
 - All commands/sub commands must implement the Command interface
  

## Roadmap (no particular order)
 - auto generate help output based on gli tags
 - Document the usage
 - Add some more examples
 - write tests
 
## Thinking about
 - allowing for setting booleans with true/false keywords
 - allow for boolean arrays
 - allow for an argument to be required, or not based on the other arguments provided