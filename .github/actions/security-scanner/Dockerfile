FROM alpine:3.18

# Install required tools
RUN apk add --no-cache \
    bash \
    curl \
    git \
    jq \
    python3 \
    py3-pip \
    nodejs \
    npm \
    go \
    && rm -rf /var/cache/apk/*

# Install security scanning tools
RUN pip3 install --no-cache-dir \
    safety \
    bandit \
    semgrep

# Install Node.js security tools
RUN npm install -g \
    npm-audit \
    retire \
    eslint-plugin-security

# Install Go security tools
RUN go install github.com/securecodewarrior/gosec/cmd/gosec@latest && \
    go install golang.org/x/vuln/cmd/govulncheck@latest

# Copy the scanner script
COPY scanner.sh /scanner.sh
RUN chmod +x /scanner.sh

# Set the entrypoint
ENTRYPOINT ["/scanner.sh"]
