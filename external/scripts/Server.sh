#! /bin/bash

cd ../bin
touch app.log
./QCaller & tailf app.log