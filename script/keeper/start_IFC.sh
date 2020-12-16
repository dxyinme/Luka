#!/bin/bash
rm -rf nohup.out
nohup ./KeeperDeployment --alsologtostderr --IFC --ClusterFile=conf/cluster.conf &