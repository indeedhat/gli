package gli

import (
  gliError "./error"
  "errors"
)

type ArgKind int
type ak struct {
  NAMED ArgKind
  KEY   ArgKind
  GROUP ArgKind
  VALUE ArgKind
}
var ARGKIND = ak{0, 1, 2, 3}


type Parser struct {
  ExpectedArgs []ExpectedArg
  Subject      Command
  Args         []string
  ValidArgs    map[string]ValidArg
}


func CreateParser(args []string, cmd Command) Parser {
  parser := Parser{}
  parser.init(args, cmd)

  return parser
}

func (p *Parser) init(args []string, cmd Command) {
  p.ExpectedArgs = extractExpected(cmd)
  p.Subject      = cmd

  p.validateArguments()
}


func (p *Parser) validateArguments() {
  // check for allowance of unexpected args
  allowUnexpected := false
  positionalMax :=
  positional := 0
  alun, okay := p.Subject.(IgnoreUnexpected)
  if okay {
    allowUnexpected = alun.IgnoreUnexpected()
  }

  for i := 0; i < len(p.Args); i++ {
    switch argToKind(p.Args[i]) {
    case ARGKIND.NAMED, ARGKIND.KEY:
      ex := p.findExpectedByKey(p.Args[i])
      if nil == ex && !allowUnexpected {
        gliError.Panic(errors.New("Unexpected Argument " + p.Args[i]), gliError.UNEXPECTED_ARG)
      }
      break
    case ARGKIND.VALUE:
      break
    case ARGKIND.GROUP:
      break
    }
  }
}


func (p *Parser) positionalCount() (count int) {
  for i := 0; i < len(p.ExpectedArgs); i++ {
    if p.ExpectedArgs[i].Type == ARGTYPE.POSITIONAL {
      p.ExpectedArgs[i].ArgType.
    }
  }
}


func (p *Parser) findExpectedByKey(key string) *ExpectedArg {
  for i := 0; i < len(p.ExpectedArgs); i++ {
    if p.ExpectedArgs[i].HasOption(key) {
      return &p.ExpectedArgs[i]
    }
  }

  return nil
}


func (p *Parser) assignArguments() {

}

func argToKind(val string) ArgKind {
  if "--" == val[:2] {
    return ARGKIND.NAMED
  } else if "-" == val[:1] && 2 == len(val) {
    return ARGKIND.KEY
  } else if "-" == val[:1] && 2 != len(val) {
    return ARGKIND.GROUP
  }

  return ARGKIND.VALUE
}