#!/bin/bash

echo "---- Send a message to recipient1 ---- "
curl -X POST -H "Content-Type: application/json" -d '{"nonce":"123456","recipientDid":"recipient1","senderDid":"sender1","payload":"Hello World"}' http://localhost:8080/
echo " ----- Test we can get it back for receipient1:"
curl http://localhost:8080/\?recipientDid\=recipient1\&nonce\=123456
echo "\n ----- Test it has been removed as read:"
curl http://localhost:8080/\?recipientDid\=recipient1\&nonce\=123456
