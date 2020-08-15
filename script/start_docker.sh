# in linux , start docker [usage] bash start_docker.sh keeper1 0.0.4
docker run -it -d -p 10137:10137 --name $1 dog1889/luka_keeper:$2