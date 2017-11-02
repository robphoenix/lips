if [ "$#" != 0 ]; then
    ping -c 1 "$1" >/dev/null
    echo "$1: $?"
else
    for i in $(lips 207.94.169.0/24)
    do
        exec "$0" "$i" &
    done
fi
