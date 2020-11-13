#!/bin/bash
mkdir $1
cp KeeperDeployment $1
cp -r conf/ $1
cp keeper/start.sh $1