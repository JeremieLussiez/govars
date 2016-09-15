package govars

import (
	"os"
	"bufio"
	"strings"
	"regexp"
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

	file, err := os.Open(initFilePath)

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			match, _ := regexp.MatchString("[-_.a-zA-Z0-9]+[ \t]*=.*", line)
			if (match) {
				entry := strings.Split(line, "=")
				config.variables[entry[0]] = entry[1]
			}
		}
	}

	for _, programArgument := range os.Args[1:] {
		entry := strings.Split(programArgument, "=")
		if (len(entry) == 2) {
			config.variables[entry[0]] = entry[1]
		}
	}

	defer file.Close()
}

func (config GoVarsConfig) Get(key string) string {
	value := config.variables[key]
	return value
}

