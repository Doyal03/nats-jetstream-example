# nats-jetstream-example
An example to demonstrate working of jetstream in NATS communication

The JetstreamSubscriber is listening to the incoming stream, so should be started first. Then JetstreamPublisher should be run, which is sending 100000 messages, with a timeout limit at 1ns.

Its important to note here that the jetstream is configured with "nats.PublishAsyncMaxPending(math.MaxInt)" otherwise default limits of sending data per stream will be implemented and error will be thrown.

A sample running of the publisher (left) and the subscriber (right) is present in the screenshot. In this case, since the timeout is reached before sending all the data, the last data received on the subscriber is 82669, instead of 99999.

![image](https://user-images.githubusercontent.com/91259249/191333311-cab09e43-fd98-4392-88cb-28eb9c1de1f9.png)
