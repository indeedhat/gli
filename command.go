package gli


// gli command, anything implementing this interface can be treated as a command
// adn handled by this lib
type Command interface {
  Run() int
}

// auto help commands will show help on any error if AutoHelp() returns true
type AutoHelper interface {
  AutoHelp() bool
}


// Unexpected input
// implementing this interface and returning true will allow the command
// to ignore unexpected arguments, by default they would cause an error
type IgnoreUnexpected interface {
  IgnoreUnexpected() bool
}