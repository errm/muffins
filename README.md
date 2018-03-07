# Muffins

2 years ago I made a patch to Kubernetes - https://github.com/kubernetes/kubernetes/pull/24292

The issue was that if you had setup a httpGet, for a liveness or readiness probe and
needed to set the Host header it didn't work.

This app provides a way to explore the issue, if this Pod restarts due to a failing
liveness probe then it may well be that setting the host header in the probe isn't
working on the version of Kubernetes you are runnning.

```
kubectl apply -f https://raw.githubusercontent.com/errm/muffins/master/muffin.yaml
kubectl describe pod/muffin
kubectl delete pod/muffin
```
