#!/bin/bash


check_launch_inprogress() {
    inProgress=true
    while [ $inProgress = true ]
    do
        output=$(sudo docker logs cpx | grep -i -c "CPX started successfully")
        if [ "$output" -eq 0 ]; then
            sleep 1

            continue
        fi
        inProgress=false
    done
}

echo "Launching CPX"
sudo docker run -dt -P --privileged=true -e NS_CPX_LITE=1 -e EULA=yes --name cpx -v $PWD/output:/cpx --ulimit core=-1 quay.io/citrix/citrix-k8s-cpx-ingress:13.1-12.51
check_launch_inprogress
echo "Inject config"
sudo docker exec cpx cli_script.sh "batch -filename /cpx/config.conf -outfile /cpx/output_docker.txt"
sudo docker stop cpx
sudo docker rm cpx
echo "Deployment errors: $(cat $PWD/output/output_docker.txt | grep ERROR | wc -l)"
return 0


