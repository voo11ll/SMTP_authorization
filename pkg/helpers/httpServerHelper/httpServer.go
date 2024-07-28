package httpServerHelper

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"auth/auth_back/pkg/globalvars"

	loggerIner "auth/auth_back/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gopkg.in/go-playground/validator.v9"
)

var l = loggerIner.Logger{}

const (
	LimitReaderLimit = 1500000 //1.5 MB
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middleware  func(next http.HandlerFunc) http.HandlerFunc
}

type Routes []Route

func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(AccessControlMiddleware)

	for _, route := range routes {
		var handler http.Handler

		if route.Middleware != nil {
			handler = route.Middleware(route.HandlerFunc)
		} else {
			handler = route.HandlerFunc
		}

		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(cors.Default().Handler(handler))
	}
	return router
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		logStr := fmt.Sprintf("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))

		log.Println(logStr)
		l.LogNotify(logStr, name)
	})
}

func ReturnErr(w http.ResponseWriter, err error, enum string) {
	w.WriteHeader(int(globalvars.Error))

	json.NewEncoder(w).Encode(&globalvars.CastomResError{
		StatusENUM: enum,
		Message:    err.Error(),
	})
}

func ReturnErrUnauthorized(w http.ResponseWriter, err error, enum string) {
	w.WriteHeader(int(globalvars.Unauthorized))

	json.NewEncoder(w).Encode(&globalvars.CastomResError{
		StatusENUM: enum,
		Message:    err.Error(),
	})
}

func AccessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ExtractBody(bodyReader io.ReadCloser, data interface{}) error {
	var validate = validator.New()

	bodyByte, err := ioutil.ReadAll(io.LimitReader(bodyReader, 1048576))

	if err != nil {
		return err
	}

	if err := bodyReader.Close(); err != nil {
		return err
	}

	if err := json.Unmarshal(bodyByte, data); err != nil {
		return err
	}

	errValidate := validate.Struct(data)

	if errValidate != nil {
		return errValidate
	}

	return nil
}
