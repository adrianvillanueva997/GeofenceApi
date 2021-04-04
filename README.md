# GeofenceAPI

Simple to use API made in Go which allows the users to upload Geofence polygons and check if certain devices are inside or outside of it.

# Installation

## Running locally
The API was developed using Go 1.16, it has not been tested in other versions, but it should probably work. By default the API uses port 3000.

If you already have Go installed just run the following commands to download the dependencies and run the project, otherwise, on GNU/Linux use your package manager program to download Go.

Currently the following libraries/frameworks are used:
- fiberprometheus
- validator
- Fiber
- Orb

```go mod download``` 
``` go run main.go```

## Docker
You can build the image by yourself using the Dockerfile provided in this repository, to do it run:

```docker build -t geofence-api .```

Once the image has been built, create a container with the following command:

```docker run -d -p 3000:3000 -v ./logs:/logs --name geofence-api-container geofence-api```


Otherwise, there is an image ready to use on Dockerhub that you can use along with the docker-compose.yml file provided in the repository:

```docker-compose up -d```

The prebuilt image currently supports x86, arm6, arm7 and arm64 architectures.

# REST API

## Versioning

Currently the API is in the version 1, future features/upgrades will be added under /api/v2/feature

## Supported methods

| HTTP Method   |      Path |
|----------|:-------------:|
| GET | /api/v1/polygon |
| PUT | /api/v1/update  |
| GET | /api/v1/devicestatus |
| GET | /api/v1/health   |
| GET | /api/v1/monitor  |
| GET | /metrics         |

### Polygon

GET /api/v1/polygon

Returns the current loaded polygon, if no polygon was loaded it will return 404.

### Update

PUT /api/v1/update

Returns status 200 if the uploaded polygon is actually a polygon, you can see an input example in the folder /data in the repository. If the polygon uploaded is not a polygon, it will return 400

### Device Status

GET /api/v1/devicestatus

Returns status 200 receives a json with Longitude and Latitude, after the coordinates validation, the requests will be answered with status 200 and a boolean. If the coordinates could not be validated, it will return 400.

Input example:
```json
{
	"longitude": 20.3,
	"latitude": 10.2 
}
```

### Health

GET /api/v1/health

Returns status 200 if the API is running.

### Monitor

GET /api/v1/monitor

Returns status 200 and an array of errors registered while the API is running.


### Monitor

Get /metrics

Returns a dataset of metrics collected by prometheus, some of the metrics collected are: 
- Number of goroutines that currently exist
- Number of bytes allocated and still in use
- Total number of bytes allocated, even if freed
- Number of bytes used for garbage collection system metadata
- Total number of scrapes by HTTP status code
- Maximum amount of virtual memory available in bytes

To see the whole set of collected metrics check /metrics/example.txt
