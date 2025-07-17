// Package "block regex urls" is a Traefik plugin to block access to certain urls using a list of regex values and return a defined status code.
package traefik_block_regex_urls

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"slices"
)

/**********************************
 *          Define types          *
 **********************************/

type traefik_block_regex_urls struct {
	next          http.Handler
	name          string
	regexps       []*regexp.Regexp
	exactMatch    []string
	silentStartUp bool
	statusCode    int
}

type Config struct {
	Regex         []string `mapstructure:"regex,omitempty"`
	ExactMatch    []string `mapstructure:"exact_match,omitempty"`
	SilentStartUp bool     `mapstructure:"silentStartUp"`
	StatusCode    int      `mapstructure:"statusCode"`
}

/**********************************
 * Define traefik related methods *
 **********************************/

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		SilentStartUp: true,
		StatusCode:    403, // https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/net/http/status.go
	}
}

// New creates a new plugin.
// Returns the configured BlockUrls plugin object.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {

	if !config.SilentStartUp {
		log.Println("Regex list: ", config.Regex)
		log.Println("ExactMatch list: ", config.ExactMatch)
		log.Println("StatusCode: ", config.StatusCode)
	}

	// regular expressions
	regexps := make([]*regexp.Regexp, len(config.Regex))

	for index, regex := range config.Regex {
		compiledRegex, compileError := regexp.Compile(regex)
		if compileError != nil {
			return nil, fmt.Errorf("error compiling regex %q: %w", regex, compileError)
		}

		regexps[index] = compiledRegex
	}

	return &traefik_block_regex_urls{
		next:          next,
		name:          name,
		regexps:       regexps,
		exactMatch:    config.ExactMatch,
		silentStartUp: config.SilentStartUp,
		statusCode:    config.StatusCode,
	}, nil
}

// This method is the middleware called during runtime and handling middleware actions.
func (blockUrls *traefik_block_regex_urls) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	fullUrl := request.Host + request.URL.RequestURI()

	if slices.Contains(blockUrls.exactMatch, fullUrl) {
		log.Printf("URL is blocked (exact match): (%s): middleware=%s", fullUrl, blockUrls.name)
		responseWriter.WriteHeader(blockUrls.statusCode)
		return
	}

	for _, regex := range blockUrls.regexps {
		if regex.MatchString(fullUrl) {
			log.Printf("URL is blocked (regex match): (%s) middleware=%s", fullUrl, blockUrls.name)
			responseWriter.WriteHeader(blockUrls.statusCode)
			return
		}
	}

	blockUrls.next.ServeHTTP(responseWriter, request)
}
