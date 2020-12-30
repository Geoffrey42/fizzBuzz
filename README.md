# fizzBuzz

A simple fizz-buzz REST server in Golang (LeBonCoin's technical test).

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Assignment](#assignment)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [First time](#first-time)
  - [Build](#build)
- [Usage](#usage)
- [Swagger](#swagger)
- [Core logic](#core-logic)
- [Contributing](#contributing)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Assignment

The iriginal fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz".

The output would look like this: **"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,..."**.

1. Your goal is to implement a web server that will expose a REST API endpoint that:

    - Accepts five parameters : three integers **int1**, **int2** and **limit**, and two strings **str1** and **str2**.
    - Returns a list of strings with numbers from 1 to **limit**, where: all multiples of **int1** are replaced by **str1**, all multiples of **int2** are replaced by **str2**, all multiples of **int1** and **int2** are replaced by **str1str2**.

2. Add a statistics endpoint allowing users to know what the most frequent request has been.

    This endpoint should:

    - Accept no parameter
    - Return the parameters corresponding to the most used request, as well as the number of hits for this request

3. The server needs to be:

    - Ready for production
    - Easy to maintain by other developers

## Prerequisites

- [make](https://www.gnu.org/software/make/manual/make.html) *version 3.81*
- [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) *version 2.26.2*
- [docker](https://docs.docker.com/get-docker/) *version 20.10.0*
- [docker-compose](https://docs.docker.com/compose/install/) *version 1.27.4*
- [gommit](https://github.com/antham/gommit) (if you want to contribute to the project) *version 2.2.0*
- go-swagger *version 0.25.0*

**Important note**: The specified versions are for information purposes only. They are the versions used to develop the project and not the minimum required to run it.

## Installation

### First time

Clone the project:

```bash
git clone git@github.com:Geoffrey42/fizzBuzz.git
cd fizzBuzz
```

Production branch is **main**, but default is **develop**.

```bash
git checkout main
```

### Build

Fetch every dependencies using ```go get -v ./...```.

Then to build **fizzbuzz-server**, run:

```bash
go install ./cmd/fizzbuzz-server
```

Run the server using any port you want. Feel free to check ```fizzbuzz-server --help```:

```bash
fizzbuzz-server --port=5000
```

## Usage

To perform some tests, you can use this Postman collection:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/4a4f3a8f7e69dc307b88)

## Swagger

The server was built using [go-swagger](https://github.com/go-swagger/go-swagger). It's a Golang implementation of Swagger 2.0 specification. The server source code has been generated from this [swagger.yml](./swagger.yml).

A Swagger UI documentation available at ```http://127.0.0.1:5000/docs#/fizzbuzz/get_api_fizzbuzz``` (IP + port may vary depending on your custom settings).

## Core logic

At the heart of the server lies a simple ```DoFizzBuzz``` function defined as the following:

```go
package fb

import (
    "errors"
    "strconv"
    "strings"
)

const start int64 = 1
const max int64 = 100
const base int = 10

func DoFizzBuzz(int1, int2, limit int64, str1, str2 string) ([]string, error) {
    result := strconv.FormatInt(start, base)
    separator := ","

    if limit < start || limit > max {
        return nil, errors.New(
            "limit must be between" + result + " and " + strconv.FormatInt(max, base))
    }

    for i := 2; i <= limit; i++ {
        if i%int1 == 0 && i%int2 == 0 {
            result += separator + str1 + str2
        } else if i%int1 == 0 {
            result += separator + str1
        } else if i%int2 == 0 {
            result += separator + str2
        } else {
            result += separator + strconv.FormatInt(i, base)
        }
    }

    return strings.Split(result, ","), nil
}
```

**Important note**: limit must be between 1 and 100.

See actual function in [fizzbuzz.go](./fb/fizzbuzz.go)

## Contributing

Pull requests are welcomed.
For more details, please refers to our [contributing file](.github/CONTRIBUTING/contributing.md).

## License

[MIT](./LICENSE)
