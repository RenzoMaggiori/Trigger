# Use Node.js 18 with Debian Bullseye as the base image
FROM node:18-bullseye

# Set environment variables for Android SDK
ENV ANDROID_SDK_ROOT=/opt/android-sdk
ENV ANDROID_HOME=${ANDROID_SDK_ROOT}
ENV PATH=${PATH}:${ANDROID_SDK_ROOT}/cmdline-tools/latest/bin:${ANDROID_SDK_ROOT}/platform-tools

# Install necessary dependencies
RUN apt-get update && apt-get install -y openjdk-17-jdk wget unzip

# Create directories for Android SDK
RUN mkdir -p ${ANDROID_SDK_ROOT}/cmdline-tools

# Download and extract Android command-line tools
RUN wget https://dl.google.com/android/repository/commandlinetools-linux-9477386_latest.zip -O /tmp/commandlinetools.zip \
    && unzip /tmp/commandlinetools.zip -d /tmp/android-sdk-tmp \
    && rm /tmp/commandlinetools.zip \
    && mv /tmp/android-sdk-tmp/cmdline-tools ${ANDROID_SDK_ROOT}/cmdline-tools/latest \
    && rm -rf /tmp/android-sdk-tmp

# Accept Android SDK licenses
RUN yes | sdkmanager --licenses

# Update SDK manager and install necessary SDK packages
RUN sdkmanager --update
RUN sdkmanager "platform-tools" "platforms;android-34" "build-tools;34.0.0"

# Set the working directory
WORKDIR /app

# Copy the application code into the container
COPY ./frontend/mobile/ /app

# Install project dependencies
RUN npm install

# Ensure gradlew has execute permissions
RUN chmod +x /app/android/gradlew

ENV NODE_ENV="production"

# Build the APK
RUN cd android && ./gradlew assembleRelease

RUN mkdir -p /shared

# Copy the generated APK to the shared directory
RUN cp android/app/build/outputs/apk/release/app-release.apk /shared/client.apk

# Declare a shared volume (placed after copying)
VOLUME ["/shared"]

# Keep the container running
CMD ["tail", "-f", "/dev/null"]
