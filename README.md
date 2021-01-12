## Overview
We present PoC demo application based on FE inner product algorithm (on top of
DDH assumption, proposed in [this paper][paper]) that demonstrates how FE can
be used to build crypto system to anonymize payments and which can be integrated
to existing blockchain without introducing new security assumptions.

[paper]: https://eprint.iacr.org/2015/017.pdf

Demo is based on [simple.DDH][gofe-ddh] from [gofe] package which requires trusted party to
perform key generation/derivation. However, we proposed a way how it can be 
easily done distributedly without trusted party.

[gofe]: https://github.com/fentec-project/gofe
[gofe-ddh]: https://github.com/spf13/cobra

## Protocol
We will briefly describe proposed protocol.

We have: 
* n senders S1,...,Sn
* m receivers R1,...,Rn
* t rounds

Protocol has 3 phases:
* Setup: key generation & derivation. Executed by trusted party, but just for simplicity of demo.
* Send. Sender Si sends a signal to receiver Rj that (s)he made a payment by (in terms of protocol)
  introducing new round. It will be published publicly (on blockchain), but no one from outside
  can tell who is receipient (signal to Ri is indistinguishable from signal to Rj).
* Search. Executes when receiver goes online. E.g. receiver was online last time and last round he saw
  was t1. Now he comes back to the Internet and sees new round t2. He can find out if someone sent a
  a signal to him just by knowing rounds t1 & t2. If he received a signal, he can find an exact round
  on which signal was made for O(log(t2-t1)).
  
### Key generation & derivation
```
# Trusted party generates master key:
mpk, msk <- Setup(SECURITY_BITS)

# For every Rj we derive key:
for j in range(m):
  y := [0; m] # array of n zeroes
  y[j] = 1
  sk[j] <- KeyDer(mpk, y)
```
We publish mpk and send sk_j privatly to corresponding receivers.

### Send signal to Rj
```
x := [0; 32]
x[j] = 1
E := Encrypt(mpk, x)

t <- CurrentRound
if t > 1:
  E_previous <- Round(t-1)
  E = E_previous + E
  
PublishRound(t+1, E)
```

### Search signal
Receiver Rj goes online at round t2 (last time it was online at round t1). It retrieves ciphertext Et 
and computes `v = Decrypt(mpk, Et, sk_j)`. If v = v' (where v' is value obtained at round t1), then
no signal was received. Otherwise, it uses binary search to find out at which round signal has been sent
(by looking up for round (t2-t1)/2, computing v'' and comparing with v', etc.).

## Run demo

### Run keygen
```bash
go run ./cmd/demo keygen -parties 5
```
It will create files:
* `stand/rounds/round0.json` containing mpk
* `stand/parties/partyN.json` N files for every receiver containing sk_j

### Send signal
```bash
go run ./cmd/demo send-signal -party 2
```

Outputs:
```
You successfully sent signal to party 2 in round 1!
```

It will create file `rounds/round1.json` containing a ciphertext which won't tell you that party 2 
received a signal if you don't know sk_2.

Send a few more signals:
```bash
go run ./cmd/demo send-signal -party 3
go run ./cmd/demo send-signal -party 4
go run ./cmd/demo send-signal -party 5
```

### Search
```bash
go run ./cmd/demo search -party 2
```

Outputs:
```
Party 2 is waking up!
Current round is 4... Party received a signal!
Looking for axact round where signal was sent...

Round 2, v != v'
Round 1, v = v'

Party received a signal at round 1.
```
