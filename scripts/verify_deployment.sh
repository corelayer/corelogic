#!/bin/sh

scp nsroot@$1:/var/tmp/output.txt .
echo "Deployment errors:\t$(cat output.txt | grep ERROR | wc -l)"
return 0