import configparser
import datetime
import json
import logging
import falco
from kubernetes import client, config
from kubernetes.client.rest import ApiException
import os


def parse_config():
    logging.info("Parsing configuration")
    parser = configparser.ConfigParser()
    parser.read("conf/falco-phoenix-sidecar.conf")
    values = parser.get("falco-phoenix-sidecar", "events")
    return json.loads(values)


def add_annotation_to_pod(pod_name, namespace, event_name):
    pod_config = api.read_namespaced_pod(
        name=pod_name,
        namespace=namespace
    )
    new_event_list = []
    if pod_config.metadata.annotations is not None:
        if pod_config.metadata.annotations[annotation_key] is not None:
            events_str = pod_config.metadata.annotations[annotation_key]
            new_event_list = json.loads(events_str)
            while len(new_event_list) >= int(max_num_of_events):
                earliest_event = None
                for e in new_event_list:
                    if earliest_event is None:
                        earliest_event = e
                    elif earliest_event['timestamp'] > e['timestamp']:
                        earliest_event = e
                new_event_list.remove(earliest_event)

    new_event_list.append({
        "timestamp": str(datetime.datetime.now()),
        "event": event_name
    })
    body = {
        'metadata': {
            'annotations': {
                annotation_key: json.dumps(new_event_list)
            }
        }
    }
    logging.info("New event item added to annotation with key %s in pod %s", annotation_key, pod_name)
    try:
        api.patch_namespaced_pod(
            name=pod_name,
            namespace=namespace,
            body=body
        )
    except ApiException as e:
        logging.warning("Exception when calling CoreV1Api->patch_namespaced_pod: %s\n", e)


max_num_of_events = os.environ.get('MAXIMUM_NUMBER_OF_EVENTS', "5")
annotation_key = os.environ.get('ANNOTATION_KEY', "phoenix.r6security.com/falcoevent")
logging.basicConfig(level=logging.INFO)
falco_client = falco.Client(endpoint="unix:///var/run/falco/falco.sock",
                            client_crt="/tmp/client.crt",
                            client_key="/tmp/client.key",
                            ca_root="/tmp/ca.crt")
conf = config.load_incluster_config()
api = client.CoreV1Api()

if __name__ == "__main__":
    event_list = parse_config()
    logging.info("Events: " + str(event_list))
    logging.info("Maximum number of events per pod:" + str(max_num_of_events))
    logging.info("Annotation key:" + str(annotation_key))
    for event in falco_client.sub():
        logging.info("#" * 5 + " New event " + "#" * 5)
        logging.info("Time: %s", str(event.time))
        logging.info("Host: %s", str(event.hostname))
        logging.info("Priority: %s", str(event.priority))
        logging.info("Source: %s", str(event.source))
        logging.info("Rule: %s", str(event.rule))
        logging.info("Output: %s", str(event.output)[:100] + "...")

        pod = str(event.output_fields['k8s.pod.name'])
        ns = str(event.output_fields['k8s.ns.name'])
        logging.info("Pod: %s", pod)

        logging.info("Fields:")
        for field in event.output_fields:
            logging.info("\t{}: {}".format(field, event.output_fields[field]))

        if pod != "<NA>":
            if str(event.rule) in event_list:
                add_annotation_to_pod(pod, ns, str(event.rule))
        else:
            logging.info("Event not contained podname")
