# dy [![Go Report Card](https://goreportcard.com/badge/github.com/sampointer/dy)](https://goreportcard.com/report/github.com/sampointer/dy) [![CircleCI](https://circleci.com/gh/sampointer/dy.svg?style=shield)](https://circleci.com/gh/sampointer/dy) [![GoDoc](https://godoc.org/github.com/sampointer/dy?status.svg)](https://godoc.org/github.com/sampointer/dy)
Construct YAML from a directory tree

## Description
The entire world seems to think declarative configuration is best represented as YAML. This is especially prevalent in the land of Kubernetes and related tools. Terrible ideas have a tendency to accumulate leading to [awful solutions](https://twitter.com/sam_pointer/status/1182321989895311362) to the wrong problems.

Whilst this tool doesn't pretend to move the mountain it does try to nudge it back in the right direction.

Put simply, `dy` allows one to build a YAML document from a directory tree containing snippets of YAML. The aim is to make the document easier to reason about and maintain.

It is useful everywhere complex YAML configuration is employed: CI pipelines, Cloudformation, Kubernetes, etc. See the [examples](https://github.com/sampointer/dy/tree/master/examples) for inspiration.

## Introducing Divvy Yaml
> **divvy** */ˈdɪvi/* - To share out. *Informal, British* - A foolish or stupid person

`dy` parses a directory tree according to the following rules:

* A directory is a text key
* A file name has contents that are rendered under a key named after the file prefix
* A file name that begins with an underscore is rendered without a key at the current indentation level

Consider the following [example](https://github.com/sampointer/dy/tree/master/examples/k8s_deployment):

```
$ tree k8s_deployment/
k8s_deployment/
├── _header.yaml
├── metadata.yaml
└── spec
    ├── _replicas.yaml
    ├── selector.yaml
    └── template
        ├── metadata
        │   └── labels.yaml
        └── spec
            └── containers.yaml

4 directories, 6 files
```

```
$ dy k8s_deployment/
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:1.7.9
          ports:
          - containerPort: 80
```

```
$ dy k8s_deployment/ | kubectl apply --validate=true --dry-run=true -f -
deployment.apps/nginx-deployment created (dry run)
```

You may pass multiple directories as arguments and they will each be parsed and
emitted as documents in their own right. In this way a single `dy` invocation
can be used to produce a valid multi-document YAML stream.

## Installing
### Homebrew
```
brew tap sampointer/dy
brew install dy
```

### Manually
Download the appropriate package for your distribution from the [releases](https://github.com/sampointer/dy/releases) page.
