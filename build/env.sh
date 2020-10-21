#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
emdir="$workspace/src/github.com/eminerchain"
if [ ! -L "$emdir/go-eminer" ]; then
    mkdir -p "$emdir"
    cd "$emdir"
    ln -s ../../../../../. go-eminer
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$emdir/go-eminer"
PWD="$emdir/go-eminer"

# Launch the arguments with the configured environment.
exec "$@"
