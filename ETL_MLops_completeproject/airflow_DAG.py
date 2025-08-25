from airflow import DAG
from airflow.operators.python import PythonOperator
from datetime import datetime

def train_model():
    # Run training script
    import os
    os.system("python train.py")

def deploy_model():
    # Deploy model
    import os
    os.system("python deploy.py")

with DAG("mlops_pipeline", start_date=datetime(2025, 7, 1), schedule_interval="@daily") as dag:
    train = PythonOperator(task_id="train_model", python_callable=train_model)
    deploy = PythonOperator(task_id="deploy_model", python_callable=deploy_model)
    train >> deploy