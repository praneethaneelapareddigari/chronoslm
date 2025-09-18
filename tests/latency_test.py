import time
from memory.faiss_layer import FaissLayer

f = FaissLayer(4)
f.add([[0.1,0.2,0.3,0.4]], ["id1"])
start = time.time()
print(f.search([0.1,0.2,0.3,0.4]))
print("Latency:", (time.time()-start)*1000, "ms")
