package cli

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
)

type Config struct {
	Root  string
	Build string
	Exec  string
}

func Parse() (*Config, error) {

	root := flag.String("root", "", "project root directory")
	build := flag.String("build", "", "build command")
	execCmd := flag.String("exec", "", "execution command")

	flag.Parse()

	if *root == "" {
		return nil, errors.New("missing --root")
	}

	if *build == "" {
		return nil, errors.New("missing --build")
	}

	if *execCmd == "" {
		return nil, errors.New("missing --exec")
	}

	absRoot, err := filepath.Abs(*root)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(absRoot)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, errors.New("root path is not a directory")
	}

	cfg := &Config{
		Root:  absRoot,
		Build: *build,
		Exec:  *execCmd,
	}

	return cfg, nil
}
