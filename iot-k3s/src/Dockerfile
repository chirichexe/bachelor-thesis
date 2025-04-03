FROM python:3.9
WORKDIR /app
COPY controller.py /app/
RUN pip install kopf kubernetes
CMD ["kopf", "run", "--standalone", "controller.py"]
