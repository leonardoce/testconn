# Test connection

This is a simple web application that just tests a PostgreSQL connection. 

It exposes just a few entrypoints:

- `/ping` pings the database
- `/readyz` pings the database
- `/livez` just returns 200

This application must not be used as a reference, but it is just meant to
explain the meaning of Kubernetes probes.
