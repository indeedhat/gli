# GLI Command Line Helper

The purpose of this package is to provide a simple struct based framework for cli argument parsing and command routing

if you need any more than that take a look at one of the other many command line frameworks available for go.

## Limitations
 - supported types int +variants, uint +variants, float32/64, string, bool, Slices of the previous types
 - All commands/sub commands must implement the Command interface
  

## Roadmap (no particular order)
 - Document the usage
 - Add some more examples
 - write tests
 - override flag
 - find a nice way of giving a description to both the app and any sub commands it has
 
## Thinking about
 - allow for an argument to be required, or not based on the other arguments provided
 
 This repo is covered under the MIT licence