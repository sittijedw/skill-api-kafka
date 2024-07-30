FROM node:lts-alpine3.20
WORKDIR /usr/src/app
COPY . .
RUN npm init playwright@latest \
    '' \
    '' \
    '' \
    ''
ENTRYPOINT [ "npx", "playwright", "test" ]