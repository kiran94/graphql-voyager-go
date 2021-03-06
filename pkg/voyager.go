package voyager

import (
	"net/http"
	"text/template"
)

var voyagerTemplate *template.Template

type voyagerOptions struct {
	Protocol string
	Endpoint string
	Host     string
	Headers  string
}

func init() {
	vt, err := template.New("voyager").Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	voyagerTemplate = vt
}

// NewVoyagerHandler Creates a new Voyager Options object
// Example:
//      http.Handle("/voyager", voyagerHandler)
func NewVoyagerHandler(endpoint string) http.Handler {
	v := voyagerOptions{
		Endpoint: endpoint,
		Headers:  "{'Accept': 'application/json', 'Content-Type': 'application/json'}",
		Host:     "window.location.host",
		Protocol: "window.location.protocol",
	}
	return v
}

// Serves the Voyager UI
func (v voyagerOptions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	voyagerTemplate.Execute(w, &v)
}

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
  <style>
    body {
      height: 100%;
      margin: 0;
      width: 100%;
      overflow: hidden;
    }

    #voyager {
      height: 100vh;
    }
  </style>

  <!--
    This GraphQL Voyager example depends on Promise and fetch, which are available in
    modern browsers, but can be "polyfilled" for older browsers.
    GraphQL Voyager itself depends on React DOM.
    If you do not want to rely on a CDN, you can host these files locally or
    include them directly in your favored resource bundler.
  -->
  <script src="https://cdn.jsdelivr.net/es6-promise/4.0.5/es6-promise.auto.min.js"></script>
  <script src="https://cdn.jsdelivr.net/fetch/0.9.0/fetch.min.js"></script>
  <script src="https://cdn.jsdelivr.net/npm/react@16/umd/react.production.min.js"></script>
  <script src="https://cdn.jsdelivr.net/npm/react-dom@16/umd/react-dom.production.min.js"></script>

  <!--
      These two files are served from jsDelivr CDN, however you may wish to
      copy them directly into your environment, or perhaps include them in your
      favored resource bundler.
  -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphql-voyager/dist/voyager.css" />
  <script src="https://cdn.jsdelivr.net/npm/graphql-voyager/dist/voyager.min.js"></script>
</head>
<body>
  <div id="voyager">Loading...</div>
  <script>

    // Defines a GraphQL introspection fetcher using the fetch API. You're not required to
    // use fetch, and could instead implement introspectionProvider however you like,
    // as long as it returns a Promise
    // Voyager passes introspectionQuery as an argument for this function
    function introspectionProvider(introspectionQuery) {
      let url = {{.Protocol}} + "//" + {{.Host}} + "{{.Endpoint}}";
      console.debug(url);

      return fetch(url, {
        method: 'post',
        headers: {{.Headers}},
        body: JSON.stringify({query: introspectionQuery}),
        credentials: 'include',
        mode: 'no-cors',
      }).then(function (response) {
        return response.text();
      }).then(function (responseBody) {
        try {
          console.debug(responseBody);
          return JSON.parse(responseBody);
        } catch (error) {
          console.log(error);
          return responseBody;
        }
      });
    }

    // Render <Voyager /> into the body.
    GraphQLVoyager.init(document.getElementById('voyager'), {
      introspection: introspectionProvider
    });
  </script>
</body>
</html>
`
