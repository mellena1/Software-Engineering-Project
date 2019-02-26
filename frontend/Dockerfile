FROM node

WORKDIR /usr/

COPY package.json /usr/

RUN npm install
COPY . /usr/

CMD ["ng","serve", "--host", "0.0.0.0"]
