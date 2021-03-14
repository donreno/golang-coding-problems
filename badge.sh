#!/bin/sh

CVPKG=$(go list ./... | grep -v mocks | tr '\n' ',')
go test -coverpkg=${CVPKG} -coverprofile=coverage.out -covermode=count  ./... > /dev/null
total=`go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+'`
if (( $(echo "$total <= 50" | bc -l) ))
then
   COLOR=red
elif (( $(echo "$total < 80" | bc -l) ))
then
   COLOR=orange
else
   COLOR=green
fi

curl "https://img.shields.io/badge/coverage-$total%25-$COLOR" > badge.svg
