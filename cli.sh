#!/usr/bin/env bash

case "$1" in
    "setup")
        # Lets create all the necessary symbolic links
        ln -s ../public run/public
    ;;
    *)
        echo "reporter build tool. This command provides a variety of commands to control the build process."
        echo ""
        echo "Commands:"
        echo "  * setup        | Sets up the reporter development environment with required symbolic links."
    ;;
esac
