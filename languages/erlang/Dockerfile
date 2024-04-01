FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

# https://rabbitmq-users.narkive.com/OtHXmuO6/rabbitmq-discuss-ll-alloc-errors-on-rabbitmq-2-1-1
ENV ERL_MAX_PORTS 1024

RUN pacman -Syu --noconfirm erlang-nox
