FROM node:latest

WORKDIR /app
COPY . .

RUN npm i

EXPOSE 3000

ENV PORT=3000

RUN npm run build

CMD ["npm", "run", "start"]