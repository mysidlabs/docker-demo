#!/bin/sh

i=0
while true; do
  echo $(hostname) - $(date) - ${i};
  sleep 1
  i=`expr ${i} + 1`
done
