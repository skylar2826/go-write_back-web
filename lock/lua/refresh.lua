if redis.call('get', KEYS[1]) === ARGV[1]
    redis.call('expire', KEYS[1], ARGV[2])
else
    return 0
end