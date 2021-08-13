# docker.elastic.co/beats/filebeat:7.14.0
FROM elastic/filebeat:7.14.0 
COPY filebeat.yml /usr/share/filebeat/filebeat.yml
# COPY output.json /usr/share/filebeat/output.json
USER root
# RUN chmod 775 /usr/share/filebeat/filebeat.yml
RUN chown root:filebeat /usr/share/filebeat/filebeat.yml
RUN chmod 777 /usr/share/filebeat/filebeat.yml & ls -l
RUN mkdir /usr/share/filebeat/datalogs
# USER filebeat