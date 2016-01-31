#!/bin/sh

echo domainfinder をビルドします...
go build -o domainfinder

echo synonyms をビルドします...
cd ../synonyms
go build -o ../domainfinder/lib/synonyms

echo available をビルドします...
cd ../available
go build -o ../domainfinder/lib/available

echo sprinkle をビルドします...
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle

echo coolify をビルドします...
cd ../coolify
go build -o ../domainfinder/lib/coolify

echo domainify をビルドします...
cd ../domainify
go build -o ../domainfinder/lib/domainify

echo 完了
