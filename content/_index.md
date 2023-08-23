---
title: "CPO (CYBERTEC-PG-Operator)"
date: 2023-03-07T14:26:51+01:00
draft: false
---
Current Release: 0.3.0 (xx.xx.xxxx) [Release Notes](/documentation/release_notes)

CPO (CYBERTEC PG Operator) allows you to create and run PostgreSQL clusters on Kubernetes. 

The operator reduces your efforts and simplifies the administration of your PostgreSQL clusters so that you can concentrate on other things. 

The following features characterise our operator: 
- Declarative mode of operation
- Takes over all the necessary steps for setting up and managing the PG cluster.
- Integrated backup solution, automatic backups and very easy restore (snapshot & PITR)
- Rolling update procedure for adjustments to the pods and minor updates
- Major upgrade with minimum interruption time
- Reduction of downtime thanks to redundancy, pod anti-affinity, auto-failover and self-healing

CPO is tested on the following platforms: 
- Kubernetes 1.21 - 1.24
- Openshift 4.8 - 4.11
- Rancher
- AWS EKS
- Azure AKS
- Google GKE 

Furthermore, CPO is basically executable on any [CSCF-certified](https://www.cncf.io/certification/software-conformance/) Kubernetes platform.

