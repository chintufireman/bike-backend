```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Incoming request:", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
```
Explanation:

func LoggingMiddleware(next http.Handler) http.Handler: This is a function declaration named LoggingMiddleware. It takes an http.Handler named next as a parameter and returns an http.Handler. In Go, a middleware typically takes a handler and returns a new handler.

return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { ... }): This is an anonymous function being returned by LoggingMiddleware. It uses http.HandlerFunc to convert a function with the signature (http.ResponseWriter, *http.Request) into an http.Handler. This anonymous function is the actual middleware.

(w http.ResponseWriter, r *http.Request): This part is the signature of the anonymous function. It specifies two parameters: w is an http.ResponseWriter, which is an interface for writing the response, and r is an *http.Request, which represents the incoming HTTP request.

{ ... }: This is the body of the anonymous function, where the actual logic of the middleware is implemented.

fmt.Println("Incoming request:", r.Method, r.URL.Path): This line prints information about the incoming request. r.Method contains the HTTP method (GET, POST, etc.), and r.URL.Path contains the path of the request.

next.ServeHTTP(w, r): This line calls the ServeHTTP method of the next handler in the chain. It allows the request to continue down the middleware chain or reach the final handler. The w parameter is the http.ResponseWriter, and the r parameter is the http.Request.