---
title: "Release Notes"
date: 2023-03-07T14:26:51+01:00
draft: false
---

### 0.3.0

Release with some improvements and stabilisation measuresm

#### Fixes
- missing pgbackrest_restore configmap fixed
- 

#### Software-Versions

- PostgreSQL: 15.1 14.7, 13.9, 12.13, 11.18 and 10.23
- Patroni: 3.0.1
- pgBackRest: 2.44
- OS: Rocky-Linux 9.1 (4.18)

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