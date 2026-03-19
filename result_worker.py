import json
import time
import redis
from sqlalchemy import create_engine, Column, String, Text, DateTime
from sqlalchemy.orm import declarative_base, sessionmaker
from datetime import datetime

DATABASE_URL = "postgresql://myuser:mypassword@postgres:5432/mydb"

# ⏳ Ждём БД
for i in range(10):
    try:
        engine = create_engine(DATABASE_URL)
        conn = engine.connect()
        conn.close()
        print("✅ Connected to DB")
        break
    except Exception as e:
        print("⏳ Waiting for DB...", e)
        time.sleep(2)
else:
    raise Exception("❌ Could not connect to DB")

SessionLocal = sessionmaker(bind=engine)
Base = declarative_base()

class Report(Base):
    __tablename__ = "reports"
    task_id = Column(String, primary_key=True, index=True)
    status = Column(String)
    result = Column(Text)
    created_at = Column(DateTime, default=datetime.utcnow)

Base.metadata.create_all(bind=engine)

# ✅ Redis тоже через сервис
r = redis.Redis(host='redis', port=6379, db=0)

def consume_results():
    print("🚀 Worker started, waiting for results...")
    while True:
        res = r.brpop("results_queue")
        data = json.loads(res[1])
        handle_result(data)

def handle_result(data):
    session = SessionLocal()
    report = Report(
        task_id=data["task_id"],
        status=data["status"],
        result=data["data"]
    )
    session.add(report)
    session.commit()
    session.close()
    print(f"✅ Saved result {data['task_id']}")

if __name__ == "__main__":
    consume_results()