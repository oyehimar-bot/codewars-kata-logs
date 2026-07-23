def initialize_names(names):
    parts = names.split()
    if len(parts) <= 2:
        # If only first and last name, return as is
        return names
    else:
        fname = parts[0]
        lname = parts[-1]
        initials = []
        # Initialize middle names
        for n in parts[1:-1]:
            if n[0].isupper():
                initials.append(n[0] + '.')
        # Join first name, initials, and last name
        return ' '.join([fname] + initials + [lname])
​