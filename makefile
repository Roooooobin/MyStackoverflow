all:
	elasticsearch &
	redis-server &
	~/./clickhouse server &
