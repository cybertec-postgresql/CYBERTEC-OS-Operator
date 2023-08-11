---
title: "Release Notes"
date: 2023-03-07T14:26:51+01:00
draft: false
---

### 0.6.1

Release with fixes

#### Fixes
- Backup-Pod now runs with "best-effort" resource definition
- Der Init-Container für die Wiederherstellung verwendet nun die gleiche Ressource-Definition wie der Datenbank-Container, wenn es keine spezifische Definition im Cluster-Manifest gibt (spec.backup.pgbackrest.resources)

#### Software-Versions

- PostgreSQL: 15.3 14.8, 13.11, 12.15
- Patroni: 3.0.4
- pgBackRest: 2.47
- OS: Rocky-Linux 9.1 (4.18)
</br></br>
___
</br></br>
### 0.6.0

Release with some improvements and stabilisation measuresm

#### Features
- Added [Pod Topology Spread Constraints](https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/)
- Added support for TDE based on the CYBERTEC PostgreSQL Enterprise Images (Licensed Container Suite)

#### Software-Versions

- PostgreSQL: 15.3 14.8, 13.11, 12.15
- Patroni: 3.0.4
- pgBackRest: 2.47
- OS: Rocky-Linux 9.1 (4.18)
</br></br>
___
</br></br>
### 0.5.0

Release with new Software-Updates and some internal Improvements
### Features
- Updated to Zalando Operator 1.9

#### Fixes
- internal Problems with Cronjobs
- updates for some API-Definitions

#### Software-Versions

- PostgreSQL: 15.2 14.7, 13.10, 12.14
- Patroni: 3.0.2
- pgBackRest: 2.45
- OS: Rocky-Linux 9.1 (4.18)
</br></br>
___
</br></br>
### 0.3.0

Release with some improvements and stabilisation measuresm

#### Fixes
- missing pgbackrest_restore configmap fixed

#### Software-Versions

- PostgreSQL: 15.1 14.7, 13.9, 12.13, 11.18 and 10.23
- Patroni: 3.0.1
- pgBackRest: 2.44
- OS: Rocky-Linux 9.1 (4.18)
</br></br>
___
</br></br>
### 0.1.0 
	
Initial Release as a Fork of the Zalando-Operator

#### Features

- Added Support for pgBackRest (PoC-State)
    - Stanza-create and Initial-Backup are executed automatically
    - Schedule automatic updates (Full/Incremental/Differential-Backup)
    - Securely store backups on AWS S3 and S3-compatible storage

#### Software-Versions

- PostgreSQL: 14.6, 13.9, 12.13, 11.18 and 10.23
- Patroni: 2.4.1
- pgBackRest: 2.42
- OS: Rocky-Linux 9.0 (4.18)