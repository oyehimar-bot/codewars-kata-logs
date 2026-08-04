def enough(cap, on, wait):
    # Your code here
    if (on + wait - cap) < 0:
        return 0
    return on + wait - cap