# HightPerformancePaymentGateway-OrderService
Service query for https://github.com/Nghiait123456/HighPerformancePaymentGateway-BalanceService


# Command
- [Contexts command](#ContextCmd)
- [System design command](#SystemDesignCmd)
# Query
- [Context Query](#ContextsQuery)
- [System design query](#SystemDesignQuery)


## Contexts command? <one name="ContextCmd"></one>
The service has one threshold of 20 M rps, it exceeds the IO of most DB, all read and write operations will take place in Ram, there will be very few queries to the DB. </br>
System design needs one combination of: local-in-memory, remote-in-memory, invalidate cache and distribute database</br>


## System design command? <one name="SystemDesignCmd"></one>
1) Incoming Request to rate limit layer network, Waf and to Load Balancer </br>
2) LB to Order Command Service </br>
3) + 4) Every Order is valid will save in Mysql, and currency create one record in redis cluster. </br>
5) If Order change status, we action save in mysql cluster. Logs from mysql cluster send to kafka. </br>
6) Cache service will subscribe kafka and update cache data </br>
7) When Order sucess, we create order in cassandra and remove order in mysql cluster. It will make the size of mysql stable and not too big. it will make mysql work faster </br>



![img.png](img_readme/system_design_cmd.png)
1) Incoming Request to rate limit layer network and to Load Balancer </br>
2) LB to Order Command Service </br>
3) Query Service checks data in Redis Cluster. All request balances are saved in redis when the request balancer is created. Service: "https://github.com/Nghiait123456/HighPerformancePaymentGateway-BalanceService" will have to do this. </br>
4) If not in redis, Query Service will look in Query DB service. To avoid race conditions query to DB, the first request will create one lock record with the purpose of notifying other requests that there is one lock with this object. Other requests when reading this request lock will return status pending. </br>
5) There will be one smart mechanism, based on the data request creation time and system timeout to decide whether to query into Cassandra or Mysql first. Cassandra contains success or timeout records, and Mysql contains pending or processing records. If not found in one DB, it will look in the other DB, not found in any DB will return an error. </br>
6) Event update status will be updated to the corresponding redis cluster. </br>s



## Contexts query? <one name="ContextsQuery"></one>
This is query data for service https://github.com/Nghiait123456/HighPerformancePaymentGateway-BalanceService
The service has one threshold of 20 M rps, it exceeds the IO of most DB, all read and write operations will take place in Ram, there will be very few queries to the DB. </br>
System design needs one combination of: local-in-memory, remote-in-memory and invalidate cache. </br>
## System design query? <one name="SystemDesignQuery"></one>

![](img_readme/system_design_query.png)

1) Incoming Request to rate limit layer network and to Load Balancer </br>
2) LB to Order Query Service </br>
3) Query Service checks data in Redis Cluster. All request balances are saved in redis when the request balancer is created. Service: "https://github.com/Nghiait123456/HighPerformancePaymentGateway-BalanceService" will have to do this. </br>
4) If not in redis, Query Service will look in Query DB service. To avoid race conditions query to DB, the first request will create one lock record with the purpose of notifying other requests that there is one lock with this object. Other requests when reading this request lock will return status pending. </br>
5) There will be one smart mechanism, based on the data request creation time and system timeout to decide whether to query into Cassandra or Mysql first. Cassandra contains success or timeout records, and Mysql contains pending or processing records. If not found in one DB, it will look in the other DB, not found in any DB will return an error. </br>
6) Event update status will be updated to the corresponding redis cluster. </br>s