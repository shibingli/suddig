## suddig {adj.} /²s'ɵdːɪg/

*(Swedish for "fuzzy")*

A fast, flexible, and modular fuzzy finder library for Go. Built for both quick, one-line matches and deep customization, suddig provides simple functions (`Match`, `Distance`, `Score`) for everyday use, and a fully configurable `matcher` package for advanced scenarios.

---

## Features

* **Modular Architecture**: Swap or extend normalization, distance, and scoring components.
* **Simple API**: Call `suddig.Match`, `suddig.Distance`, `suddig.Score`, `suddig.FindMatches` & `suddig.RankMatches` out of the box.
* **Advanced Configuration**: Instantiate a `matcher.Matcher` with custom `configs.Config` to tweak behavior.

---

## Installation

```bash
go get github.com/VincentBrodin/suddig
```

---

## Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/VincentBrodin/suddig"
)

func main() {
	args := os.Args[1:]
	argc := len(args)

	if argc < 2 {
		fmt.Println("Need 2 arguments")
		os.Exit(1)
	}
	s1, s2 := args[0], args[1]

	match := suddig.Match(s1, s2)

	if match {
		fmt.Printf("%s and %s match!\n", s1, s2)
	} else {
		fmt.Printf("%s and %s do not match\n", s1, s2)
	}
}
```

---

## Goals

The goal is to support as many different distance algorithms,
and give the tools needed for both simple and compelx usecase.

---

## Contribution

Please add any score, distance or configs you feel are missing, and submit a PR.
