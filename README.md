# 🚀 IMOTIF Tools – v1.0.9

> Generate Git commit messages interactively, quickly, and consistently.  
> Plus, run Odoo addon tests directly from your CLI.

`imotif-tools` is a cross-platform CLI tool that helps developers craft standardized commit messages, generate messages using AI, and now — run addon tests with Docker for Odoo.

---

## ✨ Features

- 🔍 **Interactive Prompt** for Task ID, commit type, and message
- 🧠 **AI-powered commit generation** with Gemini (OpenAI support coming soon)
- 🆔 Multiple Task IDs per commit (e.g. `OD-1,OD-2`)
- ✅ Optional `--no-verify` commit
- 🧪 **Run Odoo unit tests** with `imotif-tools test <addons>`
- 🛠️ Works on **macOS**, **Linux**, and **Windows**
- 📦 Built-in **self-update** mechanism
- ⚡ CLI alias support (e.g. `itcm`)

---

## 📦 Installation

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/0x-tao/imotif-tools/main/install.sh | bash
````

### Windows (PowerShell)

```powershell
irm https://raw.githubusercontent.com/0x-tao/imotif-tools/main/install.ps1 | iex
```

---

## ⚙️ Environment Setup

Create a `.env` file in the root directory (same location as your binary), and add:

```
GOOGLE_API_KEY=your_gemini_key
AI_PROVIDER=gemini
AI_MODEL=gemini-pro
```

---

## 🧪 Run Unit Tests (Odoo Addons)

Use this command to run unit tests for any Odoo addon using Docker Compose:

```bash
imotif-tools test addons
```

You can also run multiple addons by separating with commas:

```bash
imotif-tools test addons,addons
```

Behind the scenes, this uses Docker Compose with:

* Database: `odoo_test`
* Addons path: `/mnt/imbase/addons,/mnt/imbase/additional-addons`
* Coverage enabled with `pytest`

> 🔁 Automatically rebuilds and stops on container exit.

---

## 🚀 Usage

### Start interactive commit

```bash
imotif-tools commit
```

or use alias (after setup):

```bash
itcm
```

### Auto-generate commit message with AI

```bash
imotif-tools magic
```

> Suggests a commit message from staged code. You can confirm or edit it before committing.

### Update CLI

```bash
imotif-tools update
```

---

## 🛠️ Alias Setup (Optional)

To quickly access `imotif-tools`, create an alias:

### macOS / Linux (Zsh / Bash / Fish)

Add to `~/.zshrc` / `~/.bashrc` / `~/.config/fish/functions/itcm.fish`:

```bash
alias itcm='imotif-tools commit'
```

Then reload:

```bash
source ~/.zshrc
```

### Windows (PowerShell)

```powershell
function itcm {
    imotif-tools commit
}
```

Add that inside your `$PROFILE`, then run:

```powershell
. $PROFILE
```

---

## 🧠 Coming Soon

* OpenAI & Ollama support
* Auto-scan diffs
* GitHub/GitLab integration
* Git hooks and custom configs

---

## 🧪 Known Limitations

* Only supports staged files (`git add .`)
* AI commit supports Gemini only (currently)
* No rollback after commit (use `git commit --amend`)
* Docker must be installed for `test` command

---

## 🙏 Thank You

Thanks for trying `imotif-tools`!
Made with ❤️ by [@tao-thewarat](https://github.com/tao-thewarat)

```

หากคุณมีโฟลเดอร์ `docs/` อยู่ในโปรเจกต์ แนะนำให้วางไฟล์นี้ไว้ที่ `docs/README.md` ด้วยเช่นกันครับ ✍️
```
