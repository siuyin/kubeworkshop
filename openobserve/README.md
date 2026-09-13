# Configuring OpenTelemtry Collector

## Installation
```
docker pull otel/opentelemetry-collector:0.160.0
```

```bash
alias otc='docker run  --rm --name otel -h otel -v $HOME:/h -u $UID --network host -w /h otel/opentelemetry-collector:0.160.0'

otc --version
otc --help
```

## Install a trace,metric,log generator
```
go install github.com/open-telemetry/opentelemetry-collector-contrib/cmd/telemetrygen@latest

```

## Configure the collector
See `otel-external-config.yaml`.

Run it with: `otc --config=file:/h/exp/kubeworkshop/openobserve/otelconfig.yaml`

There is also a otel collector running within kubernetes. See daemon set below.
This was installed with:
```
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm install otel-collector open-telemetry/opentelemetry-collector \
   --set image.repository="otel/opentelemetry-collector-k8s" \
   --set mode=daemonset -f otel-values.yaml
```

### Configure the daemon set
See `openobserve/otel-collector-agent-daemonset.yaml` and `openobserve/otel-k8s-collector.yaml`.

Restart the daemon set with: `k rollout restart daemonset.apps/otel-collector-opentelemetry-collector-agent`

## Generate sample data
```
/h/exp/telemetrygen traces --otlp-insecure --traces 3 
/h/exp/telemetrygen metrics --otlp-insecure --metrics 5
/h/exp/telemetrygen logs --otlp-insecure --logs 3 --body "serpau merbau"
```

## Install OBI
```
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts

helm install obi -n obi --create-namespace open-telemetry/opentelemetry-ebpf-instrumentation

```

## Install OpenObserve (datadog workalike)
```
kubectl apply -f https://raw.githubusercontent.com/zinclabs/openobserve/main/deploy/k8s/statefulset.yaml
```

### Optional: Install garage (S3 object storage)
See: `https://garagehq.deuxfleurs.fr/documentation/cookbook/kubernetes/`
