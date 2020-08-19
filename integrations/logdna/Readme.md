falco-logdna
------------

A quick script to dump falco events into [logdna](https://logdna.com/).

No warranty is provided or intended.


## Quickstart

```
# maybe setup a venv
pip install -r requirements.txt
```

Install falco. Configure it to output events over grpc unix socket.


```
# # (as root)
# chmod 777 /var/run/falco.sock 
```

Get a LogDNA injestion api key.

```
python falco-logdna.py -k <mykey> -u 'https://logs.logdna.com/logs/ingest'
```
