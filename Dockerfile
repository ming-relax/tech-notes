FROM centos:7

COPY CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo

RUN yum makecache && yum update -y

RUN yum install -y man man-pages gcc gcc gcc-c++ make automake gdb wget