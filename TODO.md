# GLI
This is an experiment into creating a simple struct based framework for creating cli applications with go

## parse
- Positional args
`gli:""`
- Names args
`gli:"name,n,N"`
- Flags
`same as named but when struct val is of type bool`
- Required Args
`gli:"!"`
- Optional Args
`default behaviour`
- Chained flags
`default behaviour but all in chan except last one must be flags + will only work with single char flags`
- Default values
`default:"value"`
- Option to ignore unexpected args
`by impplementing interface`
- Auto help of fail
`by implementing interface`

`help and h will both be set to help output by default and cannot be changed`

## logic
- Command Interface
- Child commands
- propagate errors from commands to stderr
- error helpers
- output helpers
- commands must return int for return code in cli

## tbd
- allow for one of many args to be required
- Have a interface implementation for replacing default help





