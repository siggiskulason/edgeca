#!/bin/bash -e
 
EDGECA_PARAM=""

if [ "$#" -gt 1 ]; then
    EDGECA_PARAM="-d $SNAP_DATA"
fi

$SNAP/bin/edgeca $@ $EDGECA_PARAM

