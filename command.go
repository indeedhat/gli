package gli


// gli command, anything implementing this interface can be treated as a command
// adn handled by this lib
type Command interface {
    Run() int
}


// used to generate help documentation
type Helper interface {
    NeedHelp() bool
}


// Unexpected input
// implementing this interface and returning true will allow the command
// to ignore unexpected arguments, by default they would cause an error
type IgnoreUnexpected interface {
    IgnoreUnexpected() bool
}


// returns a custom help string to be passed on help call or error
type CustomHelp interface {
    Helper
    GenerateHelp(expected []*ExpectedArg) string
}