#!/bin/sh

scp nsroot@$1:/var/tmp/output.txt output/.
echo "Deployment errors:\t$(cat output/output.txt | grep ERROR | wc -l)"
return 0