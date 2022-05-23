

1. how to handle errors? panic or error
what if there is a bug that causes oom or some other panic. do we halt the network or continue with a consensus that this transaction has failed?

2. unclear GetNonce definition. why there is even a method for it
why it doesn't have access to fields of the transaction
it just reads from state

3. how to define immutability?

it looks like it needs immutable type for every complex type

4. when to charge gas?
host.Consume(amount uint64)

if it overflows the gas limit or balance is exhaused - terminate transaction 

5. do we consume funds from the failed transactions?
need to be sure that failure transactions do not modify state

if transaction passes validation phase (verify and what else?) funds will be withdrawn from account. specifically state won't be updated if transaction failed
but funds will be withdrawn.

State
---

field    | notes
---------|-----------------------------------------------------------
address  | principal address
template | nil for stub account
balance  | funds
nonce    | do we need this field or it is always a part of the state?
state    | serialized state of the contract

Tx flows
---

## Spawn

1. decode principal
2. decode selector
3. if selector is 0
4. need to decode payload into concrete type. can i do that with generics?
5. type.Spawn
6. encode P
7. save state

## Regular

1-2. the same
3. if selector is not 1
4. get lookup table for template address. no map, use switch statement 
5. consume gas inside of the methods
6 - 7. is the same


