#!/bin/bash
rm -rf nohup.out
nohup ./KeeperDeployment --alsologtostderr --ClusterFile=conf/cluster.conf &