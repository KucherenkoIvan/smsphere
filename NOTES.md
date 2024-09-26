# architecture design doc

## clause 1: client-to-end data flow
[client] -> http -> [http api controller] -> [prioritizer] ->
        priority queue -> [distributer] -> kafka -> [consumer] -> [partitioner] -> [integration-interface] -> [sms-operator]

## clause 2: client response data flow

[sms-operator] -> [integration-interface] -> [db] -> [cron] -> [status-checker] -> [http-controller]
                                                                     \--> [web-hooks]

## important features

1. ip white-list
2. rate-limiter
3. emailer
4. tg viber whatsapp notifications
5. lk gui sender
6. graphs (?)
7. private admin pannel
8. excel reports (?)
9. ip-distribution (?)
10. weighted priority queues
11. priority dependend tariff prices
12. feature flags
13. free 10 sms
14. smsphere.send(dest, text, [`email`, `viber`, `tg`, `sms`])
15. packages for `js`, `php`, `python`, `go`


