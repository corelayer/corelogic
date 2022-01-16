#!/bin/sh

scp config.txt nsroot@$1:/var/tmp/config.txt
ssh -l nsroot $1 'batch -f /var/tmp/config.txt -outfile /var/tmp/output.txt'

return 0