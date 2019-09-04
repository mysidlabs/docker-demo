#! /bin/sh

h=`hostname`
c=${1:-100}
for i in `seq 1 ${c}`; do
    echo "Hello world, from" ${h} - ${i} of ${c}
    sleep 3
done

