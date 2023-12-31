# icmp_message
`icmp_message` sends and receives custom messages through the [ICMP](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol) protocol.

For more details, please see [https://lgiki.net/post/icmp_message/](https://lgiki.net/post/icmp_message/).

# Screenshot

![icmp_message](https://github.com/LGiki/icmp_message/assets/20807713/38c5e9a9-6d0a-4c51-b17a-e4e43134461d)

# Usage

Before using icmp_message, it is necessary to disable the Linux kernel's automatic response to ICMP Echo requests:

```
# echo "1" > /proc/sys/net/ipv4/icmp_echo_ignore_all
```

The usage of icmp_message is very simple:

```
# icmp_message host
```

Assuming there are two computers, A and B, with IP addresses 192.168.1.1 and 192.168.1.2 respectively. To exchange messages using the ICMP protocol between A and B, you can execute the command `icmp_message 192.168.1.2` on A and `icmp_message 192.168.1.1`on B.


# License
MIT License.
