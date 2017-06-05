package main

const header = `
counter nginx_http_requests_total by vhost, method, code, path, backend 

counter nginx_http_request_time_milliseconds by le, vhost, method, code, path, backend 
counter nginx_http_request_time_milliseconds_sum by vhost, method, code, path, backend
counter nginx_http_request_time_milliseconds_count by vhost, method, code, path, backend

# log_format mtail '$server_name $uri $status $request_time 
#                   [$proxy_upstream_name] $upstream_addr 
#                   $upstream_response_time $upstream_status';

const REQUEST /"(?P<request_method>[A-Z]+) (?P<request_uri>\S+) (?P<http_version>HTTP\/[0-9\.]+)" /
const STATUS /(?P<status>\d{3}) /
const REQ_TIME /(?P<request_seconds>\d+)\.(?P<request_milliseconds>\d+) /
const UPS_NAME /\[(?P<upstream_name>\S+)\] /
const UPS_ADDR /(?P<upstream_addr>[0-9A-Za-z\.\-:]+) /
const UPS_RESP_TIME /(?P<ups_resp_seconds>\d+)\.(?P<ups_resp_milliseconds>\d+) /
const UPS_STATUS /(?P<ups_status>\d{3})/
`

const upstream = `
/^/ +
/({{escRegexp .Host}}) / +
/({{.PathRegexp}}) / + REQUEST + STATUS + REQ_TIME + UPS_NAME + UPS_ADDR + UPS_RESP_TIME + UPS_STATUS + 
/$/ {
  nginx_http_requests_total["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++

  ###
  # HTTP Requests with histogram buckets.
  #
  nginx_http_request_time_milliseconds_count["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  nginx_http_request_time_milliseconds_sum["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr] += $request_seconds * 1000 + $request_milliseconds

  # These statements "fall through", so the histogram is cumulative.  The
  # collecting system can compute the percentile bands by taking the ratio of
  # each bucket value over the final bucket.

  # 5ms bucket.
  $request_seconds * 1000 + $request_milliseconds < 5 {
    nginx_http_request_time_milliseconds["5"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # 10ms bucket.
  $request_seconds * 1000 + $request_milliseconds < 10 {
    nginx_http_request_time_milliseconds["10"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # 25ms bucket.
  $request_seconds * 1000 + $request_milliseconds < 25 {
    nginx_http_request_time_milliseconds["25"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # 50ms bucket.
  $request_seconds * 1000 + $request_milliseconds < 50 {
    nginx_http_request_time_milliseconds["50"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # 100ms bucket.
  $request_seconds * 1000 + $request_milliseconds < 100 {
    nginx_http_request_time_milliseconds["100"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # 250ms bucket.
  $request_seconds * 1000 + $request_milliseconds < 250 {
    nginx_http_request_time_milliseconds["250"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # 500ms bucket.
  $request_seconds * 1000 + $request_milliseconds < 500 {
    nginx_http_request_time_milliseconds["500"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # 1s bucket.
  $request_seconds * 1000 + $request_milliseconds < 1000 {
    nginx_http_request_time_milliseconds["1000"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # 2.5s bucket.
  $request_seconds * 1000 + $request_milliseconds < 2500 {
    nginx_http_request_time_milliseconds["2500"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # 5s bucket.
  $request_seconds * 1000 + $request_milliseconds < 5000 {
    nginx_http_request_time_milliseconds["5000"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # 10s bucket.
  $request_seconds * 1000 + $request_milliseconds < 10000 {
    nginx_http_request_time_milliseconds["10000"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
  }

  # "inf" bucket, also the total number of requests.
  nginx_http_request_time_milliseconds["inf"]["{{.Host}}"][$request_method][$status]["{{.PathName}}"][$upstream_addr]++
}
`
