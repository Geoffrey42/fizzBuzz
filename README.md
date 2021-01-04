# fizzBuzz

[![Build Status](https://travis-ci.com/Geoffrey42/fizzBuzz.svg?token=XpPVtxxuZC8HAHhhouZ5&branch=develop)](https://travis-ci.com/Geoffrey42/fizzBuzz)
[![codecov](https://codecov.io/gh/Geoffrey42/fizzBuzz/branch/develop/graph/badge.svg?token=7l15xpFfsz)](https://codecov.io/gh/Geoffrey42/fizzBuzz)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A simple fizz-buzz REST server in Golang (LeBonCoin's technical test).

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Assignment](#assignment)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [First time](#first-time)
  - [Build & Run](#build--run)
- [Usage](#usage)
- [Swagger](#swagger)
- [Core logic](#core-logic)
  - [For the fizz-buzz endpoint](#for-the-fizz-buzz-endpoint)
  - [For the statistics endpoint](#for-the-statistics-endpoint)
- [Monitoring](#monitoring)
- [Contributing](#contributing)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Assignment

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz".

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

To build and run this server, the following must be installed:

- [make](https://www.gnu.org/software/make/manual/make.html) *version 3.81*
- [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) *version 2.26.2*
- [docker](https://docs.docker.com/get-docker/) *version 20.10.0*
- [docker-compose](https://docs.docker.com/compose/install/) *version 1.27.4*
- [gommit](https://github.com/antham/gommit) (**only** if you want to contribute to the project) *version 2.2.0*

⚠**Important note**: The specified versions are for information purposes only. They are the versions used to develop the project and not the minimum required to run it.

## Installation

### First time

Clone the project:

```bash
git clone git@github.com:Geoffrey42/fizzBuzz.git
cd fizzBuzz
```

Production branch is **main**, but default is **develop**. Choose accordingly to your needs.

```bash
git checkout main
```

Create an **.env** file based on this [.env.sample](./.env.sample):

```bash
HTTP_PROXY=       # Your corporate proxy if applicable
HTTPS_PROXY=      # Your corporate proxy if applicable
API_PORT=         # Any available port for your API e.g my-redis
REDIS_HOSTNAME=   # Any string to identify your Redis instance
REDIS_PORT=       # Any available port for your Redis instance
REDIS_EXP_PORT=   # Any available port for redis_exporter
PROMETHEUS_PORT=  # Any available port for Prometheus
GRAFANA_PORT=     # Any available port for Grafana
```

Fill it according to your configuration needs.

### Build & Run

To build and run the server:

```bash
make # for production
```

or

```bash
make dev # for dev (no restart always and log rotation)
```

Depending on your connection, those above commands can take some time.

Once it's done, check that everything is correct by running:

```bash
$ make ps
docker-compose ps
      Name                     Command               State           Ports         
-----------------------------------------------------------------------------------
fb-api              /bin/sh -c ./fizzbuzz-serv ...   Up      0.0.0.0:5000->5000/tcp
fb-grafana          /run.sh                          Up      0.0.0.0:3000->3000/tcp
fb-prometheus       /bin/prometheus --config.f ...   Up      0.0.0.0:9090->9090/tcp
fb-redis            docker-entrypoint.sh redis ...   Up      0.0.0.0:6368->6379/tcp
fb-redis_exporter   ./redis_exporter -redis.ad ...   Up      0.0.0.0:9121->9121/tcpp
```

⚠️**Important note**: On first log in, Grafana default username and password are **admin**.

## Usage

The fizzbuzz server exposes two endpoints:

- **/api/fizzbuzz**: Hitting this endpoint and the server will perform a fizzbuzz according to the specified query parameters (to learn more about the parameters see the Swagger UI info in the [Swagger](#swagger) section).
- **/api/stats**: This endpoint returns the top fizzbuzz request, its query parameters and the number of hits.

To perform some tests on your running configuration you can still use ```curl``` but here is a Postman collection gathering the most important requests with some example:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/4a4f3a8f7e69dc307b88)

You will have to change the API port number (5000 in the Postman collection) by the API_PORT from your **.env** file.

You can also perform some tests on your Swagger UI. See [Swagger](#swagger) section.

## Swagger

The server was built using [go-swagger](https://github.com/go-swagger/go-swagger). It's a Golang implementation of Swagger 2.0 specification. The server source code has been generated from this [swagger.yml](./swagger.yml).

A Swagger UI documentation available at ```http://127.0.0.1:[API_PORT]/docs```.

## Core logic

### For the fizz-buzz endpoint

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
    result := ""
    separator := ""

    if limit < start || limit > max {
        return nil, errors.New(
            "limit must be between" + strconv.FormatInt(start, base) + " and " + strconv.FormatInt(max, base))
    }

    for i := start; i <= limit; i++ {
        if i > start {
            separator = ","
        }
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

⚠**Important note**: limit must be between 1 and 100.

See actual function in [fizzbuzz.go](./fb/fizzbuzz.go)

### For the statistics endpoint

The statistics endpoint relies heavily on Redis Sorted Set data structure. When a request hits the above **/api/fizzbuzz** endpoint:

- A Sorted Set member is created in the form of a string "int1-int2-limit-str1-str2" e.g "3-5-16-fizz-buzz" would be created for the assignment's example.
- Its score is increaseb by one.

A http middleware being a perfect fit to intercept incoming requests, here the snippet in [configure_fizzbuzz.go](./restapi/configure_fizzbuzz.go)

```go
func increaseCounterMiddleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        limit, err := strconv.Atoi(r.URL.Query()["limit"][0])

        if err == nil && limit >= 1 && limit <= 100 {
            member := utils.BuildMemberFromParams(r.URL.Query())
            client.ZIncrBy(utils.Key, 1, member)
        }
        
        handler.ServeHTTP(w, r)
    })
}
```

Then, when a client request the **/api/stats** endpoint, it returns the first element of the go-redis equivalent of [ZREVRANGE](https://redis.io/commands/zrevrange).

Here is the model for the Stat struct returned by **/api/stats** in [stat.go](./models/stat.go):

```go
type Stat struct {

    Hit int64 `json:"hit,omitempty"`
    Int1 int64 `json:"int1,omitempty"`
    Int2 int64 `json:"int2,omitempty"`
    Limit int64 `json:"limit,omitempty"`
    Str1 string `json:"str1,omitempty"`
    Str2 string `json:"str2,omitempty"`
}
```

And here is the http.handler in [configure_fizzbuzz.go](./restapi/configure_fizzbuzz.go):

```go
    api.StatsGetAPIStatsHandler = stats.GetAPIStatsHandlerFunc(func(params stats.GetAPIStatsParams) middleware.Responder {
        ok, err := client.Exists(utils.Key).Result()
        if err != nil {
            errorMessage := models.Error{Code: 500, Message: "Database isn't available: " + err.Error()}
            return stats.NewGetAPIStatsInternalServerError().WithPayload(&errorMessage)
        } else if ok == 0 {
            errorMessage := models.Error{Code: 404, Message: "No stored request can be found."}
            return stats.NewGetAPIStatsNotFound().WithPayload(&errorMessage)
        }
        topRequests, _ := client.ZRevRangeWithScores(statistics.Key, 0, -1).Result()

        topRequest, errorMessage := statistics.GetTopRequestFromList(topRequests)
        if errorMessage != nil {
            return stats.NewGetAPIStatsNotFound().WithPayload(errorMessage)
        }

        return stats.NewGetAPIStatsOK().WithPayload(topRequest)
    })
```

## Monitoring

The redis container is essential to get **/api/stats** endpoint working. Thus, a monitoring stack composed of a [Prometheus](https://prometheus.io/) instance, a [redis_exporter](https://github.com/oliver006/redis_exporter) and a [Grafana](https://grafana.com/) dashboard is available to monitor the database.

Go to your grafana instance (```http://127.0.0.1:[GRAFANA_PORT]```) to see the metrics exposed by redis_exporter. Be aware that on first log-in, both default username and password are **admin**.

Then search for **Redis Dashboard for Prometheus Redis Exporter 1.x** dashboard. See capture below:

![redis_exporter_grafana_dashboard](./assets/fb-grafana.gif)

## Contributing

Pull requests are welcomed.
For more details, please refers to our [contributing file](.github/CONTRIBUTING/contributing.md).

## License

[MIT](./LICENSE)
