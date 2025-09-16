# screenshotbase-go: Golang Website Screenshot Client

This package is a Golang wrapper for screenshotbase.com that aims to make the usage of the API as easy as possible in your project.

## Usage

Initialize the API with your API Key (get one for free at screenshotbase.com):

```go
import screenshotbase "github.com/everapihq/screenshotbase-go"

func main() {
	screenshotbase.Init("YOUR-API-KEY")
}
```

Afterwards you can make calls to the API like this:

### Status Endpoint

```go
body, err := screenshotbase.Status()
```

### Take Endpoint

```go
img, err := screenshotbase.Take(map[string]string{
	"url": "https://bbc.com",
	"format": "png",
	"full_page": "1",
})
```

Find out more about our endpoints, parameters and response data structure in the docs.

## License

The MIT License (MIT). Please see License File for more information.
