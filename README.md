kage
====

Simple keylogger in Go/C for macOS

Usage
-----

Note that root privileges are required to run a keylogger. You'll also need to
enable by going to System Settings > Privacy & Security > Accessibility and
enabling whichever terminal application you'll use to run this.

```bash
git clone https://github.com/brymck/macos-keylogger.git
cd macos-keylogger
go build
sudo ./macos-keylogger
```
