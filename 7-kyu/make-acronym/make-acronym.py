def to_acronym(inp):
    acronym = "".join(word[0].upper() for word in inp.split())
    return acronym
    pass