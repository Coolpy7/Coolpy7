FROM alpine
MAINTAINER coolpy7 <www.coolpy.net>
ENV TZ=Asia/Shanghai
ENV LANG=zh_CN.UTF-8
ADD go_build_Coolpy7_go_linux /
CMD ["/go_build_Coolpy7_go_linux"]