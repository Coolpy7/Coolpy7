FROM alpine
MAINTAINER Coolpy7 <Coolpy.net>
ENV TZ=Asia/Shanghai
ENV LANG=zh_CN.UTF-8

ADD go_build_Coolpy7_go_linux.zip /
RUN unzip /go_build_Coolpy7_go_linux.zip
RUN chmod -R 777 /go_build_Coolpy7_go_linux
ENTRYPOINT ["/go_build_Coolpy7_go_linux"]

#ADD go_build_Coolpy7_ws_go_linux.zip /
#RUN unzip /go_build_Coolpy7_ws_go_linux.zip
#RUN chmod -R 777 /go_build_Coolpy7_ws_go_linux
#ENTRYPOINT ["/go_build_Coolpy7_ws_go_linux"]