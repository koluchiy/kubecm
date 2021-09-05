kubecm: a tool for managing kubernetes configs
=========================================================

### Commands

#### Create

Create new config file by merging source config files.

`kubecm create -f <result config file path> <source config path> <source config path>`

For example `kubecm create -f /path/config /path1/config1 /path2/config2` 
creates new config file `/path/config` that contains content of config files
`/path1/config1` and `/path2/config2`.

If source config files have users, clusters or contexts with same names, 
kubecm make it unique. For example if two config files has context with
name `context`, in result config file second `context` will be renamed to 
`context1`.

You can specify aliases for contexts from source config files. 
For example argument `/path/to/config::alias1,alias2` tells kubecm to
rename contexts from config file `/path/to/config` to `alias1` and `alias2`.
Number of aliases must be equal to number of contexts in config file.

#### Add

Add content of new config files to exist config file.

`kubecm add -f <result config file path> <source config path> <source config path>`

For example `kubecm add -f /path/config /path1/config1 /path2/config2` 
adds content of config files `/path1/config1` and `/path2/config2` to
exist config file `/path/config`.

#### Rename

Rename context in exist config file.

`kubecm rename -f <result config file path> <old context name> <new context name>`

For example `kubecm rename -f /path/config name1 name2` rename context 
with name `name1` to `name2`.

#### Rm

Remove context from exist config file.

`kubecm rm -f <result config file path> <context name to remove> <context name to remove>`

For example `kubecm rm -f /path/config name1 name2` remove contexts with names
`name1` and `name2` from exist config file `/path/config`.
