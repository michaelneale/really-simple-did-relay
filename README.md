# Really simple DID Relay (rdr)

This is a simple DID relay that can be used to relay messages between dids. 

## Running

`go run rdr.go`


## Usage

Sender: Post a message to a specified recipient. The nonce is used to prove the senderDid is legit (but doesn't affect recipient payload).

`curl -X POST -H "Content-Type: application/json" -d '{"nonce":"123456","recipientDid":"recipient1","senderDid":"sender1","payload":"Hello World"}' http://localhost:8080/`


Receiver: Check if you have any messages do your did: 

`curl http://localhost:8080/\?recipientDid\=recipient1\&nonce\=123456`

This will return a list of payloads if there are any. Once it is fetched they are removed from the relay automatically. 

The nonce is to prove you controll the  recipientDid. If that fails you get nothing.