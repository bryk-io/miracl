#!/bin/bash
# Remove older build
rm -rf core

# Clone fork repository and setup build contents
git clone https://github.com/bryk-io/miracl-core.git temp
for i in {1..43}
do
   echo "$i" >> temp/go/build.txt
done
echo "0" >> temp/go/build.txt

## Build library and adjust custom path used
cd temp/go || exit
UPSTREAM_VERSION=$(git log --pretty=format:'%H' -n1)
python3 config64.py < build.txt
find . -type f -name '*.go' -exec sed -i '' "s|github.com/miracl/core/go|go.bryk.io/miracl|g" {} +
cp -Rv ./core ../../.

## Copy test files
cp Test* ../../.
cp Bench* ../../.

## Final clean-up
cd ../..
printf "package core\nconst Version=\"%s\"" "${UPSTREAM_VERSION}" > core/info.go
printf "1234\n1234\n1234\n1234\n" > pins.txt
rm -rf temp

## Run tests and remove files
go run ./TestECC.go
go run ./TestBLS.go
go run ./TestHPKE.go
go run ./TestNHS.go
go run ./TestMPIN.go < pins.txt
rm Test* pins.txt

## Run benchmarks and remove files
go run ./BenchtestALL.go
rm BenchtestALL.go
