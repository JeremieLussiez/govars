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

##### String
~~~~
configuration.GetString("MY_VARIABLE_FROM_SOMEWHERE")
~~~~

##### Integer
~~~~
configuration.GetInt("MY_VARIABLE_FROM_SOMEWHERE")
~~~~

##### Float 32
~~~~
configuration.GetFloat32("MY_VARIABLE_FROM_SOMEWHERE")
~~~~

##### Float 64
~~~~
configuration.GetFloat64("MY_VARIABLE_FROM_SOMEWHERE")
~~~~

### Setup variables

#### Environment

Just set you variable like you usually do with your favorite OS.

#### .INI File

~~~~
# You can put comments in you .ini file
# You can define string variables (surrounded by double quotes)
stringVariable = "This is a "string" variable"

# You can define empty string variables too
emptyStringVariable = ""

# You also can define integer variables
intVariable = 42

# Or float variables
floatVariable = 3.14
~~~~

#### Command line

~~~~
myWonderfulGoProgram myCuteVariable=5 myCuteOtherVariable=42
~~~~

