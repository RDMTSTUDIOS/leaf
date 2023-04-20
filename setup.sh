go build -o leaf main.go

if [ $? -eq 0 ]; then
    echo "Build done"
else
    echo "Build failed"
    exit 1
fi

sudo mv leaf /usr/local/bin
sudo chmod +x /usr/local/bin/leaf

if [[ ":$PATH:" != *":/usr/local/bin:"* ]]; then
    echo "Warning: /usr/local/bin is not in the PATH"
    echo "Add the following line to your shell startup file (e.g. .bashrc, .zshrc):"
    echo "export PATH=\$PATH:/usr/local/bin"
else
    echo "Installation successful"
fi
