package orderedmap

import (
	"errors"
	"sort"
)

type Pair struct {
	Key   string
	Value string
}

type OrderedMap struct {
	keyToPair map[string]Pair
	keyToCounter map[string]int
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		keyToPair: make(map[string]Pair),
		keyToCounter: make(map[string]int),
	}
}

func (o *OrderedMap) Set(key, value string) error {
	if (key == "") {
		return errors.New("key cannot be empty")
	}
	o.keyToCounter[key] = len(o.keyToPair)
	o.keyToPair[key] = Pair{key, value}

	return nil
}

func (o *OrderedMap) Get(key string) string {
	return o.keyToPair[key].Value
}

func (o *OrderedMap) Delete(key string) {
	delete(o.keyToPair, key)
	delete(o.keyToCounter, key)
}

type KeyCounter struct {
	Key     string
	Counter int
}

func (o *OrderedMap) GetAll() []Pair {
	counters := make([]KeyCounter, 0, len(o.keyToCounter))
	for key, counter := range o.keyToCounter {
		counters = append(counters, KeyCounter{
			Key:     key,
			Counter: counter,
		})
	}
	sort.Slice(counters, func(i, j int) bool {
		return counters[i].Counter < counters[j].Counter
	})

	pairs := make([]Pair, 0, len(o.keyToPair))
	
	for _, counter := range counters {
		pairs = append(pairs, o.keyToPair[counter.Key])
	}
	return pairs
}
