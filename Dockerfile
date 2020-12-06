FROM python@sha256:e4e54a385b186ebc7e5f49b072185be8179ddbf7c05ea7c50009818b864e522a

LABEL maintainer="Sebastián Chamena sebachamena@gmail.com"

WORKDIR /app

COPY requirements.txt .

RUN pip3 install --upgrade pip && pip3 install -r ./requirements.txt

COPY . .

ENTRYPOINT [ "python3", "manage.py" ]

CMD [ "runserver", "0.0.0.0:8000" ]