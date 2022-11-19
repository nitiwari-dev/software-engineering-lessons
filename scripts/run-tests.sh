#!/usr/bin/env sh

## run tdd tests
cd tdd || { echo "Cannot find path /tdd"; exit 1;}
mvn test
cd .. || { echo "Cannot go back from /tdd"; exit 1;}

## run dsa tests
cd dsa || { echo  "Cannot find path /dsa"; exit 1;}
mvn test
cd .. || { echo "Cannot go back from /dsa"; exit 1;}