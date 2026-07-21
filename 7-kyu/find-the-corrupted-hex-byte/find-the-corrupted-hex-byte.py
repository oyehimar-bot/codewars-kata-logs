HEX = set("0123456789ABCDEF")
​
def find_corrupted_byte(dump):
    for i, dum in enumerate(dump):
        if len(dum) != 2:
            return i
        for ch in dum:
            if ch not in HEX:
                return i
    return -1