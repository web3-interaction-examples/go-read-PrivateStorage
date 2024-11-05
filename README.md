# go-read-PrivateStorage


## Format
```
			Slot 0 (32 bytes):
			+---------------+------------------+--------------------------------+
			| Padding       | startTime        | user                           |
			| (4 bytes)     | (8 bytes)       | (20 bytes)                      |
			+---------     +------------------+---------------------------------+

			Slot 1 (32 bytes):
			+------------------------------------------------------------------+
			| amount                                                           |
			| (32 bytes)                                                       |
			+------------------------------------------------------------------+
```

## Logs
```
arrayLength: 11
locks[0]:
  user: 0x0000000000000000000000000000000000000001
  startTime: 2079-09-10 15:41:20 (3461557280)
  amount: 1.00 ETH
  Debug - slot1 data: 0000000000000000ce532c200000000000000000000000000000000000000001
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c20
        Offset: 12 - user: 0000000000000000000000000000000000000001
  Debug - slot2 data (amount): 0000000000000000000000000000000000000000000000000de0b6b3a7640000
----------------------------------------
locks[1]:
  user: 0x0000000000000000000000000000000000000002
  startTime: 2079-09-10 15:41:19 (3461557279)
  amount: 2.00 ETH
  Debug - slot1 data: 0000000000000000ce532c1f0000000000000000000000000000000000000002
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c1f
        Offset: 12 - user: 0000000000000000000000000000000000000002
  Debug - slot2 data (amount): 0000000000000000000000000000000000000000000000001bc16d674ec80000
----------------------------------------
locks[2]:
  user: 0x0000000000000000000000000000000000000003
  startTime: 2079-09-10 15:41:18 (3461557278)
  amount: 3.00 ETH
  Debug - slot1 data: 0000000000000000ce532c1e0000000000000000000000000000000000000003
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c1e
        Offset: 12 - user: 0000000000000000000000000000000000000003
  Debug - slot2 data (amount): 00000000000000000000000000000000000000000000000029a2241af62c0000
----------------------------------------
locks[3]:
  user: 0x0000000000000000000000000000000000000004
  startTime: 2079-09-10 15:41:17 (3461557277)
  amount: 4.00 ETH
  Debug - slot1 data: 0000000000000000ce532c1d0000000000000000000000000000000000000004
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c1d
        Offset: 12 - user: 0000000000000000000000000000000000000004
  Debug - slot2 data (amount): 0000000000000000000000000000000000000000000000003782dace9d900000
----------------------------------------
locks[4]:
  user: 0x0000000000000000000000000000000000000005
  startTime: 2079-09-10 15:41:16 (3461557276)
  amount: 5.00 ETH
  Debug - slot1 data: 0000000000000000ce532c1c0000000000000000000000000000000000000005
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c1c
        Offset: 12 - user: 0000000000000000000000000000000000000005
  Debug - slot2 data (amount): 0000000000000000000000000000000000000000000000004563918244f40000
----------------------------------------
locks[5]:
  user: 0x0000000000000000000000000000000000000006
  startTime: 2079-09-10 15:41:15 (3461557275)
  amount: 6.00 ETH
  Debug - slot1 data: 0000000000000000ce532c1b0000000000000000000000000000000000000006
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c1b
        Offset: 12 - user: 0000000000000000000000000000000000000006
  Debug - slot2 data (amount): 00000000000000000000000000000000000000000000000053444835ec580000
----------------------------------------
locks[6]:
  user: 0x0000000000000000000000000000000000000007
  startTime: 2079-09-10 15:41:14 (3461557274)
  amount: 7.00 ETH
  Debug - slot1 data: 0000000000000000ce532c1a0000000000000000000000000000000000000007
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c1a
        Offset: 12 - user: 0000000000000000000000000000000000000007
  Debug - slot2 data (amount): 0000000000000000000000000000000000000000000000006124fee993bc0000
----------------------------------------
locks[7]:
  user: 0x0000000000000000000000000000000000000008
  startTime: 2079-09-10 15:41:13 (3461557273)
  amount: 8.00 ETH
  Debug - slot1 data: 0000000000000000ce532c190000000000000000000000000000000000000008
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c19
        Offset: 12 - user: 0000000000000000000000000000000000000008
  Debug - slot2 data (amount): 0000000000000000000000000000000000000000000000006f05b59d3b200000
----------------------------------------
locks[8]:
  user: 0x0000000000000000000000000000000000000009
  startTime: 2079-09-10 15:41:12 (3461557272)
  amount: 9.00 ETH
  Debug - slot1 data: 0000000000000000ce532c180000000000000000000000000000000000000009
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c18
        Offset: 12 - user: 0000000000000000000000000000000000000009
  Debug - slot2 data (amount): 0000000000000000000000000000000000000000000000007ce66c50e2840000
----------------------------------------
locks[9]:
  user: 0x000000000000000000000000000000000000000A
  startTime: 2079-09-10 15:41:11 (3461557271)
  amount: 10.00 ETH
  Debug - slot1 data: 0000000000000000ce532c17000000000000000000000000000000000000000a
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c17
        Offset: 12 - user: 000000000000000000000000000000000000000a
  Debug - slot2 data (amount): 0000000000000000000000000000000000000000000000008ac7230489e80000
----------------------------------------
locks[10]:
  user: 0x000000000000000000000000000000000000000b
  startTime: 2079-09-10 15:41:10 (3461557270)
  amount: 11.00 ETH
  Debug - slot1 data: 0000000000000000ce532c16000000000000000000000000000000000000000b
        Offset: 0 - padding: 00000000
        Offset: 4 - startTime: 00000000ce532c16
        Offset: 12 - user: 000000000000000000000000000000000000000b
  Debug - slot2 data (amount): 00000000000000000000000000000000000000000000000098a7d9b8314c0000
----------------------------------------
```
