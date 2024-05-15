---
title: "Backup"
date: 2023-12-28T14:26:51+01:00
draft: false
---

Backups are essential for databases. From broken storage to deployments gone wrong, backups often save the day. Starting with pg_dump, which was released in the late 1990s, to the archiving of WAL files (PostgreSQL 8.0 / 2005) and pg_basebackup (PostgreSQL 9.0 / 2010), PostgreSQL already offers built-in options for backups and restores based on logical and physical backups. 

## A short journey into the topic of backups

CPO relies on [pgBackRest](www.pgbackrest.org) as its backup solution, a tried-and-tested tool with extensive backup and restore options.
The backup is based on two elements: 
- Snapshots in the form of physical backups
- WAL archive: Continuous archiving of the WAL files 

### Snapshots

Snapshots represent a snapshot of the database in the form of pyhsical files. This contains all relevant information that PostgreSQL holds in its data folder.
With pgBackRest it is possible to create different types of snapshots: 
- full Snapshot: This captures and saves all files at the time of the backup
- Differential backup: Only captures all files that have been changed since the last full snapshot
- Incremental backup: Only records the files that have been changed since the last backup (of any kind). 

When restoring using differential or incremental snapshots, it is necessary to also use the previous snapshots that provide the basis for the selected snapshot. 

> **_HINT:_**  The choice of snapshot type depends on factors such as the size of the database, the time available for backups and the restore. 

### WAL-Archive

The WAL (Write-Ahead-Log) refers to log files which record all changes to the database data before they are written to the actual files. The basic idea here is to guarantee the consistency and recoverability of the comitted data even in the event of failures. 

PostgreSQL normally cleans up or recycles the WAL files that are no longer required. By using WAL archiving, the WAL files are saved to a different location before this process so that they can be used for various activities in the future. 
These activities include
- Providing the WAL files for replicas to keep them up to date
- Restoring instances that have lost parts of the WAL files in the event of a failure and cannot return to a consistent state without them without losing data
- Point-In-Time-Recovery (PITR): In contrast to snapshots, which map a fixed point in time, WAL files make it possible to jump dynamically to a desired point in time and restore the database to the closest available consistent data point

> **_HINT:_**  WAL archiving is an indispensable tool for data availability, recoverability and the continuous availability of PostgreSQL.

## Backup your Cluster

With pgBackRest, backups can be stored on different types of storage: 
- Block storage (PVC)
- S3 / S3-compatible storage
- Azure blob storage
- GCS

### Backups on PVC

When using block storage, the operator creates an additional pod that acts as a repo host. Based on a TLS connection, the repo host obtains the data for the snapshot from the current primary of the cluster, which is compressed before being sent.
WAL archives are pushed from the primary pod to the repo host.

```
apiVersion: cpo.opensource.cybertec.at/v1
kind: postgresql
metadata:
  name: cluster
  namespace: cpo
spec:
  backup:
    pgbackrest:
      global:
        repo1-retention-full: '7'
        repo1-retention-full-type: count
      image: 'docker.io/cybertecpostgresql/cybertec-pg-container-dev:pgbackrest-16.3-1'
      repos:
        - name: repo1
          schedule:
            full: 30 2 * * *
          storage: pvc
          volume:
            size: 15Gi
            storageClass: default
```

This example creates backups based on a repo host with a daily full snapshot at 2:30 am. In addition, pgBackRest is instructed to keep a maximum of 7 full snapshots. The oldest one is always removed when a new snapshot is created. 

> **_HINT:_**  In addition, further configurations for pgBackRest can be defined in the global object. Information on possible configurations can be found in the p[gBackRest documentation](https://pgbackrest.org/configuration.html)


### Backups on S3
pgBackRest can be used directly with AWS S3 or S3-compatible storage such as MinIO, Cloduian HyperStore or SwiftStack.

```
apiVersion: cpo.opensource.cybertec.at/v1
kind: postgresql
metadata:
  name: cluster
  namespace: cpo
spec:
  backup:
    pgbackrest:
      configuration:
        secret: cluster-s3-credential
      global:
        repo1-path: /cluster/repo1/
        repo1-retention-full: '7'
        repo1-retention-full-type: count
      image: 'docker.io/cybertecpostgresql/cybertec-pg-container-dev:pgbackrest-16.3-1'
      repos:
        - endpoint: 'https://s3-zurich.cyberlink.cloud:443'
          name: repo1
          region: zurich
          resource: cpo-cluster-bucket
          schedule:
            full: 30 2 * * *
            incr: '*/30 * * * *'
          storage: s3
```
This example creates a backup in an S3 bucket. In addition to the above configurations, a secret is also required which contains the access data for the S3 storage. The name of the secret must be stored in the `spec.backup.pgbackrest.configuration.secret` object and the secret must be located in the same namespace as the cluster.
Information required to address the S3 bucket:
- `Endpoint`: S3 api endpoint
- `Region`: Region of the bucket
- `resource`: Name of the bucket

The secret must be defined as follows for the use of S3 storage: 
```
kind: Secret
apiVersion: v1
metadata:
  name: cluster-s3-credential
  namespace: cpo
stringData:
  s3.conf |
    [global]
    repo1-s3-key=YOUR_S3_KEY
    repo1-s3-key-secret=YOUR_S3_KEY_SECRET
```
An [example](https://github.com/cybertec-postgresql/CYBERTEC-operator-tutorials/tree/main/cluster-tutorials/pgbackrest_with_s3) with a sercret generator is also available in the tutorials. Enter your access data in the s3.conf file and transfer the tutorial to your Kubernetes with kubectl apply -k cluster-tutorials/pgbackrest_with_s3/.

## Encrypt your backup client-side
pgBackRest also allows you to encrypt your backups on the client side before uploading them. This is possible with any type of storage and is very easy to activate. 

Firstly, we need to define an encryption key. This must be specified separately for each repo and stored in the same secret that is defined in the `spec.backup.pgbackrest.configuration.secret` object.
```
kind: Secret
apiVersion: v1
metadata:
  name: cluster-s3-credential
  namespace: cpo
stringData:
  s3.conf |
    [global]
    repo1-s3-key=YOUR_S3_KEY
    repo1-s3-key-secret=YOUR_S3_KEY_SECRET
    repo1-cipher-pass=YOUR_ENCRYPTION_KEY
```

We also need to configure the type of encryption for pgBackRest. This is done via the cipher-type parameter, which must also be specified for each repo.  You can find the available values for the parameter [here](https://pgbackrest.org/configuration.html#section-repository/option-repo-cipher-type) 

```
apiVersion: cpo.opensource.cybertec.at/v1
kind: postgresql
metadata:
  name: cluster
  namespace: cpo
spec:
  backup:
    pgbackrest:
      configuration:
        secret: cluster-s3-credential
      global:
        repo1-path: /cluster/repo1/
        repo1-retention-full: '7'
        repo1-retention-full-type: count
        repo1-cipher-type: aes-256-cbc
      image: 'docker.io/cybertecpostgresql/cybertec-pg-container-dev:pgbackrest-16.3-1'
      repos:
        - endpoint: 'https://s3-zurich.cyberlink.cloud:443'
          name: repo1
          region: zurich
          resource: cpo-cluster-bucket
          schedule:
            full: 30 2 * * *
            incr: '*/30 * * * *'
          storage: s3
```
