package utils

import (
	"log"

	"github.com/subosito/gotenv"
)

func LoadVars() error {
	paths := []string{
		"/config/config",
		"./config/config",
		"../config/config",
		"../../config/config",
		"../../../config/config",
		"../../../../config/config",
		"../../../../../config/config",
		"../../../../../../config/config",
		"../../../../../../../config/config",
	}
	var loaded bool
	for _, path := range paths {
		// err := gotenv.Load(filepath.Join(os.Getenv("PWD"), path))
		err := gotenv.Load(path)
		if err == nil {
			loaded = true
		}
	}

	if !loaded {
		log.Fatal("load-vars: config not found")
	}
	return nil
}
