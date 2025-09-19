import json, os, time, pathlib
CP_DIR = pathlib.Path("memory/checkpoint/store")
CP_DIR.mkdir(parents=True, exist_ok=True)

def write_checkpoint(offset: int, meta: dict):
    ts = int(time.time())
    path = CP_DIR / f"cp_{ts}_{offset}.json"
    with open(path, "w") as f:
        json.dump({"offset": offset, "meta": meta, "ts": ts}, f, indent=2)
    return str(path)

def list_checkpoints():
    return sorted([p.name for p in CP_DIR.glob("cp_*.json")])

def replay_from(offset: int):
    duration_ms = min(5000, 10 + offset // 1000)
    time.sleep(duration_ms / 1000.0)
    return {"replayed_from_offset": offset, "duration_ms": duration_ms}
