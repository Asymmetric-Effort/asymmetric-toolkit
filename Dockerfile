# Dockerfile
# (c) 2018 Sam Caldwell.  See LICENSE.txt for more information.
#
FROM ubuntu:20.04 AS base_image
ENV APT_OPTS="--no-install-recommends -y"
RUN apt-get update -y --fix-missing && apt-get upgrade -y

FROM base_image AS toolkit_container
RUN apt-get install ${APT_OPTS} apt-utils curl wget git man openvpn net-tools vim nano \
            unzip hexedit hexdiff hexcurse hexer binwalk ncurses-hexedit python3 python3-pip \
            golang ruby ruby-dev gnupg2 gawk ssh openssl bison g++ gcc make nasm build-essential \
            ufw openvpn zlib1g-dev

FROM base_image AS command_and_control
#
# Run the command and control server to make attack decisions and execute attack strategies.
#

FROM base_image AS observability
#
# Run the observability server to collect metrics.
#

FROM base_image AS attack_dnsenum
RUN apt-get install ${APT_OPTS} dnsenum
#
# Setup dnsenum as CMD.
#

#FROM base_image AS attack_dnsrecon
#RUN apt-get install ${APT_OPTS} dnsrecon
#
# Setup dnsrecon as CMD.
#

#FROM base_image AS attack_nikto
#RUN apt-get install ${APT_OPTS} nikto
#
# Setup nikto as CMD.
#

#FROM base_image AS attack_arping
#RUN apt-get install ${APT_OPTS} arping \
#
# Setup arping as CMD.
#

#FROM base_image AS attack_fping
#RUN apt-get install ${APT_OPTS} fping \
#
# Setup fping as CMD.
#

#FROM base_image AS attack_traceroute
#RUN apt-get install ${APT_OPTS} traceroute \
#
# Setup traceroute as CMD.
#

#FROM base_image AS attack_tcptraceroute
#RUN apt-get install ${APT_OPTS} tcptraceroute \
#
# Setup tcptraceroute as CMD.
#

#FROM base_image AS attack_nmap
#RUN apt-get install ${APT_OPTS} nmap \
#
# Setup nmap as CMD.
#

#FROM base_image AS attack_masscan
#RUN apt-get install ${APT_OPTS} masscan \
#
# Setup nmap as CMD.
#

#
#
#
#FROM base_image AS attack_tools
#RUN apt-get install ${APT_OPTS} netmask wafw00f nbtscan smbmap sqlmap ssldump sslh recon-ng dnsrecon \
#             macchanger dirb cadaver burp mutt swaks dirmngr udptunnel netcat tshark tcpdump \
#             udptunnel netcat aircrack-ng reaver kismet wifite pixiewps hashdeep samdump2 \
#             aircrack-ng reaver kismet wifite pixiewps hashdeep hashcat ncrack ophcrack hydra john \
#             whois
#RUN gem install wpscan
#
#COPY tools/scripts/dirbuster /usr/local/bin/
#RUN cd /usr/local/ && git clone https://github.com/sam-caldwell/dirbuster.git
#
#RUN wget https://raw.githubusercontent.com/rapid7/metasploit-omnibus/master/config/templates/metasploit-framework-wrappers/msfupdate.erb && \
#    mv msfupdate.erb /usr/local/bin/msfinstall && \
#    chmod 755 /usr/local/bin/msfinstall && \
#    msfinstall
#
#COPY LICENSE.txt /etc/LICENSE.txt
#COPY tools/scripts/banner /etc/banner
#COPY tools/scripts/env_startup /usr/local/bin/
#CMD [ "/usr/local/bin/env_startup" ]
##
## ------------OSINT TOOLS------------
##
#FROM attack_tools AS osint_tools
#RUN cd /usr/local && \
#    git clone https://github.com/offensive-security/exploitdb.git && \
#    ln -sf /usr/local/exploitdb/searchsploit /usr/local/bin/searchsploit
#COPY tools/scripts/update_exploitdb /usr/local/bin/
#RUN ( \
#        cd /opt/ && mkdir /opt/wordlists && \
#        git clone https://github.com/Asymmetric-Effort/hacked-passwords \
#    )
#
##RUN pip3 install theharvester
#
## ToDo: Deduplicate word lists and generate as a database (SQL) dump.  More efficient.
#COPY LICENSE.txt /etc/LICENSE.txt
#COPY tools/scripts/banner /etc/banner
#COPY tools/scripts/env_startup /usr/local/bin/
#CMD [ "/usr/local/bin/env_startup" ]
