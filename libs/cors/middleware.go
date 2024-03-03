package cors

import (
	"net/http"
	"net/url"
	"strings"

	libCors "github.com/rs/cors"
)

type option struct {
	isDev          bool
	allowedDomains []string
}

type Option func(*option)

// LocalEnvironment enable AllowAll CORS for local development
func LocalEnvironment() Option {
	return func(o *option) {
		o.isDev = true
	}
}

// WithAllowedDomains set the domains to be allowed
func WithAllowedDomains(domains ...string) Option {
	return func(o *option) {
		o.allowedDomains = domains
	}
}

// NewHandler generate a cors.Cors object using Option parameter
func NewHandler(opts ...Option) *libCors.Cors {
	df := &option{
		isDev:          false,
		allowedDomains: []string{},
	}
	for _, opt := range opts {
		opt(df)
	}
	if df.isDev {
		return libCors.AllowAll()
	}
	return libCors.New(
		libCors.Options{
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodConnect,
				http.MethodOptions,
				http.MethodTrace,
			},
			AllowCredentials: true,
			AllowedHeaders:   []string{"*"},
			AllowOriginVaryRequestFunc: func(r *http.Request, origin string) (bool, []string) {
				if strings.TrimSpace(origin) == "" {
					return false, []string{}
				}
				parsedOrigin, err := url.Parse(origin)
				if err != nil {
					return false, []string{}
				}

				origin = parsedOrigin.Hostname()
				for _, v := range df.allowedDomains {
					if v == origin {
						return true, []string{"Origin"}
					}
				}

				return false, []string{}
			},
		},
	)
}

// DomainAllowed checks if the domain from host is allowed using the options set in cors.Cors object
func DomainAllowed(c *libCors.Cors, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !c.OriginAllowed(r) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		c.Handler(next)
	})
}
