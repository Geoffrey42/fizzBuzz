// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-redis/redis"

	"github.com/Geoffrey42/fizzbuzz/fb"
	"github.com/Geoffrey42/fizzbuzz/models"
	"github.com/Geoffrey42/fizzbuzz/restapi/operations"
	"github.com/Geoffrey42/fizzbuzz/restapi/operations/fizzbuzz"
	"github.com/Geoffrey42/fizzbuzz/restapi/operations/stats"
	"github.com/Geoffrey42/fizzbuzz/utils"
)

const forbiddenChars string = "-"

var client *redis.Client = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_HOSTNAME") + ":6379",
	Password: "",
	DB:       0,
})

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

//go:generate swagger generate server --target ../../fizzBuzz --name Fizzbuzz --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.FizzbuzzAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.FizzbuzzAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.FizzbuzzGetAPIFizzbuzzHandler = fizzbuzz.GetAPIFizzbuzzHandlerFunc(func(params fizzbuzz.GetAPIFizzbuzzParams) middleware.Responder {
		if strings.Contains(params.Str1, forbiddenChars) || strings.Contains(params.Str2, forbiddenChars) {
			errorMessage := models.Error{Code: 422, Message: "Request can't contain any of the following characters: " + forbiddenChars}
			return fizzbuzz.NewGetAPIFizzbuzzUnprocessableEntity().WithPayload(&errorMessage)
		}
		res, _ := fb.DoFizzBuzz(params.Int1, params.Int2, params.Limit, params.Str1, params.Str2)
		return fizzbuzz.NewGetAPIFizzbuzzOK().WithPayload(res)
	})

	api.StatsGetAPIStatsHandler = stats.GetAPIStatsHandlerFunc(func(params stats.GetAPIStatsParams) middleware.Responder {
		ok, err := client.Exists(utils.Key).Result()
		if err != nil {
			errorMessage := models.Error{Code: 500, Message: "Database isn't available: " + err.Error()}
			return stats.NewGetAPIStatsInternalServerError().WithPayload(&errorMessage)
		} else if ok == 0 {
			errorMessage := models.Error{Code: 404, Message: "No stored request can be found."}
			return stats.NewGetAPIStatsNotFound().WithPayload(&errorMessage)
		}
		val, _ := client.ZRevRangeWithScores(utils.Key, 0, -1).Result()

		res := models.Stat{}

		if str, ok := val[0].Member.(string); ok {
			p := strings.Split(str, "-")
			res.Hit = int64(val[0].Score)
			res.Int1, _ = strconv.ParseInt(p[0], 10, 64)
			res.Int2, _ = strconv.ParseInt(p[1], 10, 64)
			res.Limit, _ = strconv.ParseInt(p[2], 10, 64)
			res.Str1 = p[3]
			res.Str2 = p[4]
		}
		return stats.NewGetAPIStatsOK().WithPayload(&res)
	})

	api.AddMiddlewareFor("GET", "/api/fizzbuzz", increaseCounterMiddleware)

	if api.FizzbuzzGetAPIFizzbuzzHandler == nil {
		api.FizzbuzzGetAPIFizzbuzzHandler = fizzbuzz.GetAPIFizzbuzzHandlerFunc(func(params fizzbuzz.GetAPIFizzbuzzParams) middleware.Responder {
			return middleware.NotImplemented("operation fizzbuzz.GetAPIFizzbuzz has not yet been implemented")
		})
	}
	if api.StatsGetAPIStatsHandler == nil {
		api.StatsGetAPIStatsHandler = stats.GetAPIStatsHandlerFunc(func(params stats.GetAPIStatsParams) middleware.Responder {
			return middleware.NotImplemented("operation stats.GetAPIStats has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
