clean_config:
	sh scripts/clean_config.sh
	make remove_protocols
	make add_protocols

generate_config:
	make clean_config
	go run main.go > config.txt
	make verify_config

verify_config:
	sh scripts/count_lines.sh config.txt

deploy_config:
	sh scripts/deploy_config.sh $(DEVVPX)

verify_deployment:
	sh scripts/verify_deployment.sh $(DEVVPX)

run:
	make remove_protocols
	make add_protocols
	make clean_config
	make generate_config
	make deploy_config
	make verify_deployment

regenerate_protocols:
	make remove_protocols
	sleep 2
	make add_protocols

add_protocols:
	sh scripts/add_protocol_ipfilter.sh 11.0 http http
	sh scripts/add_protocol_ipfilter.sh 11.0 ssl http
	sh scripts/add_protocol_ipfilter.sh 11.0 tcp tcp
	sh scripts/add_protocol_ipfilter.sh 11.0 ssltcp tcp
	sh scripts/add_protocol_ipfilter.sh 11.0 udp udp

remove_protocols:
	sh scripts/remove_protocol_ipfilter.sh 11.0 http
	sh scripts/remove_protocol_ipfilter.sh 11.0 ssl
	sh scripts/remove_protocol_ipfilter.sh 11.0 tcp
	sh scripts/remove_protocol_ipfilter.sh 11.0 ssltcp
	sh scripts/remove_protocol_ipfilter.sh 11.0 udp
	