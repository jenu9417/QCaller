#! /bin/bash

echo "Initialising Aerospike server v3.13"

cd ../../as/
tar -xvf aerospike.tar.gz

mkdir AS_HOME

aerospike-server/bin/aerospike init --home ./AS_HOME


