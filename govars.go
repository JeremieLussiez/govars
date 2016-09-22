package govars

import (
	"os"
	"bufio"
	"strings"
	"regexp"
	"strconv"
)

type GoVarsConfig struct {
	variables map[string]string
}

func (config *GoVarsConfig) Load(initFilePath string) {

	config.variables = make(map[string]string)

	for _, envVariable := range os.Environ() {
		entry := strings.Split(envVariable, "=")
		config.variables[entry[0]] = entry[1]
	}

	file, fileOpeningError := os.Open(initFilePath)
	defer file.Close()

	if fileOpeningError == nil {
		lineMatcher := regexp.MustCompile(`^(?P<Name>[-_.a-zA-Z0-9]+)[ \t]*=[ \t]*"?(?P<Value>.*?)"?$`)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			matched := lineMatcher.FindStringSubmatch(line)
			if (matched != nil && len(matched) >= 2) {
				config.variables[matched[1]] = matched[2]
			}
		}
	}

	for _, programArgument := range os.Args[1:] {
		entry := strings.Split(programArgument, "=")
		if (len(entry) == 2) {
			config.variables[entry[0]] = entry[1]
		}
	}

}

func (config GoVarsConfig) GetString(key string) string {
	value := config.variables[key]
	return value
}

func (config GoVarsConfig) GetFloat32(key string) float32 {
	value, present := config.variables[key]
	if (!present) {
		return 0
	}
	output, conversionError := strconv.ParseFloat(value, 32)
	if (conversionError != nil) {
		return 0
	}
	return float32(output)
}

func (config GoVarsConfig) GetFloat64(key string) float64 {
	value, present := config.variables[key]
	if (!present) {
		return 0
	}
	output, conversionError := strconv.ParseFloat(value, 64)
	if (conversionError != nil) {
		return 0
	}
	return output
}

func (config GoVarsConfig) GetInt(key string) int {
	value, present := config.variables[key]
	if (!present) {
		return 0
	}
	output, conversionError := strconv.Atoi(value)
	if (conversionError != nil) {
		return 0
	}
	return output
}