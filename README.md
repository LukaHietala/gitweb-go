### Git server web interface

Simple web interface for git server that allows inspection of repositories, branches, commits, and files. This is a personal side project to learn more about git and golang.

#### Setting up a Git Server

On Debian-based systems:

```bash
sudo apt-get install git
sudo adduser git
sudo mkdir /home/git/repositories
cd /home/git/repositories
sudo git init --bare example.git
sudo chown -R git:git /home/git/repositories
```

Once configured, you can clone this repository (requires SSH setup):

```bash
git clone git@your-server:/home/git/repositories/example.git
```

For more details, refer to the [official Git documentation](https://git-scm.com/book/en/v2/Git-on-the-Server-The-Protocols).
