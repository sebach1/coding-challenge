FROM node:14.9.0-alpine3.10 as build-stage

WORKDIR /frontend

ENV NODE_ENV development

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:1.19.2-alpine as prd-server

RUN mkdir /app

COPY --from=build-stage /frontend/dist /app

EXPOSE 8000

COPY nginx.conf /etc/nginx/nginx.conf
