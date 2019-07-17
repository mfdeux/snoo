# snoo
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmfdeux%2Fsnoo.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmfdeux%2Fsnoo?ref=badge_shield)


A Go client for the [Reddit API](https://github.com/reddit/reddit/wiki/API). This package aims to implement every endpoint exposed according to the [documentation](https://www.reddit.com/dev/api) in a user friendly.

## Installation

Install the package with

`go get -u github.com/mfdeux/snoo`

## Authentication

Many endpoints in the Reddit API require OAuth2 authentication to access. To get started, register an app at https://www.reddit.com/prefs/apps and be sure to note the ID, secret, and redirect URI. These values will be used to construct the Authenticator to generate a client with OAuth access. The following is an example of creating an authenticated client using a manual approach:

```Go
package main

import (
  "fmt"
  "github.com/mfdeux/snoo"
)

func main() {
  // Create a new authenticator with your app's client ID, secret, and redirect URI
  // A random string representing state and a list of requested OAuth scopes are required
  authenticator := reddit.NewAuthenticator("<client_id>", "<client_secret>", "<redirect_uri>",
     "<platform>:<app ID>:<version string> (by /u/<snoo username>)", "<random_string>", snoo.ScopeIdentity)

  // Instruct your user to visit the URL retrieved from GetAuthenticationURL in their web browser
  url := authenticator.GetAuthenticationURL()
  fmt.Printf("Please proceed to %s\n", url)

  // After the user grants permission for your client, they will be redirected to the supplied redirect_uri with a code and state as URL parameters
  // Gather these values from the user on the console
  // Note: this can be automated by having a web server listen on the redirect_uri and parsing the state and code params
  fmt.Print("Enter state: ")
  var state string
  fmt.Scanln(&state)

  fmt.Print("Enter code: ")
  var code string
  fmt.Scanln(&code)

  // Exchange the code for an access token
  token, err := authenticator.GetToken(state, code)

  // Create a new client using the access token and a user agent string to identify your application
  client := authenticator.GetAuthClient(token)
}
```

## Examples

```Go
// Returns a new unauthenticated client for invoking the API
client := snoo.NoAuthClient

// Retrives a listing of default subreddits
client.GetDefaultSubreddits()

// Retrives a listing of hot links for the "news" subreddit
client.GetHotLinks("news")
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmfdeux%2Fsnoo.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmfdeux%2Fsnoo?ref=badge_large)