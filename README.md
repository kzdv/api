# ZDV API

The new ZDV Monolithic API.

## Introduction

This monolithic API is meant to break apart the previous Laravel web application into a backend and frontend.  This backend will provide data to the frontend as well as other web components in the future, ie: IDS.

This API is designed to be run in Kubernetes with a few helper tools. Ideally, we would be using Vault to feed Redis and Database credentials, and external secrets for other information, however, for our purpose this is not necessary so we will be using normal secrets. These secrets can be replaced with External Secrets, Sealed Secrets, etc. in the future or if necessary.

The API follows standard RESTful API designs. The API documentation is accessible through root, and the methods are standard:

- GET
- POST (create)
- PUT (replace)
- PATCH (update)
- DELETE

