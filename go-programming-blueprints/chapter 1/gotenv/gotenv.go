package gotenv

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Env type to store key and var
type Env map[string]string

// Load to load env file without override already set variables
func Load(filename string) (Env, error) {
	return loadenv(false, filename)
}

// func OverLoad(filename string) {
// 	loadenv(true, filename)
// }

func loadenv(override bool, filename string) (Env, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	env, err := parseFile(f, override)

	if err != nil {
		return nil, err
	}

	f.Close()

	return env, nil
}

func parseFile(f io.Reader, override bool) (Env, error) {
	env, err := strictParse(f)

	if err != nil {
		return nil, err
	}

	for key, val := range env {
		fmt.Println(key, " ", val, " ", override)
		// setenv(key, val, override)
	}

	return env, nil
}

func strictParse(r io.Reader) (Env, error) {
	env := make(Env)

	scanner := bufio.NewScanner(r)

	i := 1

	bom := string([]byte{239, 187, 191})

	for scanner.Scan() {
		line := scanner.Text()

		if i == 1 {
			line = strings.TrimPrefix(line, bom)
		}

		i++

		err := parseLine(line, env)

		if err != nil {
			return nil, err
		}
	}

	return env, nil
}

func parseLine(s string, env Env) error {
	splited := strings.Split(s, "=")

	if len(splited) != 2 {
		return errors.New("line incorrect in env file")
	}

	env[splited[0]] = splited[1]

	return nil
}

func setenv(key, val string, override bool) {
	if override {
		os.Setenv(key, val)
	} else {
		if _, present := os.LookupEnv(key); !present {
			os.Setenv(key, val)
		}
	}
}
