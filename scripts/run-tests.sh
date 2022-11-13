#!/usr/bin/env sh

## run tdd tests
cd tdd || { echo "Cannot find path /tdd"; exit 1;}
mvn test
cd .. || { echo "Cannot go back from /tdd"; exit 1;}