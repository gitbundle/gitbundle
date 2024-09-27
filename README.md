# GitBundle

GitBundle is a modern DevOps service. Which is based on git, built with Golang, Docker, K8s, Svelte, Tailwind ...

The target for GitBundle is to make everything become more efficient and easy for DevOps.

## A Simple Config

```console
GITBUNDLE_APP_NAME="GitBundle"
GITBUNDLE_DEBUG=true
GITBUNDLE_CALLER=true

# Some names are limited because of GitBundle needed
GITBUNDLE_USER_ADMIN_UID=root
GITBUNDLE_USER_ADMIN_EMAIL=root@example.com
# The super admin user password, admin pages are coming soon
GITBUNDLE_USER_ADMIN_PASSWORD=root

GITBUNDLE_GIT_ROOT=/data/gitbundle
GITBUNDLE_BLOBSTORE_BUCKET=/data/gitbundle/blob

# Default database is sqlite3, also supports mysql postgres
GITBUNDLE_DATABASE_DEBUG=false
GITBUNDLE_DATABASE_DRIVER=sqlite3
GITBUNDLE_DATABASE_DATASOURCE=database.sqlite3
```

## Notes

Currently, GitBundle is only support linux/amd64, linux/arm64, other platforms are not available, as we have limited server resources. And a lot of features are still developing or testing.
