from memory.checkpoint.manager import write_checkpoint, list_checkpoints, replay_from
p1 = write_checkpoint(100000, {"note": "startup"})
p2 = write_checkpoint(500000, {"note": "mid-run"})
p3 = write_checkpoint(1250000, {"note": "pre-upgrade"})
print("Checkpoints:", list_checkpoints())
print("Replay:", replay_from(500000))
