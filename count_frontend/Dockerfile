FROM node

WORKDIR /usr/

COPY package.json /usr/

RUN npm install
RUN npm install -g @angular/cli@1.7.1
COPY . /usr/

CMD ["ng","serve", "--host", "0.0.0.0"]
