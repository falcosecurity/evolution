# falco-logdna

# This utility connects to falco over grpc/unix socket then submits events
# into logDNA

# (c) LogDNA
# (c) Falco Authors
# (c) IBM

# LICENSE: Apache2(mostly)
# Portions are MIT and from https://github.com/logdna/python
# Some is pulled from https://github.com/falcosecurity/client-py/


import logging
import json
from logdna import LogDNAHandler

import argparse

import falco

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--output-format", "-o", dest="output_format", default=None, help="output_format")
    parser.add_argument("--logdna-key", "-k", dest="logdna_key", default=None, help="Key for Log DNA Authentication")
    parser.add_argument("--logdna-url", "-u", dest="logdna_url", default="https://logs.logdna.com/logs/ingest", help="URL for Log DNA Loghost url e.g. https://logs.logdna.com/logs/ingest'")
    args = parser.parse_args()

    log = logging.getLogger('logdna')
    log.setLevel(logging.INFO)

    # Probably change these
    options = {
              'hostname': 'pytest',
              'ip': '10.0.1.1',
              'mac': 'C0:FF:EE:C0:FF:EE',
              'url': args.logdna_url
             }

    options['index_meta'] = True
    test = LogDNAHandler(args.logdna_key, options)

    log.addHandler(test)

    c = falco.Client(
        endpoint="unix:///var/run/falco.sock",
        output_format=args.output_format,
    )

    for event in c.sub():
        print(event)

        event_options = {
                "level": event.priority.value,
                "app": "falco",
                "meta": json.loads(event.to_json())
                }
        log.info(event.output, event_options)
