# Phoenix operator sidecar for Falco

The main purpose of this sidecar is to create an invisible connection between Falco
and the Phoenix operator. The one-way flow of information is achieved by modifying 
the affected Pod's annotation whenever a new Falco event occurs. 

## Getting started

### Prerequisites
Have [Helm](https://helm.sh/) and [Kustomize](https://kustomize.io/) installed

### Installing the chart
The following commands will install the latest version of the Falco Helm chart, but with a small twist as the installer
uses Kustomize as a Helm post renderer to add our sidecar to the installation.
```
# first add falco to your local helm repo list and update it
helm repo add falcosecurity https://falcosecurity.github.io/charts
helm repo update
# from the root directory of the falco-phoenix-sidecar repository
cd kustomize
helm install falco-phoenix falcosecurity/falco \
--set falco.grpc.enabled=true --set falco.grpcOutput.enabled=true \
--post-renderer ./kustomize-wrapper.sh
```

### Edit configuration
To edit the default configuration of the sidecar head into the `kustomize` directory.
#### Event list
By default, the annotating process will be triggered only for a limited number of Falco events
listed in the `config/falco-phoenix-sidecar.conf` file. If you want to 'listen' on other events you can
add or delete events before you install the chart.

#### Annotation key
By default, the key of the annotation that is added to the affected Pod is `phoenix.r6security.com/falcoevent`.
It can be controlled by modifying the `ANNOTATION_KEY` environment variable's value in `patches/daemonset.json`

#### Maximum number of events stored in the Pod's annotations list
By default, the maximum number of events stored in the affected Pod's annotation list is `5`.
If the limit is reached, the event with the earliest timestamp is removed.
It can be controlled by modifying the `MAXIMUM_NUMBER_OF_EVENTS` environment variable's value in `patches/daemonset.json`


### Cleanup

```
helm uninstall falco-phoenix
```

### Test functionality
To test the functionality you can easily trigger `Read sensitive file untrusted` Falco event by
``kubectl exec -it <example-pod-name> -- cat /etc/shadow``. By reading the sidecar's log you can see as
the desired event occurs and being added to the `<example-pod-name>`'s annotations. 