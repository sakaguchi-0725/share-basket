FROM node:23.11.0-slim AS develop

ENV TZ=Asia/Tokyo

WORKDIR /app

COPY package*.json ./

RUN npm install

CMD ["npm", "run", "dev"]