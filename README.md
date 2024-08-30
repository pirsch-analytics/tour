# Tour of Pirsch Analytics

Take the tour of Pirsch Analytics and learn how to implement and interpret your web analytics data!

[pirsch.io/tour](https://pirsch.io/tour)

Please note that the code in this repository is for demonstration purposes only! The server implemented here should not be used in production.

## Running the Demo Locally

Running the demo locally is easy! You only need to have [Go](https://go.dev) installed. After you've installed it, you can clone the repository and start the server.

```
make run
```

The server accepts the configuration path as parameter in case you would like to build and run it afterward.

## Building for Production

The repository provides a Docker image that can be used for production. It exposes port 8080. Everything else is static. The version number needs to be provided.

```
VERSION=<version_number> make build
```

The server can then be started using Docker:

```
docker run -p 8080:8080 ghcr.io/pirsch-analytics/tour:<version_number>
```

To run it locally or for your own uses, modify the image path in the `Makefile`.
