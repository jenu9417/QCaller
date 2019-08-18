# QCaller

## Improvements & Future Scope
------------------------------

QCaller is in v1 stage and has a vast improvement potential. It has been designed with potential to accomodate this.

For a Caller ID App, the two main use cases are:

- **`Input filtering`**
	we can add more proper filtering to the api's to accept only proper clean records. Because a lot on people store more than phone numbers in their phone and it would be meaningless and costly to add all those.

- **`Search`**
	This is a critical usecase and needs more research to be done on its own. We have gone with ES based on the popular community idea and a calculated guess. But we need data to quantify this and adding metrics to QCaller is critical. We can improve search results by adding more granular filters like, giving preference to contacts with a proper name than generic names such a spam, banker etc., A bare minimum version of this is already implemented, which will give preference to contacts without a 'spam' keyword when multiple contacts match for same number. We can build on top of that and ES is very much powerful in this regard


## Authentication
Add better authentication. For simplicity, have gone for basic auth. Can go for token based auth.


## Metrics
The one critical thing to do is metrics. We have built a platform. But inorder to improve it, we need proper metrics.
Some of the metrics params we could analyse are:

**`App Metrics`**
- Number of requests
- Number of 2xx vs 4xx vs 5xx
- Request latency
- DB Operations latency

**`AS Metrics`**
- Data Size
- Cache miss rate
- Number of requests served
- Operation latency
- High Watermark thresholds
- Cluster nodes

**`ES Metrics`**
- Data Size
- Operation latency
- Cluster health
		
and along with this common system metrics such as ram, diskusage, cpu usage etc.,


## Throttling
This is another critical feature to be done. QCaller will be public facing and its better to have a control at our end on how much we can serve than to flood ourselves and choke over it.


## Context Based Logging
 Currently the logs doesn't have any indicators to request id. Hence it might be bit more tedious while debugging. We already have context with request id. We can pass it all along, and for each log, add corresponding request-id to it.


## Automate Build Jobs with Ansible
We can write ansible jobs to automate build and deployment.


## Add Containerization support
Horizontal scaling is one of the strong points of QCaller. We can add containerization support to leverage this to max extent.
