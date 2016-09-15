# govars
Very basic variable loader for go language.

You can either retrieve your variable from:
* environment
* .ini configuration file
* command line

## Usage

### Go program

#### Initialize

~~~~
import (
	"github.com/JeremieLussiez/govars"
)
var configuration = govars.GoVarsConfig{}
func main {
    configuration.Load("anEventuallyExistingConfiurationFile.ini");
}
~~~~

#### Get variable

~~~~
configuration.Get("MY_VARIABLE_FROM_SOMEWHERE")
~~~~

### Setup variables

#### Environment

Just set you variable like you usually do with your favorite OS.

#### .INI File

~~~~
# You can put comments in you .ini file
myBeautiful_variable.so-nice=Something very beautiful
anotherVariableWithSpaces=      Look, some spaces before the string !
~~~~

#### Command line

~~~~
myWonderfulGoProgram myCuteVariable=kitten myCuteOtherVariable=puppy
~~~~

