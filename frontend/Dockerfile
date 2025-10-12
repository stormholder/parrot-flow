# Multi-stage build for React + Vite frontend

# Stage 1: Build stage
FROM node:18-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci --only=production=false

# Copy source code
COPY . .

# Build argument for API URL
ARG VITE_API_URL=/api
ENV VITE_API_URL=$VITE_API_URL

# Build the application
RUN npm run build

# Stage 2: Runtime stage (nginx will serve these files)
FROM scratch AS dist

# Copy built files to be used by nginx
COPY --from=builder /app/dist /app/dist

# This image won't run standalone - nginx will copy from it
