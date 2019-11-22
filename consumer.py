from kafka import KafkaConsumer
consumer = KafkaConsumer('test', bootstrap_servers='localhost:9092', group_id='my-group')
while True:
    for msg in consumer:
        print (msg)
