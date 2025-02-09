FROM node:alpine AS build_stage
WORKDIR /container_monitoring
COPY ./package*.json ./
RUN npm install
ARG REACT_APP_API_URL
ENV REACT_APP_API_URL $REACT_APP_API_URL
COPY ./ ./
RUN npm run build

FROM nginx
COPY --from=build_stage /container_monitoring/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf