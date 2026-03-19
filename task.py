import json
import redis

r = redis.Redis(host='redis', port=6379, db=0)

def send_task(data: dict):
    r.lpush("tasks_queue", json.dumps(data))