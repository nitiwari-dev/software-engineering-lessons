#!/usr/bin/env sh

## run tdd tests
cd tdd || { echo "Cannot find path /tdd"; exit 1;}
mvn test || { echo "Test failed for 'tdd' project "; exit 1;}
cd .. || { echo "Cannot go back from /tdd"; exit 1;}

## run dsa tests
cd dsa-kotlin || { echo  "Cannot find path /dsa-kotlin"; exit 1;}
mvn test || { echo "Test failed for 'dsa-kotlin' project"; exit 1;}
cd .. || { echo "Cannot go back from /dsa-kotlin"; exit 1;}