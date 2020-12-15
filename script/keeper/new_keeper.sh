#!/bin/bash
mkdir $1
cp KeeperDeployment $1
cp -r conf/ $1
rm -rf $1/start_ICC.sh
touch $1/start_ICC.sh
echo '#!/bin/bash' >> $1/start_ICC.sh
echo 'rm -rf nohup.out' >> $1/start_ICC.sh
echo 'nohup ./KeeperDeployment --alsologtostderr --ICC' '--HostAddr='$2 '--KeeperID='$3 '&' >> $1/start_ICC.sh