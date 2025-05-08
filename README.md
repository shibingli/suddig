## suddig {adj.} /²s'ɵdːɪg/

*(Swedish for "fuzzy")*

A fast, flexible, and modular fuzzy finder library for Go. Built for both quick, one-line matches and deep customization, suddig provides simple functions (`Match`, `Distance`, `Score`) for everyday use, and a fully configurable `matcher` package for advanced scenarios.

---

## Features

* **Modular Architecture**: Swap or extend normalization, distance, and scoring components.
* **Simple API**: Call `suddig.Match`, `suddig.Distance`, `suddig.Score`, `suddig.FindMatches` & `suddig.RankMatches` out of the box.
* **Advanced Configuration**: Instantiate a `matcher.Matcher` with custom `Config` to tweak behavior.
* **Lightweight**: No external dependencies, minimal overhead.

---

## Installation

```bash
go get github.com/yourusername/suddig
```
