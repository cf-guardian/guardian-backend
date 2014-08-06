package config_builder

/*

Configuration items to be defined for compatibility with warden-linux:

* networkPool - a CIDR from which per-container /30 CIDRs are allocated.
* portPoolStart - start of the ephemeral port range for mapped container ports
* portPoolSize - size of ephemeral port range for mapped container ports
* denyNetworks - CIDR blocks representing networks to blacklist
* allowNetworks - CIDR blocks representing networks to whitelist

Warden protocol operations:

* NetIn - map a host port to a container port
* NetOut - whitelist a network or IP address and optional port.

Questions about network configuration, based on the configuration parameters of warden-linux:

1. networkPool seems to be used to allocate IP addresses on the host and in the container. Why
is it necessary to allocate an IP address in the container given that the container has its own
network namespace? Couldn't it, at least in principle, use the same IP address as on the host?

2. networkPool is a CIDR. Does this provide a limited range of IP addresses? nextIP does not seem
to enforce an upper bound:
https://github.com/cloudfoundry-incubator/warden-linux/blob/master/linux_backend/network/network.go#L73.

3. Why is a /30 CIDR used as the container's IP address? Wouldn't a /32 CIDR be sufficient?

4. If each container has its own IP address on the host, why does "net in" allow distinct
values of host and container ports to be set?

5. If each container has its own IP address on the host, why do containers appear to share a range
of ephemeral ports?

 */
