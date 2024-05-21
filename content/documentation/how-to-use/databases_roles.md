---
title: "Databases & Users"
date: 2023-12-28T14:26:51+01:00
draft: false
---

CPO not only supports you in deploying your cluster, it also supports you in setting it up in terms of the database and users. 
CPO offers you three different options for this: 
- Create roles
- Create databases
- preapared databases

## Create Roles
The creation of users is based on the definition of the user name and the definition of the required rights for this user. Available rights are
- `superuser`
- `inherit`
- `login`
- `nologin`
- `createrole`
- `createdb`
- `replication`
- `bypassrls`

Unless explicitly defined via `nologin`, a created user automatically receives the `login` permission. 

```
spec:
  users:
    db_owner:
    - login
    - createdb
    appl_user:
    - login
```

For each user created, CPO automatically creates a secret with `username` and `password` in the namespace of the cluster, which follows the following naming convention: 
[USERNAME].[CLUSTERNAME].credentials.postgresql.cpo.opensource.cybertec.at 

If the secrets for an application are to be stored in a different namespace, for example, it is necessary to define the setting enable_cross_namespace_secret as true in the operator configuration. You can find more information about the operator configuration [here](documentation/how-to-use/operator_configuration/).

The namespace must then be written before the user name.
```
spec:
  users:
    db_owner:
    - login
    - createdb
    app_namespace.appl_user:
    - login
```

## Create Databases 

Databases are basically created in a very similar way to users.
The definition is based on the database name and the database owner. 

```
spec:
  users:
    db_owner:
    - login
    - createdb
    app_namespace.appl_user:
    - login
  databases;
    app_db: app_namespace.appl_user
```

> **_HINT:_**  Be aware that the user name must be defined for the database owner in the same way as it is done in the users object. 

## Prepared Databases

The `preparedDatabases` object is available for a much more extensive setup of databases and users. 
In addition to the creation of `databases` and `users`, this also enables the creation of `schemas` and `extensions`. A more detailed rights management is also available.