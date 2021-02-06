FROM python:3.8.4

COPY requirements.txt /

RUN pip install -r requirements.txt

COPY main.py /
COPY utils /utils

CMD [ "python", "./main.py" ]