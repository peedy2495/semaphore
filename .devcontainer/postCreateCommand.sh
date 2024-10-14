go install github.com/go-task/task/v3/cmd/task@latest

(cd ./web && npm install)

python3 -m venv .venv

./.venv/bin/pip3 install ansible

task build

./bin/semaphore user add \
    --admin \
    --login admin \
    --name Admin \
    --email admin@example.com \
    --password changeme \
    --config ./.devcontainer/config.json