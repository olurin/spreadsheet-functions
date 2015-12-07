# Formulas

[![Go Report Card](http://goreportcard.com/badge/TaperBox/formulas)](http://goreportcard.com/report/TaperBox/formulas) [![Build Status](https://travis-ci.org/TaperBox/formulas.svg?branch=master)](https://travis-ci.org/TaperBox/formulas)
[![Build Status](https://semaphoreci.com/api/v1/projects/2d77940c-e0f6-4f5f-a4c6-f92548f19cff/624394/badge.svg)](https://semaphoreci.com/muyiwaolurin/formulas)

Go implementation of most Spreadsheet | Excel Formulas, (:bulb:) inspired by [excelfunctions.net](http://www.excelfunctions.net), [formula.js](https://github.com/sutoiku/formula.js)

This library is currently under development.

## Current Functions 
- [x] Math and Trigonometry functions
-
	- [x] Basic Numeric Information Abs, Sign, GCD, and LCM
	- [x] Basic Mathematical Operations Sum, Product, Power, Sqrt, Quotient, Mod e.t.c.
	

- [ ] Statistical functions
- [ ] Financial functions
- [ ] Date and time functions
- [ ] Logical Functions
- [ ] Engineering Functions
- [ ] Information Functions
- [ ] Text Functions
- [ ] Web Functions

## Installation

use go get

```
go get github.com/TaperBox/formulas
```

or to update

```
go get -u github.com/TaperBox/formulas
```


## Error

ERROR.TYPE is to test for specific errors and display a relevant message (instead of error values)
when certain error conditions exist.

Error Code Key: 

- 1 = #NULL!
- 2 = #DIV/0!
- 3 = #VALUE!
- 4 = #REF!
- 5 = #NAME?
- 6 = #NUM!
- 7 = #N/A
- 8 = #GETTING_DATA
 ... 

## Usage

Value and Common Error 

```go
import fm "github.com/TaperBox/formulas"

data := []float64{2, 1, 6, 4, 3, 5}

// Percentile 
p, err := fm.Percentile(60, nums...) // 60th Percentile

if err != nil {
	fmt.Println(err.Error())
	// Common Errors could be either

	// "#VALUE! - Occurred because the supplied value of k is non-numeric" 
	// or 
	// "#NUM!   - Occurred because the supplied value of k is less than 0 or greater than 100 or the array is empty

	return 
}

fmt.Println("Percentile: ", p) // Percentile: 4

```

Example

```go
import fm "github.com/TaperBox/formulas"


```

## License
Distributed under MIT [License](../formulas/blob/master/LICENSE), please see license file in code for more details.