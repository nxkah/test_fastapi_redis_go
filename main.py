# main.py
from fastapi import FastAPI
from pydantic import BaseModel
import uuid
from task import send_task
from sqlalchemy.orm import sessionmaker
from models import engine, Report

SessionLocal = sessionmaker(bind=engine)
app = FastAPI()

# --- отправка задачи ---
class TaskRequest(BaseModel):
    user_id: str

@app.post("/create_report/")
def create_report(task: TaskRequest):
    task_id = str(uuid.uuid4())
    send_task({
        "task_id": task_id,
        "type": "generate_report",
        "user_id": task.user_id
    })
    return {"status": "processing", "task_id": task_id}

# --- проверка статуса ---
@app.get("/get_report/{task_id}")
def get_report(task_id: str):
    session = SessionLocal()
    report = session.query(Report).filter(Report.task_id == task_id).first()
    session.close()

    if report:
        return {"status": report.status, "result": report.result}
    return {"status": "processing"}