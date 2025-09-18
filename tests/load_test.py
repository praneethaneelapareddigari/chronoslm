import time
from confluent_kafka import Producer

producer = Producer({'bootstrap.servers': 'kafka:9092'})
for i in range(50):
    producer.produce("multimodal-stream", f"token-{i}".encode("utf-8"))
    if i % 10 == 0:
        print(f"Sent {i} messages")
    time.sleep(0.05)
producer.flush()
