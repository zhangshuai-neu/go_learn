#! /bin/bash

if ! [ -e cgroup_test ];then
	mkdir cgroup_test
fi

sudo -A mount -t cgroup -o none,name=cgroup_test cgroup_test ./cgroup_test
ls ./cgroup_test


# 创建子cgroup
sudo -A mkdir cgroup_test/cgroup_1
sudo -A mkdir cgroup_test/cgroup_2
