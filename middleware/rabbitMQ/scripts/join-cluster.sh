#!/bin/sh

# 替代 rabbitmq-server 的默认启动命令，加入集群后再启动 rabbitmq-server
# 仅用于 node2 node3 等非主节点，主节点 node1 不需要加入集群

set -eu

if [ -z "${RABBITMQ_JOIN_CLUSTER:-}" ]; then
	exec docker-entrypoint.sh rabbitmq-server
fi

rabbitmq-server -detached

until rabbitmqctl await_startup >/dev/null 2>&1; do
	sleep 2
done

rabbitmqctl stop_app
rabbitmqctl reset

until rabbitmqctl join_cluster "rabbit@${RABBITMQ_JOIN_CLUSTER}"; do
	sleep 2
done

rabbitmqctl start_app
rabbitmqctl stop

exec docker-entrypoint.sh rabbitmq-server
