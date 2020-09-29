This executable can be invoked with multiple programs to execute. The first one will exec'd into - the rest will be 
forked to backround. Separator is `--`. To represent `--` repeat it 2 times  (`----` gets rewritten to `--`).

## Details
 * forks and execs secondary programs
   * Sets forground to false
   * Sets Setsid to true
 * execs into arguments
 
 ## Example
 ```
 ./launcher /bin/bash -- /bin/touch hello -- /bin/sleep 1000
 ```
 The following will happen, in order:
 
 * Fork and Exec `/bin/touch hello` in background
 * Fork and Exec `/bin/sleep 1000` in background
 * Exec `/bin/bash`