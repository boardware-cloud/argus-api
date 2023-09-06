# BoardWare uptime api

## Configuration

Example

## Generate model from openapi

```bash
openapi-go-generator -o . -p argusapi
```

## Generate Go SDK

```bash
openapi-generator generate -i openapi.yaml -g go \
  -o ./go-sdk
```

## Generate typescript SDK

```bash
openapi-generator generate -i openapi.yaml -g typescript-fetch -o ./ts-sdk \
   --additional-properties=npmName=@boardware/argus-ts-sdk
```

```
GOPRIVATE=gitea.svc.boardware.com/bwc/uptime-api go get -u -f gitea.svc.boardware.com/bwc/uptime-api
```

```
npm config set registry https://gitea.svc.boardware.com/api/packages/bwc/npm/
npm config set -- '//gitea.svc.boardware.com/api/packages/bwc/npm/:_authToken' "Token"
npm publish
npm config set registry https://registry.npmjs.org/
```

npm version $(npm view athena_core_apis version)
