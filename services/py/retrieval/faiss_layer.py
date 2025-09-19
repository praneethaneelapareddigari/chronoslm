"""
FAISS-like ANN layer (mock).
"""
import time, random
def search(query: str):
    start = time.time()
    time.sleep(random.uniform(0.008, 0.018))
    latency_ms = int((time.time() - start) * 1000)
    return {"hits":[{"id":"doc_42","score":0.965},{"id":"doc_7","score":0.951}],"latency_ms":latency_ms}
