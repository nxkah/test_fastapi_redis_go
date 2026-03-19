from sqlalchemy import create_engine, Column, String, Text, DateTime
from sqlalchemy.orm import declarative_base
from datetime import datetime

DATABASE_URL = "postgresql://myuser:mypassword@postgres:5432/mydb"
engine = create_engine(DATABASE_URL)

Base = declarative_base()

class Report(Base):
    __tablename__ = "reports"

    task_id = Column(String, primary_key=True, index=True)
    status = Column(String, nullable=False)
    result = Column(Text, nullable=True)
    created_at = Column(DateTime, default=datetime.utcnow)

def init_db():
    Base.metadata.create_all(bind=engine)