FROM nginx:1.27.4

COPY nginx.conf /etc/nginx/nginx.conf

RUN rm -rf /usr/share/nginx/html/*

COPY  connect.js /usr/share/nginx/html/
COPY  main.html /usr/share/nginx/html/

WORKDIR /usr/share/nginx/html

RUN chown -R nginx:nginx /usr/share/nginx/html

EXPOSE 80