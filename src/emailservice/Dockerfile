FROM python:3.12

RUN python -m venv venv
#    . venv/bin/activate not needed as i am setting the env path

ENV PATH="$PATH:venv/bin/"

COPY requirements.txt ./

RUN pip install --upgrade-pip && \
    pip install -r requirements.txt

COPY . .

CMD ["python3", "email_server.py"]