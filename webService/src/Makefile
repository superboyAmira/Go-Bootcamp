
zip:
	tar -czf binary.tar.gz web/ cmd/day06

unzip:
	tar -xzf binary.tar.gz

test_race_linit: 
	/bin/bash -c 'for i in {1..105}; do curl -X GET http://localhost:8888/page/1; done'

test_basic_auth:
	echo -e "\033[0;32m-----------\n\n With Auth \n\n-----------\033[0m"
	curl -i -u superboy:superboy "http://localhost:8888/admin"
	echo -e "\033[0;32m-----------\n\n Without Auth \n\n-----------\033[0m"
	curl -i http://localhost:8888/admin