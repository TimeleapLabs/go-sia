# Sia

[![Build Status](https://github.com/logicalangel/go-sia/actions/workflows/test.yml/badge.svg?branch=master)][actions]

Sia - Binary serialisation and deserialisation with built-in compression. You can consider Sia a strongly typed, statically typed domain specific binary language for constructing data. Sia preserves data types and supports custom ones.

[actions]: https://github.com/logicalangel/go-sia

## Install

`go get github.com/pouya-eghbali/go-sia/v2`

## Basic Usage

To serialize multiple values, first create a sia object and then you can add values in order. Note that the order of adding values should be considered when you want to read them again.

```go
package main

import (
	"fmt"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

type person struct {
	Name string
	Age  uint8
}

func (p *person) Sia() sia.Sia {
	return sia.New().
		AddString8(p.Name).
		AddUInt8(p.Age)
}
func (p *person) FromSia(s sia.Sia) {
	p.Name = s.ReadString8()
	p.Age = s.ReadUInt8()
}
func main() {
	p := person{Name: "Pouya", Age: 33}
	payload := p.Sia().Bytes()
	fmt.Println(payload) // [5 80 111 117 121 97 33]
	var p2 person
	p2.FromSia(sia.New().EmbedBytes(payload))
	fmt.Println(p2) // {Pouya 33}
}
```

## Serializers

Sia provides a set of serializers for each primitive type. They are most useful for cases where you're adding an array of values. Instead of writing the function yourself, just use the exported utility functions.

```go
import (
	sializer "github.com/pouya-eghbali/go-sia/v2/pkg"
)

const sia = sializer.New();
var words = []string{"Hello", "World"}
sia.AddArray8(words, sializer.SerializeString8ArrayItem);
```

The `SerializeString8ArrayItem` runs for each item in the array and adds the item to the Sia content.

## Deserializers

For the opposite scenario, where you want to read an array of values from the Sia content, you can use the `ReadArray8` method in combination with the deserializers.

```go
import (
	sializer "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func main() {
	sia := sializer.New()
	var words = []string{"Hello", "World"}
	sia.AddArray8(words, sializer.DeserializeString8ArrayItem)

	const desia = sializer.New().EmbedBytes(sia.Bytes());
	const deserialized = desia.ReadArray8(sializer.DeserializeString8ArrayItem);

	fmt.Println(deserialized); // ["Hello", "World"]
}
```
