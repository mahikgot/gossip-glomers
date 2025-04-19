## 🌀 Gossip Glomers

A series of hands-on **distributed systems challenges** that run on a **simulated, fault-prone network** — built by [Fly.io](https://fly.io) and [Jepsen’s](https://aphyr.com) Kyle Kingsbury.  
**This repository contains my solutions** to these challenges, which test and explore key distributed systems concepts using the [Maelstrom](https://github.com/jepsen-io/maelstrom) simulator.

---

### 🌐 Multi-Node Broadcast

#### First solution (naive):

When a node receives a broadcast call, it iterates through the topology and sends a Broadcast call (blocking) to each node that is not itself nor the sender.

To prevent infinite broadcast loops, the node first checks if it has already received the message by checking its message store.

---

