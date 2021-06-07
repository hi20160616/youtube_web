package configs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type configuration struct {
	Title, Address, RootPath, DBPath, TmplPath string
}

var Value = &configuration{}

func init() {
	root, err := os.Getwd()
	if err != nil {
		log.Printf("config Getwd: %#v", err)
	}
	// root = "../../../" // for test handler
	f, err := os.ReadFile(filepath.Join(root, "configs/configs.json"))
	if err != nil {
		log.Printf("config ReadFile: %#v", err)
	}
	if err = json.Unmarshal(f, Value); err != nil {
		log.Printf("config Unmarshal err: %#v", err)
	}
	Value.RootPath = root
	Value.DBPath = filepath.Join(root, Value.DBPath)
	Value.TmplPath = filepath.Join(root, Value.TmplPath)
}
