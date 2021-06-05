# StunIP

A very (very) simple program to connect to a STUN server and get your host's public IP. In fact, it's almost the example provided for the [stun Go library](https://github.com/pion/stun)

Useful to check if you can reach a STUN server and what is the public IP presented to this server.

# Usage

With the default STUN server (stun.l.google.com:19302):

    ./stunip

Or, in a Windows environment:

    .\stunip.exe

Output will be the [XOR Mapped Address](https://datatracker.ietf.org/doc/html/draft-ietf-behave-rfc3489bis-02#section-10.2.12) returned by the STUN server when binding to it. This IP is the public IP address of your host (or your Internet router).

With a custom STUN server:

    ./stunip -server some_stun_server:port

