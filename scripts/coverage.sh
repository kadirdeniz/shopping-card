#!/bin/bash

total_coverage=$(go tool cover -func=coverage.out | grep 'total:' | awk '{printf "%s", $NF}')

echo "Total Coverage: $total_coverage"