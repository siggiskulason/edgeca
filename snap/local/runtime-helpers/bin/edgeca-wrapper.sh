#!/bin/bash -e
 
EDGECA_PARAM=""

if [ "$#" -gt 1 ]; then
    if [[ $@ != *"-d "* ]]; then
        EDGECA_PARAM="-d $SNAP_DATA"
    fi
fi

$SNAP/bin/edgeca $@ $EDGECA_PARAM

