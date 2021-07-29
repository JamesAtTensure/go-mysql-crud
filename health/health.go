package health

import (
    "github.com/alexliesenfeld/health"
    "fmt"
    "github.com/alexliesenfeld/health/interceptors"
    "time"
    "context"
    "log"
    "net/http"
    "github.com/alexliesenfeld/health/middleware"


)
func HealthChecker() http.HandlerFunc {
    checker := health.NewChecker(
// Configure a global timeout that will be applied to all checks.
        health.WithTimeout(10),

        // Set for how long check responses are cached.
        health.WithDisabledCache(),
        health.WithCacheDuration(2*time.Second),

        // Cut error message length.
        health.WithMaxErrorMessageLength(500),

        // Configure a global timeout that will be applied to all checks.
        health.WithTimeout(10*time.Second),

        // Disable the details in the health check results.
        // This will configure the checker to only return the aggregated
        // health status but no component details.
        health.WithDisabledDetails(),

        // Disable automatic Checker start. By default, the Checker is started automatically.
        // This configuration option disables this behaviour, so you can delay startup. You need to start
        // the Checker explicitly though see Checker.Start).
        health.WithDisabledAutostart(),

        // This listener will be called whenever system health status changes (e.g., from "up" to "down").
        health.WithStatusListener(onSystemStatusChanged),

        // A list of interceptors that will be applied to every check function. Interceptors may intercept the function
        // call and do some pre- and post-processing, having the check state and check function result at hand.
        health.WithInterceptors(interceptors.BasicLogger()),

        // A simple successFunc to see if a fake database connection is up.
        health.WithCheck(health.Check{
            // Each check gets a unique name.
            Name: "database",
            // The check function. Return an error if the service is unavailable.
            Check: func(ctx context.Context) error {
                return nil // no error.
            },
            // A successFunc specific timeout.
            Timeout: 2 * time.Second,
            // A status listener that will be called if status of this component changes.
            StatusListener: onComponentStatusChanged,
            // A check specific interceptor that pre- and post-processes each call to the check function.
            Interceptors: []health.Interceptor{interceptors.BasicLogger()},
            // The check is allowed to fail up to 5 times in a row
            // until considered unavailable.
            MaxContiguousFails: 5,
            // Check is allowed to stay for up to 1 minute in an error
            // state until considered unavailable.
            MaxTimeInError: 1 * time.Minute,
        }),
    )
    checker.Start()
    handler := health.NewHandler(checker,
        // A result writer writes a check result into an HTTP response.
        // JSONResultWriter is used by default.
        health.WithResultWriter(health.NewJSONResultWriter()),

        // A list of middlewares to pre- and post-process health check requests.
        health.WithMiddleware(
            middleware.BasicLogger(),                   // This middleware will log incoming requests.
            middleware.BasicAuth("user", "password"),   // Removes check details based on basic auth.
            middleware.FullDetailsOnQueryParam("full"), // Disables health details unless request contains query param "full".
        ),

        // Set a custom HTTP status code that should be used if the system is considered "up".
        health.WithStatusCodeUp(200),

        // Set a custom HTTP status code that should be used if the system is considered "down".
        health.WithStatusCodeDown(503),
    )
    return handler
}

func onComponentStatusChanged(_ context.Context, name string, state health.CheckState) {
    log.Println(fmt.Sprintf("component %s changed status to %s", name, state.Status))
}

func onSystemStatusChanged(_ context.Context, state health.CheckerState) {
    log.Println(fmt.Sprintf("system status changed to %s", state.Status))
}

