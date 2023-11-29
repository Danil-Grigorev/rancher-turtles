---
title: "CAPIProvider"
draft: false
toc: true
---

`turtles-capi.cattle.io/v1alpha1`
Type|Link
----|----
GoDoc|[Common-controller/apis/v1alpha1#CAPIProvider](https://pkg.go.dev/github.com/rancher-sandbox/rancher-turtles/Common-controller/apis/v1alpha1#CAPIProvider)

## Metadata
Property|Value
--------|-----
Scope|Namespaced
Kind|`CAPIProvider`
ListKind|`CAPIProviderList`
Plural|`capiproviders`
Singular|`capiprovider`



## Spec


```yaml
additionalManifests: 
  name: string
  namespace: string
configSecret: 
  name: string
  namespace: string
credentials: 
  rancherCloudCredential: string
deployment: 
  affinity: 
    nodeAffinity: 
      preferredDuringSchedulingIgnoredDuringExecution:
        preference: 
          matchExpressions:
          - key: string
            operator: string
            values:
            - string
          matchFields:
          - key: string
            operator: string
            values:
            - string
        weight: integer
      requiredDuringSchedulingIgnoredDuringExecution: 
        nodeSelectorTerms:
          matchExpressions:
          - key: string
            operator: string
            values:
            - string
          matchFields:
          - key: string
            operator: string
            values:
            - string
    podAffinity: 
      preferredDuringSchedulingIgnoredDuringExecution:
        podAffinityTerm: 
          labelSelector: 
            matchExpressions:
            - key: string
              operator: string
              values:
              - string
            matchLabels: {}
          namespaceSelector: 
            matchExpressions:
            - key: string
              operator: string
              values:
              - string
            matchLabels: {}
          namespaces:
          - string
          topologyKey: string
        weight: integer
      requiredDuringSchedulingIgnoredDuringExecution:
        labelSelector: 
          matchExpressions:
          - key: string
            operator: string
            values:
            - string
          matchLabels: {}
        namespaceSelector: 
          matchExpressions:
          - key: string
            operator: string
            values:
            - string
          matchLabels: {}
        namespaces:
        - string
        topologyKey: string
    podAntiAffinity: 
      preferredDuringSchedulingIgnoredDuringExecution:
        podAffinityTerm: 
          labelSelector: 
            matchExpressions:
            - key: string
              operator: string
              values:
              - string
            matchLabels: {}
          namespaceSelector: 
            matchExpressions:
            - key: string
              operator: string
              values:
              - string
            matchLabels: {}
          namespaces:
          - string
          topologyKey: string
        weight: integer
      requiredDuringSchedulingIgnoredDuringExecution:
        labelSelector: 
          matchExpressions:
          - key: string
            operator: string
            values:
            - string
          matchLabels: {}
        namespaceSelector: 
          matchExpressions:
          - key: string
            operator: string
            values:
            - string
          matchLabels: {}
        namespaces:
        - string
        topologyKey: string
  containers:
    args: {}
    command:
    - string
    env:
    - name: string
      value: string
      valueFrom: 
        configMapKeyRef: 
          key: string
          name: string
          optional: boolean
        fieldRef: 
          apiVersion: string
          fieldPath: string
        resourceFieldRef: 
          containerName: string
          divisor: integer
          resource: string
        secretKeyRef: 
          key: string
          name: string
          optional: boolean
    imageUrl: string
    name: string
    resources: 
      claims:
      - name: string
      limits: {}
      requests: {}
  imagePullSecrets:
  - name: string
  nodeSelector: {}
  replicas: integer
  serviceAccountName: string
  tolerations:
  - effect: string
    key: string
    operator: string
    tolerationSeconds: integer
    value: string
features: 
  clusterResourceSet: boolean
  clusterTopology: boolean
  machinePool: boolean
fetchConfig: 
  selector: 
    matchExpressions:
    - key: string
      operator: string
      values:
      - string
    matchLabels: {}
  url: string
manager: 
  cacheNamespace: string
  controller: 
    cacheSyncTimeout: integer
    groupKindConcurrency: {}
    recoverPanic: boolean
  featureGates: {}
  gracefulShutDown: string
  health: 
    healthProbeBindAddress: string
    livenessEndpointName: string
    readinessEndpointName: string
  leaderElection: 
    leaderElect: boolean
    leaseDuration: string
    renewDeadline: string
    resourceLock: string
    resourceName: string
    resourceNamespace: string
    retryPeriod: string
  maxConcurrentReconciles: integer
  metrics: 
    bindAddress: string
  profilerAddress: string
  syncPeriod: string
  verbosity: integer
  webhook: 
    certDir: string
    host: string
    port: integer
manifestPatches:
- string
name: string
type: string
variables: {}
version: string
```



| Field | Description | Details |
| ----- | ----------- | ------- |
| **additionalManifests**<br/>Optional | **object**<br/>AdditionalManifests is reference to configmap that contains additional manifests that will be applied together with the provider components. The key for storing these manifests has to be `manifests`. The manifests are applied only once when a certain release is installed/upgraded. If namespace is not specified, the namespace of the provider will be used. There is no validation of the yaml content inside the configmap. |  |
| **additionalManifests.name**<br/>Required | **string**<br/>Name defines the name of the configmap. |  |
| **additionalManifests.namespace**<br/>Optional | **string**<br/>Namespace defines the namespace of the configmap. |  |
| **configSecret**<br/>Optional | **object**<br/>ConfigSecret is the object with name and namespace of the Secret providing the configuration variables for the current provider instance, like e.g. credentials. Such configurations will be used when creating or upgrading provider components. The contents of the secret will be treated as immutable. If changes need to be made, a new object can be created and the name should be updated. The contents should be in the form of key:value. This secret must be in the same namespace as the provider. |  |
| **configSecret.name**<br/>Required | **string**<br/>Name defines the name of the secret. |  |
| **configSecret.namespace**<br/>Optional | **string**<br/>Namespace defines the namespace of the secret. |  |
| **credentials**<br/>Optional | **object**<br/>Credentials is the structure holding the credentials to use for the provider. Only one credential type could be set at a time. | {'rancherCloudCredential': 'user-credential'} |
| **credentials.rancherCloudCredential**<br/>Optional | **string**<br/>RancherCloudCredential is the Rancher Cloud Credential name |  |
| **deployment**<br/>Optional | **object**<br/>Deployment defines the properties that can be enabled on the deployment for the provider. |  |
| **deployment.affinity**<br/>Optional | **object**<br/>If specified, the pod's scheduling constraints |  |
| **deployment.affinity.nodeAffinity**<br/>Optional | **object**<br/>Describes node affinity scheduling rules for the pod. |  |
| **deployment.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution**<br/>Optional | **array**<br/>The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred. |  |
| **deployment.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution**<br/>Optional | **object**<br/>If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node. |  |
| **deployment.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms**<br/>Required | **array**<br/>Required. A list of node selector terms. The terms are ORed. |  |
| **deployment.affinity.podAffinity**<br/>Optional | **object**<br/>Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)). |  |
| **deployment.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution**<br/>Optional | **array**<br/>The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred. |  |
| **deployment.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution**<br/>Optional | **array**<br/>If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied. |  |
| **deployment.affinity.podAntiAffinity**<br/>Optional | **object**<br/>Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)). |  |
| **deployment.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution**<br/>Optional | **array**<br/>The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred. |  |
| **deployment.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution**<br/>Optional | **array**<br/>If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied. |  |
| **deployment.containers**<br/>Optional | **array**<br/>List of containers specified in the Deployment |  |
| **deployment.imagePullSecrets**<br/>Optional | **array**<br/>List of image pull secrets specified in the Deployment |  |
| **deployment.nodeSelector**<br/>Optional | **object**<br/>NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |  |
| **deployment.replicas**<br/>Optional | **integer**<br/>Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1. |  |
| **deployment.serviceAccountName**<br/>Optional | **string**<br/>If specified, the pod's service account |  |
| **deployment.tolerations**<br/>Optional | **array**<br/>If specified, the pod's tolerations. |  |
| **features**<br/>Optional | **object**<br/>Features is a collection of features to enable. | {'clusterResourceSet': True, 'clusterTopology': True, 'machinePool': True} |
| **features.clusterResourceSet**<br/>Optional | **boolean**<br/>ClusterResourceSet if set to true will enable the cluster resource set feature. |  |
| **features.clusterTopology**<br/>Optional | **boolean**<br/>ClusterTopology if set to true will enable the clusterclass feature. |  |
| **features.machinePool**<br/>Optional | **boolean**<br/>MachinePool if set to true will enable the machine pool feature. |  |
| **fetchConfig**<br/>Optional | **object**<br/>FetchConfig determines how the operator will fetch the components and metadata for the provider. If nil, the operator will try to fetch components according to default embedded fetch configuration for the given kind and `ObjectMeta.Name`. For example, the infrastructure name `aws` will fetch artifacts from https://github.com/kubernetes-sigs/cluster-api-provider-aws/releases. |  |
| **fetchConfig.selector**<br/>Optional | **object**<br/>Selector to be used for fetching provider’s components and metadata from ConfigMaps stored inside the cluster. Each ConfigMap is expected to contain components and metadata for a specific version only. Note: the name of the ConfigMap should be set to the version or to override this add a label like the following: provider.cluster.x-k8s.io/version=v1.4.3 |  |
| **fetchConfig.selector.matchExpressions**<br/>Optional | **array**<br/>matchExpressions is a list of label selector requirements. The requirements are ANDed. |  |
| **fetchConfig.selector.matchLabels**<br/>Optional | **object**<br/>matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed. |  |
| **fetchConfig.url**<br/>Optional | **string**<br/>URL to be used for fetching the provider’s components and metadata from a remote Github repository. For example, https://github.com/{owner}/{repository}/releases You must set `providerSpec.Version` field for operator to pick up desired version of the release from GitHub. |  |
| **manager**<br/>Optional | **object**<br/>Manager defines the properties that can be enabled on the controller manager for the provider. |  |
| **manager.cacheNamespace**<br/>Optional | **string**<br/>CacheNamespace if specified restricts the manager's cache to watch objects in the desired namespace Defaults to all namespaces <br/> Note: If a namespace is specified, controllers can still Watch for a cluster-scoped resource (e.g Node).  For namespaced resources the cache will only hold objects from the desired namespace. |  |
| **manager.controller**<br/>Optional | **object**<br/>Controller contains global configuration options for controllers registered within this manager. |  |
| **manager.controller.cacheSyncTimeout**<br/>Optional | **integer**<br/>CacheSyncTimeout refers to the time limit set to wait for syncing caches. Defaults to 2 minutes if not set. |  |
| **manager.controller.groupKindConcurrency**<br/>Optional | **object**<br/>GroupKindConcurrency is a map from a Kind to the number of concurrent reconciliation allowed for that controller. <br/> When a controller is registered within this manager using the builder utilities, users have to specify the type the controller reconciles in the For(...) call. If the object's kind passed matches one of the keys in this map, the concurrency for that controller is set to the number specified. <br/> The key is expected to be consistent in form with GroupKind.String(), e.g. ReplicaSet in apps group (regardless of version) would be `ReplicaSet.apps`. |  |
| **manager.controller.recoverPanic**<br/>Optional | **boolean**<br/>RecoverPanic indicates if panics should be recovered. |  |
| **manager.featureGates**<br/>Optional | **object**<br/>FeatureGates define provider specific feature flags that will be passed in as container args to the provider's controller manager. Controller Manager flag is --feature-gates. |  |
| **manager.gracefulShutDown**<br/>Optional | **string**<br/>GracefulShutdownTimeout is the duration given to runnable to stop before the manager actually returns on stop. To disable graceful shutdown, set to time.Duration(0) To use graceful shutdown without timeout, set to a negative duration, e.G. time.Duration(-1) The graceful shutdown is skipped for safety reasons in case the leader election lease is lost. |  |
| **manager.health**<br/>Optional | **object**<br/>Health contains the controller health configuration |  |
| **manager.health.healthProbeBindAddress**<br/>Optional | **string**<br/>HealthProbeBindAddress is the TCP address that the controller should bind to for serving health probes It can be set to "0" or "" to disable serving the health probe. |  |
| **manager.health.livenessEndpointName**<br/>Optional | **string**<br/>LivenessEndpointName, defaults to "healthz" |  |
| **manager.health.readinessEndpointName**<br/>Optional | **string**<br/>ReadinessEndpointName, defaults to "readyz" |  |
| **manager.leaderElection**<br/>Optional | **object**<br/>LeaderElection is the LeaderElection config to be used when configuring the manager.Manager leader election |  |
| **manager.leaderElection.leaderElect**<br/>Required | **boolean**<br/>leaderElect enables a leader election client to gain leadership before executing the main loop. Enable this when running replicated components for high availability. |  |
| **manager.leaderElection.leaseDuration**<br/>Required | **string**<br/>leaseDuration is the duration that non-leader candidates will wait after observing a leadership renewal until attempting to acquire leadership of a led but unrenewed leader slot. This is effectively the maximum duration that a leader can be stopped before it is replaced by another candidate. This is only applicable if leader election is enabled. |  |
| **manager.leaderElection.renewDeadline**<br/>Required | **string**<br/>renewDeadline is the interval between attempts by the acting master to renew a leadership slot before it stops leading. This must be less than or equal to the lease duration. This is only applicable if leader election is enabled. |  |
| **manager.leaderElection.resourceLock**<br/>Required | **string**<br/>resourceLock indicates the resource object type that will be used to lock during leader election cycles. |  |
| **manager.leaderElection.resourceName**<br/>Required | **string**<br/>resourceName indicates the name of resource object that will be used to lock during leader election cycles. |  |
| **manager.leaderElection.resourceNamespace**<br/>Required | **string**<br/>resourceName indicates the namespace of resource object that will be used to lock during leader election cycles. |  |
| **manager.leaderElection.retryPeriod**<br/>Required | **string**<br/>retryPeriod is the duration the clients should wait between attempting acquisition and renewal of a leadership. This is only applicable if leader election is enabled. |  |
| **manager.maxConcurrentReconciles**<br/>Optional | **integer**<br/>MaxConcurrentReconciles is the maximum number of concurrent Reconciles which can be run. |  |
| **manager.metrics**<br/>Optional | **object**<br/>Metrics contains thw controller metrics configuration |  |
| **manager.metrics.bindAddress**<br/>Optional | **string**<br/>BindAddress is the TCP address that the controller should bind to for serving prometheus metrics. It can be set to "0" to disable the metrics serving. |  |
| **manager.profilerAddress**<br/>Optional | **string**<br/>ProfilerAddress defines the bind address to expose the pprof profiler (e.g. localhost:6060). Default empty, meaning the profiler is disabled. Controller Manager flag is --profiler-address. |  |
| **manager.syncPeriod**<br/>Optional | **string**<br/>SyncPeriod determines the minimum frequency at which watched resources are reconciled. A lower period will correct entropy more quickly, but reduce responsiveness to change if there are many watched resources. Change this value only if you know what you are doing. Defaults to 10 hours if unset. there will a 10 percent jitter between the SyncPeriod of all controllers so that all controllers will not send list requests simultaneously. |  |
| **manager.verbosity**<br/>Optional | **integer**<br/>Verbosity set the logs verbosity. Defaults to 1. Controller Manager flag is --verbosity.<br/><br/>Default: 1 |  |
| **manager.webhook**<br/>Optional | **object**<br/>Webhook contains the controllers webhook configuration |  |
| **manager.webhook.certDir**<br/>Optional | **string**<br/>CertDir is the directory that contains the server key and certificate. if not set, webhook server would look up the server key and certificate in {TempDir}/k8s-webhook-server/serving-certs. The server key and certificate must be named tls.key and tls.crt, respectively. |  |
| **manager.webhook.host**<br/>Optional | **string**<br/>Host is the hostname that the webhook server binds to. It is used to set webhook.Server.Host. |  |
| **manager.webhook.port**<br/>Optional | **integer**<br/>Port is the port that the webhook server serves at. It is used to set webhook.Server.Port. |  |
| **manifestPatches**<br/>Optional | **array**<br/>ManifestPatches are applied to rendered provider manifests to customize the provider manifests. Patches are applied in the order they are specified. The `kind` field must match the target object, and if `apiVersion` is specified it will only be applied to matching objects. This should be an inline yaml blob-string https://datatracker.ietf.org/doc/html/rfc7396 |  |
| **name**<br/>Required | **string**<br/>Name is the name of the provider to enable | aws |
| **type**<br/>Required | **string**<br/>Type is the type of the provider to enable | infrastructure |
| **variables**<br/>Optional | **object**<br/>Variables is a map of environment variables to add to the content of the ConfigSecret | {'CLUSTER_TOPOLOGY': 'true', 'EXP_CLUSTER_RESOURCE_SET': 'true', 'EXP_MACHINE_POOL': 'true'} |
| **version**<br/>Optional | **string**<br/>Version indicates the provider version. |  |

## Status

```yaml
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  severity: string
  status: string
  type: string
contract: string
installedVersion: string
observedGeneration: integer
state: string
variables: {}
```

| Field | Description | Details |
| ----- | ----------- | ------- |
| **conditions**<br/>Optional | **array**<br/>Conditions define the current service state of the provider. |  |
| **contract**<br/>Optional | **string**<br/>Contract will contain the core provider contract that the provider is abiding by, like e.g. v1alpha4. |  |
| **installedVersion**<br/>Optional | **string**<br/>InstalledVersion is the version of the provider that is installed. |  |
| **observedGeneration**<br/>Optional | **integer**<br/>ObservedGeneration is the latest generation observed by the controller. |  |
| **state**<br/>Optional | **string**<br/>Indicates the provider status<br/><br/>Default: Pending |  |
| **variables**<br/>Optional | **object**<br/>Variables is a map of environment variables added to the content of the ConfigSecret<br/><br/>Default: {'CLUSTER_TOPOLOGY': 'true', 'EXP_CLUSTER_RESOURCE_SET': 'true', 'EXP_MACHINE_POOL': 'true'} |  |

