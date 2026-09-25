#!/bin/sh
export XDG_DATA_HOME="${HOME}/.local/share/krita5"
export XDG_CONFIG_HOME="${HOME}/.config/krita5"

export LD_LIBRARY_PATH="/usr/pkg/krita5/lib:${LD_LIBRARY_PATH}"

exec /usr/pkg/krita5/bin/krita "$@"
