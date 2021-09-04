kubecm: a tool for managing kubernetes configs
=========================================================

Command `kubecm create -f result.conf kubeconfig1.conf kubeconfig.conf` 
creates new config `result.conf` that contains content of source configs
`kubeconfig1.conf` and `kubeconfig.conf`.

`kubecm` deduplicate users, clusters and contexts names with rename equals one.
For example, if there are two user with same name `john`, the second one
will be renamed in `john1`.

You can specify aliases for context in `create` command.
For example, command `kubecm create -f result.conf kubeconfig1.conf::ctx1 kubeconfig.conf::ctx2`
will create contexts with names `ctx1` and `ctx2` instead of original contexts names.