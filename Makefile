make pre_commit:
	sh scripts/clean_config.sh
	make remove_protocols
	
clean_config:
	sh scripts/clean_config.sh
	make remove_protocols

init_config:
	make add add_protocols

generate_config:
	go run main.go
	make verify_config

verify_config:
	sh scripts/count_lines.sh config.conf

deploy_config:
	sh scripts/deploy_config.sh $(DEVVPX)

verify_deployment:
	sh scripts/verify_deployment.sh $(DEVVPX)

run_and_deploy:
	make run
	make deploy

run:
	make clean_config
	make add_protocols
	make generate_config

deploy:
	make deploy_config
	make verify_deployment

add_protocols:
	bash scripts/templates/contentswitching/ipfilter.sh 11.0 fake
	bash scripts/templates/contentswitching/ipfilter_blocklist.sh 11.0 fake
	
	sleep 2
	sh scripts/add_protocol_ipfilter.sh 11.0 core http http
	sh scripts/add_protocol_ipfilter.sh 11.0 core ssl http
	sh scripts/add_protocol_ipfilter.sh 11.0 core tcp tcp
	sh scripts/add_protocol_ipfilter.sh 11.0 core ssltcp tcp
	sh scripts/add_protocol_ipfilter.sh 11.0 core udp udp

	sh scripts/add_protocol_ipfilter.sh 11.0 contentswitching http http
	sh scripts/add_protocol_ipfilter.sh 11.0 contentswitching ssl http
	sh scripts/add_protocol_ipfilter.sh 11.0 contentswitching tcp tcp
	sh scripts/add_protocol_ipfilter.sh 11.0 contentswitching ssltcp tcp
	sh scripts/add_protocol_ipfilter.sh 11.0 contentswitching udp udp

	sh scripts/add_protocol_ipfilter.sh 11.0 loadbalancers http http
	sh scripts/add_protocol_ipfilter.sh 11.0 loadbalancers ssl http
	sh scripts/add_protocol_ipfilter.sh 11.0 loadbalancers tcp tcp
	sh scripts/add_protocol_ipfilter.sh 11.0 loadbalancers ssltcp tcp
	sh scripts/add_protocol_ipfilter.sh 11.0 loadbalancers udp udp

	sh scripts/add_protocol_ipfilter.sh 11.0 responders http http
	sh scripts/add_protocol_ipfilter.sh 11.0 responders ssl http
	sh scripts/add_protocol_ipfilter.sh 11.0 responders tcp othertcp
	sh scripts/add_protocol_ipfilter.sh 11.0 responders ssltcp othertcp


remove_protocols:
	sh scripts/remove_protocol_ipfilter.sh 11.0 http
	sh scripts/remove_protocol_ipfilter.sh 11.0 ssl
	sh scripts/remove_protocol_ipfilter.sh 11.0 tcp
	sh scripts/remove_protocol_ipfilter.sh 11.0 ssltcp
	sh scripts/remove_protocol_ipfilter.sh 11.0 udp	