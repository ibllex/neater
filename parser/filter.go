package parser

// FilterConfig The structure of the filter in the configuration file
type FilterConfig struct {
	Whitelist []string
	Blacklist []string
	Rule      string
}

// Filter The filter matches the rules for processing this type of file by file suffix
type Filter struct {
	whitelist []string
	blacklist []string
	rule      *Rule
}

// GetRule Get the filter rule
func (filter *Filter) GetRule() *Rule {
	return filter.rule
}

// NewFilter Construct a read-only Filter
func NewFilter(config FilterConfig, rule *Rule) *Filter {
	return &Filter{
		whitelist: config.Whitelist,
		blacklist: config.Blacklist,
		rule:      rule,
	}
}
