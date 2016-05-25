package types

import "time"

// Set is used for storing aggregated values for sets.
type Set struct {
	Values   map[string]int64 // The number of occurrences for a specific value
	Interval                  // The flush and expiration interval information
}

// NewSet initialises a new set.
func NewSet(timestamp time.Time, flushInterval time.Duration, values map[string]int64) Set {
	return Set{Values: values, Interval: Interval{Timestamp: timestamp, Flush: flushInterval}}
}

// Sets stores a map of sets by tags.
type Sets map[string]map[string]Set

// MetricsName returns the name of the aggregated metrics collection.
func (s Sets) MetricsName() string {
	return "Sets"
}

// Delete deletes the metrics from the collection.
func (s Sets) Delete(k string) {
	delete(s, k)
}

// DeleteChild deletes the metrics from the collection for the given tags.
func (s Sets) DeleteChild(k, t string) {
	delete(s[k], t)
}

// HasChildren returns whether there are more children nested under the key.
func (s Sets) HasChildren(k string) bool {
	return len(s[k]) != 0
}

// Each iterates over each set while f return true.
// Returns true if all items, if any, were visited.
func (s Sets) EachWhile(f func(string, string, Set) bool) bool {
	for key, value := range s {
		for tags, set := range value {
			if !f(key, tags, set) {
				return false
			}
		}
	}
	return true
}

// Each iterates over each set.
func (s Sets) Each(f func(string, string, Set)) {
	for key, value := range s {
		for tags, set := range value {
			f(key, tags, set)
		}
	}
}
