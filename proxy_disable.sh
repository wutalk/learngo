#!/bin/bash

echo "disable system proxy"
cp -pf /etc/environment.no-proxy /etc/environment

echo "disable docker proxy"
cp -pf /etc/systemd/system/docker.service.d/http-proxy.conf.noproxy /etc/systemd/system/docker.service.d/http-proxy.conf
systemctl daemon-reload
systemctl restart docker

echo "swith git config to github"
cp -pf /home/owuv/.gitconfig-github /home/owuv/.gitconfig

echo "logout to take effect"

