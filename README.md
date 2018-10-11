# GLI Command Line Helper

The purpose of this package is to provide a simple struct based framework for cli argument parsing and command routing

if you need any more than that take a look at one of the other many command line frameworks available for go.

## Limitations
 - supported types int +variants, uint +variants, float32/64, string, bool, Slices of the previous types
 - All commands/sub commands must implement the Command interface
  

## Roadmap (no particular order)
 - Add some more examples
 - write tests
 - find a nice way of giving a description to both the app and any sub commands it has
 
## Thinking about
 - allow for an argument to be required, or not based on the other arguments provided
 
 This repo is covered under the MIT licence
 
## Usage

### Define a command
The simples form of application poosible would be just a single command with no arguments, if that is your
 goal then i dont know why you are looking at this package but carry on.
 
This is done simply by defining a struct that implements the Command interface
```go
import "github.com/indeedhat/gli"

type Application struct {}

func (app *Application) Run() int {
    fmt.Println("Hello, World!")
    
    return 0
}

func main() {
    app := gli.NewApp(&Application{})
    app.Run()
}
```

### arguments
Arguments can be of almost any simple scalar type: 
`int/uint +variations`, `float32/64`, `string`, `bool` or an array/slice of these types\
At this time chars are not supported

Args are created by adding the `gli` tag to a struct field

#### Named arguments
```go
type Application struct {
    SomeArg string `gli:"somearg"`
    OtherArg string `gli:"other"`
}
```
they would then be assigned as follows

`./app --somearg val1 --other val2`
this is also valid
`/app --somearg=val1 --other "val with spaces"`

#### Name aliases
you may add aliases to argument names for other options on how to set them

```go
type Application struct {
    SomeArg string `gli:"somearg,sa,s"`
    OtherArg string `gli:"other,o"`
}
```

you can then use any of the aliases to assign to that field\
Note: single character aliases can use a single dash
`./app -o "someval" --sa "and another one"`
when using a single dash argument you can ommit the space before the value if you wish
`./app -ovalue -s"another one"`

#### Flags
If the struct field is a bool it will be classed as a flag\
by default all flags are false until they are called. 
```go
type Application struct {
    Flag bool `gli:"flag,f"`
    AnotherFlag bool `gli:"a"`
}
```
flags do not require a value to be set when called
`./app --flag -a`
however if you wish you can set a value as with a named argument
`./app --flag true -afalse`

#### Flag Groups
When using flags and their single character aliases you may group two or more together
`./app -a -f`
is identical to
`./app -af`

when using groups all arguments must be flags (bool type) except the last one
```go
type Application struct {
    Flag bool          `gli:"flag,f"`
    AnotherFlag bool   `gli:"a"`
    StringVal   string `gli:"s"`
}
```
this could be called as follows
`./app -afs "value for string"`
however if called in this order
`./app -saf "value for string"` StringVal would get the value "af" and the flags would not be set

#### Positional arguments
positional arguments can be defined by no setting a name in the gli tag
```go
type Application struct {
    Positional string `gli:""`
    Pos2       string `gli:""`
}
```
In the case of this call 
`./app val1 val2`
"val1" would be assigned to Positional and "val2" would be assigned to Pos2

#### Arrays/Slices
Any argument can be a slice of the other possible types
```go
type Application struct {
    Named      []string `gli:"name,n"`
    Positional string   `gli:""`
    Pos2       []string `gli:""`
}
```
for named arguments this allows the argument to be called multiple times
`./app --name jon -n frank -n jenny`
and the value of Named would be \["jon" "frank" "jenny"]\
in the case of positional arguments the values will still be assigned in order of the fields in the struct 
but once a slice filed is found all following positional arguments will be assigned
`./app arg1 arg2 arg3`
Positional would get the value "arg1" and Pos2 would get \["arg2" "arg3"]\
any positional arguments defined in the struct after Pos2 would never get a value

### Required
Any argument named, flag or positional can be marked as required by starting the gli tag with an !
```go
type Application struct {
    Required string `gli:"!required,r,R"`
}
```
If a required argument isnt given then an error will be returned and the command will not be run

### Override
Any argument may also be marked as an override argument with the use of ^
```go
type Application struct {
    Override bool   `gli:"^o"`
    Required string `gli:"!required,r,R"`
}
```
Whenever an override argument is found all other arguments will be discarded and the command will be run\
`./app` would error out because the required argument was not given\
`./app -o` would run without error because the override argument was set\
`./app -r somedata -o` would discard the argument for -r and run in the same way as the above example

### Sub commands
Sub commands can be added to give you application more functionality this is done by adding a struct field with the 
value of another struct that implements command

```go
type Application struct {
    Override   bool   `gli:"^o"`
    Required   string `gli:"!required,r,R"`
    SubCommand Sub    `gli:"sub,s"`
}
...

type Sub struct {
    Arg string `gli:'a'`
}
func (s *Sub) Run() int {
    fmt.Println("This is the sub command")
    
    return 0
}
``` 
Sub commands can be called by name in the argument string but they have some gotchas:
- Sub commands must appear before any other arguments
- Sub commands must be named, they cannot be positonal
- Sub commands cannot be called as a single or double dash argument
As such the following would be valid:\
`./app sub`\
`./app s`\
`./app sub -a val`\
however these would not\
`./app -r val sub`\
`./app -s`\
`./app --sub -a val`\
Sub commands can be nested as far as you like but the chain must come before any other arguments.
All none command arguments will be passed to the final command in the chain and only the Run method of the final 
command will be called
```go
type Application struct {
    Sub Sc1 `gli:"sub"`
}
...
type Sc1 struct {
    Sub Sc2 `gli:"s2"`
}
...
type Sc2 struct {
    Val string `gli:""`
}
func (cmd *Sc2) Run() int {
    fmt.Println(cmd.Val)
    
    return 0
}
```
in this case the following call\
`./app sub s2 "My Value"`\
Would call the run method of Sc2 resulting in "My Value" being printed to the screen

### Help
Gli provides a help generation feature out of the box. To make use of this you must implement the `Helper` 
interface as follows
```go
type AppWithHelp struct {
    Help bool `gli:"^help,h"`
}
func (app *AppWithHelp) NeedHelp() bool {
    return app.Help
}
```
Whenever the `NeedHelp` method returns true the help will be shown\
Note: it is recommended that you make your help argument an override to stop other argument checks
but it is not actually required. Your run method will still be called prior to showing the help response.

### description tag
To aid with help generation the documentation builder will read any `description` tags assigned to your struct 
fields and include them in the help output
```go
type AppWithHelp struct {
    Help bool `gli:"^help,h"`
    Arg  string `gli:"arg" description:"this is an arg"`
}
...
``` 

### Default Values
If you wish for default values to be assigned to your struct fields when they are not set by the command arguments
this can be done with the `default` tag
```go
type AppWithCustomHelp struct {
    Help       bool   `gli:"^help,h"`
    Arg        string `gli:"arg" description:"this is an arg"`
    HasDefault int32  `gli:"" description:"this is an int argument" default:"384"`
}
...
``` 
If the default value is not compatible with the type of the field an error will occur whenever the value is 
not set by the command line args

### Custom Help
If you wish to write your own help docs then these can be incorporated into the application by implementing
the `CustomHelp` interface
```go
type AppWithCustomHelp struct {
    Help bool `gli:"^help,h"`
    ...
}
func (app *AppWithCustomHelp) NeedHelp() bool {
    return app.Help
}
func (app *AppWithCustomHelp) GenerateHelp(expected []*ExpectedArg) string {
    ...
}
```
An ExpectedArg will be generated for each struct field in the current command

### Unexpected Arguments
By default unexpected arguments will cause a parser error and your app will not run.\
If you would like to ignore unexpected arguments that can be done by implementing the `IgnoreUnexpected`
interface
```go
type AppThatIgnores struct {
    ...
}
func (app *AppThatIgnores) IgnoreUnexpected() bool {
    return true
}
```
If you want it to actually ignore the arguments then the `IgnoreUnexpected()` method must return true

### Note on all interfaces
Any functionality added by implementing a gli interface will only be present on the command it is implemented on.\
Implementing `Helper` on the root command for example will not enable it in sub commands unless they also implement
`Helper`
