# dsys-hw5

*Distributed Systems Homework 5*

Uses active replication to hold a one-item auction.

### Run

* Start docker with `docker compose up --build --scale frontend=3 --scale replica=5` (or use any other values instead of `3` and `5`).
* Build client artefact with `make client`
* Start client with `./bin/client`

The TCP proxy will log a message "relay TCP/IP connections on :50050 to frontend:50050" when it is ready to receive client requests. At this point use the commands displayed when starting the client to interact with the system.

### What's actually running

The compose file describes four services, of which the frontend and replica are scalable. Besides the scalable services, there's also a discovery service that facilitates connections between the frontends and replicas. Lastly, there's a TCP proxy that forwards incoming traffic (from clients) to dockers DNS resolver, which in turn will distribute the requests among the frontends in a round-robin fashion.

When the system starts all replicas register themselves with the discovery service, and all frontends subscribe to updates of which replicas exist (and connect directly to them when such an update is received).

The flow of a client request is as follows:
1. A client makes a request
2. It goes to the TCP proxy
3. It gets forwarded to a frontend
4. The frontend sends it to all replicas, in order
5. The response travels back up in the reverse order 

*What is this TCP proxy even?* - A convenience mostly. We could easily do away with it, but then we'd have to pass the port of a frontend when starting a client.

### Logs

Logs from each system component are written to log files with the name of their host containers in the ./logs/ directory.

It's probably easiest to read the logs as they're being echoed by docker. If however that is not an option, some postprocessing can be done to make a combined log with correct container names - just run this simple command `docker ps --all | tail -n+2 | sed -r 's/^([^ ]+).* ([^ ]+)$/s\/\1\/\2\//' | sed -f - logs/* | sort -k3 > combined.log`
