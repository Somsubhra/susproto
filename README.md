The SUS Protocol
=================


The SUS protocol is a consensus protocol to arrive at a consensus to label certain people as sus along with the sus index.

### Example Usage

1. In order to arrive at a consensus we need to have multiple *gossipers* who elect a lead *gossiper*. To spin up multiple gossipers run the following commands

```shell
 ADDRESS="127.0.0.1:8001" BOOTSTRAP_SERVERS="127.0.0.1:8000" STORE="raft-log-1.bolt" go run ./...
```


```shell
 ADDRESS="127.0.0.1:8000" BOOTSTRAP_SERVERS="127.0.0.1:8001" STORE="raft-log-2.bolt" go run ./...
```

You can spin up as many gossipers as needed.

### Sample output

Gossiper 1

```shell
ADDRESS="127.0.0.1:8001" BOOTSTRAP_SERVERS="127.0.0.1:8000" STORE="raft-log-1.bolt" go run ./...
2024-09-26T12:21:57.891+0530 [INFO]  raft: initial configuration: index=0 servers=[]
2024-09-26T12:21:57.892+0530 [INFO]  raft: entering follower state: follower="Node at 127.0.0.1:8001 [Follower]" leader-address= leader-id=
2024/09/26 12:21:57 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:58 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:58 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:58 I am just the follower, waiting for the leader to come and declare the sus list :P
2024-09-26T12:21:58.611+0530 [DEBUG] raft: received a requestPreVote with a newer term, grant the pre-vote
2024-09-26T12:21:58.621+0530 [DEBUG] raft: lost leadership because received a requestVote with a newer term
2024-09-26T12:22:00.808+0530 [WARN]  raft: heartbeat timeout reached, starting election: last-leader-addr=127.0.0.1:8000 last-leader-id=127.0.0.1:8000
2024-09-26T12:22:00.808+0530 [INFO]  raft: entering candidate state: node="Node at 127.0.0.1:8001 [Candidate]" term=3
2024-09-26T12:22:00.809+0530 [DEBUG] raft: asking for pre-vote: term=3 from=127.0.0.1:8000 address=127.0.0.1:8000
2024-09-26T12:22:00.809+0530 [DEBUG] raft: pre-voting for self: term=3 id=127.0.0.1:8001
2024-09-26T12:22:00.809+0530 [DEBUG] raft: calculated votes needed: needed=2 term=3
2024-09-26T12:22:00.809+0530 [DEBUG] raft: pre-vote received: from=127.0.0.1:8001 term=3 tally=0
2024-09-26T12:22:00.809+0530 [DEBUG] raft: pre-vote granted: from=127.0.0.1:8001 term=3 tally=1
2024-09-26T12:22:00.809+0530 [ERROR] raft: failed to make requestVote RPC: target="{Voter 127.0.0.1:8000 127.0.0.1:8000}" error="dial tcp 127.0.0.1:8000: connect: connection refused" term=3
2024-09-26T12:22:00.809+0530 [DEBUG] raft: pre-vote received: from=127.0.0.1:8000 term=3 tally=1
2024-09-26T12:22:00.809+0530 [DEBUG] raft: pre-vote denied: from=127.0.0.1:8000 term=3 tally=1
```


Gossiper 2

```shell
ADDRESS="127.0.0.1:8000" BOOTSTRAP_SERVERS="127.0.0.1:8001" STORE="raft-log-2.bolt" go run ./...
2024-09-26T12:21:56.667+0530 [INFO]  raft: initial configuration: index=0 servers=[]
2024-09-26T12:21:56.669+0530 [INFO]  raft: entering follower state: follower="Node at 127.0.0.1:8000 [Follower]" leader-address= leader-id=
2024/09/26 12:21:56 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:56 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:57 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:57 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:57 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:57 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:57 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:58 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:58 I am just the follower, waiting for the leader to come and declare the sus list :P
2024/09/26 12:21:58 I am just the follower, waiting for the leader to come and declare the sus list :P
2024-09-26T12:21:58.610+0530 [WARN]  raft: heartbeat timeout reached, starting election: last-leader-addr= last-leader-id=
2024-09-26T12:21:58.610+0530 [INFO]  raft: entering candidate state: node="Node at 127.0.0.1:8000 [Candidate]" term=2
2024-09-26T12:21:58.610+0530 [DEBUG] raft: asking for pre-vote: term=2 from=127.0.0.1:8001 address=127.0.0.1:8001
2024-09-26T12:21:58.610+0530 [DEBUG] raft: pre-voting for self: term=2 id=127.0.0.1:8000
2024-09-26T12:21:58.610+0530 [DEBUG] raft: calculated votes needed: needed=2 term=2
2024-09-26T12:21:58.610+0530 [DEBUG] raft: pre-vote received: from=127.0.0.1:8000 term=2 tally=0
2024-09-26T12:21:58.610+0530 [DEBUG] raft: pre-vote granted: from=127.0.0.1:8000 term=2 tally=1
2024-09-26T12:21:58.611+0530 [DEBUG] raft: pre-vote received: from=127.0.0.1:8001 term=2 tally=1
2024-09-26T12:21:58.611+0530 [DEBUG] raft: pre-vote granted: from=127.0.0.1:8001 term=2 tally=2
2024-09-26T12:21:58.611+0530 [INFO]  raft: pre-vote successful, starting election: term=2 tally=2 refused=0 votesNeeded=2
2024-09-26T12:21:58.621+0530 [DEBUG] raft: asking for vote: term=2 from=127.0.0.1:8001 address=127.0.0.1:8001
2024-09-26T12:21:58.621+0530 [DEBUG] raft: voting for self: term=2 id=127.0.0.1:8000
2024-09-26T12:21:58.649+0530 [DEBUG] raft: vote granted: from=127.0.0.1:8000 term=2 tally=1
2024-09-26T12:21:58.660+0530 [DEBUG] raft: vote granted: from=127.0.0.1:8001 term=2 tally=2
2024-09-26T12:21:58.660+0530 [INFO]  raft: election won: term=2 tally=2
2024-09-26T12:21:58.660+0530 [INFO]  raft: entering leader state: leader="Node at 127.0.0.1:8000 [Leader]"
2024-09-26T12:21:58.660+0530 [INFO]  raft: added peer, starting replication: peer=127.0.0.1:8001
2024-09-26T12:21:58.660+0530 [INFO]  raft: pipelining replication: peer="{Voter 127.0.0.1:8001 127.0.0.1:8001}"
2024/09/26 12:21:58 This node is the leader.
2024/09/26 12:21:58 Updated state with sus names: [sus0: Charlie sus1: Eve sus2: Bob]
2024/09/26 12:21:58 I am the leader and here is the sus list I declare :P - [Charlie Eve Bob]
```