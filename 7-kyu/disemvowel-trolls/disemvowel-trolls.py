def disemvowel(string_):
    res = ''
    for s in string_:
        if s == 'a' or s == 'e' or s == 'i' or s == 'o' or s == 'u' or s == 'A' or s == 'E' or s == 'I' or s == 'O' or s == 'U':
            continue
        res += str(s)
    return res