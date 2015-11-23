# Formulas

[![Go Report Card](http://goreportcard.com/badge/TaperBox/formulas)](http://goreportcard.com/report/TaperBox/formulas) [![Build Status](https://travis-ci.org/TaperBox/formulas.svg?branch=master)](https://travis-ci.org/TaperBox/formulas)

Go implementation of most Spreadsheet | Excel Formulas, :bulb: by [excelfunctions.net](http://www.excelfunctions.net), [formula.js](https://github.com/sutoiku/formula.js)

This library is currently under development.

## Overview of functions
- [x] [Math and Trigonometry functions](../master/math/README.md)
- [x] [Statistical functions](../master/statistics/README.md)
- [x] [Financial functions](../master/financial/README.md)
- [ ] Date and time functions
- [ ] Logical Functions
- [ ] Engineering Functions
- [ ] Information Functions
- [ ] Text Functions
- [ ] Web Functions

## Installation

use go get. 

```
go get github.com/TaperBox/formulas
```

or to update
go get -u github.com/TaperBox/formulas


## Usage

Value and Common Error 

```go
import fm "github.com/TaperBox/formulas"

data := []float64{2, 1, 6, 4, 3, 5}

// Percentile 
percentile, err := fm.Percentile(60, nums...) // 60th Percentile

if err != nil {
	fmt.Println(err.Error())
	// Common Errors 
	// "#VALUE! - Occurred because the supplied value of k is non-numeric" or 
	// "#NUM!   - Occurred because the supplied value of k is less than 0 or greater than 100 or the array is empty
}

fmt.Println(percentile) // 4

```


## License
Distributed under MIT [License](../formulas/blob/master/LICENSE), please see license file in code for more details.