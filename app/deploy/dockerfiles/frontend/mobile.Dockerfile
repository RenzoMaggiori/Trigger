# Use Node.js 18 with Debian Bullseye as the base image for amd64
FROM node:18-bullseye

# Set environment variables for Android SDK
ENV ANDROID_SDK_ROOT=/opt/android-sdk
ENV ANDROID_HOME=${ANDROID_SDK_ROOT}
ENV PATH=${PATH}:${ANDROID_SDK_ROOT}/cmdline-tools/latest/bin:${ANDROID_SDK_ROOT}/platform-tools

# Install necessary dependencies
RUN apt-get update && apt-get install -y openjdk-17-jdk wget unzip

# Create directories for Android SDK
RUN mkdir -p ${ANDROID_SDK_ROOT}/cmdline-tools/latest

# Download and extract Android command-line tools
RUN wget https://dl.google.com/android/repository/commandlinetools-linux-9477386_latest.zip -O /tmp/commandlinetools.zip \
    && unzip /tmp/commandlinetools.zip -d /tmp/ \
    && mv /tmp/cmdline-tools/* ${ANDROID_SDK_ROOT}/cmdline-tools/latest/ \
    && rm -rf /tmp/cmdline-tools /tmp/commandlinetools.zip

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

# Set environment variables
# ENV GRADLE_OPTS="-Dorg.gradle.daemon=false -Dorg.gradle.jvmargs=-Xmx2048m -Dfile.encoding=UTF-8"

# Ensure gradlew has execute permissions
# RUN chmod +x /app/android/gradlew

# Clean the project before building
# RUN cd android && ./gradlew clean

# Build the APK with detailed logs and disable parallel execution
# RUN cd android && ./gradlew assembleRelease --no-daemon --no-parallel --stacktrace --info

# To generate all the Android and IOS files
RUN npx expo prebuild

# If you want to sign the APK and publish to Google Play Store.
RUN npx react-native build-android --mode=release

# Create a directory to store the built APK and copy it there
RUN mkdir -p /app/dist
RUN cp /app/android/app/build/outputs/apk/release/app-release.apk /app/dist/

# Install a simple HTTP server to serve the APK
RUN npm install -g http-server

CMD ["tail", "-f", "/dev/null"]

