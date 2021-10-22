package main

import (
  "fmt"
)


type InvertedIndex struct {
  Map map[string]map[string][]uint16
  WordCount map[string]uint16
}

func InitInvertedIndex() * InvertedIndex {
  index := InvertedIndex{}
  index.Map = map[string]map[string][]uint16{}
  index.WordCount = map[string]uint16{}
  return &index
}

func (index *InvertedIndex) getWordCount(location string) uint16 {
  return index.WordCount[location]
}

func (index *InvertedIndex) Add(word string, location string, position uint16) {
  if(index.Map[word] == nil) {
    index.Map[word] = make(map[string][]uint16)
  }
  index.Map[word][location] = append(index.Map[word][location], position)
  index.WordCount[location] = position
  fmt.Println("Added word to index")
}

func (index *InvertedIndex) getLocationsOfWord(word string) []string {
  keys := make([]string, 0, len(index.Map[word]))
  for k := range index.Map[word] {
      keys = append(keys, k)
  }
  return keys
}

func (index *InvertedIndex) getWords() []string {
  keys := make([]string, 0, len(index.Map))
  for k := range index.Map {
      keys = append(keys, k)
  }
  return keys
}

func (index *InvertedIndex) getLocations(word string, location string) []uint16 {
  keys := make([]uint16, 0, len(index.Map[word][location]))
  for _, k := range index.Map[word][location] {
      keys = append(keys, k)
  }
  return keys
}

func (index *InvertedIndex) containsWord(word string) bool {
  if _, ok := index.Map[word]; ok {
    return true
  }
  return false
}

func (index *InvertedIndex) containsWordLocation(word string, location string) bool {
  if _, ok := index.Map[word][location]; ok {
    return true
  }
  return false
}

func (index *InvertedIndex) contains(word string, location string, position uint16) bool {
  for _, item := range index.Map[word][location] {
    if item == position {
      return true
    }

  }
  return false
}

func (index *InvertedIndex) wordSize() int {
  return len(index.Map)
}

func (index *InvertedIndex) locationSize(word string) int {
  return len(index.Map[word])
}

func (index *InvertedIndex) positionSize(word string, location string) int {
  return len(index.Map[word][location])
}

func main() {
  index := InitInvertedIndex()
  index.Add("a", "simple.txt", 2);
  index.Add("c", "simple.txt", 3);
  index.Add("b", "simple.txt", 5);
  index.Add("a", "simple.txt", 3);
  index.Add("c", "simple2.txt", 1);
  // test 
  fmt.Println(index)
  fmt.Println(index.getLocationsOfWord("c"))
  fmt.Println(index.getWords())
  fmt.Println(index.getLocations("a","simple.txt"))
  fmt.Println(index.contains("a","simple.txt",2))
  fmt.Println(index.containsWord("a"))
  fmt.Println(index.containsWordLocation("a","simple.txt"))
  fmt.Println(index.contains("t","simple.txt",2))
  fmt.Println(index.containsWord("u"))
  fmt.Println(index.containsWordLocation("a","simpley.txt"))
  fmt.Println(index.wordSize())
  fmt.Println(index.locationSize("c"))
  fmt.Println(index.positionSize("a","simple.txt"))
}
