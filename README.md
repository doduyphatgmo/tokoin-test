# Basic search application 
# Introduction
Simple command line application to search the data and return the results in a human readable format
- Where the data exists, values from any related entities MUST be included in the results, i.e.  o Searching organization MUST return its ticket subject and users name o Searching users MUST return his/her assignee ticket subject and submitted ticket subject and his/her organization name o Searching tickets MUST return its assignee name, submitter name, and organization name. 
- The user should be able to search on any field, full value matching is fine (e.g. "mar" won't return "mary"). 
- The user should also be able to search for empty values, e.g. where description is empty. 
# Prerequisites
Before you continue, ensure you meet the following requirements:
- You have installed golang.

# Installation
```bash
$ cd tokoin-test
```
```bash
$ go mod download
```

# Quick Start
```bash
$ go run main.go
```

# Tests
```bash
$ go test ./test
```

# Note
For user convenience, the fields with array type just need input item in the array 