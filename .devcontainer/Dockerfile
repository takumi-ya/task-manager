FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

USER root

RUN apt-get update \
    && apt-get install -y postgresql-client \
    && apt-get clean

# vscode(UID=1000) が /go に書き込めるように所有者を変更
RUN chown -R vscode:vscode /go

USER vscode

# root以外のユーザで go install (書き込み権限が付与されている)
RUN go install github.com/air-verse/air@latest

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
