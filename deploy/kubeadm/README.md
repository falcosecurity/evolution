# Example Kubeadm configuration for Kubernetes dynamic auditing

This Kubeadm set the related required API Server in order to:
- enable the dynamic audit configuration
- enable the `DynamicAuditing` feature gate
- enable the `auditregistration.k8s.io/v1alpha1` API

Anyway, these `kube-apiserver` flags are required whether or not you set up Kubernetes with `kubeadm` in order to [audit Kubernetes Events](https://kubernetes.io/docs/tasks/debug-application-cluster/audit/) and send them to the Falco dynamic Audit Backend.
