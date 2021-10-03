FROM node:alpine
WORKDIR '/app'
COPY ./package.json ./
COPY ./package-lock.json ./
COPY ./nuxt.config.js ./
COPY ./static ./
ENV HOST=0.0.0.0
ENV API_URL=http://tennis-env-1.eba-dr2xznbr.ap-northeast-1.elasticbeanstalk.com

RUN npm install
COPY . .
RUN npm run build
CMD ["npm","run","start"]