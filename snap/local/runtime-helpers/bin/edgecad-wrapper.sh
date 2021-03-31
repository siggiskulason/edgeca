#!/bin/bash -e

POLICY_OPT=""

policy=$(snapctl get "policy")

if [ ! -z "$policy" ]; then
    POLICY_OPT="--policy $policy" 
fi

TPP_URL=""
TPP_ZONE=""
TPP_TOKEN=""

tppurl=$(snapctl get "tpp.url")
tppzone=$(snapctl get "tpp.zone")
tpptoken=$(snapctl get "tpp.token")


if [[ ! -z "$tppurl" ]] && [[ ! -z "$tppzone" ]] && [[ ! -z "$tpptoken" ]]; then
    TPP="--url $tppurl --zone $tppzone --token $tpptoken" 
fi

$SNAP/bin/edgecad $POLICY_OPT $TPP

