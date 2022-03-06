FROM centos
RUN mkdir /home/server
ENV PROJECT_PATH=/home/server
WORKDIR $PROJECT_PATH

ADD start.sh $PROJECT_PATH
ADD httpServer $PROJECT_PATH
EXPOSE 80

ENTRYPOINT [ "/bin/bash", "start.sh" ]