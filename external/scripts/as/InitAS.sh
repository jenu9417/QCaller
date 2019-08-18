#! /bin/bash

cd ../../tools/as/
tar -xvf aerospike.tar.gz

mkdir AS_HOME

aerospike-server/bin/aerospike init --home ./AS_HOME

echo "Initialized Aerospike server v3.13"