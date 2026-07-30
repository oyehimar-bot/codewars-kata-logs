def remove(s):
    if s == '':
        return s
    elif s[-1] == '!':
        return s[:-1]
    else:
        return s
    pass