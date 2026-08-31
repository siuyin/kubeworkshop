# Gateway API

## Check if Gateway API is already installed
```
kubectl get crd gateways.gateway.networking.k8s.io -o jsonpath='{.metadata.annotations.gateway\.networking\.k8s\.io/bundle-version}'
```
If not install proceed to install below.

### Install API
```
kubectl apply --server-side -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.1.0/standard-install.yaml
```

## Check if any Gateway Classes are available
```
$ kubectl get gatewayclasses
NAME      CONTROLLER                      ACCEPTED   AGE
traefik   traefik.io/gateway-controller   True       27s
```

If no classes are returned, proceed to define a gateway class below.

### Define Gateway Class for traefic (used in k3s)
```
/var/lib/rancher/k3s/server/manifests/traefik-config.yaml:

apiVersion: helm.cattle.io/v1
kind: HelmChartConfig
metadata:
  name: traefik
  namespace: kube-system
spec:
  valuesContent: |-
    providers:
      kubernetesGateway:
        enabled: true
```

## Define gateway and route for a service
For example, see `podinfo_gateway_and_route.yaml` and `podinfo.yaml`.

### Apply gateway and route
```
$ kubectl apply -f podinfo_gateway_and_route.yaml 
gateway.gateway.networking.k8s.io/podinfo created
httproute.gateway.networking.k8s.io/podinfo created

$ kubectl -ngeekshacking get gateway,httproute,svc
NAME                                        CLASS     ADDRESS   PROGRAMMED   AGE
gateway.gateway.networking.k8s.io/podinfo   traefix             Unknown      5m9s

NAME                                          HOSTNAMES           AGE
httproute.gateway.networking.k8s.io/podinfo   ["podinfo.local"]   5m9s

NAME              TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
service/podinfo   ClusterIP   10.43.151.75   <none>        80/TCP    15h

$ curl -H "Host: podinfo.local" http://10.43.151.75/
{
  "hostname": "podinfo-78679b6fc7-cvz54",
  "version": "6.14.1",
  "revision": "eec06d1ea459af4cb4e10e806f8be7c7bd58b361",
  "color": "#cc0099",
  "logo": "https://raw.githubusercontent.com/stefanprodan/podinfo/gh-pages/cuddle_clap.gif",
  "message": "Gerbau Terpau",
  "goos": "linux",
  "goarch": "amd64",
  "runtime": "go1.26.5",
  "num_goroutine": "6",
  "num_cpu": "8"
}
```
You could also point your browser to `http://10.43.151.75/`
