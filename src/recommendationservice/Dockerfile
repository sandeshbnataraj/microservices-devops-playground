FROM python:3.12

RUN python3 -m venv venv

ENV PATH="$PATH:venv/bin/"

COPY requirements.txt ./

RUN pip install --upgrade pip && \
    pip install -r requirements.txt

COPY . ./

CMD ["python3", "recommendation_server.py"]