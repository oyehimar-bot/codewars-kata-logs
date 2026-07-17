def divisors(integer):
    list = []
    for i in range (2, integer):
        if integer % i == 0:
            list.append(i)
    if not list:
        return f'{integer} is prime'
    return list