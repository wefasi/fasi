# FASI

Serve static files from S3 using Go.

## Why

If you create a site with react, you are probably using server side rendering. It's a node server where you can add your custom logic. If you create a static site with nextjs, gatsby, gohugo, zola... you don't have a server.

I handle multiple static sites that, in general, needs the same config, features and optimizations. FASI let me solve all of this problems once, so I can serve all the sites with the same tool.

## How it works

Fasi is a reverse proxy with extra features, such us cache or releases. You can serve multiple sites using a single FASI server, because it uses the `Host header` in order to know which site to serve.

Features:

- Serve multiple sites with a single binary using the `Host` header.
- Releases: cleanup cache or rollback.
- Local cache for improved performance.

## Usage

### Development

```
go run ./cmd/fasi <host> <release>
```

### Production 

You need to build the project locally:

```
go build ./cmd/fasi
```
