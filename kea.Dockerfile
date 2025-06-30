FROM ubuntu
RUN apt-get update
RUN apt-get -y install kea-common kea-dhcp4-server kea-ctrl-agent
ADD --chown=_kea:_kea ./test-data/kea/kea-dhcp4.conf /etc/kea/kea-dhcp4.conf
ADD --chown=_kea:_kea ./test-data/kea/kea-ctrl-agent.conf /etc/kea/kea-ctrl-agent.conf
RUN mkdir -p /var/lib/kea
RUN chown -R _kea. /var/lib/kea
RUN mkdir -p /var/log/kea /run/kea
RUN chown -R _kea. /var/log/kea /run/kea
USER _kea
EXPOSE 8000/tcp
COPY ./entrypoint.kea.sh /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]
