# Multi-stage build for Agent with Playwright

# Stage 1: Build stage
FROM node:18-alpine AS builder

WORKDIR /build

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci --only=production=false

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Stage 2: Runtime stage with Playwright
FROM node:18-bookworm-slim

# Install Playwright system dependencies and Chromium
# Using bookworm-slim (Debian) for better Playwright compatibility
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    wget \
    ca-certificates \
    fonts-liberation \
    libasound2 \
    libatk-bridge2.0-0 \
    libatk1.0-0 \
    libatspi2.0-0 \
    libcups2 \
    libdbus-1-3 \
    libdrm2 \
    libgbm1 \
    libgtk-3-0 \
    libnspr4 \
    libnss3 \
    libwayland-client0 \
    libxcomposite1 \
    libxdamage1 \
    libxfixes3 \
    libxkbcommon0 \
    libxrandr2 \
    xdg-utils && \
    rm -rf /var/lib/apt/lists/*

# Create non-root user
RUN groupadd -r appuser && useradd -r -g appuser -G audio,video appuser && \
    mkdir -p /home/appuser/Downloads && \
    chown -R appuser:appuser /home/appuser

WORKDIR /app

# Copy package files and install production dependencies
COPY package*.json ./
RUN npm ci --only=production

# Copy built application from builder
COPY --from=builder /build/dist ./dist
COPY --from=builder /build/index.ts ./index.ts

# Install Playwright and Chromium browser
# This downloads Chromium binary (~280MB)
RUN npx playwright install chromium --with-deps

# Change ownership
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Set Playwright cache directory
ENV PLAYWRIGHT_BROWSERS_PATH=/ms-playwright

# Run the agent
CMD ["node", "dist/index.js"]
