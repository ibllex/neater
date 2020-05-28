package parser

import (
	"neater/utils"
)

// FilterContainer Contains a series of filters,
// and executes the corresponding filters according to the filtering conditions
type FilterContainer struct {
	filters []*Filter
}

// Add a New filter into container
func (c *FilterContainer) Add(filter *Filter) *FilterContainer {
	c.filters = append(c.filters, filter)
	return c
}

// Match Get a filter by given file extension name
// return nil if no filter match
func (c *FilterContainer) Match(ext string) *Filter {
	for _, filter := range c.filters {

		// Whitelist takes precedence, if there are whitelist rules, the blacklist becomes invalid

		// Enable whitelist rules
		if len(filter.whitelist) > 0 {
			if i := utils.IndexOfStringList(ext, filter.whitelist); i != -1 {
				return filter
			}
		} else if len(filter.blacklist) > 0 { // Enable blacklist rules
			if i := utils.IndexOfStringList(ext, filter.blacklist); i == -1 {
				return filter
			}
		}
	}
	return nil
}
