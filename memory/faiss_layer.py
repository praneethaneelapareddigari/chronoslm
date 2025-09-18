import faiss
import numpy as np

class FaissLayer:
    def __init__(self, dim=128):
        self.index = faiss.IndexFlatL2(dim)
        self.ids = []

    def add(self, vecs, ids):
        np_vecs = np.array(vecs).astype('float32')
        self.index.add(np_vecs)
        self.ids.extend(ids)

    def search(self, query, k=5):
        q = np.array(query).astype('float32').reshape(1, -1)
        D, I = self.index.search(q, k)
        return [(self.ids[i], float(D[0][j])) for j, i in enumerate(I[0]) if i != -1]
