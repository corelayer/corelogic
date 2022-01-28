#!/bin/sh

scp config.conf nsroot@$1:/var/tmp/config.conf
ssh -l nsroot $1 'batch -f /var/tmp/config.conf -outfile /var/tmp/output.txt'

return 0