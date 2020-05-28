package parser

import (
	"log"
	"neater/utils"
	"os"
	"path"
	"strings"
)

// RuleConfig The structure of a rule in the configuration file
type RuleConfig struct {
	Rule   string
	Target string
}

// Rule The rule matches the file name and returns the output path
type Rule struct {
	name   string
	config []RuleConfig
}

// GetName Get rule's name
func (rule *Rule) GetName() string {
	return rule.name
}

// Exec Match the rules and copy files
func (rule *Rule) Exec(fullpath string, base string, move bool) bool {
	for _, item := range rule.config {
		if item.Rule == "*" || strings.Contains(path.Base(fullpath), item.Rule) {
			destination := path.Join(base, item.Target, path.Base(fullpath))
			os.MkdirAll(path.Dir(destination), os.ModePerm)
			log.Println("from " + fullpath + " to: " + destination)

			if move {
				os.Rename(fullpath, destination)
			} else {
				_, err := utils.Copy(fullpath, destination)
				if err != nil {
					log.Println(err)
				}
			}

			return true
		}
	}
	return false
}

// NewRule Construct a read-only Rule
func NewRule(name string, config []RuleConfig) *Rule {
	return &Rule{
		name:   name,
		config: config,
	}
}
