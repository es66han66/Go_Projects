1- Run kafka using command ```docker-compose up```.  
  
2- You can find your IP using ```ifconfig -a``` in your terminal.  
  
3- To create a topic example foo, run ```docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic foo --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:2181```  
  