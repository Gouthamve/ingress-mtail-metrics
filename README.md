# ingress mtail metrics


## Running

Deploy a sample service to be exposed via ingress. One has been provided to you at: ```manifests/sample-app```. Make sure you change the ingress domain in the manifest before deploying.

To deploy the ingress with metrics, modify the config in ```manifests/ingress/nginx-mtail-config-gen.yaml``` and run :
```
./hack/ingress/deploy.sh
```
The Prometheus annotations will automatically make prometheus discover this pod and scrape it.

To teardown
```
./hack/ingress/teardown.sh
```

## Architecture

The nginx controller will be writing access logs in the following logformat to /mtail/logs/access.log:
 
log_format metricinfo '$server_name $uri "$request" $status $request_time [$proxy_upstream_name] $upstream_addr $upstream_response_time $upstream_status';
 
Which will be tailed by mtail, which captures and exposes the metrics.
 
An example config for mtail can be viewed here: https://github.com/Gouthamve/ingress-mtail-metrics/blob/master/config-gen/nginx.mtail

## Config
 
We need a way to specify the paths to monitor and once we have that, we can parse the access.log to get the metrics. The following config is proposed for each path of a host, the metrics mentioned below are tracked.
 
Config format:
```
backends:
  - host: ingress.goutham.local
    paths:
      - name: /hello
        regexp: \/hello\S+
 ```
 
Metrics exposed:
counter nginx_http_requests_total by vhost, method, code, path, backend 
 
counter nginx_http_request_time_milliseconds_bucket by le, vhost, method, code, path, backend 
counter nginx_http_request_time_milliseconds_sum by vhost, method, code, path, backend
counter nginx_http_request_time_milliseconds_count by vhost, method, code, path, backend
 
https://gist.github.com/Gouthamve/0882256494a55ee8736b8969a1e1403c
