# Use the Node.js 18 image
FROM node:20.13

# Set the working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json package-lock.json ./
RUN yarn install

# Copy all project files
COPY . .

# Expose Vite's development server port
EXPOSE 3000

# Default command
CMD ["yarn", "run", "dev"]
