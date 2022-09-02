all:
	elasticsearch &
	redis-server &
	~/./clickhouse server &

mq:
	nohup sh ~/Downloads/rocketmq-all-4.9.4-bin-release/bin/mqnamesrv &
	nohup sh ~/Downloads/rocketmq-all-4.9.4-bin-release/bin/mqbroker -n localhost:9876 autoCreateTopicEnable=true &