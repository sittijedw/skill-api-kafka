FROM node:lts-alpine3.20
WORKDIR /usr/src/app
COPY . .
RUN npm install playwright@1.45.3
ENTRYPOINT [ "npx", "playwright", "test" ]