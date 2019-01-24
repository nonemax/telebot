FROM alpine
ENV LANGUAGE="en"
COPY ./ .
CMD [ "./telebot" ]