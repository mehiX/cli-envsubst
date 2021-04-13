# CLI envsubst
Substitutes environment variables in shell format strings. A simplified GOLANG version of `envsubst`

Implemented as a finite state machine (for practice...)

## Run 

```bash
go get -u ./...
go install ./cmd/envsubst

$(go env GOPATH)/bin/envsubst -h
```