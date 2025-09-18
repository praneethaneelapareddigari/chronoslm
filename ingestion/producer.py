import json
from confluent_kafka import Producer
import time

conf = {'bootstrap.servers': 'kafka:9092'}
producer = Producer(conf)

def delivery_report(err, msg):
    if err:
        print(f"Delivery failed: {err}")
    else:
        print(f"Produced to {msg.topic()} [{msg.partition()}] @ {msg.offset()}")

if __name__ == "__main__":
    for i in range(5):
        data = {"id": i, "text": f"hello world {i}", "timestamp": time.time()}
        producer.produce("multimodal-stream", json.dumps(data).encode("utf-8"), callback=delivery_report)
        producer.flush()
        time.sleep(1)
