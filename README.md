# check

[![Go Report Card](https://goreportcard.com/badge/github.com/imulab/check)](https://goreportcard.com/report/github.com/imulab/check)
[![Version](https://img.shields.io/badge/version-v0.1.0-green)](https://img.shields.io/badge/version-v0.1.0-green)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://img.shields.io/badge/license-MIT-blue)
[![Go Reference](https://pkg.go.dev/badge/github.com/imulab/check.svg)](https://pkg.go.dev/github.com/imulab/check)

functional, fluent and extensible validation toolkit for Go.

__This library is under active development, feature PRs are welcomed__

## Goals

When validating large objects with sophisticated business logic and intertwined value dependencies, imperative
programming often makes it hard to tell the actual logic underneath, thus making it hard to maintain. To add salt to
injury, Go's verbose error handling makes it even harder to maintain.

This project aims at providing a validation API that is *almost* **fluent** in natural language, **composable** to
represent sophisticated logic, and **extensible** wherever the defaults fall short.

In addition, with the assumption that the validation code already knows how the validated value should be like, this
library would try to stay away from reflection as much as possible, while having a **minimal API** surface area.

## Install

```bash
go get -u github.com/imulab/check
```

## Main Concepts

The library is designed with minimal API surface, it has only a few concepts:

1. `check.Step` is the core concept, it takes in any `interface{}` value with assumed type, and applies validation to
   it. It returns an error when validation fails.
2. `check.Skip` is the special error to be returned by `check.Step`, to skip all remaining steps.
3. `check.That` is the main entrypoint for validation, it accepts multiple `check.Step` to execute sequentially.
4. `check.AnyErr` can chain multiple `check.That` together to eagerly return any error.

## Usage

```go
// Check str variable is not empty.
check.That(str, stringz.IsNotEmpty)

// Check str is "a" or "b", as long as it is not empty. If str is
// not "a" and "b", return a customErr.
check.That(str,
    check.Optional.When(stringz.IsEmpty),
    stringz.In("a", "b").Err(customErr),
)

// Check all elements of the string slice is "a", "b", or "c", or
// return a customErr.
check.That(slice,
    slicez.OfString.All(stringz.In("a", "b", "c")),
).Err(customErr)
```
