FROM node:18-alpine

RUN apk add --no-cache python3 make g++

WORKDIR /app

COPY package*.json ./

RUN npm audit fix --force && npm ci --only=production

COPY . ./

ENV PORT=7001
ENV DISABLE_PROFILER=true

EXPOSE 7001

CMD ["node", "index.js"]