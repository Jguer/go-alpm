# Use the jguer/yay-builder image as a parent image with archlinux
FROM docker.io/jguer/yay-builder

# Install extra packages (pacman-contrib and fish)
# Note: Using -Sy instead of -Syu to avoid full system upgrade which fails
# due to read-only /etc/resolv.conf and /etc/hosts during Docker build
RUN sudo pacman -Sy --noconfirm pacman-contrib fish git-delta openssh bat go

# Set passwordless sudo for the docker user
RUN echo "docker ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/docker

# Add /app/bin to the PATH
ENV PATH="/app/bin:$PATH"

# Create a non-root user and switch to it
USER docker

# Set the working directory
WORKDIR /workspace

# Command to run when starting the container
CMD ["bash"]