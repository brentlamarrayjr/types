Types
=========

Types uses reflection and other internal packages to extends the functionality of existing data structures and provide access to other common data structures

## What data structures are supported by Structures?

Types will support struct, map, set, linked list, stack, and queue

## Installation (Package)

```bash

$ go get github.com/brentlrayjr/types

```

## Usage

```go

import "github.com/brentlrayjr/types/structs"

type S struct {

  Id int
  Name string

}

s := &S{Id: 1, Name: "Brent"}

//Valid. Returns *field, error
f1, err := types.Struct(s).FieldByName("Name")

//Also valid. Also returns *field, error
f2, err := structs.Struct(s).FieldByIndex(0)

err = f1.Set("B. Ray")

fmt.Printf("Name of struct field at index 0: %s", f2.Name())

```

## Development:

Types is under active development. Alpha, beta, whatever. Most things will remain the same, but nonetheless minor changes will occur.

|   Contributors                              |
|---------------------------------------------|
| [Brent Ray](https://github.com/brentlamarrayjr) |
