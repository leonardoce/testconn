# Test connection

This is a simple web application that just tests a PostgreSQL connection. 

It exposes just a few entrypoints:

- `/ping` pings the database
- `/readyz` pings the database
- `/livez` returns 200 if the probe is not failed, otherwise it will raise
  an internal server error
- `/fake` set the liveness probe to be faked

This application must not be used as a reference, but it is just meant to
explain the meaning of Kubernetes probes.

Some example Kubernetes definitions can be found in the `k8s` directory.
