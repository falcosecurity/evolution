# Example Kubernetes Deployments for Falco

This directory gives you the required YAML files to stand up Falco on Kubernetes only for audit purpose as a Deployment.

## Deploying to Kubernetes

Since v1.8 RBAC has been available in Kubernetes, and running with RBAC enabled is considered the best practice. This directory provides the YAML to create a Service Account for Falco, as well as the ClusterRoles and bindings to grant the appropriate permissions to the Service Account.

```
kubectl create -f ./rbac.yaml
```

We also create a service that allows other services to reach the embedded webserver in falco, which listens on https port 8765:

```
kubectl create -f ./service.yaml
```

The Deployment also relies on a Kubernetes ConfigMap to store the Falco configuration and make the configuration available to the Falco Pods. This allows you to manage custom configuration without rebuilding and redeploying the underlying Pods. In order to create the ConfigMap you'll first need to copy the required configuration from their location in this GitHub repo to the `falco-config/` directory (please note that you will need to create the /falco-config directory). Any modification of the configuration should be performed on these copies rather than the original files.

```
mkdir -p falco-config
cp ./falco.yaml ./falco-config/
cp ./rules/k8s_audit_rules.yaml ./falco-config/
```

If you want to send Falco alerts to a Slack channel, you'll want to modify the `falco.yaml` file to point to your Slack webhook. For more information on getting a webhook URL for your Slack team, refer to the [Slack documentation](https://api.slack.com/incoming-webhooks). Add the below to the bottom of the `falco.yaml` config file you just copied to enable Slack messages.

```
program_output:
  enabled: true
  keep_alive: false
  program: "jq '{text: .output}' | curl -d @- -X POST https://hooks.slack.com/services/see_your_slack_team/apps_settings_for/a_webhook_url"
```

You will also need to enable JSON output. Find the `json_output: false` setting in the `falco.yaml` file and change it to read `json_output: true`. Any custom rules for your environment can be added to into the `falco_rules.local.yaml` file and they will be picked up by Falco at start time. You can now create the ConfigMap in Kubernetes.

```
kubectl create configmap falco-config --from-file=./falco-config
```

Now that we have the requirements for our Deployment in place, we can create our Deployment.

```
kubectl create -f ./deployment.yaml
```

### Deploy AuditSink objects

[audit-sink.yaml.in](./audit-sink.yaml.in), in this directory, is a template audit sink configuration that defines the dynamic audit policy and webhook to route Kubernetes audit events to Falco.

Run the following to fill in the template file with the `ClusterIP` IP address you created with the `falco-service` service above. Although services like `falco-service.default.svc.cluster.local` can not be resolved from the kube-apiserver, the ClusterIPs associated with those services are routable.

```
FALCO_SERVICE_CLUSTERIP=$(kubectl get service falco-k8s-audit -o=jsonpath={.spec.clusterIP}) envsubst < audit-sink.yaml.in > audit-sink.yaml
kubectl create -f audit-sink.yaml
```

> The example above is not intended to be used in production. To register the webhook using a service reference please see the [Kubernetes Documentation](https://kubernetes.io/docs/tasks/debug-application-cluster/audit/#service-reference) and enable SSL for `webserver` feature in `falco.yaml`.

## Verifying the installation

In order to test that Falco is working correctly you can deploy the [event-generator](https://github.com/falcosecurity/event-generator) deployement to have events automatically generated. Please note that this Deployment will generate a large number of events.

You can follow [these steps](https://github.com/falcosecurity/event-generator#with-kubernetes) in order to deploy it in Kubernetes.
