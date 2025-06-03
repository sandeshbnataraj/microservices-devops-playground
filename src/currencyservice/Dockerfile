FROM node:23-alpine

RUN apk add --no-cache python3 make g++

WORKDIR /app

COPY package*.json ./

RUN npm install --only=production

COPY . ./

ENV PORT=7000
ENV DISABLE_PROFILER=true

EXPOSE 7000

CMD ["node", "server.js"]