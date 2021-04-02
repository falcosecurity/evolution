# Falco

This directory gives you the required YAML files to stand up Falco on Kubernetes as a Daemon Set. This will result in a Falco Pod being deployed to each node, and thus the ability to monitor any running containers for abnormal behavior.

## Deploying to Kubernetes

```
kubectl apply -k falco
```

## Verifying the installation

In order to test that Falco is working correctly, you can launch a shell in a Pod. You should see a message in your Slack channel (if configured), or in the logs of the Falco pod.

```
$ kubectl get pods
NAME          READY     STATUS    RESTARTS   AGE
falco-74htl   1/1       Running   0          13h
falco-fqz2m   1/1       Running   0          13h
falco-sgjfx   1/1       Running   0          13h

$ kubectl exec -it falco-74htl bash
root@falco-74htl:/# exit

$ kubectl logs falco-74htl
{"output":"17:48:58.590038385: Notice A shell was spawned in a container with an attached terminal (user=root k8s.pod=falco-74htl container=a98c2aa8e670 shell=bash parent=<NA> cmdline=bash  terminal=34816)","priority":"Notice","rule":"Terminal shell in container","time":"2017-12-20T17:48:58.590038385Z", "output_fields": {"container.id":"a98c2aa8e670","evt.time":1513792138590038385,"k8s.pod.name":"falco-74htl","proc.cmdline":"bash ","proc.name":"bash","proc.pname":null,"proc.tty":34816,"user.name":"root"}}
```

Alternatively, you can deploy the [event-generator](https://github.com/falcosecurity/event-generator) deployement to have events automatically generated. Please note that this Deployment will generate a large number of events.

You can follow [these steps](https://github.com/falcosecurity/event-generator#with-kubernetes) in order to deploy it in Kubernetes.

## Falco integration with Slack

If you want to send Falco alerts to a Slack channel, you'll want to modify the `falco.yaml` file to point to your Slack webhook. For more information on getting a webhook URL for your Slack team, refer to the [Slack documentation](https://api.slack.com/incoming-webhooks). Add the below to the bottom of the `falco.yaml` `configMap` key to enable Slack messages.

```
program_output:
  enabled: true
  keep_alive: false
  program: "jq '{text: .output}' | curl -d @- -X POST https://hooks.slack.com/services/see_your_slack_team/apps_settings_for/a_webhook_url"
```

You will also need to enable JSON output. Find the `json_output: false` setting in the `falco.yaml` `configMap` key and change it to read `json_output: true`. Any custom rules for your environment can be added to into the `falco_rules.local.yaml` `configMap` key and they will be picked up by Falco at start time.

# Pod Security Policy for Falco

We create a least privilege PodSecurityPolicy tailored to Falco. The Falco service account is authorized to use it, as configured in `rbac.yaml` manifest.

```
$ kubectl apply -k falco-psp
```

# [DEPRECATED] Kubernetes AuditSink for Falco

**Important**: this implementation for dyanmic auditing in Kubernetes has been removed in Kubernetes 1.19. More info [here](https://groups.google.com/g/kubernetes-sig-auth/c/aV_nXpa5uWU?pli=1). More info on current implementation proposal [here](https://docs.google.com/document/d/16cy_ZD94ooBAvlH-rFOel8RPDWRiGFg4Cz11l4sfEII) and [here](https://docs.google.com/document/d/1MqA-RR_wUrMNMbPB6eDyghn9z3z6CDgKK2lsQcciSE8).

In [falco-k8s-audit-sink](./falco-k8s-audit-sink) directory, there is a `AuditSink` API object configuration template that defines:
- the dynamic audit policy 
- and webhook to route Kubernetes audit events to Falco.

## Deploy

```
kubectl apply -f ./falco-k8s-audit-sink/service.yaml
FALCO_SERVICE_IP=$(kubectl get service falco-k8s-audit -o=jsonpath={.spec.clusterIP}) envsubst < ./falco-k8s-audit-sink/audit-sink.yaml.in > ./falco-k8s-audit-sink/audit-sink.yaml
kubectl apply -f ./falco-k8s-audit-sink/audit-sink.yaml
```

> The example above is not intended to be used in production. To register the webhook using a service reference please see the [Kubernetes Documentation](https://kubernetes.io/docs/tasks/debug-application-cluster/audit/#service-reference) and enable SSL for `webserver` feature in `falco.yaml`.

# Falco Sidekick

```
kubectl apply -k falcosidekick
```

# Falco Exporter

```
kubectl apply -k falco-exporter
```
