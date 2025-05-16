FROM node:20-slim AS builder
LABEL maintainer="Maciej Ołdakowski"
WORKDIR /app

COPY package*.json ./
ENV NODE_ENV=production
RUN npm install --omit=dev

FROM node:20-slim
LABEL org.opencontainers.image.authors="Maciej Ołdakowski"
WORKDIR /app

COPY --from=builder /app/node_modules ./node_modules
COPY app.js package*.json ./

ENV NODE_ENV=production
ENV PORT=3000
EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:3000 || exit 1

CMD ["node", "app.js"]
