# GoPanic: Panic hotkeys for Windows
If you have sensitive stuff on your computer, you probably already have made advances to secure that sensitive data from prying eyes, but if your computer gets yanked from your lap while you're logged in, almost all other security mechanisms, like full disk encryption, become void. This is what GoPanic aims to solve. GoPanic allows you to set up a hotkey that, when pressed, executes your specified action, for example, shutting down the computer or securely deleting files.

# Installation
Requirements:
- Windows
- Golang compiler in PATH

How to install:
1. `git clone https://github.com/SpoofIMEI/GoPanic` (or download as zip)
2. run install.bat
3. follow instructions

# Usage
### 1. Configure what the panic hotkey does.
Open panic.json and tell GoPanic what should happen when you press the hotkey.

Example1:
```json
{
  "commands": [
    "erase somefile.txt",
    "shutdown /s /t 10"
  ]
}
```


Example2:
```json
{
  "presets": {
    "sdelete": "C:/sensitivedir"
  }
}
```

Presets:
- sdelete (arg: directory or file)
- kill (arg: process name)
- shutdown

### 2. Run gopanic.exe
### 3. Test if it works 
