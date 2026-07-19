def stray(arr):
    stray = 0
    for a in arr:
        stray ^= a
    return stray