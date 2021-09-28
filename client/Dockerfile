FROM node:alpine as builder
WORKDIR '/app'
COPY ./package.json ./
COPY ./package-lock.json ./
COPY ./nuxt.config.js ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx
EXPOSE 3000
COPY --from=builder /app/.next /usr/share/nginx/html