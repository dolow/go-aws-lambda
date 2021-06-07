# Prerequisites

- go
- zip
- aws CLI
- go-swagger

NOTE: aliasing swagger to docker run causes command reference error in Makefile

# Instructions

## Generate go-swagger server

```
% make swagger-generate
```

## Test

```
% make test
```

# Build

```
% make build
```

## Clean

```
% make clean
```

# Deploy

```
% AWS_PROFILE=xxx make deploy
```
