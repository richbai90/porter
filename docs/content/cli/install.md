---
title: "porter install"
slug: porter_install
url: /cli/porter_install/
---
## porter install

Create a new installation of a bundle

### Synopsis

Create a new installation of a bundle.

The first argument is the name of the installation to create. This defaults to the name of the bundle. 

Porter uses the Docker driver as the default runtime for executing a bundle's invocation image, but an alternate driver may be supplied via '--driver/-d'.
For example, the 'debug' driver may be specified, which simply logs the info given to it and then exits.

```
porter install [INSTALLATION] [flags]
```

### Examples

```
  porter install
  porter install MyAppFromReference --reference getporter/kubernetes:v0.1.0
  porter install --reference localhost:5000/getporter/kubernetes:v0.1.0 --insecure-registry --force
  porter install MyAppInDev --file myapp/bundle.json
  porter install --parameter-set azure --param test-mode=true --param header-color=blue
  porter install --cred azure --cred kubernetes
  porter install --driver debug

```

### Options

```
      --allow-docker-host-access    Controls if the bundle should have access to the host's Docker daemon with elevated privileges. See https://porter.sh/configuration/#allow-docker-host-access for the full implications of this flag.
      --cnab-file string            Path to the CNAB bundle.json file.
  -c, --cred stringArray            Credential to use when installing the bundle. May be either a named set of credentials or a filepath, and specified multiple times.
  -d, --driver string               Specify a driver to use. Allowed values: docker, debug (default "docker")
  -f, --file string                 Path to the porter manifest file. Defaults to the bundle in the current directory.
      --force                       Force a fresh pull of the bundle
  -h, --help                        help for install
      --insecure-registry           Don't require TLS for the registry
      --no-logs                     Do not persist the bundle execution logs
      --param stringArray           Define an individual parameter in the form NAME=VALUE. Overrides parameters otherwise set via --parameter-set. May be specified multiple times.
  -p, --parameter-set stringArray   Name of a parameter set file for the bundle. May be either a named set of parameters or a filepath, and specified multiple times.
  -r, --reference string            Use a bundle in an OCI registry specified by the given reference.
```

### Options inherited from parent commands

```
      --debug           Enable debug logging
      --debug-plugins   Enable plugin debug logging
```

### SEE ALSO

* [porter](/cli/porter/)	 - I am porter 👩🏽‍✈️, the friendly neighborhood CNAB authoring tool

